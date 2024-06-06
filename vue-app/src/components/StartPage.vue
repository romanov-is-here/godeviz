<script>
import {ref, watch, computed, reactive} from "vue";
import {VNetworkGraph} from "v-network-graph";
import * as vNG from "v-network-graph"
import "v-network-graph/lib/style.css"
import axios from 'axios';
import FilterBar from "@/components/FilterBar.vue";
import NodeCard from "@/components/NodeCard.vue";
import ErrorNotification from "@/components/ErrorNotification.vue";

export default {
  name: "StartPage",
  components: {
    FilterBar,
    VNetworkGraph,
    NodeCard,
    ErrorNotification
  },
  mounted() {
    this.showGraph()
  },
  setup: function () {
    const isLoaderVisible = ref(false)
    const isGraphVisible = ref(false)
    const selectedNodes = ref([])
    const selectedNode = ref()
    const errorMessage = ref('')

    const data = reactive({nodes : {}, edges:{}, moduleName: "loading"})
    const layouts = ref({})

    const showGraph = () => {
      isLoaderVisible.value = true
      fetchGraph()
    };

    const onFilterChanged = function (){
      showGraph()
    }

    const fetchGraph = async () => {
      try {
        let focusType = 0
        if (focusFilter.focusIns && focusFilter.focusOuts) {
          focusType = 2
        } else if (focusFilter.focusIns) {
          focusType = 1
        }

        let focusPackage = ""
        if (focusFilter.doFocus) {
          focusPackage = selectedNode.value?.id ?? ""
        }
        const response = await axios.get('/api/graph', {
          params: {
            show_standard : filter.showStandard,
            show_platform : filter.showPlatform,
            show_outer : filter.showOuter,
            focus_package: focusPackage,
            focus_type: focusType,
            show_ratio: filter.showRatio,
          }
        });

        data.nodes = response.data.nodes
        data.edges = response.data.edges
        data.moduleName = response.data.moduleName
        isGraphVisible.value = true
      } catch (error) {
        showErrorNotification("Cannot fetch graph: "+ error.message)
      }

      isLoaderVisible.value = false
    }

    const showErrorNotification = function(text) {
      const formattedTime = new Date().toLocaleTimeString('ru-ru', {hour12: false});
      errorMessage.value = formattedTime + " " + text
    }

    /////////////////// configs

    const nodeSize = 20
    const NODE_RADIUS = nodeSize / 2
    const letterWidth = 8
    const emojiWidth = 25
    const snap = 10

    const getNodeWidth = function (node) {
      let width = (node.name.length -1 ) * letterWidth + emojiWidth
      if (!filter.showRatio) {
        width += (node.fanInCount.toString().length + node.fanOutCount.toString().length + 2) * letterWidth
      }
      return Math.ceil(width / snap) * snap
    }

    const configs = vNG.defineConfigs({
      view: {
        layoutHandler: new vNG.GridLayout({ grid: snap }), // snap to grid
      },
      node: {
        selectable: true,
        normal: {
          radius: node => node.size ?? NODE_RADIUS,
          color: node => node.color,
          type: "rect",
          width: node => getNodeWidth(node)
        },
        hover: {
          color: node => node.color,
        },
        label: { color: "#000000", fontSize: 14 },
      },
      edge: {
        type: "curve",
        normal: {
          color: edge => edge.color,
          width: 2,
        },
        hover: {
          color: edge => edge.color,
        },
        margin: 4,
        marker: {
          target: {
            type: "arrow",
            width: 4,
            height: 4,
          },
        },
      },
    })

    ////////////// tooltips

    const graph = ref()
    const tooltip = ref()
    const targetNodeId = ref("")
    const tooltipOpacity = ref(0) // 0 or 1
    const tooltipPos = ref({ left: "0px", top: "0px" })

    const edgetip = ref()
    const edgetipOpacity = ref(0) // 0 or 1
    const edgetipPos = ref({ left: "0px", top: "0px" })
    const EDGE_MARGIN_TOP = 2
    const targetEdgeId = ref("")

    const targetNodePos = computed(() => {
      if (!targetNodeId.value) {
        return {x: 0, y: 0}
      }

      const nodePos = layouts.value.nodes[targetNodeId.value]
      return nodePos || {x: 0, y: 0}
    })

    const edgeCenterPos = computed(() => {
      const edge = data.edges[targetEdgeId.value]
      if (!edge) {
        return { x: 0, y: 0 }
      }

      const sourceNode = data.edges[targetEdgeId.value].source
      const targetNode = data.edges[targetEdgeId.value].target
      return {
        x: (layouts.value.nodes[sourceNode].x + layouts.value.nodes[targetNode].x) / 2,
        y: (layouts.value.nodes[sourceNode].y + layouts.value.nodes[targetNode].y) / 2,
      }
    })
    /////////// tooltip
    watch(
        () => [targetNodePos.value, tooltipOpacity.value],
        () => {
          if (!graph.value || !tooltip.value) return

          // translate coordinates: SVG -> DOM
          const domPoint = graph.value.translateFromSvgToDomCoordinates(targetNodePos.value)
          // calculates top-left position of the tooltip.
          tooltipPos.value = {
            left: domPoint.x - tooltip.value.offsetWidth / 2 + "px",
            top: domPoint.y - NODE_RADIUS - tooltip.value.offsetHeight - 10 + "px",
          }
        },
        { deep: true }
    )
    ////////////// edgetip
    watch(
        () => [edgeCenterPos.value, edgetipOpacity.value],
        () => {
          if (!graph.value || !edgetip.value) {
            return { x: 0, y: 0 }
          }
          if (!targetEdgeId.value) {
            return { x: 0, y: 0 }
          }

          // translate coordinates: SVG -> DOM
          const domPoint = graph.value.translateFromSvgToDomCoordinates(edgeCenterPos.value)
          // calculates top-left position of the tooltip.
          edgetipPos.value = {
            left: domPoint.x - edgetip.value.offsetWidth / 2 + "px",
            top: domPoint.y - EDGE_MARGIN_TOP - edgetip.value.offsetHeight - 10 + "px",
          }
        },
        { deep: true }
    )

    /////////// selectedNodesChanged
    watch(
        () => [selectedNodes.value],
        () => {
          if (selectedNodes.value.length === 0) {
            selectedNode.value = null
          } else {
            selectedNode.value = data.nodes[selectedNodes.value[0]]
          }
        },
        { deep: true }
    )

    //////////// events

    const oldColorSource = ref('')
    const oldColorTarget = ref('')
    const eventHandlers = {
      "node:pointerover": ({ node }) => {
        targetNodeId.value = node
        tooltipOpacity.value = 1 // show
      },
      "node:pointerout": ({node}) => {
        targetNodeId.value = node
        tooltipOpacity.value = 0 // hide
      },
      "edge:pointerover": ({ edge }) => {
        targetEdgeId.value = edge ?? ""
        const ed = data.edges[edge]
        oldColorSource.value = data.nodes[ed.source].color
        data.nodes[ed.source].color = "red"
        data.nodes[ed.source].size = NODE_RADIUS * 2
        oldColorTarget.value = data.nodes[ed.target].color
        data.nodes[ed.target].color = "red"
        data.nodes[ed.target].size = NODE_RADIUS * 2
        edgetipOpacity.value = 1 // show
      },
      "edge:pointerout": ({edge}) => {
        targetEdgeId.value = edge ?? ""
        edgetipOpacity.value = 0 // hide
        const ed = data.edges[edge]
        data.nodes[ed.source].color = oldColorSource.value
        delete data.nodes[ed.source].size
        data.nodes[ed.target].color = oldColorTarget.value
        delete data.nodes[ed.target].size
      },
    }

    const filter = reactive({
      showStandard: false,
      showPlatform: true,
      showOuter: true,
      showRatio: false
    })

    const focusFilter = reactive({
      focusIns: true,
      focusOuts: true,
      doFocus: false
    })

    return {
      isLoaderVisible,
      isGraphVisible,
      showGraph,
      data,
      layouts,
      configs,
      eventHandlers,
      tooltipPos,
      tooltipOpacity,
      targetNodeId,
      graph,
      tooltip,
      edgetip,
      targetEdgeId,
      edgetipPos,
      edgetipOpacity,
      edgeCenterPos,
      filter,
      onFilterChanged,
      selectedNodes,
      selectedNode,
      focusFilter,
      errorMessage,
    };
  }
}

</script>

<template>
  <ErrorNotification :error-message="errorMessage" />
  <div class="header">
    <div v-if="isGraphVisible">
      <div>
        <h5>
          Dependency graph for: {{data.moduleName}}
        </h5>
      </div>
    </div>

    <div v-if="isGraphVisible">
      <div class="filter-zone boxed">
        <FilterBar
            :filter = "filter"
            @filter-changed = "onFilterChanged"
        />
      </div>

      <div v-if="selectedNode" class="node-card boxed">
        <NodeCard
            :focusFilter = "focusFilter"
            :selectedNode = "selectedNode"
            @focus-request = "onFilterChanged"/>
      </div>
    </div>

    <Transition>
      <div v-if="isLoaderVisible" class="loader" v-bind:class="{ 'fade-enter': isLoaderVisible, 'fade-leave-to': !isLoaderVisible }">
        <div class="alert alert-success" role="alert">
          Loading data
        </div>
      </div>
    </Transition>
  </div>
  <div v-if="isGraphVisible" class="tooltip-wrapper">
    <div class="graph-zone boxed">
      <VNetworkGraph
          class="graph"
          ref="graph"
          v-model:layouts="layouts"
          v-model:selected-nodes="selectedNodes"
          :nodes="data.nodes"
          :edges="data.edges"
          :configs="configs"
          :event-handlers="eventHandlers"
      >
        <template
            #override-node-label="{
           scale, text
        }"
        >
          <text
              x="0"
              y="0"
              :font-size="14 * scale"
              text-anchor="middle"
              dominant-baseline="central"
              fill="#000000"
              font-family="monospace"
          >{{ text }}</text>
        </template>
      </VNetworkGraph>
    </div>
    <div
        ref="tooltip"
        class="tooltip"
        :style="{ ...tooltipPos, opacity: tooltipOpacity }"
    >
      <div>
        {{ data.nodes[targetNodeId]?.name ?? "" }}
        <br>
        FanIn: {{data.nodes[targetNodeId]?.fanInCount ?? -1}}
        <br/>
        FanOut: {{data.nodes[targetNodeId]?.fanOutCount ?? -1}}
      </div>
    </div>

    <div
        ref="edgetip"
        class="edgetip"
        :style="{ ...edgetipPos, opacity: edgetipOpacity }"
    >
      <div>
        {{data.nodes[data.edges[targetEdgeId]?.source]?.name ?? " ??"}} ->
        {{data.nodes[data.edges[targetEdgeId]?.target]?.name ?? " ??"}}
      </div>
    </div>
  </div>
</template>

<style scoped>
.boxed {
  border-radius: 5px;
  box-shadow: inset 0 0 1px rgba(0, 0, 0, 1);
}
.loader {
  width: 500px;
  height: 50px;
  margin: 0 auto;
}
.header {
  height: 185px;
  width: 100%;
  margin: 0 10px;
}
.filter-zone {
  float: left;
  width: 200px;
  height: 145px;
  margin: 0 5px 0 0;
}
.node-card {
  width: calc(100% - 225px); /* width - around filter width */
  height: 145px;
  float:left;
}
.graph-zone {
  height: calc(100vh - 200px);
  width: calc(100% - 20px);
  margin: 0 10px;
}
.tooltip-wrapper {
  position: relative;
}
.tooltip {
  top: 0;
  left: 0;
  opacity: 0;
  position: absolute;
  padding: 5px;
  width: auto;
  height: auto;
  display: grid;
  place-content: center;
  font-size: 12px;
  background-color: #fff0bd;
  border: 1px solid #ffb950;
  box-shadow: 2px 2px 2px #aaa;
  transition: opacity 0.2s linear;
  pointer-events: none;
  text-align: left;
}
.edgetip {
  top: 0;
  left: 0;
  opacity: 0;
  position: absolute;
  padding: 5px;
  width: auto;
  height: auto;
  display: grid;
  place-content: center;
  text-align: center;
  font-size: 12px;
  background-color: #fff0bd;
  border: 1px solid #ffb950;
  box-shadow: 2px 2px 2px #aaa;
  transition: opacity 0.2s linear;
  pointer-events: none;
}

.v-enter-active,
.v-leave-active {
  transition: opacity 0.5s ease;
}

.v-enter-from,
.v-leave-to {
  opacity: 0;
}
</style>
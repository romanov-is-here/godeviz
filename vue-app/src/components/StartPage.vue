<script>
import {ref, watch, computed, reactive} from "vue";
import {VNetworkGraph} from "v-network-graph";
import * as vNG from "v-network-graph"
import "v-network-graph/lib/style.css"
import axios from 'axios';
import FilterBar from "@/components/FilterBar.vue";

export default {
  name: "StartPage",
  components: {
    FilterBar,
    VNetworkGraph
  },
  setup: function () {
    const path = ref('/Users/ilyaromanov/Projects/platform-go/');
    const isInputsVisible = ref(true);
    const isLoaderVisible = ref(false);
    const isGraphVisible = ref(false)

    const data = reactive({nodes : {}, edges:{}})
    const layouts = ref({})

    const showGraph = () => {
      isInputsVisible.value = false
      isLoaderVisible.value = true
      fetchGraph()
    };

    const reset = () => {
      isLoaderVisible.value = false
      isGraphVisible.value = false
      isInputsVisible.value = true
    };

    const onFilterChanged = function (){
      showGraph()
    }

    const fetchGraph = async () => {
      try {
        const response = await axios.get('/api/graph', {
          params: {
            path: encodeURIComponent(path.value),
            show_standard : filter.showStandard,
            show_platform : filter.showPlatform,
            show_outer : filter.showOuter,
          }
        });

        data.nodes = response.data.nodes
        data.edges = response.data.edges
        isGraphVisible.value = true
      } catch (error) {
        console.error('Ошибка при получении данных:', error);
      }

      isLoaderVisible.value = false
    }

    /////////////////// configs

    const nodeSize = 20
    const NODE_RADIUS = nodeSize / 2
    const letterWidth = 8
    const emojiWidth = 20

    const configs = vNG.defineConfigs({
      view: {
        layoutHandler: new vNG.GridLayout({ grid: 10 }), // snap to grid
      },
      node: {
        normal: {
          radius: node => node.size ?? NODE_RADIUS,
          color: node => node.color,
          type: "rect",
          width: node => node.name.length * letterWidth + emojiWidth
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
      showOuter: true
    })

    return {
      path,
      isInputsVisible,
      isLoaderVisible,
      isGraphVisible,
      showGraph,
      reset,
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
      onFilterChanged
    };
  }
}

</script>

<template>
  <div class="header">
    <div v-if="isGraphVisible">
      <FilterBar
          :filter = "filter"
          @filter-changed = "onFilterChanged"
      />
    </div>
    <div v-if="isInputsVisible" class="centered-div">
      <div class="alert alert-success" role="alert">
        Enter a path to the package:
      </div>

      <input v-model="path" type="text" class="form-control" placeholder="Your path here" aria-label="Path">

      <button type="button" class="btn btn-primary show-btn" @click="showGraph">Show graph</button>
    </div>

    <div v-if="!isInputsVisible && !isLoaderVisible">
      <h4>{{path}}</h4>
      <br>
      <button type="button" class="btn btn-primary" @click="reset">Reset</button>
    </div>

    <div v-if="isLoaderVisible" class="centered-div">
      <div class="alert alert-success" role="alert">
        Loading data
      </div>
    </div>
  </div>
  <div v-if="isGraphVisible" class="tooltip-wrapper graph-zone">
    <VNetworkGraph
        class="graph"
        ref="graph"
        v-model:layouts="layouts"
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
.centered-div {
  width: 500px;
  margin: 0 auto;
}
.show-btn {
  margin: 16px
}
.graph-zone {
  height: calc(100vh - 300px);
  width: calc(100% - 20px);
  border: 3px solid gray; /* Толщина 3 пикселя, серый цвет */
  border-radius: 10px; /* Скругленные края */
  margin: 10px;
}
.header {
  height: 200px;
  width: 100%;
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
</style>
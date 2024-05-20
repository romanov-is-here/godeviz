<script>
import {ref, watch, computed, reactive} from "vue";
import {VNetworkGraph} from "v-network-graph";
import * as vNG from "v-network-graph"
import "v-network-graph/lib/style.css"
import axios from 'axios';
import dagre from "dagre"

export default {
  name: "StartPage",
  components: {
    VNetworkGraph
  },
  setup: function () {
    const path = ref('/Users/ilyaromanov/Projects/platform-go/cmd/devp');
    const isInputsVisible = ref(true);
    const isLoaderVisible = ref(false);
    const isGraphVisible = ref(false)

    const data = reactive({nodes : {}, edges:{}})
    const layouts = reactive({
      nodes: {},
    })

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

    const fetchGraph = async () => {
      try {
        const response = await axios.get('/api/graph', {
          params: {
            path: encodeURIComponent(path.value)
          }
        });

        data.nodes = response.data.nodes
        data.edges = response.data.edges
        layout("TB")
        isGraphVisible.value = true
      } catch (error) {
        console.error('Ошибка при получении данных:', error);
      }

      isLoaderVisible.value = false
    }

    /////////////////// configs

    const nodeSize = 20
    const NODE_RADIUS = nodeSize / 2

    const configs = vNG.defineConfigs({
      view: {
        autoPanAndZoomOnLoad: "fit-content",
        onBeforeInitialDisplay: () => layout("TB"),
      },
      node: {
        normal: {
          radius: node => node.size ?? NODE_RADIUS,
          color: node => node.color,

        },
        hover: {
          color: node => node.color,
        },
        label: { direction: "south", color: "#000000", directionAutoAdjustment: true, fontSize: 14 },
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
      const nodePos = layouts.nodes[targetNodeId.value]
      return nodePos || {x: 0, y: 0}
    })

    const edgeCenterPos = computed(() => {
      const edge = data.edges[targetEdgeId.value]
      if (!edge) {
        console.log("edge not found")
        return { x: 0, y: 0 }
      }

      console.log("edge found")
      const sourceNode = data.edges[targetEdgeId.value].source
      const targetNode = data.edges[targetEdgeId.value].target
      return {
        x: (layouts.nodes[sourceNode].x + layouts.nodes[targetNode].x) / 2,
        y: (layouts.nodes[sourceNode].y + layouts.nodes[targetNode].y) / 2,
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
            console.log("watch graph or edgetip not found")
            return { x: 0, y: 0 }
          }
          if (!targetEdgeId.value) {
            console.log("targetEdgeId no value")
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

    ////////////// layout

    function layout(direction) {
      if (Object.keys(data.nodes).length <= 1 || Object.keys(data.edges).length == 0) {
        return
      }

      // convert graph
      // ref: https://github.com/dagrejs/dagre/wiki
      const g = new dagre.graphlib.Graph()
      // Set an object for the graph label
      g.setGraph({
        rankdir: direction,
        nodesep: nodeSize * 5,
        edgesep: nodeSize * 2,
        ranksep: nodeSize * 5,
      })
      // Default to assigning a new object as a label for each new edge.
      g.setDefaultEdgeLabel(() => ({}))

      // Add nodes to the graph. The first argument is the node id. The second is
      // metadata about the node. In this case we're going to add labels to each of
      // our nodes.
      Object.entries(data.nodes).forEach(([nodeId, node]) => {
        g.setNode(nodeId, { label: node.name, width: nodeSize, height: nodeSize })
      })

      // Add edges to the graph.
      Object.values(data.edges).forEach(edge => {
        g.setEdge(edge.source, edge.target)
      })

      dagre.layout(g)

      g.nodes().forEach((nodeId) => {
        // update node position
        const x = g.node(nodeId).x
        const y = g.node(nodeId).y
        layouts.nodes[nodeId] = { x, y }
      })
    }

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
      edgeCenterPos
    };
  }
}

</script>

<template>
  <div v-if="isInputsVisible" class="centered-div">
    <div class="alert alert-success" role="alert">
      Enter a path to the package:
    </div>

    <input v-model="path" type="text" class="form-control" placeholder="Your path here" aria-label="Path">

    <button type="button" class="btn btn-primary show-btn" @click="showGraph">Show graph</button>

    <h3>{{path}}</h3>
  </div>

  <div v-if="!isInputsVisible && !isLoaderVisible">
    <button type="button" class="btn btn-primary" @click="reset">Reset</button>
  </div>

  <div v-if="isLoaderVisible" class="centered-div">
    <div class="alert alert-success" role="alert">
      Loading data
    </div>
  </div>

  <div v-if="isGraphVisible" class="tooltip-wrapper full-page">
    <VNetworkGraph
        class="graph"
        ref="graph"
        :nodes="data.nodes"
        :edges="data.edges"
        :layouts="layouts"
        :configs="configs"
        :event-handlers="eventHandlers"
    />
    <div
        ref="tooltip"
        class="tooltip"
        :style="{ ...tooltipPos, opacity: tooltipOpacity }"
    >
      <div>
        {{ data.nodes[targetNodeId]?.name ?? "" }}
        <br>
        Imports: {{data.nodes[targetNodeId]?.outDeps ?? 0}}
        <br>
        In deps: {{data.nodes[targetNodeId]?.inDeps ?? 0}}
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
.full-page {
  height: 100vh;
  width: 100%;
  border: 3px solid gray; /* Толщина 3 пикселя, серый цвет */
  border-radius: 10px; /* Скругленные края */
  margin: 10px
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
<script>
import {ref} from "vue";
import {reactive} from "vue";
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

    ///////////////////

    const nodeSize = 20

    const configs = vNG.defineConfigs({
      view: {
        autoPanAndZoomOnLoad: "fit-content",
        onBeforeInitialDisplay: () => layout("TB"),
      },
      node: {
        normal: {
          radius: nodeSize / 2,
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
      configs
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

  <div v-if="isGraphVisible" class="full-page">
    <VNetworkGraph
        class="graph"
        :nodes="data.nodes"
        :edges="data.edges"
        :layouts="layouts"
        :configs="configs"
    />
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
}
</style>
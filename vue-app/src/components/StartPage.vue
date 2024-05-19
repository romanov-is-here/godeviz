<script>
import {ref} from "vue";
import {reactive} from "vue";
import {VNetworkGraph} from "v-network-graph";
import "v-network-graph/lib/style.css"
import axios from 'axios';

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

    const graph = reactive({nodes : {}, edges:{}})
    // const nodes = {
    //   node1: { name: "Node 1" },
    //   node2: { name: "Node 2" },
    // }
    //
    // const edges = {
    //   edge1: { source: "node1", target: "node2" },
    // }

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

        //this.myData = response.data;
        console.log("received", response.data)
        graph.nodes = response.data.nodes
        graph.edges = response.data.edges
        isGraphVisible.value = true
      } catch (error) {
        console.error('Ошибка при получении данных:', error);
      }

      isLoaderVisible.value = false
    }

    return {
      path,
      isInputsVisible,
      isLoaderVisible,
      isGraphVisible,
      showGraph,
      reset,
      graph
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
        :nodes="graph.nodes"
        :edges="graph.edges"
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
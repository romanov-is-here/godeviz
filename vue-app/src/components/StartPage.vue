<script>
import {ref} from "vue";
import {VNetworkGraph} from "v-network-graph";
import "v-network-graph/lib/style.css"
export default {
  name: "StartPage",
  components: {
    VNetworkGraph
  },
  setup: function () {
    const path = ref('');
    const isInputsVisible = ref(true);
    const isLoaderVisible = ref(false);
    const isGraphVisible = ref(false)

    const nodes = {
      node1: { name: "Node 1" },
      node2: { name: "Node 2" },
    }

    const edges = {
      edge1: { source: "node1", target: "node2" },
    }

    const showGraph = () => {
      hideInputs()
      isLoaderVisible.value = true
      isGraphVisible.value = true
    };

    const reset = () => {
      isLoaderVisible.value = false
      isGraphVisible.value = false
      showInputs()
    };

    const showInputs = () => {
      isInputsVisible.value = true
      isGraphVisible.value = true
    };

    const hideInputs = () => {
      isInputsVisible.value = false
    };

    return {
      path,
      isInputsVisible,
      isLoaderVisible,
      isGraphVisible,
      showGraph,
      reset,
      nodes,
      edges
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

  <div v-if="isGraphVisible">
    <VNetworkGraph
        class="graph"
        :nodes="nodes"
        :edges="edges"
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
</style>
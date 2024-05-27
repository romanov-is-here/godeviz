<script>
export default {
  name: "NodeCard",
  props: ["focusFilter", "selectedNode"],
  computed: {
    filterLocal: {
      get: function() {
        return this.focusFilter
      }
    }
  },
  methods:{
    fireFocusRequest(filterValue, doFocus) {
      filterValue.doFocus = doFocus
      this.$emit('focus-request', filterValue);
    }
  }
}
</script>

<template>
  <div v-if="selectedNode">
    <div class="header">
      {{selectedNode.name}}
    </div>

    <div class="filter-container">
      <form>
        <div>
          <div class="form-check form-switch">
            <input class="form-check-input" type="checkbox" v-model="filterLocal.focusIns" id="focusIns" required>
            <label class="form-check-label" for="focusIns">Show dependents</label>
          </div>

          <div class="form-check form-switch">
            <input class="form-check-input" type="checkbox" v-model="filterLocal.focusOuts" id="focusOuts" required>
            <label class="form-check-label" for="focusOuts">Show imports</label>
          </div>
        </div>
        <button type="button" class="btn btn-primary" @click="fireFocusRequest(filterLocal, true)">Focus</button>
        <button type="button" class="btn btn-primary" @click="fireFocusRequest(filterLocal, false)">Reset focus</button>
      </form>
    </div>
  </div>
</template>

<style scoped>
.btn {
  margin: 5px
}
.filter-container {
  display: flex
}
.form-check {
   margin: 5px;
   text-align: left;
}
.header {
  text-align: left;
  margin: 5px;
}
</style>
<script>

export default {
  name: "FilterBar",
  props: ["filter"],
  computed: {
    filterLocal: {
      get: function() {
        return this.filter
      }
    }
  },
  created() {
    this.$watch('filterLocal', newValue => {
      this.$emit('filter-changed', newValue);
    }, { deep: true });
  }
}
</script>

<template>
  <div class="filter-container">
    <form>
      <div class="form-check form-switch">
        <input class="form-check-input" type="checkbox" v-model="filterLocal.showStandard" id="showStandard">
        <label class="form-check-label" for="showStandard">Show standard</label>
      </div>

      <div class="form-check form-switch">
        <input class="form-check-input" type="checkbox" v-model="filterLocal.showPlatform" id="showPlatform">
        <label class="form-check-label" for="showPlatform">Show platform</label>
      </div>

      <div class="form-check form-switch">
        <input class="form-check-input" type="checkbox" v-model="filterLocal.showOuter" id="showOuter">
        <label class="form-check-label" for="showOuter">Show outer</label>
      </div>

      <div class="form-check form-switch">
        <input class="form-check-input" type="checkbox" v-model="filterLocal.showRatio" id="showRatio">
        <label class="form-check-label" for="showRatio">Show ratio</label>
      </div>
    </form>
  </div>
</template>


<style scoped>
.filter-container {
  display: flex
}

.form-check {
  margin: 5px;
  text-align: left;
}
</style>
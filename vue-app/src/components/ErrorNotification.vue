<script>
import {ref} from "vue";

export default {
  name: "ErrorNotification",
  props: {
    errorMessage: {
      required: true
    }
  },
  watch:{
    errorMessage(newValue){
      if (newValue.length > 0)
      {
        this.showError()
      }
    }
  },
  setup: function() {
    const showMessage = ref(false)

    const closeError = function() {
      showMessage.value = false
    }

    const timeout = ref(5000)

    const showError = function(){
      showMessage.value = true
      setTimeout(() => {
        closeError();
      }, timeout.value);
    }

    return {
      showMessage,
      closeError,
      showError,
      timeout,
    }
  },
};
</script>

<template>
  <div v-if="showMessage">
    <div class="error-notification alert alert-danger">
      <button class="btn btn-close top-right" @click="closeError"/>
      <p>{{ $props.errorMessage }}</p>
    </div>
  </div>
</template>

<style scoped>
.error-notification {
  position: fixed;
  top: 20px;
  right: 20px;
  padding: 10px;
  border: 1px solid #ff0000;
  max-width: 300px;
  z-index: 999;
}
.error-notification p {
  padding-top: 30px;
  word-wrap: break-word;
}
.top-right {
  position: absolute;
  top: 5px;
  right: 5px;
  background-color: #ccc;
  color: #fff;
  border: none;
  padding: 5px 10px;
  cursor: pointer;

}
</style>

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
  <div v-if="showMessage" class="error-notification">
    <p>{{ $props.errorMessage }}</p>
    <button @click="closeError">Закрыть</button>
  </div>
</template>

<style scoped>
.error-notification {
  position: fixed;
  top: 20px;
  right: 20px;
  background-color: #ffcccc;
  padding: 10px;
  border: 1px solid #ff0000;
}
</style>

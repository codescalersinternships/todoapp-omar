<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
  props: ['errMsg']
})
</script>

<template>
  <div>
    <h1>{{ errMsg }}</h1>
  </div>
</template>

<style lang="css" scoped>
h1 {
  color: red;
}
</style>

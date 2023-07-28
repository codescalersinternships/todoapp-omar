<script setup lang="ts">
import TodoComponent from '../components/TodoComponent.vue'
</script>

<template>
  <main>
    <h1>Todo App</h1>
    <TodoComponent></TodoComponent>
  </main>
</template>

<style lang="css" scoped>
main {
  display: flex;
  flex-direction: column;
  max-width: 500px;
  width: 100%;
  margin: 0 auto;
  padding: 40px 16px;

  h1 {
    margin-bottom: 16px;
    text-align: center;
  }
}
</style>

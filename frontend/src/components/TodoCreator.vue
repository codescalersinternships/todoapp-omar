<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
  data() {
    return {
      taskTitle: ''
    }
  },
  methods: {
    async handleAddTask() {
      if (this.taskTitle == '') return
      try {
        const response = await this.axios({
          method: 'POST',
          url: 'http://localhost:8080/task',
          data: {
            title: this.taskTitle
          }
        })

        this.taskTitle = ''
        this.$emit('addTask', response.data)
      } catch {
        this.$emit('showErr', 'could not add new task')
      }
    }
  }
})
</script>

<template>
  <div class="input-wrap">
    <input type="text" v-model="taskTitle" />
    <button @click="handleAddTask">Add</button>
  </div>
</template>

<style lang="css" scoped>
.input-wrap {
  display: flex;
}

input {
  width: 100%;
  padding: 8px 6px;
  border: 2px solid #171d1b;
  &:focus {
    outline: none;
  }
}

button {
  padding: 8px 16px;
  font-weight: 800;
  border: none;
}
</style>

<script lang="ts">
import { defineComponent, ref } from 'vue'
import TodoCreator from './TodoCreator.vue'
import TodoTask from './TodoTask.vue'
import ErrorComponent from './ErrorComponent.vue'

declare interface Task {
  id: number
  title: string
  is_completed: boolean
}

export default defineComponent({
  data() {
    return {
      tasks: [] as Task[],
      showErrMsg: ref(false),
      errMsg: ''
    }
  },
  async beforeMount() {
    try {
      const response = await this.axios.get('http://localhost:8080/task')

      this.tasks = response.data
    } catch {
      this.showErr('could not get your tasks')
    }
  },
  methods: {
    editTask(updatedTask: Task) {
      this.tasks = this.tasks.map((task: Task) => (task.id === updatedTask.id ? updatedTask : task))
    },
    deleteTask(taskId: number) {
      this.tasks = this.tasks.filter((task: Task) => task.id !== taskId)
    },
    addTask(newTask: Task) {
      this.tasks.unshift(newTask)
    },
    showErr(errMsg: string) {
      this.errMsg = errMsg
      this.showErrMsg = true
    }
  },
  components: { TodoCreator, TodoTask, ErrorComponent }
})
</script>

<template>
  <div>
    <TodoCreator @addTask="addTask" @showErr="showErr"></TodoCreator>
    <ErrorComponent v-if="showErrMsg" :errMsg="errMsg"></ErrorComponent>
    <ul>
      <li v-for="task in tasks" :key="task.id">
        <TodoTask
          :task="task"
          @editTask="editTask"
          @deleteTask="deleteTask"
          @showErr="showErr"
        ></TodoTask>
      </li>
    </ul>
  </div>
</template>

<style lang="css" scoped>
ul {
  list-style-type: none;
  padding: 0px;
}
</style>

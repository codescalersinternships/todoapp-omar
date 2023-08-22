<script lang="ts">
import { defineComponent, ref } from 'vue'
import TaskCreator from '@/components/TaskCreator.vue'
import TaskTile from '@/components/TaskTile.vue'
import ErrorIndicator from '@/components/ErrorIndicator.vue'

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
      errMsg: '',
      apiBaseUrl: import.meta.env.VITE_API_BASE_URL
    }
  },
  async beforeMount() {
    try {
      const response = await this.axios.get(`${this.apiBaseUrl}/task`)

      this.tasks = response.data.tasks
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
  components: { TaskCreator, TaskTile, ErrorIndicator }
})
</script>

<template>
  <main>
    <h1>Todo App</h1>
    <div>
      <TaskCreator @addTask="addTask" @showErr="showErr"></TaskCreator>
      <ErrorIndicator v-if="showErrMsg" :errMsg="errMsg"></ErrorIndicator>
      <ul>
        <li v-for="task in tasks" :key="task.id">
          <TaskTile
            :task="task"
            @editTask="editTask"
            @deleteTask="deleteTask"
            @showErr="showErr"
          ></TaskTile>
        </li>
      </ul>
    </div>
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
}
h1 {
  margin-bottom: 16px;
  text-align: center;
}
ul {
  list-style-type: none;
  padding: 0px;
}
</style>

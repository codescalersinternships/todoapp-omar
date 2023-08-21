<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
  data() {
    return {
      currentTaskTitle: this.task.title,
      isTaskCompleted: this.task.is_completed,
      isTaskEditEnabled: false,
      apiBaseUrl: import.meta.env.VITE_API_BASE_URL
    }
  },
  props: ['task'],
  methods: {
    async deleteTask() {
      try {
        await this.axios({
          method: 'DELETE',
          url: this.apiBaseUrl + `/task/${this.task.id}`
        })

        this.$emit('deleteTask', this.task.id)
      } catch {
        this.$emit('showErr', 'could not delete task')
      }
    },
    async editTask() {
      this.isTaskEditEnabled = false

      const updatedTask = {
        id: this.task.id,
        title: this.currentTaskTitle,
        is_completed: this.isTaskCompleted
      }
      try {
        await this.axios({
          method: 'PUT',
          url: this.apiBaseUrl + `/task/${this.task.id}`,
          data: {
            id: this.task.id,
            title: this.currentTaskTitle,
            is_completed: this.isTaskCompleted
          }
        })

        this.$emit('editTask', updatedTask)
      } catch {
        this.$emit('showErr', 'could not edit task')
      }
    },
    handleToEditTask() {
      this.isTaskEditEnabled = true
    }
  }
})
</script>

<template>
  <div class="task-wrapper">
    <input type="checkbox" v-model="isTaskCompleted" @change="editTask" />

    <input v-if="isTaskEditEnabled" type="text" v-model="currentTaskTitle" />
    <span v-else :class="{ 'task-done': isTaskCompleted }" class="task-title">{{
      currentTaskTitle
    }}</span>

    <div class="divider"></div>

    <span v-if="isTaskEditEnabled">
      <svg
        xmlns="http://www.w3.org/2000/svg"
        height="1.5em"
        viewBox="0 0 448 512"
        @click="editTask"
      >
        <path
          d="M438.6 105.4c12.5 12.5 12.5 32.8 0 45.3l-256 256c-12.5 12.5-32.8 12.5-45.3 0l-128-128c-12.5-12.5-12.5-32.8 0-45.3s32.8-12.5 45.3 0L160 338.7 393.4 105.4c12.5-12.5 32.8-12.5 45.3 0z"
        />
      </svg>
    </span>
    <span v-else>
      <svg
        xmlns="http://www.w3.org/2000/svg"
        height="1.5em"
        viewBox="0 0 512 512"
        @click="handleToEditTask"
      >
        <path
          d="M471.6 21.7c-21.9-21.9-57.3-21.9-79.2 0L362.3 51.7l97.9 97.9 30.1-30.1c21.9-21.9 21.9-57.3 0-79.2L471.6 21.7zm-299.2 220c-6.1 6.1-10.8 13.6-13.5 21.9l-29.6 88.8c-2.9 8.6-.6 18.1 5.8 24.6s15.9 8.7 24.6 5.8l88.8-29.6c8.2-2.7 15.7-7.4 21.9-13.5L437.7 172.3 339.7 74.3 172.4 241.7zM96 64C43 64 0 107 0 160V416c0 53 43 96 96 96H352c53 0 96-43 96-96V320c0-17.7-14.3-32-32-32s-32 14.3-32 32v96c0 17.7-14.3 32-32 32H96c-17.7 0-32-14.3-32-32V160c0-17.7 14.3-32 32-32h96c17.7 0 32-14.3 32-32s-14.3-32-32-32H96z"
        />
      </svg>
    </span>

    <span>
      <svg
        xmlns="http://www.w3.org/2000/svg"
        height="1.5em"
        viewBox="0 0 448 512"
        @click="deleteTask"
      >
        <path
          d="M135.2 17.7L128 32H32C14.3 32 0 46.3 0 64S14.3 96 32 96H416c17.7 0 32-14.3 32-32s-14.3-32-32-32H320l-7.2-14.3C307.4 6.8 296.3 0 284.2 0H163.8c-12.1 0-23.2 6.8-28.6 17.7zM416 128H32L53.2 467c1.6 25.3 22.6 45 47.9 45H346.9c25.3 0 46.3-19.7 47.9-45L416 128z"
        />
      </svg>
    </span>
  </div>
</template>

<style lang="css" scoped>
.task-wrapper {
  display: flex;
  gap: 8px;
  padding: 0px 60px 18px 0px;
}
.divider {
  flex: 1;
}

.task-done {
  text-decoration: line-through;
}

.task-title {
  font-size: 1.5em;
}
</style>

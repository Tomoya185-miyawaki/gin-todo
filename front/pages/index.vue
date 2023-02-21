<template>
  <div>
    <p v-if="$fetchState.pending">Loading...</p>
    <p v-else-if="$fetchState.error">Error: {{ $fetchState.error.message }}</p>
    <div v-else>
      <ul>
        <li v-for="(todo, index) in todos" :key="index">
          <nuxt-link :to="`/todo/${todo.todoId}`">{{ todo.todoTitle }}</nuxt-link>
        </li>
      </ul>
      <button><nuxt-link to="/todo/create">作成する</nuxt-link></button>
    </div>
  </div>
</template>

<script>
export default {
  name: 'IndexPage',
  data() {
    return {
      todos: []
    }
  },
  fetchOnServer: false,
  async fetch() {
    this.todos = await this.$axios.get(`${process.env.baseAPIUrl}/todos`).then(res =>
      res.data
    )
  }
}
</script>

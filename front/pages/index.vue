<template>
  <p v-if="$fetchState.pending">Loading...</p>
  <p v-else-if="$fetchState.error">Error: {{ $fetchState.error.message }}</p>
  <ul v-else>
    <li v-for="(todo, index) in todos" :key="index">
      <nuxt-link :to="`/todo/${todo.todoId}`">{{ todo.todoTitle }}</nuxt-link>
    </li>
  </ul>
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
    this.todos = await fetch(`${process.env.baseAPIUrl}/todos`).then(res =>
      res.json()
    )
  }  
}
</script>

<template>
  <div>
    <p v-if="$fetchState.pending">Loading...</p>
    <p v-else-if="$fetchState.error">Error: {{ $fetchState.error.message }}</p>
    <ul v-else>
      <li>
        ID: {{ todo.todoId }}
      </li>
      <li>
        Title: {{ todo.todoTitle }}
      </li>
    </ul>
    <nuxt-link to="/">一覧へ戻る</nuxt-link>
  </div>
</template>

<script>
export default {
  name: 'TodoIdPage',
  data() {
    return {
      todo: []
    }
  },
  fetchOnServer: false,
  async fetch() {
    this.todo = await fetch(`${process.env.baseAPIUrl}/todo/${this.$route.params.id}`).then(res =>
      res.json()
    )
  }  
}
</script>

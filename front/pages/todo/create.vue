<template>
  <div>
    <div style="color: red;">{{ message }}</div>
    <form method="POST" @submit.prevent="create">
      <label for="title">タイトル：</label>
      <input type="title" v-model="title" placeholder="タイトル" id="title" required>
      <br/>
      <input type="submit" value="作成">
    </form>
    <nuxt-link to="/">一覧へ戻る</nuxt-link>
  </div>
</template>

<script>
export default {
  name: 'TodoIdPage',
  data() {
    return {
      title: '',
      message: ''
    }
  },
  methods: {
    async create() {
      let params = {
        'title': this.title
      }
      await this.$axios.post(`${process.env.baseAPIUrl}/todo/create`, params).then(res => {
        this.message = 'TODOの作成に成功しました'
        this.title = ''
      }).catch(err => {
        this.message = 'TODOの作成に失敗しました'
        console.warn(err.response)
      })
    }  
  },
}
</script>

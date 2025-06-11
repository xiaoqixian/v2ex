<template>
  <div class="register-container">
    <h2>注册</h2>
    <form @submit.prevent="handleRegister">
      <div class="form-group">
        <label>用户名</label>
        <input v-model="form.username" type="text" required />
      </div>
      <div class="form-group">
        <label>邮箱</label>
        <input v-model="form.email" type="email" required />
      </div>
      <div class="form-group">
        <label>密码</label>
        <input v-model="form.password" type="password" required />
      </div>
      <button type="submit">注册</button>
    </form>
  </div>
</template>

<script setup>
import { reactive } from 'vue'

const form = reactive({
  username: '',
  email: '',
  password: '',
})

function handleRegister() {
  console.log('注册信息:', form)
  // 在这里你可以调用 API 进行注册提交
}
</script>

<style scoped>
.register-container {
  max-width: 400px;
  margin: 50px auto;
  padding: 2em;
  border: 1px solid #ccc;
  border-radius: 8px;
}
.form-group {
  margin-bottom: 1em;
}
label {
  display: block;
  margin-bottom: 0.25em;
}
input {
  width: 100%;
  padding: 0.5em;
  box-sizing: border-box;
}
button {
  width: 100%;
  padding: 0.75em;
  background-color: #42b983;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}
</style>

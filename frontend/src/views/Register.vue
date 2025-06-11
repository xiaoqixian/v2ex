<template>
  <form class="login" @submit.prevent="handleRegister">
    <input v-model="form.username" type="text" placeholder="用户名" required />
    <input v-model="form.email" type="email" placeholder="邮箱" required />
    <input v-model="form.password" type="password" placeholder="密码" required />
    <input v-model="form.confirmPassword" type="password" placeholder="再次确认密码" required />
    <button type="submit">Register</button>
  </form>
</template>

<script setup>
import { reactive } from 'vue'
import axios from 'axios'

const form = reactive({
  username: '',
  email: '',
  password: '',
  confirmPassword: ''
})

async function handleRegister() {
  if (form.password !== form.confirmPassword) {
    alert('两次密码不一致')
    return
  }

  try {
    const res = await axios.post('/api/register', {
      username: form.username,
      email: form.email,
      password: form.password
    })
    alert('注册成功')
    console.log(res.data)
  } catch (err) {
    console.error(err)
    alert('注册失败')
  }
}
</script>

<style scoped>
body {
  background-color: #f45b69;
  font-family: "Asap", sans-serif;
}

.login {
  overflow: hidden;
  background-color: white;
  padding: 40px 30px 30px 30px;
  border-radius: 10px;
  position: absolute;
  top: 50%;
  left: 50%;
  width: 400px;
  transform: translate(-50%, -50%);
  transition: transform 300ms, box-shadow 300ms;
  box-shadow: 5px 10px 10px rgba(2, 128, 144, 0.2);
}
.login::before, .login::after {
  content: "";
  position: absolute;
  width: 600px;
  height: 600px;
  border-top-left-radius: 40%;
  border-top-right-radius: 45%;
  border-bottom-left-radius: 35%;
  border-bottom-right-radius: 40%;
  z-index: -1;
}
.login::before {
  left: 40%;
  bottom: -130%;
  background-color: rgba(69, 105, 144, 0.15);
  animation: wawes 6s infinite linear;
}
.login::after {
  left: 35%;
  bottom: -125%;
  background-color: rgba(2, 128, 144, 0.2);
  animation: wawes 7s infinite;
}
.login > input {
  font-family: "Asap", sans-serif;
  display: block;
  border-radius: 5px;
  font-size: 16px;
  background: white;
  width: 100%;
  border: 1px solid #000;
  padding: 10px 10px;
  margin: 15px -10px;
}
.login > button {
  font-family: "Asap", sans-serif;
  cursor: pointer;
  color: #fff;
  font-size: 14px;
  text-transform: uppercase;
  width: 80px;
  border: 0;
  padding: 10px 0;
  margin-top: 10px;
  margin-left: -5px;
  border-radius: 5px;
  background-color: #f45b69;
  transition: background-color 300ms;
}
.login > button:hover {
  background-color: #f24353;
}

@keyframes wawes {
  from {
    transform: rotate(0);
  }
  to {
    transform: rotate(360deg);
  }
}

a {
  text-decoration: none;
  color: rgba(255, 255, 255, 0.6);
  position: absolute;
  right: 10px;
  bottom: 10px;
  font-size: 12px;
}
</style>

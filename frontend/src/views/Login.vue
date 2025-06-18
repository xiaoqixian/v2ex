<template>
  <form v-if="!isRegistered" class="login" @submit.prevent="handleLogin">
    <input v-model="form.username" type="text" placeholder="用户名" required />
    <input v-model="form.password" type="password" placeholder="密码" required />
    <button type="submit">登录</button>
  </form>
</template>

<script setup>
import { reactive, ref } from 'vue'
import axios from 'axios'
import { useRouter } from 'vue-router'
import { useUserStore } from '@/stores/user'

const form = reactive({
  username: '',
  email: '',
  password: '',
  confirmPassword: ''
})
const router = useRouter()
const userStore = useUserStore()

async function handleLogin() {
  try {
    const res = await axios.post('/api/login', {
      username: form.username,
      password: form.password
    }, { withCredentials: true })

    scheduleRefresh(res.data.expires_in)

    userStore.login(res.data.user.id)
    router.push("/")
  } catch (err) {
    console.error(err)
    alert(err.response.data.error)
  }
}

async function scheduleRefresh(expires_in) {
  const delay = expires_in > 60 ? expires_in - 60 : expires_in;
  setTimeout(async () => {
    try {
      const res = await axios.post("/api/auth/refresh", {}, { withCredentials: true })
      const new_expires_in = res.data.expires_in || expires_in
      scheduleRefresh(new_expires_in)
    } catch (err) {
      console.error('刷新 token 失败', err)
    }
  }, delay * 1000)
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

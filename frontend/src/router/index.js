// Date:   Tue Jun 10 19:44:28 2025
// Mail:   lunar_ubuntu@qq.com
// Author: https://github.com/xiaoqixian

import { createRouter, createWebHistory } from 'vue-router'
import Register from '../views/Register.vue'

const routes = [
  { path: '/register', component: Register },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router

// Date:   Tue Jun 10 19:44:28 2025
// Mail:   lunar_ubuntu@qq.com
// Author: https://github.com/xiaoqixian

import { createRouter, createWebHistory } from 'vue-router'
import { useUserStore } from '@/stores/user'

import Register from '@/views/Register.vue'
import Login from '@/views/Login.vue'
import Home from '@/views/Home.vue'
import Write from '@/views/Write.vue'

const routes = [
  { path: '/', component: Home },
  { path: '/register', component: Register },
  { path: '/login', component: Login },
  { 
    path: '/write', 
    component: Write,
    meta: {
      requiresAuth: true
    }
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

router.beforeEach((to, _from, next) => {
  const userStore = useUserStore()

  if (to.meta.requiresAuth && !userStore.isLoggedIn) {
    next('/login')
  } else {
    next()
  }
})

export default router

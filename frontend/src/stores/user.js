// Date:   Thu Jun 12 19:55:03 2025
// Mail:   lunar_ubuntu@qq.com
// Author: https://github.com/xiaoqixian

import axios from "axios";
import { defineStore } from "pinia";

export const useUserStore = defineStore("user", {
  state: () => ({
    isLoggedIn: false,
    userInfo: null,
    checked: false
  }),
  actions: {
    async checkLogin() {
      try {
        const res = await axios.get("/api/auth/me", { withCredentials: true })
        this.isLoggedIn = true
        this.userInfo = res.data
      } catch (err) {
        this.isLoggedIn = false
        this.userInfo = null
      } finally {
        this.checked = true
      }
    },
    login() {
      this.isLoggedIn = true
    },
    logout() {
      this.isLoggedIn = false
      this.userInfo = null
    }
  }
})

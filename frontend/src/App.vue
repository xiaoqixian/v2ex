<template>
  <div id="app">
    <router-view />
  </div>
</template>

<script setup>
import { useUserStore } from "@/stores/user"
const userStore = useUserStore()
userStore.checkLogin()
</script>

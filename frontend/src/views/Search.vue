<template>
  <div class="app-container" :class="currentTheme">
    <Header />
    <div class="container">
      <div class="main-content">
        <div class="left-content">
          <PostList :posts="posts" />
        </div>
        <div class="right-content">
          <Sidebar />
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, provide, watch, onMounted, computed } from 'vue';
import { useRoute } from 'vue-router';
import PostList from '@/components/PostList.vue';
import Sidebar from '@/components/Sidebar.vue';
import Header from '@/components/Header.vue';
import axios from "axios";

import { themes, defaultTheme } from '@/themes.js';

const route = useRoute()
const keyword = computed(() => route.query.keyword || '')

const getPosts = async (keyword) => {
  try {
    const res = await axios.get(`/api/search?keyword=${encodeURIComponent(keyword)}`, { 
      withCredential: true
    })
    console.log("res.data = ")
    console.log(res.data)
    return res.data.posts
  } catch(err) {
    console.error(err)
  }
}

const posts = ref([])

onMounted(async () => {
  posts.value = await getPosts(keyword.value)
})

watch(() => route.query.keyword, async (newKeyword) => {
  if (!newKeyword) return
  posts.value = await getPosts(newKeyword)
}, { immediate: true })

</script>

<style>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

body {
  font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Oxygen, Ubuntu, Cantarell, "Fira Sans", "Droid Sans", "Helvetica Neue", sans-serif;
  background-color: var(--secondary);
  color: var(--text);
  font-size: 14px;
}

.app-container {
  min-height: 100vh;
  height: 100%;
  background-color: var(--secondary);
  color: var(--text);
  transition: background-color 0.3s ease, color 0.3s ease;
  display: flex;
  flex-direction: column;
}

.container {
  max-width: 1100px;
  width: 100%;
  margin: 0 auto;
  padding: 0 15px;
  margin-top: 44px; /* 只为固定头部留出空间 */
  flex: 1;
  display: flex;
  flex-direction: column;
}

.main-content {
  display: flex;
  margin-top: 20px;
  flex: 1;
}

.left-content {
  flex: 1;
  margin-right: 20px;

  display: flex;
  flex-direction: column;

  /* 允许子元素在垂直方向上收缩（避免撑开） */
  min-height: 0;
  max-height: 93vh;
}

.right-content {
  width: 270px;
}
</style>

<style scoped>
.logo {
  height: 6em;
  padding: 1.5em;
  will-change: filter;
  transition: filter 300ms;
}
.logo:hover {
  filter: drop-shadow(0 0 2em #646cffaa);
}
.logo.vue:hover {
  filter: drop-shadow(0 0 2em #42b883aa);
}
</style>

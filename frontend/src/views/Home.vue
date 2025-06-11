<template>
  <div class="app-container" :class="currentTheme">
    <Header />
    <div class="container">
      <div class="main-content">
        <div class="left-content">
          <TabNav v-model:selectedCategory="selectedCategory" />
          <TopicList :category="selectedCategory" />
        </div>
        <div class="right-content">
          <Sidebar />
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, provide, watch, onMounted } from 'vue';
import Header from '@/components/Header.vue';
import TopicList from '@/components/TopicList.vue';
import Sidebar from '@/components/Sidebar.vue';
import TabNav from '@/components/TabNav.vue';
import { themes, defaultTheme } from '@/themes.js';

// 主题状态管理
const currentTheme = ref(localStorage.getItem('theme') || defaultTheme);

// 当前选中的分类
const selectedCategory = ref('技术');

// 提供给所有组件使用的主题相关函数和状态
provide('theme', {
  current: currentTheme,
  themes,
  setTheme: (theme) => {
    currentTheme.value = theme;
    localStorage.setItem('theme', theme);
    applyTheme(theme);
  }
});

// 应用主题到CSS变量
const applyTheme = (themeName) => {
  const theme = themes[themeName];
  if (!theme) return;
  
  const root = document.documentElement;
  Object.entries(theme.colors).forEach(([key, value]) => {
    root.style.setProperty(`--${key}`, value);
  });
};

// 监听主题变化
watch(currentTheme, (newTheme) => {
  applyTheme(newTheme);
});

// 初始化主题
onMounted(() => {
  applyTheme(currentTheme.value);
});
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

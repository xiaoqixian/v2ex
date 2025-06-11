<template>
  <header class="header">
    <div class="header-inner">
      <div class="logo">
        <a href="#">V2EX</a>
      </div>
      <div class="search">
        <input type="text" class="search-input" placeholder="搜索" />
      </div>
      <nav class="main-nav">
        <div class="theme-switcher">
          <button 
            v-for="(theme, name) in themes" 
            :key="name"
            @click="setTheme(name)"
            :class="{ active: current === name }"
            class="theme-btn"
          >
            {{ theme.name }}
          </button>
        </div>
        <a href="#" class="nav-item">首页</a>
        <router-link to="/register" class="nav-item">注册</router-link>
        <a href="#" class="nav-item">登录</a>
      </nav>
    </div>
  </header>
</template>

<script setup>
import { inject } from 'vue';

// 注入主题相关状态和方法
const { current, themes, setTheme } = inject('theme');
</script>

<style scoped>
.header {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  background-color: var(--headerBg);
  color: var(--headerText);
  height: 44px;
  z-index: 100;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
  transition: background-color 0.3s ease, color 0.3s ease;
}

.header-inner {
  max-width: 1100px;
  margin: 0 auto;
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 100%;
  padding: 0 15px;
}

.logo a {
  font-size: 20px;
  font-weight: bold;
  color: var(--headerText);
  text-decoration: none;
  transition: color 0.2s ease;
}

.logo a:hover {
  opacity: 0.8;
}

.search {
  flex: 1;
  max-width: 300px;
  margin: 0 20px;
}

.search-input {
  width: 100%;
  padding: 5px 10px;
  border: none;
  border-radius: 3px;
  font-size: 14px;
  background-color: var(--searchBg);
  color: var(--searchText);
  transition: background-color 0.2s ease;
}

.search-input::placeholder {
  color: var(--searchPlaceholder);
}

.search-input:focus {
  outline: none;
  background-color: var(--searchBg);
  opacity: 0.9;
}

.main-nav {
  display: flex;
  align-items: center;
}

.nav-item {
  margin-left: 15px;
  color: var(--headerText);
  opacity: 0.8;
  text-decoration: none;
  font-size: 14px;
  transition: color 0.2s ease, opacity 0.2s ease;
}

.nav-item:hover {
  opacity: 1;
  color: var(--navHover);
}

.theme-switcher {
  display: flex;
  margin-right: 10px;
}

.theme-btn {
  background: none;
  border: none;
  color: var(--headerText);
  opacity: 0.6;
  padding: 2px 6px;
  margin: 0 2px;
  cursor: pointer;
  font-size: 12px;
  border-radius: 3px;
  transition: all 0.2s ease;
}

.theme-btn:hover {
  opacity: 0.8;
}

.theme-btn.active {
  opacity: 1;
  background-color: var(--accent);
  color: #fff;
}
</style>

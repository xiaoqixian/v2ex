<template>
  <div class="sidebar">
    <div class="sidebar-box">
      <div class="sidebar-title">V2EX = way to explore</div>
      <div class="sidebar-content">
        <p>V2EX 是一个关于分享和探索的地方</p>
      </div>
      <div class="sidebar-footer">
        <div class="current-online">当前在线 <span class="count">3580</span></div>
        <div class="member-count">已注册用户 <span class="count">582,262</span></div>
      </div>
    </div>
    
    <div class="sidebar-box">
      <div v-if="!userStore.isLoggedIn" class="sidebar-content">
        <router-link to="/login" class="sign-button">现在登录</router-link>
        <p>未注册用户请 <router-link to="/register">注册</router-link></p>
      </div>

      <div v-if="userStore.isLoggedIn" class="sidebar-content" id="sidebar-profile-box">
        <div class="avatar-box">
          <img src="@/assets/default_avatar.png" alt="Avatar" />
        </div>
        <div class="username">Kaleidopink</div>
      </div>
    </div>

    <div v-if="userStore.isLoggedIn" class="sidebar-box" id="write">
      <router-link to="/write">写点什么</router-link>
    </div>
    
    <div class="sidebar-box">
      <div class="sidebar-title">更为详尽的信息</div>
      <div class="sidebar-content">
        <div class="hot-topics">
          <p>创造新的内容会在这里显示给大家，让更多人看到你的观点。</p>
          <div class="topic-list">
            <div class="hot-topic">
              <a href="#">有什么值得买的 2.1 音箱推荐吗？</a>
            </div>
            <div class="hot-topic">
              <a href="#">为什么现在一个 DOS 界面的加油站客户端，都要装 Linux 虚拟机啊</a>
            </div>
            <div class="hot-topic">
              <a href="#">iOS 16 怎么了？大家更新了吗？</a>
            </div>
            <div class="hot-topic">
              <a href="#">有没有，有没有做好了？</a>
            </div>
            <div class="hot-topic">
              <a href="#">阿里云 AliDNS 下载链接不能用了？</a>
            </div>
            <div class="hot-topic">
              <a href="#">有没有可以在国内下不了？</a>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
// 注入主题
import { inject, ref } from 'vue';
import { useUserStore } from '@/stores/user'
const { current } = inject('theme', { current: ref('v2ex') });

const userStore = useUserStore()

</script>

<style scoped>
.sidebar-box {
  background-color: var(--primary);
  border-radius: 3px;
  box-shadow: 0 2px 3px rgba(0, 0, 0, 0.1);
  margin-bottom: 20px;
  overflow: hidden;
  transition: background-color 0.3s ease, box-shadow 0.3s ease;
}

.sidebar-title {
  padding: 10px;
  font-size: 14px;
  font-weight: bold;
  border-bottom: 1px solid var(--border);
  color: var(--text);
  transition: border-color 0.3s ease, color 0.3s ease;
}

.sidebar-content {
  padding: 10px;
  color: var(--text);
  transition: color 0.3s ease;
}

.sidebar-footer {
  background-color: var(--secondary);
  padding: 10px;
  font-size: 12px;
  color: var(--textSecondary);
  border-top: 1px solid var(--border);
  transition: background-color 0.3s ease, color 0.3s ease, border-color 0.3s ease;
}

.count {
  color: var(--accent);
  transition: color 0.3s ease;
}

.sidebar-content a.sign-button {
  display: block;
  width: 100%;
  padding: 8px 0;
  background-color: var(--accent);
  color: #fff;
  border: none;
  border-radius: 3px;
  font-size: 14px;
  cursor: pointer;
  margin-bottom: 10px;
  transition: background-color 0.3s ease;
}

.sidebar-content a.sign-button:hover {
  opacity: 0.9;
}

.sidebar-content p {
  font-size: 12px;
  color: var(--textSecondary);
  text-align: center;
  transition: color 0.3s ease;
}

.sidebar-content a {
  color: var(--textSecondary);
  text-decoration: none;
  transition: color 0.3s ease;
}

#write {
  padding: 10px;
}
#write a {
  color: var(--textSecondary);
  font-size: 16px;
}

.hot-topics {
  font-size: 13px;
  color: var(--text);
  transition: color 0.3s ease;
}

.topic-list {
  margin-top: 10px;
}

.hot-topic {
  margin-bottom: 8px;
}

.hot-topic a {
  color: var(--textSecondary);
  text-decoration: none;
  font-size: 13px;
  transition: color 0.3s ease;
}

.hot-topic a:hover {
  text-decoration: underline;
  color: var(--accent);
}

.avatar-box {
  flex: 8;
  display: flex;
  align-items: center;
  justify-content: center;
}

.avatar-box img {
  width: 96px;
  height: 96px;
  object-fit: cover;
  padding: 12px;
}

.username {
  flex: 2;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 14px;
  color: #333;
}
</style>

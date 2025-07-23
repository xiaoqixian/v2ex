<template>
  <div class="v2ex-container">
    <header class="v2ex-header">
      <nav class="v2ex-nav">
        <a href="/" class="v2ex-logo">V2EX</a>
      </nav>
    </header>

    <main class="main">

      <div class="post">
        <article>
          <div class="post-header">
            <h1>{{ title }}</h1>
            <div class="post-meta">
              <div class="post-info">
                <span class="author">{{ author }}</span>
                <span class="time">{{ createTime }}</span>
                <span class="views">{{ views }}次点击</span>
              </div>
              <div class="author-avatar">
                <img :src="authorAvatar" @error="onAvatarError">
              </div>
            </div>
          </div>

          <div class="post-content">
            <p>{{ content }}</p>
          </div>

          <div class="post-actions">
            <button class="v2ex-btn">加入收藏</button>
            <button class="v2ex-btn">Tweet</button>
            <button class="v2ex-btn">忽略主题</button>
            <button class="v2ex-btn">感谢</button>
          </div>

          <div class="post-stats">
            <span>{{ views }} 次点击</span>
          </div>

        </article>
      </div>

      <div v-if="userStore.isLoggedIn" class="comment-box">
        <div class="comment-editor">
          <div class="comment-title">回复</div>
          <textarea v-model="myComment" class="comment-textarea" placeholder="写下你的回复..."></textarea>
          <div class="comment-submit">
            <button class="submit-btn" @click="submitComment">提交回复</button>
          </div>
        </div>
      </div>

      <div v-if="!userStore.isLoggedIn" class="login-hint">
        请<router-link to="/login">登录</router-link>以后再提交回复
      </div>

      <div class="comments">
        <div class="comments-header">
          <h2>{{ commentsSize }}条回复</h2>
          <span class="comments-time">2025-06-20 13:21:34 +08:00</span>
        </div>
        
        <div class="comment" v-for="(comment, index) in comments" :key="index">
          <div class="comment-left">
            <img class="comment-avatar" src="@/assets/default_avatar.png" :alt="comment.author">
          </div>
          <div class="comment-right">
            <div class="comment-header">
              <span class="comment-author">{{ comment.user_name }}</span>
              <span class="comment-time">{{ timeEval(comment.created_at.seconds) }}</span>
              <span class="comment-floor">{{ index + 1 }}楼</span>
            </div>
            <div class="comment-content" v-html="comment.content"></div>
          </div>
        </div>
      </div>
    </main>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue';
import { useRoute } from 'vue-router';
import axios from 'axios';
import { useUserStore } from '@/stores/user';
import { timeEval } from '@/utils/time';
import { onAvatarError } from '@/utils/img-load-err';

const route = useRoute()
const userStore = useUserStore()

const title = ref('')
const author = ref('')
const authorAvatar = ref('')
const views = ref(1234)
const createdAt = ref(1750412014)
const createTime = computed(() => {
  return timeEval(createdAt.value)
})

const content = ref('')
const comments = ref([])
const myComment = ref('')

const commentsSize = computed(() => {
  return comments.value.length
})

async function fetchContent() {
  try {
    const res = await axios.get(`/api/posts/${route.params.id}`, {}, { withCredentials: true })
    console.log(res.data)
    title.value = res.data.title
    author.value = res.data.author
    authorAvatar.value = res.data.avatar
    createdAt.value = res.data.created_at.seconds
    content.value = res.data.content
  } catch (err) {
    console.error(err)
  }
}

async function fetchComments() {
  try {
    const res = await axios.get(`/api/comments/${route.params.id}`, {}, { withCredentials: true })
    console.log(res.data)
    comments.value = res.data
  } catch (err) {
    console.error(err)
  }
}

async function submitComment() {
  try {
    if (myComment.value == "") {
      alert("评论内容不可为空")
      return
    }

    const res = await axios.post(`/api/comments/${route.params.id}`, {
      user_id: userStore.userInfo.id,
      content: myComment.value,
    }, {
      withCredentials: true
    })

    myComment.value = ""
    await fetchComments()
  } catch (err) {
    console.error(err)
  }
}

onMounted(() => {
  fetchContent()
  fetchComments()
})

</script>

<style>
.v2ex-container {
  text-align: left;
  max-width: 1000px;
  margin: 0 auto;
  font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Helvetica, Arial, sans-serif;
  font-size: 14px;
  line-height: 1.6;
  color: #333;
  padding: 0 15px;
  background-color: #f5f5f5;
}

.v2ex-header {
  padding: 10px 0;
  background-color: #f9f9f9;
  border-bottom: 1px solid #e2e2e2;
}

.v2ex-nav {
  display: flex;
  align-items: center;
}

.v2ex-logo {
  font-weight: bold;
  color: #333;
  text-decoration: none;
  font-size: 18px;
}

.v2ex-nav-separator {
  margin: 0 5px;
  color: #ccc;
}

.v2ex-nav-current {
  color: #666;
}

.main {
  margin-top: 20px;
  background-color: white;
  border-radius: 3px;
  box-shadow: 0 1px 2px rgba(0,0,0,0.1);
  padding: 20px;
}

.post {
  margin-bottom: 10px;
}

.post-header h1 {
  font-size: 22px;
  margin-bottom: 5px;
  font-weight: 500;
  color: #333;
}

.post-meta {
  display: flex;
  justify-content: space-between;
  align-items: center;
  color: #999;
  font-size: 12px;
  margin-bottom: 20px;
  padding-bottom: 15px;
  border-bottom: 1px solid #f0f0f0;
}

.author {
  color: #666;
  font-weight: bold;
  margin-right: 10px;
}

.author-avatar img {
  width: 48px;
  height: 48px;
  border-radius: 3px;
}

.post-info .time {
  margin-right: 10px;
}

.post-content {
  font-size: 16px;
  margin-bottom: 20px;
  line-height: 1.8;
  color: #333;
}

.post-append {
  border-top: 1px dashed #e2e2e2;
  padding-top: 15px;
  margin-top: 15px;
  color: #666;
}

.v2ex-append-title {
  font-size: 12px;
  color: #999;
  margin-bottom: 10px;
}

.post-actions {
  margin: 20px 0;
}

.v2ex-btn {
  background: none;
  border: 1px solid #ddd;
  padding: 3px 10px;
  margin-right: 10px;
  border-radius: 3px;
  cursor: pointer;
  color: #666;
  font-size: 12px;
}

.v2ex-btn:hover {
  background-color: #f5f5f5;
}

.post-stats {
  color: #999;
  font-size: 12px;
  margin-bottom: 30px;
  padding-bottom: 15px;
  border-bottom: 1px solid #f0f0f0;
}

.comments {
  margin-top: 30px;
}

.comments-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding-bottom: 10px;
  border-bottom: 1px solid #f0f0f0;
  margin-bottom: 15px;
}

.comments-header h2 {
  font-size: 14px;
  color: #666;
  font-weight: normal;
}

.comments-time {
  font-size: 12px;
  color: #999;
}

.comment {
  display: flex;
  padding: 15px 0;
  border-bottom: 1px solid #f0f0f0;
}

.comment-left {
  margin-right: 15px;
  display: flex;
  justify-content: center;
  align-items: center;
}

.comment-avatar {
  width: 36px;
  height: 36px;
  border-radius: 3px;
}

.comment-right {
  flex: 1;
}

.comment-header {
  color: #999;
  font-size: 12px;
  margin-bottom: 8px;
}

.comment-author {
  color: #666;
  font-weight: bold;
  margin-right: 10px;
}

.comment-floor {
  float: right;
  color: #ccc;
}

.comment-content {
  line-height: 1.7;
  color: #333;
}

.comment-box {
  margin-top: 30px;
  background-color: white;
  border-radius: 3px;
  box-shadow: 0 1px 2px rgba(0,0,0,0.1);
}

.comment-actions {
  display: flex;
  padding: 10px 15px;
  border-bottom: 1px solid #f0f0f0;
  background-color: #f9f9f9;
  border-radius: 3px 3px 0 0;
}

.comment-btn {
  background: none;
  border: 1px solid #ddd;
  padding: 5px 12px;
  margin-right: 10px;
  border-radius: 3px;
  cursor: pointer;
  color: #666;
  font-size: 13px;
  transition: all 0.2s;
}

.comment-btn:hover {
  background-color: #f5f5f5;
  border-color: #ccc;
}

.comment-editor {
  padding: 15px;
}

.comment-title {
  font-size: 14px;
  color: #666;
  margin-bottom: 10px;
  font-weight: bold;
}

.comment-textarea {
  width: 100%;
  min-height: 120px;
  padding: 10px;
  border: 1px solid #ddd;
  border-radius: 3px;
  font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Helvetica, Arial, sans-serif;
  font-size: 14px;
  line-height: 1.6;
  resize: vertical;
  outline: none;
}

.comment-textarea:focus {
  border-color: #999;
}

.comment-submit {
  margin-top: 15px;
  text-align: right;
}

.submit-btn {
  background-color: #06c;
  color: white;
  border: none;
  padding: 6px 15px;
  border-radius: 3px;
  cursor: pointer;
  font-size: 13px;
  transition: background-color 0.2s;
}

.submit-btn:hover {
  background-color: #005bb7;
}

.login-hint {
  text-align: center;
}
</style>

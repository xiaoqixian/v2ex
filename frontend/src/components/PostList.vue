<template>
  <div class="post-list">
    <div v-for="(post, index) in posts" :key="index" class="post-item">
      <div class="avatar">
        <img :src="post.avatar" :alt="post.author" @error="onAvatarError"/>
      </div>
      <div class="post-content">
        <div class="post-title">
          <a :href="`/post/${post.post_id}`">{{ post.title }}</a>
        </div>
        <div class="postentry-meta">
          <a href="#" class="node">{{ post.node }}</a>
          <span class="separator">•</span>
          <a href="#" class="author">{{ post.author }}</a>
          <span class="separator">•</span>
          <span class="time">{{ timeEvalTimeStr(post.created_at) }}</span>
          <span class="separator">•</span>
          <span class="last-reply">最后回复来自 <a href="#">{{ post.lastReplyFrom }}</a></span>
        </div>
      </div>
      <div class="reply-count">
        <span>{{ post.replyCount }}</span>
      </div>
    </div>
    <div v-if="posts.length === 0" class="empty-state">
      <p>暂无内容</p>
    </div>
  </div>
</template>

<script setup>
import { timeEvalTimeStr } from '@/utils/time'
import { onAvatarError } from '@/utils/img-load-err'

// 父组件传进来的 posts
const props = defineProps({
  posts: {
    type: Array,
    default: () => []
  }
})
</script>

<style scoped>
.post-list {
  background-color: var(--primary);
  border-radius: 0 0 3px 3px;
  box-shadow: 0 2px 3px rgba(0, 0, 0, 0.1);
  transition: background-color 0.3s ease, box-shadow 0.3s ease;
  flex: 1 1 auto;
  overflow-y: auto;
  -webkit-overflow-scrolling: touch;
  min-height: 0;
}

.post-item {
  display: flex;
  padding: 15px;
  border-bottom: 1px solid var(--border);
  align-items: flex-start;
  text-align: left;
}

.avatar {
  width: 48px;
  height: 48px;
  margin-right: 15px;
}

.avatar img {
  width: 100%;
  height: 100%;
  border-radius: 4px;
}

.post-content {
  flex: 1;
}

.post-title {
  margin-bottom: 5px;
}

.post-title a {
  color: var(--text);
  text-decoration: none;
  font-size: 15px;
}

.post-title a:hover {
  color: var(--accent);
}

.postentry-meta {
  color: var(--textSecondary);
  font-size: 12px;
  text-align: left;
}

.node, .author, .last-reply a {
  color: var(--textSecondary);
  text-decoration: none;
}

.node:hover, .author:hover, .last-reply a:hover {
  color: var(--accent);
}

.separator {
  margin: 0 5px;
}

.reply-count {
  min-width: 30px;
  height: 30px;
  background-color: var(--secondary);
  border-radius: 15px;
  color: var(--textSecondary);
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: 700;
  font-size: 14px;
  padding: 0 10px;
  margin-top: 5px;
}

.empty-state {
  padding: 30px;
  text-align: center;
  color: var(--textSecondary);
}
</style>

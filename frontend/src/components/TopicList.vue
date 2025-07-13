<template>
  <div class="topic-list">
    <div v-for="(topic, index) in filteredTopics" :key="index" class="topic-item">
      <div class="avatar">
        <img :src="topic.avatar" :alt="topic.author" @error="onAvatarError"/>
      </div>
      <div class="topic-content">
        <div class="topic-title">
          <a :href="`/post/${topic.postid}`">{{ topic.title }}</a>
        </div>
        <div class="topic-meta">
          <a href="#" class="node">{{ topic.node }}</a>
          <span class="separator">•</span>
          <a href="#" class="author">{{ topic.author }}</a>
          <span class="separator">•</span>
          <span class="time">{{ timeEval(topic.time.seconds) }}</span>
          <span class="separator">•</span>
          <span class="last-reply">最后回复来自 <a href="#">{{ topic.lastReplyFrom }}</a></span>
        </div>
      </div>
      <div class="reply-count">
        <span>{{ topic.replyCount }}</span>
      </div>
    </div>
    <div v-if="filteredTopics.length === 0" class="empty-state">
      <p>该分类下暂无内容</p>
    </div>
  </div>
</template>

<script setup>
import { ref, inject, computed, watch, onMounted } from 'vue';
import { useUserStore } from '@/stores/user';
import { timeEval } from '@/utils/time';
import axios from 'axios';

const { current } = inject('theme', { current: ref('v2ex') });

const props = defineProps({
  category: {
    type: String,
    default: null
  }
});

const getTopics = async () => {
  try {
    const res = await axios.get('/api/home_posts', { withCredential: true })
    console.log("res.data = ")
    console.log(res.data)
    return res.data
  } catch(err) {
    console.error(err)
  }
}
const topics = ref([])

const filteredTopics = computed(() => {
  return topics.value;
});

watch(() => props.category, (newCategory) => {
  console.log(`加载 ${newCategory} 分类的数据`);
});

onMounted(async () => {
  topics.value = await getTopics()
})

const defaultAvatar = new URL("@/assets/default_avatar.png", import.meta.url).href

function onAvatarError(event) {
  event.target.src = defaultAvatar
}
</script>

<style scoped>
.topic-list {
  background-color: var(--primary);
  border-radius: 0 0 3px 3px;
  box-shadow: 0 2px 3px rgba(0, 0, 0, 0.1);
  transition: background-color 0.3s ease, box-shadow 0.3s ease;
}

.topic-item {
  display: flex;
  padding: 15px;
  border-bottom: 1px solid var(--border);
  transition: border-color 0.3s ease;
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

.topic-content {
  flex: 1;
}

.topic-title {
  margin-bottom: 5px;
}

.topic-title a {
  color: var(--text);
  text-decoration: none;
  font-weight: normal;
  font-size: 15px;
  transition: color 0.3s ease;
}

.topic-title a:hover {
  color: var(--accent);
}

.topic-meta {
  color: var(--textSecondary);
  font-size: 12px;
  transition: color 0.3s ease;
}

.node, .author, .last-reply a {
  color: var(--textSecondary);
  text-decoration: none;
  transition: color 0.3s ease;
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
  text-decoration: none;
  text-align: center;
  transition: background-color 0.3s ease, color 0.3s ease;
  align-self: flex-start;
  margin-top: 5px;
}

.empty-state {
  padding: 30px;
  text-align: center;
  color: var(--textSecondary);
}
</style>

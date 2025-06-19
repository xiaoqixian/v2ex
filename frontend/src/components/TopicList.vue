<template>
  <div class="topic-list">
    <div v-for="(topic, index) in filteredTopics" :key="index" class="topic-item">
      <div class="avatar">
        <img :src="topic.avatar" :alt="topic.author" />
      </div>
      <div class="topic-content">
        <div class="topic-title">
          <a :href="topic.link">{{ topic.title }}</a>
        </div>
        <div class="topic-meta">
          <a href="#" class="node">{{ topic.node }}</a>
          <span class="separator">•</span>
          <a href="#" class="author">{{ topic.author }}</a>
          <span class="separator">•</span>
          <span class="time">{{ topic.time }}</span>
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
import { ref, inject, computed, watch } from 'vue';
import { useUserStore } from '@/stores/user'

// 注入主题
const { current } = inject('theme', { current: ref('v2ex') });

// 接收分类参数
const props = defineProps({
  category: {
    type: String,
    default: null
  }
});

const userStore = useUserStore()

// 模拟数据
const allTopics = ref([
  {
    avatar: 'https://cdn.v2ex.com/avatar/c4ca/4238/1_normal.png?m=1630513007',
    title: 'OOMOL Studio 更新：更完善的工作流 IDE',
    node: '分享创造',
    author: 'BlackHole1',
    time: '1 小时 43 分钟前',
    lastReplyFrom: 'LinusWong',
    link: '#',
    replyCount: 23,
    category: '技术'
  },
  {
    avatar: 'https://cdn.v2ex.com/avatar/c4ca/4238/2_normal.png?m=1630513007',
    title: '最近在学日语，自己也顺手做了一个日语句子的解析器，分享给大家使用',
    node: '分享创造',
    author: 'howenhuang',
    time: '1 小时前',
    lastReplyFrom: 'howenhuang',
    link: '#',
    replyCount: 11,
    category: '技术'
  },
  {
    avatar: 'https://cdn.v2ex.com/avatar/c4ca/4238/3_normal.png?m=1630513007',
    title: 'AiEnglish：在阅读中渐进式学习英语',
    node: '分享创造',
    author: 'llcj',
    time: '1 小时前',
    lastReplyFrom: 'llcj',
    link: '#',
    replyCount: 16,
    category: '创意'
  },
  {
    avatar: 'https://cdn.v2ex.com/avatar/c4ca/4238/4_normal.png?m=1630513007',
    title: 'Casdoor: 具有 Web UI 界面的开源身份认证、单点登录平台，支持 GitHub、Gitee、QQ、微信、钉钉登录等',
    node: '分享创造',
    author: 'Casbin',
    time: '1 小时前',
    lastReplyFrom: 'supermama',
    link: '#',
    replyCount: 378,
    category: '技术'
  },
  {
    avatar: 'https://cdn.v2ex.com/avatar/c4ca/4238/5_normal.png?m=1630513007',
    title: '分享一个 V2EX 更好的讨论插件脚本，支持自定义表情和快速上传图片',
    node: '分享创造',
    author: 'Dogxi',
    time: '1 小时前',
    lastReplyFrom: 'Pipecraft',
    link: '#',
    replyCount: 2,
    category: '创意'
  },
  {
    avatar: 'https://cdn.v2ex.com/avatar/c4ca/4238/6_normal.png?m=1630513007',
    title: '15M CN2带宽限量300/月',
    node: '优惠信息',
    author: '3613N3544',
    time: '1 小时前',
    lastReplyFrom: '3613N3544',
    link: '#',
    replyCount: 35,
    category: '交易'
  },
  {
    avatar: 'https://cdn.v2ex.com/avatar/c4ca/4238/7_normal.png?m=1630513007',
    title: '兄弟们，救救手了，听说 jav-play 移动版改版本来了，不充钱下载不了广告，jav-play-go 让你安装不了"啊"浏览器',
    node: '分享发现',
    author: 'Alimov01',
    time: '1 小时前',
    lastReplyFrom: 'dodadada',
    link: '#',
    replyCount: 36,
    category: '好玩'
  },
  {
    avatar: 'https://cdn.v2ex.com/avatar/c4ca/4238/8_normal.png?m=1630513007',
    title: '新模型 Flux Kontext，API 价格比谷歌 8 倍',
    node: '分享发现',
    author: 'sickrimax',
    time: '1 小时前',
    lastReplyFrom: 'KIW1',
    link: '#',
    replyCount: 4,
    category: '好玩'
  },
  {
    avatar: 'https://cdn.v2ex.com/avatar/c4ca/4238/9_normal.png?m=1630513007',
    title: 'draw.io 的前前后',
    node: '设计',
    author: 'RushYoung',
    time: '1 小时前',
    lastReplyFrom: 'kimwang',
    link: '#',
    replyCount: 1,
    category: '创意'
  }
]);

const myTopics = ref([
  {
    avatar: 'https://cdn.v2ex.com/avatar/d41d/8cd9/10_normal.png?m=1640000000',
    title: '开源 ChatGPT 网页客户端 v3 发布，支持插件扩展',
    node: '开源项目',
    author: 'opengptdev',
    time: '2 小时前',
    lastReplyFrom: 'coderlong',
    link: '#',
    replyCount: 42,
    category: '技术'
  },
  {
    avatar: 'https://cdn.v2ex.com/avatar/d41d/8cd9/11_normal.png?m=1640000001',
    title: 'Figma 最新插件推荐：AI 布局优化器',
    node: '工具推荐',
    author: 'designhub',
    time: '3 小时前',
    lastReplyFrom: 'figmastar',
    link: '#',
    replyCount: 9,
    category: '创意'
  },
  {
    avatar: 'https://cdn.v2ex.com/avatar/d41d/8cd9/12_normal.png?m=1640000002',
    title: '自建 NAS 一周体验：软路由 + ZFS 方案分享',
    node: '硬件',
    author: 'techguy2024',
    time: '30 分钟前',
    lastReplyFrom: 'naslove',
    link: '#',
    replyCount: 27,
    category: '技术'
  },
  {
    avatar: 'https://cdn.v2ex.com/avatar/d41d/8cd9/13_normal.png?m=1640000003',
    title: '使用 Astro + Tailwind 快速构建博客站点',
    node: '前端开发',
    author: 'astrodev',
    time: '1 小时 15 分钟前',
    lastReplyFrom: 'mdxlover',
    link: '#',
    replyCount: 6,
    category: '技术'
  },
  {
    avatar: 'https://cdn.v2ex.com/avatar/d41d/8cd9/14_normal.png?m=1640000004',
    title: '低成本买到 iPhone 的另类方式：拼团盲盒？',
    node: '奇思妙想',
    author: 'thinker09',
    time: '4 小时前',
    lastReplyFrom: 'shoptips',
    link: '#',
    replyCount: 14,
    category: '好玩'
  },
  {
    avatar: 'https://cdn.v2ex.com/avatar/d41d/8cd9/15_normal.png?m=1640000005',
    title: '分享一个提升打字效率的 Vim 插件',
    node: 'Vim',
    author: 'vimkicker',
    time: '20 分钟前',
    lastReplyFrom: 'insertgod',
    link: '#',
    replyCount: 3,
    category: '创意'
  },
  {
    avatar: 'https://cdn.v2ex.com/avatar/d41d/8cd9/16_normal.png?m=1640000006',
    title: '618 快到了，大家都打算买啥数码产品？',
    node: '数码',
    author: 'salehunter',
    time: '5 小时前',
    lastReplyFrom: 'smartman',
    link: '#',
    replyCount: 66,
    category: '交易'
  },
  {
    avatar: 'https://cdn.v2ex.com/avatar/d41d/8cd9/17_normal.png?m=1640000007',
    title: '最近哪个 AI 聊天工具体验最好？分享一下',
    node: '人工智能',
    author: 'aiwatcher',
    time: '2 小时 30 分钟前',
    lastReplyFrom: 'huggingdev',
    link: '#',
    replyCount: 18,
    category: '技术'
  },
  {
    avatar: 'https://cdn.v2ex.com/avatar/d41d/8cd9/18_normal.png?m=1640000008',
    title: '我做了一个 RSS 自动翻译工具（支持 DeepL）',
    node: '自动化',
    author: 'rssflow',
    time: '1 小时前',
    lastReplyFrom: 'linguist',
    link: '#',
    replyCount: 5,
    category: '创意'
  }
]);

function formatTime(seconds) {
  const now = Date.now();
  const diffMillis = now - seconds * 1000;
  const diffSec = Math.floor(diffMillis / 1000);

  if (diffSec < 60) return diffSec + ' 秒前';
  const diffMin = Math.floor(diffSec / 60);
  if (diffMin < 60) return diffMin + ' 分钟前';
  const diffHour = Math.floor(diffMin / 60);
  if (diffHour < 24) return diffHour + ' 小时前';
  const diffDay = Math.floor(diffHour / 24);
  return diffDay + ' 天前';
}

const getTopics = async () => {
  let params = {}

  if (userStore.isLoggedIn) {
    params.userid = userStore.user.id
  }

  const res = await axios.get('/posts', { params })
  return res.data.map(post => {
    return {
      avatar: 'https://cdn.v2ex.com/avatar/c4ca/4238/1_normal.png?m=1630513007',
      title: post.title || '',
      node: post.node || '',
      author: `user_${post.author_id}` || '',
      time: formatTime(post.created_at?.seconds || 0),
      lastReplyFrom: 'Linus',
      replyCount: 23,
      category: '技术'
    }
  })
}
const topics = await getTopics()

const filteredTopics = computed(() => {
  if (!props.category) return topics;
  
  return topics.filter(topic => {
    if (topic.category === props.category) return true;
    if (topic.node === props.category) return true;
    return false;
  });
});

watch(() => props.category, (newCategory) => {
  console.log(`加载 ${newCategory} 分类的数据`);
});
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

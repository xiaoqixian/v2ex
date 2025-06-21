<template>
  <div class="v2ex-container">
    <header class="v2ex-header">
      <nav class="v2ex-nav">
        <a href="#" class="v2ex-logo">V2EX</a>
        <span class="v2ex-nav-separator">›</span>
        <span class="v2ex-nav-current">Java</span>
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
                <img src="../assets/default_avatar.png" alt="用户头像">
              </div>
            </div>
          </div>

          <div class="post-content">
            <p>{{ content }}</p>
          </div>

          <div class="post-append">
            <div class="v2ex-append-title">第 1 条附言 · 1 天前</div>
            <p>我没加方法上，之前是挺快的，拉取代码后就慢了，所以拉取代码后有个能断点位置变到方法定义上了？</p>
          </div>

          <div class="post-actions">
            <button class="v2ex-btn">加入收藏</button>
            <button class="v2ex-btn">Tweet</button>
            <button class="v2ex-btn">忽略主题</button>
            <button class="v2ex-btn">感谢</button>
          </div>

          <div class="post-stats">
            <span>4571 次点击 · 4 人收藏</span>
          </div>
        </article>
      </div>

      <div class="reply-box">
        <div class="reply-editor">
          <div class="reply-title">回复</div>
          <textarea class="reply-textarea" placeholder="写下你的回复..."></textarea>
          <div class="reply-submit">
            <button class="submit-btn">提交回复</button>
          </div>
        </div>
      </div>

      <div class="replies">
        <div class="replies-header">
          <h2>25 条回复</h2>
          <span class="replies-time">2025-06-20 13:21:34 +08:00</span>
        </div>
        
        <div class="reply" v-for="(reply, index) in replies" :key="index">
          <div class="reply-left">
            <img class="reply-avatar" src="@/assets/avatar1.png" :alt="reply.author">
          </div>
          <div class="reply-right">
            <div class="reply-header">
              <span class="reply-author">{{ reply.author }}</span>
              <span class="reply-time">{{ reply.time }}</span>
              <span class="reply-floor">{{ index + 1 }}楼</span>
            </div>
            <div class="reply-content" v-html="reply.content"></div>
          </div>
        </div>
      </div>
    </main>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue';

const title = ref('测试标题')
const author = ref('测试作者')
const views = ref(1234)
const createAt = 1750412014
const createTime = computed(() => {
  const now = Date.now()
  const ts = String(createAt).length === 10 ? createAt * 1000 : createAt // 兼容秒/毫秒
  const diff = Math.floor((now - ts) / 1000) // 差值（秒）

  if (diff < 60) return '刚刚'
  if (diff < 3600) return `${Math.floor(diff / 60)} 分钟前`
  if (diff < 86400) return `${Math.floor(diff / 3600)} 小时前`
  if (diff < 2592000) return `${Math.floor(diff / 86400)} 天前`
  if (diff < 31536000) return `${Math.floor(diff / 2592000)} 个月前`
  return `${Math.floor(diff / 31536000)} 年前`
})

const content = ref('同样的项目，idea 启动贼卡（一个小时） eclipse 启动两分钟。捣鼓了半天，加内存，换 idea ，换 jdk 都不行。最后把调试模式断点关闭，两分钟就起来了...醉了，删了所有断点就好使了。\n大家有类似的经历吗？说出来避避坑')

const replies = ref([
  {
    author: 'layxy',
    time: '1 天前',
    href: "@/assets/avatar1.png",
    content: '断点打到方法上或者某些特殊的地方,会影响启动,之前遇到过,而且现在还容易卡断点,尤其是查数据库的地方'
  },
  {
    author: 'opengps',
    time: '1 天前',
    href: "@/assets/avatar6.png",
    content: '所以说正式发布不应该带着调试文件 (@db 之类)'
  },
  {
    author: 'kk2syc',
    time: '1 天前',
    href: "@/assets/avatar5.png",
    content: '你就没想过为什么吗？因为有断点打在方法上面了。<br><br>IntelliJ IDEA Help: Note that using method breakpoints can slow down the application you are debugging@'
  },
  {
    author: 'miael@K',
    time: '1 天前 via Android',
    href: "@/assets/avatar4.png",
    content: '@layxy 对，我好像打在 aop 上了，但之前没事，拉了一段时间的代码就不行了，当时查代码也没看出啥来，现在也无从查证了'
  },
  {
    author: 'miael@K',
    time: '1 天前 via Android',
    href: "@/assets/avatar3.png",
    content: '@kk2syc 还真没注意过，下次留意下，感谢'
  },
  {
    author: 'Kiriri',
    time: '1 天前',
    href: "@/assets/avatar2.png",
    content: '断点上加 condition 也会影响'
  }
]);
</script>

<style>
.v2ex-container {
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

.post-content {
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

.replies {
  margin-top: 30px;
}

.replies-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding-bottom: 10px;
  border-bottom: 1px solid #f0f0f0;
  margin-bottom: 15px;
}

.replies-header h2 {
  font-size: 14px;
  color: #666;
  font-weight: normal;
}

.replies-time {
  font-size: 12px;
  color: #999;
}

.reply {
  display: flex;
  padding: 15px 0;
  border-bottom: 1px solid #f0f0f0;
}

.reply-left {
  margin-right: 15px;
}

.reply-avatar {
  width: 36px;
  height: 36px;
  border-radius: 3px;
}

.reply-right {
  flex: 1;
}

.reply-header {
  color: #999;
  font-size: 12px;
  margin-bottom: 8px;
}

.reply-author {
  color: #666;
  font-weight: bold;
  margin-right: 10px;
}

.reply-floor {
  float: right;
  color: #ccc;
}

.reply-content {
  line-height: 1.7;
  color: #333;
}

.reply-box {
  margin-top: 30px;
  background-color: white;
  border-radius: 3px;
  box-shadow: 0 1px 2px rgba(0,0,0,0.1);
}

.reply-actions {
  display: flex;
  padding: 10px 15px;
  border-bottom: 1px solid #f0f0f0;
  background-color: #f9f9f9;
  border-radius: 3px 3px 0 0;
}

.reply-btn {
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

.reply-btn:hover {
  background-color: #f5f5f5;
  border-color: #ccc;
}

.reply-editor {
  padding: 15px;
}

.reply-title {
  font-size: 14px;
  color: #666;
  margin-bottom: 10px;
  font-weight: bold;
}

.reply-textarea {
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

.reply-textarea:focus {
  border-color: #999;
}

.reply-submit {
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
</style>

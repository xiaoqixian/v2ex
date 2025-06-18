<template>
  <div class="editor-container">
    <div class="header">
      <input
        v-model="title"
        type="text"
        class="title-input"
        placeholder="请输入主题标题，如果标题能够表达完整内容，则正文可以为空"
      />
    </div>

    <div class="editor-body">
      <div class="tabs">
        <button :class="{ active: tab === 'text' }" @click="tab = 'text'">正文</button>
        <button :class="{ active: tab === 'preview' }" @click="tab = 'preview'">预览</button>
      </div>

      <div class="syntax-selector">
        <span>Syntax</span>
        <button :class="{ active: syntax === 'v2ex' }" @click="syntax = 'v2ex'">V2EX 原生格式</button>
        <button :class="{ active: syntax === 'markdown' }" @click="syntax = 'markdown'">Markdown</button>
      </div>

      <textarea v-model="content" class="editor-textarea" rows="15"></textarea>
    </div>

    <div class="footer">
      <select v-model="node">
        <option disabled value="">请选择一个节点</option>
        <option value="tech">技术</option>
        <option value="life">生活</option>
      </select>

      <button class="submit-btn" @click="submitPost">
        🚀 发布主题
      </button>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useUserStore } from '@/stores/user'

const title = ref('')
const content = ref('')
const node = ref('')
const tab = ref('text')
const syntax = ref('v2ex')

const userStore = useUserStore()

async function submitPost() {
  try {
    const res = await axios.post("/api/submit/post", {
      userid: userStore.userInfo.id,
      title: title.value,
      node: node.value,
      content: content.value
    }, { withCredentials: true })
  } catch (err) {
    console.error("发布失败: ", err)
    alert("发布失败: ", err)
  }
}
</script>

<style>
.editor-container {
  width: 700px;
  margin: 30px auto;
  background: #fff;
  border: 1px solid #ddd;
  padding: 16px;
  font-family: system-ui, sans-serif;
}

.header {
  margin-bottom: 12px;
}

.breadcrumb {
  display: block;
  margin-bottom: 8px;
  color: #888;
}

.title-input {
  width: 100%;
  padding: 10px;
  font-size: 16px;
  box-sizing: border-box;
  border: none;
  border-bottom: 1px solid #ccc;
}

.editor-body {
  margin-top: 12px;
}

.tabs {
  display: flex;
  border-bottom: 1px solid #ccc;
  margin-bottom: 8px;
}

.tabs button {
  padding: 8px 12px;
  border: none;
  background: none;
  cursor: pointer;
  color: #333;
}

.tabs .active {
  border-bottom: 2px solid #333;
  font-weight: bold;
}

.syntax-selector {
  display: flex;
  align-items: center;
  gap: 8px;
  margin-bottom: 8px;
}

.syntax-selector span {
  color: #666;
}

.syntax-selector button {
  border: 1px solid #ccc;
  padding: 4px 8px;
  background: #f8f8f8;
  cursor: pointer;
}

.syntax-selector .active {
  background: #ddd;
}

.editor-textarea {
  width: 100%;
  padding: 10px;
  box-sizing: border-box;
  resize: vertical;
  font-family: monospace;
  font-size: 14px;
  border: 1px solid #ccc;
}

.footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-top: 16px;
}

.footer select {
  padding: 6px;
  font-size: 14px;
}

.submit-btn {
  background: #eee;
  border: 1px solid #ccc;
  padding: 8px 16px;
  cursor: pointer;
}

.submit-btn:hover {
  background: #ddd;
}
</style>

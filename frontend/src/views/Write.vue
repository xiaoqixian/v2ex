<template>
  <div class="editor-container">
    <div v-if="submitted" class="success-container">
      <h2 class="success-message">🎉 发布成功！</h2>
      <div class="success-actions">
        <button @click="goToPost" class="view-btn">立即查看</button>
        <button @click="resetForm" class="repost-btn">再次发布</button>
      </div>
    </div>
    <div v-else>
      <div class="header">
        <input
          v-model="title"
          type="text"
          class="title-input"
          placeholder="请输入主题标题，如果标题能够表达完整内容，则正文可以为空"
        />
      </div>

      <div class="editor-body">
        <textarea v-model="content" class="editor-textarea" rows="15"></textarea>
      </div>

      <div class="footer">
        <select v-model="node">
          <option disabled value="">请选择一个节点</option>
            <option 
              v-for="(label, key) in nodeTypes" 
              :key="key" 
              :value="key"
            >
              {{ label }}
            </option>
        </select>

        <button class="submit-btn" @click="submitPost">
          🚀 发布主题
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useUserStore } from '@/stores/user'
//import Alert from '@/components/Alert.vue'
import axios from 'axios'

const userStore = useUserStore()
const router = useRouter()

const title = ref('')
const content = ref('')
const node = ref('')
const tab = ref('text')
const syntax = ref('v2ex')
const submitted = ref(false)
const postid = ref(null)

const nodeTypes = {
  "apple": "Apple", 
  "frontend-dev": "前端开发", 
  "backend-dev": "后端技术", 
  "machine-learning": "机器学习", 
  "game": "游戏", 
  "life": "生活"
}

async function submitPost() {
  try {
    if (node.value == "") {
      throw new Error("请选择一个节点!")
    }

    const res = await axios.post("/api/posts", {
      title: title.value,
      node: node.value,
      content: content.value
    }, { withCredentials: true })

    submitted.value = true
    postid.value = res.data.postid
  } catch (err) {
    console.error(err)
    // 在这里抛出一个 alert，内容为 err
    alert(err)
  }
}

function goToPost() {
  router.push(`/post/${postid.value}`)
}

function resetForm() {
  title.value = ''
  content.value = ''
  node.value = ''
  tab.value = 'text'
  syntax.value = 'v2ex'
  submitted.value = false
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
  border: 1px solid #ccc;
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

.success-message {
  background-color: #f0fdf4;
  border: 2px solid #bbf7d0;
  padding: 24px;
  margin: 20px auto;
  border-radius: 12px;
  text-align: center;
  max-width: 600px;
  font-family: 'Segoe UI', sans-serif;
  color: #15803d;
  font-size: 20px;
  font-weight: 600;
}

.success-actions {
  display: flex;
  justify-content: center;
  gap: 24px;
  margin-top: 30px;
}

.success-actions button {
  background-color: white;
  color: black;
  border: 2px solid #d4d4d4;
  padding: 10px 24px;
  font-size: 16px;
  border-radius: 8px;
  cursor: pointer;
  transition: all 0.2s ease;
}

.success-actions button:hover {
  background-color: #3b82f6;
  border-color: #2563eb;
  color: white;
}
</style>

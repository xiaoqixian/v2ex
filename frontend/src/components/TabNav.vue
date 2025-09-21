<template>
  <div class="tab-nav">
    <div class="tab-nav-inner">
      <div class="primary-tabs">
        <a 
          v-for="(tab, index) in primaryTabs" 
          :key="index"
          href="#" 
          class="tab-item" 
          :class="{ active: activeTab === index }"
          @click.prevent="setActiveTab(index)"
        >
          {{ tab.name }}
        </a>
      </div>
      <div class="secondary-tabs" v-if="false && activeTab !== null && secondaryTabs.length > 0">
        <a 
          v-for="(tab, index) in secondaryTabs" 
          :key="index"
          href="#" 
          class="secondary-tab-item" 
          :class="{ active: activeSecondaryTab === index }"
          @click.prevent="setActiveSecondaryTab(index)"
        >
          {{ tab }}
        </a>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, watch, inject } from 'vue';

// 注入主题
const { current } = inject('theme');

// 定义props和emit
const props = defineProps({
  selectedTopic: {
    type: String,
    default: null
  }
});

const emit = defineEmits(['update:selectedCategory']);

// 主标签数据
const primaryTabs = [
  {
    name: "推荐",
    categories: []
  },
  {
    name: "热门",
    categories: []
  }
];

// 当前激活的主标签
const activeTab = ref(0); // 默认选中第一个标签

// 当前激活的次级标签
const activeSecondaryTab = ref(0);

// 计算当前显示的次级标签
const secondaryTabs = computed(() => {
  if (activeTab.value === null) return [];
  return primaryTabs[activeTab.value].categories;
});

// 计算当前选中的分类
const selectedCategory = computed(() => {
  if (activeTab.value === null || secondaryTabs.value.length === 0) {
    return primaryTabs[activeTab.value].name;
  }
  return secondaryTabs.value[activeSecondaryTab.value];
});

// 设置激活的主标签
const setActiveTab = (index) => {
  activeTab.value = index;
  activeSecondaryTab.value = 0; // 重置次级标签选择
  updateSelectedCategory();
};

// 设置激活的次级标签
const setActiveSecondaryTab = (index) => {
  activeSecondaryTab.value = index;
  updateSelectedCategory();
};

// 更新选中的分类并触发事件
const updateSelectedCategory = () => {
  emit('update:selectedTopic', selectedCategory.value);
};

// 初始化时触发一次更新
watch(selectedCategory, (newValue) => {
  if (newValue && !props.selectedCategory) {
    updateSelectedCategory();
  }
}, { immediate: true });
</script>

<style scoped>
.tab-nav {
  background-color: var(--primary);
  border-radius: 3px 3px 0 0;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
  margin-bottom: 0;
  transition: background-color 0.3s ease, box-shadow 0.3s ease;
  border-bottom: 1px solid var(--border);
  /* 确保不是固定定位 */
  position: relative;
}

.tab-nav-inner {
  padding: 0 15px;
}

.primary-tabs {
  display: flex;
  flex-wrap: wrap;
  padding: 10px 0;
}

.tab-item {
  color: var(--text);
  text-decoration: none;
  font-size: 14px;
  margin-right: 20px;
  padding: 5px 0;
  position: relative;
  transition: color 0.2s ease;
}

.tab-item:hover {
  color: var(--accent);
}

.tab-item.active {
  color: var(--accent);
  font-weight: 500;
}

.tab-item.active::after {
  content: '';
  position: absolute;
  bottom: -2px;
  left: 0;
  width: 100%;
  height: 2px;
  background-color: var(--accent);
  transition: background-color 0.3s ease;
}

.secondary-tabs {
  display: flex;
  flex-wrap: wrap;
  padding: 8px 0;
  border-top: 1px solid var(--border);
  transition: border-color 0.3s ease;
}

.secondary-tab-item {
  color: var(--textSecondary);
  text-decoration: none;
  font-size: 13px;
  margin-right: 15px;
  padding: 3px 0;
  transition: color 0.2s ease;
}

.secondary-tab-item:hover {
  color: var(--accent);
}

.secondary-tab-item.active {
  color: var(--accent);
}

/* 动画效果 */
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s ease, transform 0.3s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
  transform: translateY(-10px);
}

.fade-enter-to,
.fade-leave-from {
  opacity: 1;
  transform: translateY(0);
}
</style>

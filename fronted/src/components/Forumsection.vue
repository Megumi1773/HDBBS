<script setup>
import { ref, reactive, onMounted } from 'vue'
import axios from 'axios';
let Classify = reactive([])
let get = () => {
  axios.get('/api/cate/getAllCate')
    .then(response => {
      Classify.length = 0
      response.data.data.forEach(item => Classify.push(item))
    })
    .catch(error => {
      console.error('Error fetching posts:', error);
    });
}
onMounted(() => {
  get();
});

</script>
<template>
  <div class="sidebar-card">
    <!-- 导航栏标题 -->
    <div class="sidebar-header">
      <svg width="20" height="20" fill="currentColor" viewBox="0 0 20 20">
        <path
          d="M7 3a1 1 0 000 2h6a1 1 0 100-2H7zM4 7a1 1 0 011-1h10a1 1 0 110 2H5a1 1 0 01-1-1zM2 11a2 2 0 012-2h12a2 2 0 012 2v4a2 2 0 01-2 2H4a2 2 0 01-2-2v-4z" />
      </svg>
      <h2>论坛版块</h2>
    </div>

    <!-- 导航菜单列表 -->
    <nav class="nav-menu">
      <a href="#" class="nav-item" :class="{ active: activeNavItem === 'home' }" @click="setActiveNav('home')"
        v-for="item in Classify" :key="item.id">
        <svg width="20" height="20" fill="currentColor" viewBox="0 0 20 20">
          <path
            d="M10.707 2.293a1 1 0 00-1.414 0l7 7a1 1 0 001.414 1.414L4 10.414V17a1 1 0 001 1h2a1 1 0 001-1v-2a1 1 0 011-1h2a1 1 0 011 1v2a1 1 0 001 1h2a1 1 0 001-1v-6.586l.293.293a1 1 0 001.414-1.414l-7-7z" />
        </svg>
        <span>{{ item.CategoryName }}</span>
      </a>
    </nav>
  </div>

</template>
<style scoped>
/* ==================== 左侧导航栏样式 ==================== */
.sidebar-card {
  background: white;
  border-radius: 0.75rem;
  box-shadow: 0 1px 3px 0 rgba(0, 0, 0, 0.1);
  padding: 1.5rem;
  position: sticky;
  top: 6rem;
}

.sidebar-header {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  margin-bottom: 1.5rem;
  padding-bottom: 1rem;
  border-bottom: 1px solid #f3f4f6;
}

.sidebar-header svg {
  color: #8b5cf6;
}

.sidebar-header h2 {
  font-weight: 600;
  color: #1f2937;
  font-size: 1rem;
}

.nav-menu {
  display: flex;
  flex-direction: column;
  gap: 0.25rem;
}

.nav-item {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  padding: 0.75rem 1rem;
  border-radius: 0.5rem;
  text-decoration: none;
  color: #6b7280;
  transition: all 0.2s;
  font-size: 0.875rem;
}

.nav-item:hover {
  background: #f9fafb;
  color: #374151;
  transform: translateX(2px);
}

.nav-item.active {
  background: #8b5cf6;
  color: white;
  font-weight: 500;
  box-shadow: 0 2px 4px rgba(139, 92, 246, 0.2);
}

.nav-item svg {
  width: 1.25rem;
  height: 1.25rem;
  flex-shrink: 0;
}
</style>
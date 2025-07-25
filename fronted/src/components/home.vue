<template>
  <div class="app">
    <!-- ==================== 头部导航区域 ==================== -->
    <header class="header">
      <div class="header-container">
        <div class="header-content">
          <!-- 左侧：Logo和标题 -->
          <div class="logo-section">
            <div class="logo-icon">
              <svg width="20" height="20" fill="currentColor" viewBox="0 0 20 20">
                <path
                  d="M10.394 2.08a1 1 0 00-.788 0l-7 3a1 1 0 000 1.84L5.25 8.051a.999.999 0 01.356-.257l4-1.714a1 1 0 11.788 1.84L7.667 9.088l1.94.831a1 1 0 00.787 0l7-3a1 1 0 000-1.84l-7-3zM3.31 9.397L5 10.12v4.102a8.969 8.969 0 00-1.05-.174 1 1 0 01-.89-.89 11.115 11.115 0 01.25-3.762zM9.3 16.573A9.026 9.026 0 007 14.935v-3.957l1.818.78a3 3 0 002.364 0l5.508-2.361a11.026 11.026 0 01.25 3.762 1 1 0 01-.89.89 8.968 8.968 0 00-5.35 2.524 1 1 0 01-1.4 0zM6 18a1 1 0 001-1v-2.065a8.935 8.935 0 00-2-.712V17a1 1 0 001 1z" />
              </svg>
            </div>
            <h1 class="title">校园论坛</h1>
          </div>

          <!-- 中间：搜索框 -->
          <div class="search-container">
            <div class="search-wrapper">
              <input type="text" placeholder="搜索帖子、用户或话题..." class="search-input" v-model="searchQuery"
                @keyup.enter="handleSearch">
              <svg class="search-icon" width="20" height="20" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                  d="m21 21-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
              </svg>
            </div>
          </div>

          <!-- 右侧：用户操作区域 -->
          <div class="user-actions">
            <!-- 通知按钮 -->
            <div class="notification-wrapper">
              <button class="notification-btn" @click="handleNotificationClick">
                <svg width="20" height="20" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                    d="M15 17h5l-5 5v-5zM9 7H4l5-5v5zm6 10V7a1 1 0 00-1-1H5a1 1 0 00-1 1v10a1 1 0 001 1h9a1 1 0 001-1z" />
                </svg>
              </button>
              <span class="notification-badge">{{ notificationCount }}</span>
            </div>
            <!-- 用户头像 -->
            <div class="user-avatar" @click="handleUserClick">张</div>
          </div>
        </div>
      </div>
    </header>

    <!-- ==================== 主要内容区域 ==================== -->
    <div class="main-container">
      <div class="content-grid">

        <!-- ==================== 左侧导航栏 ==================== -->
        <div class="sidebar">
          <Forumsection></Forumsection>
        </div>

        <!-- 中间主内容 -->
        <div class="main-content">
          <Message></Message>
        </div>

        <!-- ==================== 右侧信息栏 ==================== -->
        <div class="right-sidebar">

          <!-- 热门话题模块 -->
          <RightSideTaskbar></RightSideTaskbar>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import RightSideTaskbar from './RightSideTaskbar.vue'
import Right from './RightSideTaskbar.vue'


// ==================== 响应式数据定义 ====================

// 搜索相关
const searchQuery = ref('')

// 通知相关
const notificationCount = ref(3)

// 响应式导航项
const activeNavItem = ref('home')

const setActiveNav = (navItem) => {
  activeNavItem.value = navItem
  // 可根据需要添加内容区联动逻辑
}

// ==================== 方法定义 ====================

// 搜索处理
const handleSearch = () => {
  console.log('搜索内容:', searchQuery.value)
  // 这里可以添加搜索逻辑
}

// 通知点击处理
const handleNotificationClick = () => {
  console.log('点击通知')
  // 这里可以显示通知列表
}

// 用户头像点击处理
const handleUserClick = () => {
  console.log('点击用户头像')
  // 这里可以显示用户菜单
}


// ==================== 生命周期钩子 ====================
onMounted(() => {
  console.log('校园论坛组件已挂载')
  // 这里可以添加初始化逻辑，如获取数据等
})
</script>

<style scoped>
/* ==================== 全局基础样式 ==================== */
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

.app {
  min-height: 100vh;
  background-color: #f8fafc;
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
  line-height: 1.6;
}

/* ==================== 头部导航样式 ==================== */
.header {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1);
  position: sticky;
  top: 0;
  z-index: 100;
}

.header-container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 1rem;
}

.header-content {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 1rem 0;
}

/* Logo区域样式 */
.logo-section {
  display: flex;
  align-items: center;
  gap: 0.75rem;
}

.logo-icon {
  width: 2rem;
  height: 2rem;
  background: rgba(255, 255, 255, 0.2);
  border-radius: 0.5rem;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: background-color 0.2s;
}

.logo-icon:hover {
  background: rgba(255, 255, 255, 0.3);
}

.title {
  font-size: 1.25rem;
  font-weight: bold;
  letter-spacing: 0.5px;
}

/* 搜索框样式 */
.search-container {
  flex: 1;
  max-width: 28rem;
  margin: 0 2rem;
}

.search-wrapper {
  position: relative;
}

.search-input {
  width: 100%;
  padding: 0.5rem 1rem 0.5rem 2.5rem;
  border-radius: 9999px;
  background: rgba(255, 255, 255, 0.1);
  backdrop-filter: blur(4px);
  border: 1px solid rgba(255, 255, 255, 0.2);
  color: white;
  outline: none;
  transition: all 0.2s;
  font-size: 0.875rem;
}

.search-input::placeholder {
  color: rgba(255, 255, 255, 0.7);
}

.search-input:focus {
  background: rgba(255, 255, 255, 0.15);
  border-color: rgba(255, 255, 255, 0.4);
  box-shadow: 0 0 0 2px rgba(255, 255, 255, 0.1);
}

.search-icon {
  position: absolute;
  left: 0.75rem;
  top: 50%;
  transform: translateY(-50%);
  color: rgba(255, 255, 255, 0.7);
  pointer-events: none;
}

/* 用户操作区域样式 */
.user-actions {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.notification-wrapper {
  position: relative;
}

.notification-btn {
  padding: 0.5rem;
  border-radius: 9999px;
  background: rgba(255, 255, 255, 0.1);
  border: none;
  color: white;
  cursor: pointer;
  transition: background-color 0.2s;
  display: flex;
  align-items: center;
  justify-content: center;
}

.notification-btn:hover {
  background: rgba(255, 255, 255, 0.2);
}

.notification-badge {
  position: absolute;
  top: -0.25rem;
  right: -0.25rem;
  background: #ef4444;
  color: white;
  font-size: 0.75rem;
  border-radius: 9999px;
  width: 1.25rem;
  height: 1.25rem;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: bold;
  animation: pulse 2s infinite;
}

@keyframes pulse {

  0%,
  100% {
    opacity: 1;
  }

  50% {
    opacity: 0.7;
  }
}

.user-avatar {
  width: 2rem;
  height: 2rem;
  background: #fb923c;
  border-radius: 9999px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: bold;
  font-size: 0.875rem;
  cursor: pointer;
  transition: transform 0.2s;
}

.user-avatar:hover {
  transform: scale(1.05);
}



/* ==================== 主要内容区域样式 ==================== */
.posts-container {
  background: white;
  border-radius: 0.75rem;
  box-shadow: 0 1px 3px 0 rgba(0, 0, 0, 0.1);
  overflow: hidden;
}

.posts-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 1.5rem;
  border-bottom: 1px solid #f3f4f6;
  background: #fafafa;
}

.posts-title {
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.posts-title svg {
  color: #3b82f6;
}

.posts-title h2 {
  font-size: 1.125rem;
  font-weight: 600;
  color: #1f2937;
}

.new-post-btn {
  padding: 0.5rem 1rem;
  background: #10b981;
  color: white;
  border: none;
  border-radius: 0.5rem;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s;
  font-size: 0.875rem;
}

.new-post-btn:hover {
  background: #059669;
  transform: translateY(-1px);
  box-shadow: 0 4px 8px rgba(16, 185, 129, 0.2);
}

/* 帖子列表样式 */
.posts-list {
  divide-y: 1px solid #f3f4f6;
}

.post-item {
  padding: 1.5rem;
  transition: background-color 0.2s;
}

.post-item:hover {
  background: #fafafa;
}

.post-content {
  display: flex;
  gap: 1rem;
}

.post-avatar {
  width: 2.5rem;
  height: 2.5rem;
  border-radius: 9999px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  font-weight: 500;
  flex-shrink: 0;
  font-size: 0.875rem;
}

.post-avatar.purple {
  background: linear-gradient(135deg, #8b5cf6, #7c3aed);
}

.post-avatar.blue {
  background: linear-gradient(135deg, #3b82f6, #2563eb);
}

.post-avatar.green {
  background: linear-gradient(135deg, #10b981, #059669);
}

.post-main {
  flex: 1;
  min-width: 0;
}

.post-meta {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 0.5rem;
}

.post-username {
  font-weight: 500;
  color: #1f2937;
  font-size: 0.875rem;
}

.post-time {
  font-size: 0.75rem;
  color: #6b7280;
  margin-top: 0.125rem;
}

.post-category {
  padding: 0.25rem 0.75rem;
  border-radius: 9999px;
  font-size: 0.75rem;
  font-weight: 500;
  flex-shrink: 0;
}

.post-category.purple {
  background: #ede9fe;
  color: #7c3aed;
}

.post-category.blue {
  background: #dbeafe;
  color: #2563eb;
}

.post-category.green {
  background: #d1fae5;
  color: #059669;
}

.post-title {
  font-size: 1.125rem;
  font-weight: 600;
  color: #1f2937;
  margin-bottom: 0.5rem;
  cursor: pointer;
  transition: color 0.2s;
  line-height: 1.4;
}

.post-title:hover {
  color: #3b82f6;
}

.post-excerpt {
  color: #6b7280;
  margin-bottom: 1rem;
  line-height: 1.5;
  font-size: 0.875rem;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.post-stats {
  display: flex;
  align-items: center;
  gap: 1.5rem;
  font-size: 0.875rem;
  color: #6b7280;
}

.stat-item {
  display: flex;
  align-items: center;
  gap: 0.25rem;
  cursor: pointer;
  transition: color 0.2s;
  padding: 0.25rem 0;
}

.stat-item:hover {
  color: #3b82f6;
}

.stat-item svg {
  width: 1rem;
  height: 1rem;
}


/* ==================== 滚动条样式 ==================== */
::-webkit-scrollbar {
  width: 6px;
}

::-webkit-scrollbar-track {
  background: #f1f1f1;
  border-radius: 3px;
}

::-webkit-scrollbar-thumb {
  background: #c1c1c1;
  border-radius: 3px;
}

::-webkit-scrollbar-thumb:hover {
  background: #a8a8a8;
}

.main-container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 1.5rem 1rem;
}

.content-grid {
  display: grid;
  grid-template-columns: 1fr 2fr 1fr;
  gap: 1.5rem;
}
</style>
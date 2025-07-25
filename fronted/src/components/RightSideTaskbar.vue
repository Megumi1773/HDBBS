<script setup>

// ==================== 响应式数据定义 ====================
// 响应式导航项
const activeNavItem = ref('home')
// 热门话题数据
const hotTopics = reactive([
  { id: 1, name: '#期末复习攻略#', discussions: '1.2万', views: '3.5万' },
  { id: 2, name: '#校园美食推荐#', discussions: '8.9k', views: '2.1万' },
  { id: 3, name: '#实习求职经验#', discussions: '6.7k', views: '1.8万' },
  { id: 4, name: '#宿舍生活日常#', discussions: '5.4k', views: '1.5万' },
  { id: 5, name: '#社团活动分享#', discussions: '4.2k', views: '1.2万' }
])

// 社区统计数据
const communityStats = reactive({
  onlineUsers: '1,247',
  todayPosts: '156',
  totalUsers: '25.6k'
})

// 活跃用户数据
const activeUsers = reactive([
  { id: 1, name: '李明同学', avatar: '李', avatarColor: 'purple', lastActive: '刚刚活跃', status: 'online' },
  { id: 2, name: '王小美', avatar: '王', avatarColor: 'blue', lastActive: '5分钟前', status: 'online' },
  { id: 3, name: '陈大厨', avatar: '陈', avatarColor: 'green', lastActive: '10分钟前', status: 'away' }
])


// 话题点击处理
const handleTopicClick = (topic) => {
  console.log('点击话题:', topic.name)
  // 这里可以跳转到话题页面
}

// 用户资料点击处理
const handleUserProfileClick = (user) => {
  console.log('点击用户资料:', user.name)
  // 这里可以跳转到用户资料页面
}

// ==================== 生命周期钩子 ====================
onMounted(() => {
  console.log('校园论坛组件已挂载')
  // 这里可以添加初始化逻辑，如获取数据等
})
</script>
<template>

  <!-- 热门话题模块 -->
  <div class="widget">
    <div class="widget-header">
      <svg width="20" height="20" fill="currentColor" viewBox="0 0 20 20">
        <path fill-rule="evenodd"
          d="M12.395 2.553a1 1 0 00-1.45-.385c-.345.23-.614.558-.822.88-.214.33-.403.713-.57 1.116-.334.804-.614 1.768-.84 2.734a31.365 31.365 0 00-.613 3.58 2.64 2.64 0 01-.945-1.067c-.328-.68-.398-1.534-.398-2.654A1 1 0 005.05 6.05 6.981 6.981 0 003 11a7 7 0 1011.95-4.95c-.592-.591-.98-.985-1.348-1.467-.363-.476-.724-1.063-1.207-2.03z"
          clip-rule="evenodd" />
      </svg>
      <h3>热门话题</h3>
    </div>

    <div class="topics-list">
      <div class="topic-item" v-for="topic in hotTopics" :key="topic.id" @click="handleTopicClick(topic)">
        <span class="topic-name">{{ topic.name }}</span>
        <div class="topic-stats">
          <div>{{ topic.discussions }}讨论</div>
          <div>{{ topic.views }}阅读</div>
        </div>
      </div>
    </div>
  </div>

  <!-- 社区统计模块 -->
  <div class="widget">
    <div class="widget-header">
      <svg width="20" height="20" fill="currentColor" viewBox="0 0 20 20">
        <path
          d="M2 11a1 1 0 011-1h2a1 1 0 011 1v5a1 1 0 01-1 1H3a1 1 0 01-1-1v-5zM8 7a1 1 0 011-1h2a1 1 0 011 1v9a1 1 0 01-1 1H9a1 1 0 01-1-1V7zM14 4a1 1 0 011-1h2a1 1 0 011 1v12a1 1 0 01-1 1h-2a1 1 0 01-1-1V4z" />
      </svg>
      <h3>社区统计</h3>
    </div>

    <!-- 主要统计数据 -->
    <div class="stats-main">
      <div class="main-stat">{{ communityStats.onlineUsers }}</div>
      <div class="main-stat-label">在线用户</div>
    </div>

    <!-- 次要统计数据网格 -->
    <div class="stats-grid">
      <div class="stat-box">
        <div class="stat-number blue">{{ communityStats.todayPosts }}</div>
        <div class="stat-label">今日新帖</div>
      </div>
      <div class="stat-box">
        <div class="stat-number purple">{{ communityStats.totalUsers }}</div>
        <div class="stat-label">总用户</div>
      </div>
    </div>
  </div>

  <!-- 活跃用户模块 -->
  <div class="widget">
    <div class="widget-header">
      <svg width="20" height="20" fill="currentColor" viewBox="0 0 20 20">
        <path d="M13 6a3 3 0 11-6 0 3 3 0 016 0zM18 8a2 2 0 11-4 0 2 2 0 014 0zM14 15a4 4 0 00-8 0v3h8v-3z" />
      </svg>
      <h3>活跃用户</h3>
    </div>

    <div class="users-list">
      <div class="user-item" v-for="user in activeUsers" :key="user.id" @click="handleUserProfileClick(user)">
        <div class="user-avatar" :class="user.avatarColor">{{ user.avatar }}</div>
        <div class="user-info">
          <div class="user-name">{{ user.name }}</div>
          <div class="user-status">{{ user.lastActive }}</div>
        </div>
        <div class="status-dot" :class="user.status"></div>
      </div>
    </div>
  </div>

</template>
<style scoped>
/* ==================== 右侧信息栏样式 ==================== */
.right-sidebar {
  display: flex;
  flex-direction: column;
  gap: 1.5rem;
}

.widget {
  background: white;
  border-radius: 0.75rem;
  box-shadow: 0 1px 3px 0 rgba(0, 0, 0, 0.1);
  padding: 1.5rem;
  position: sticky;
  top: 6rem;
}

.widget-header {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  margin-bottom: 1rem;
  padding-bottom: 0.75rem;
  border-bottom: 1px solid #f3f4f6;
}

.widget-header h3 {
  font-weight: 600;
  color: #1f2937;
  font-size: 0.875rem;
}

.widget-header svg {
  width: 1.25rem;
  height: 1.25rem;
}

/* 热门话题样式 */
.topics-list {
  display: flex;
  flex-direction: column;
  gap: 0.75rem;
}

.topic-item {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0.5rem;
  border-radius: 0.375rem;
  cursor: pointer;
  transition: background-color 0.2s;
}

.topic-item:hover {
  background: #f9fafb;
}

.topic-name {
  color: #1f2937;
  font-weight: 500;
  font-size: 0.875rem;
}

.topic-stats {
  text-align: right;
  font-size: 0.75rem;
  color: #6b7280;
  line-height: 1.3;
}

/* 社区统计样式 */
.stats-main {
  text-align: center;
  margin-bottom: 1.5rem;
  padding: 1rem;
  background: linear-gradient(135deg, #f0fdf4, #ecfdf5);
  border-radius: 0.5rem;
}

.main-stat {
  font-size: 2rem;
  font-weight: bold;
  color: #10b981;
  margin-bottom: 0.25rem;
}

.main-stat-label {
  color: #6b7280;
  font-size: 0.875rem;
}

.stats-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 1rem;
  text-align: center;
}

.stat-box {
  padding: 0.75rem;
  background: #f9fafb;
  border-radius: 0.375rem;
}

.stat-number {
  font-size: 1.125rem;
  font-weight: 600;
  margin-bottom: 0.25rem;
}

.stat-number.blue {
  color: #3b82f6;
}

.stat-number.purple {
  color: #8b5cf6;
}

.stat-label {
  font-size: 0.75rem;
  color: #6b7280;
}

/* 活跃用户样式 */
.users-list {
  display: flex;
  flex-direction: column;
  gap: 0.75rem;
}

.user-item {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  padding: 0.5rem;
  border-radius: 0.375rem;
  cursor: pointer;
  transition: background-color 0.2s;
}

.user-item:hover {
  background: #f9fafb;
}

.user-info {
  flex: 1;
  min-width: 0;
}

.user-name {
  font-weight: 500;
  color: #1f2937;
  font-size: 0.875rem;
}

.user-status {
  font-size: 0.75rem;
  color: #6b7280;
  margin-top: 0.125rem;
}

.status-dot {
  width: 0.5rem;
  height: 0.5rem;
  border-radius: 9999px;
  flex-shrink: 0;
}

.status-dot.online {
  background: #10b981;
  box-shadow: 0 0 0 2px rgba(16, 185, 129, 0.2);
}

.status-dot.away {
  background: #f59e0b;
  box-shadow: 0 0 0 2px rgba(245, 158, 11, 0.2);
}
</style>

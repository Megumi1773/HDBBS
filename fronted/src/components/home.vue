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
              <a href="#" class="nav-item" :class="{ active: activeNavItem === 'home' }" @click="setActiveNav('home')">
                <svg width="20" height="20" fill="currentColor" viewBox="0 0 20 20">
                  <path
                    d="M10.707 2.293a1 1 0 00-1.414 0l-7 7a1 1 0 001.414 1.414L4 10.414V17a1 1 0 001 1h2a1 1 0 001-1v-2a1 1 0 011-1h2a1 1 0 011 1v2a1 1 0 001 1h2a1 1 0 001-1v-6.586l.293.293a1 1 0 001.414-1.414l-7-7z" />
                </svg>
                <span>首页</span>
              </a>

              <a href="#" class="nav-item" :class="{ active: activeNavItem === 'study' }"
                @click="setActiveNav('study')">
                <svg width="20" height="20" fill="currentColor" viewBox="0 0 20 20">
                  <path d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
                </svg>
                <span>学习交流</span>
              </a>

              <a href="#" class="nav-item" :class="{ active: activeNavItem === 'job' }" @click="setActiveNav('job')">
                <svg width="20" height="20" fill="currentColor" viewBox="0 0 20 20">
                  <path d="M13 6a3 3 0 11-6 0 3 3 0 016 0zM18 8a2 2 0 11-4 0 2 2 0 014 0zM14 15a4 4 0 00-8 0v3h8v-3z" />
                </svg>
                <span>求职招聘</span>
              </a>

              <a href="#" class="nav-item" :class="{ active: activeNavItem === 'entertainment' }"
                @click="setActiveNav('entertainment')">
                <svg width="20" height="20" fill="currentColor" viewBox="0 0 20 20">
                  <path fill-rule="evenodd"
                    d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7-4a1 1 0 11-2 0 1 1 0 012 0zM9 9a1 1 0 000 2v3a1 1 0 001 1h1a1 1 0 100-2v-3a1 1 0 00-1-1H9z"
                    clip-rule="evenodd" />
                </svg>
                <span>娱乐休闲</span>
              </a>

              <a href="#" class="nav-item" :class="{ active: activeNavItem === 'sports' }"
                @click="setActiveNav('sports')">
                <svg width="20" height="20" fill="currentColor" viewBox="0 0 20 20">
                  <path fill-rule="evenodd"
                    d="M3 6a3 3 0 013-3h10a1 1 0 01.8 1.6L14.25 8l2.55 3.4A1 1 0 0116 13H6a1 1 0 00-1 1v3a1 1 0 11-2 0V6z"
                    clip-rule="evenodd" />
                </svg>
                <span>运动健身</span>
              </a>

              <a href="#" class="nav-item" :class="{ active: activeNavItem === 'food' }" @click="setActiveNav('food')">
                <svg width="20" height="20" fill="currentColor" viewBox="0 0 20 20">
                  <path fill-rule="evenodd"
                    d="M6 2a2 2 0 00-2 2v12a2 2 0 002 2h8a2 2 0 002-2V4a2 2 0 00-2-2H6zm1 2a1 1 0 000 2h6a1 1 0 100-2H7z"
                    clip-rule="evenodd" />
                </svg>
                <span>美食分享</span>
              </a>

              <a href="#" class="nav-item" :class="{ active: activeNavItem === 'startup' }"
                @click="setActiveNav('startup')">
                <svg width="20" height="20" fill="currentColor" viewBox="0 0 20 20">
                  <path fill-rule="evenodd"
                    d="M11.49 3.17c-.38-1.56-2.6-1.56-2.98 0a1.532 1.532 0 01-2.286.948c-1.372-.836-2.942.734-2.106 2.106.54.886.061 2.042-.947 2.287-1.561.379-1.561 2.6 0 2.978a1.532 1.532 0 01.947 2.287c-.836 1.372.734 2.942 2.106 2.106a1.532 1.532 0 012.287.947c.379 1.561 2.6 1.561 2.978 0a1.533 1.533 0 012.287-.947c1.372.836 2.942-.734 2.106-2.106a1.533 1.533 0 01.947-2.287c1.561-.379 1.561-2.6 0-2.978a1.532 1.532 0 01-.947-2.287c.836-1.372-.734-2.942-2.106-2.106a1.532 1.532 0 01-2.287-.947z"
                    clip-rule="evenodd" />
                </svg>
                <span>创业创新</span>
              </a>

              <a href="#" class="nav-item" :class="{ active: activeNavItem === 'help' }" @click="setActiveNav('help')">
                <svg width="20" height="20" fill="currentColor" viewBox="0 0 20 20">
                  <path fill-rule="evenodd"
                    d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-8-3a1 1 0 00-.867.5 1 1 0 11-1.731-1A3 3 0 0113 8a3.001 3.001 0 01-2 2.83V11a1 1 0 11-2 0v-1a1 1 0 011-1 1 1 0 100-2zm0 8a1 1 0 100-2 1 1 0 000 2z"
                    clip-rule="evenodd" />
                </svg>
                <span>问答求助</span>
              </a>
            </nav>
          </div>
        </div>

        <!-- 中间主内容 -->
        <div class="main-content">
          <Message></Message>
        </div>

        <!-- ==================== 右侧信息栏 ==================== -->
        <div class="right-sidebar">

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
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import axios from 'axios';
import Message from './message.vue';
let Classify = reactive([])
let get = () => {
  axios.get('/api/categories')
    .then(response => {
      Classify = response.data.data
      console.log(`主题数据请求成功：`);
      console.log(Classify);

    })
    .catch(error => {
      console.error('Error fetching posts:', error);
    });
}
onMounted(() => {
  get();
});

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

// 帖子数据
const posts = reactive([
  {
    id: 1,
    username: '李明同学',
    avatar: '李',
    avatarColor: 'purple',
    time: '2小时前',
    category: '学习交流',
    categoryColor: 'purple',
    title: '期末复习资料分享 - 高等数学重点知识点整理',
    excerpt: '马上就要期末考试了，整理了一份高等数学的重点知识点，包括极限、导数、积分等核心内容，希望能帮助大家，一起加油！附件中有详细的思维导图和练习题...',
    likes: 128,
    comments: 45,
    views: '1.2k'
  },
  {
    id: 2,
    username: '王小美',
    avatar: '王',
    avatarColor: 'blue',
    time: '4小时前',
    category: '求职招聘',
    categoryColor: 'blue',
    title: '【实习机会】某知名互联网公司前端开发实习生招聘',
    excerpt: '公司正在招聘前端开发实习生，要求熟悉React、Vue等框架，有项目经验优先。实习期间会有导师指导，表现优秀可转正，感兴趣的同学可以私信我...',
    likes: 89,
    comments: 67,
    views: '2.1k'
  },
  {
    id: 3,
    username: '陈大厨',
    avatar: '陈',
    avatarColor: 'green',
    time: '6小时前',
    category: '美食分享',
    categoryColor: 'green',
    title: '宿舍简易版麻辣香锅制作教程，零厨艺也能做！',
    excerpt: '作为一个资深吃货，终于研究出了适合宿舍制作的简易麻辣香锅，用料简单，步骤清晰，保证你们都能学会。图片教程已上传，快来试试吧...',
    likes: 156,
    comments: 32,
    views: '892'
  }
])

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

// ==================== 事件处理函数 ====================

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

// 导航项切换
// const setActiveNav = (navItem) => {
//   activeNavItem.value = navItem
//   console.log('切换到导航项:', navItem)
//   // 这里可以添加路由跳转或内容过滤逻辑
// }

// 发布新帖
const handleNewPost = () => {
  console.log('发布新帖')
  // 这里可以打开发帖弹窗或跳转到发帖页面
}

// 帖子点击处理
const handlePostClick = (post) => {
  console.log('点击帖子:', post.title)
  // 这里可以跳转到帖子详情页
}

// 点赞处理
const handleLike = (post) => {
  post.likes++
  console.log('点赞帖子:', post.title)
}

// 评论处理
const handleComment = (post) => {
  console.log('评论帖子:', post.title)
  // 这里可以打开评论弹窗
}

// 收藏处理
const handleFavorite = (post) => {
  console.log('收藏帖子:', post.title)
  // 这里可以添加收藏逻辑
}

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

.sidebar {
  /* 左侧导航栏样式 */
}

.main-content {
  /* 中间内容区样式 */
}

.right-sidebar {
  /* 右侧信息栏样式 */
}
</style>
<script setup>
import { ref, reactive, onMounted } from 'vue'
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


// ==================== 生命周期钩子 ====================
onMounted(() => {
  console.log('校园论坛组件已挂载')
  // 这里可以添加初始化逻辑，如获取数据等
})
</script>

<template>
  <!-- ==================== 中间主要内容区域 ==================== -->
  <div class="main-content">
    <div class="posts-container">
      <!-- 帖子列表头部 -->
      <div class="posts-header">
        <div class="posts-title">
          <svg width="20" height="20" fill="currentColor" viewBox="0 0 20 20">
            <path fill-rule="evenodd"
              d="M10 18a8 8 0 100-16 8 8 0 000 16zm1-12a1 1 0 10-2 0v4a1 1 0 00.293.707l2.828 2.829a1 1 0 101.415-1.415L11 9.586V6z"
              clip-rule="evenodd" />
          </svg>
          <h2>最新帖子</h2>
        </div>
        <button class="new-post-btn" @click="handleNewPost">+ 发布新帖</button>
      </div>

      <!-- 帖子列表 -->
      <div class="posts-list">
        <!-- 帖子项目 1 - 学习交流类 -->
        <div class="post-item" v-for="post in posts" :key="post.id">
          <div class="post-content">
            <!-- 用户头像 -->
            <div class="post-avatar" :class="post.avatarColor">{{ post.avatar }}</div>

            <!-- 帖子主要内容 -->
            <div class="post-main">
              <!-- 帖子元信息：用户名、时间、分类 -->
              <div class="post-meta">
                <div class="post-user-info">
                  <h3 class="post-username">{{ post.username }}</h3>
                  <p class="post-time">{{ post.time }}</p>
                </div>
                <span class="post-category" :class="post.categoryColor">{{ post.category }}</span>
              </div>

              <!-- 帖子标题 -->
              <h4 class="post-title" @click="handlePostClick(post)">{{ post.title }}</h4>

              <!-- 帖子摘要 -->
              <p class="post-excerpt">{{ post.excerpt }}</p>

              <!-- 帖子统计信息：点赞、评论、浏览、收藏 -->
              <div class="post-stats">
                <div class="stat-item" @click="handleLike(post)">
                  <svg width="16" height="16" fill="currentColor" viewBox="0 0 20 20">
                    <path
                      d="M2 10.5a1.5 1.5 0 113 0v6a1.5 1.5 0 01-3 0v-6zM6 10.333v5.43a2 2 0 001.106 1.79l.05.025A4 4 0 008.943 18h5.416a2 2 0 001.962-1.608l1.2-6A2 2 0 0015.56 8H12V4a2 2 0 00-2-2 1 1 0 00-1 1v.667a4 4 0 01-.8 2.4L6.8 7.933a4 4 0 00-.8 2.4z" />
                  </svg>
                  <span>{{ post.likes }}</span>
                </div>
                <div class="stat-item" @click="handleComment(post)">
                  <svg width="16" height="16" fill="currentColor" viewBox="0 0 20 20">
                    <path fill-rule="evenodd"
                      d="M18 13V5a2 2 0 00-2-2H4a2 2 0 00-2 2v8a2 2 0 002 2h3l3 3 3-3h3a2 2 0 002-2zM5 7a1 1 0 011-1h8a1 1 0 110 2H6a1 1 0 01-1-1zm1 3a1 1 0 100 2h3a1 1 0 100-2H6z"
                      clip-rule="evenodd" />
                  </svg>
                  <span>{{ post.comments }}</span>
                </div>
                <div class="stat-item">
                  <svg width="16" height="16" fill="currentColor" viewBox="0 0 20 20">
                    <path d="M10 12a2 2 0 100-4 2 2 0 000 4z" />
                    <path fill-rule="evenodd"
                      d="M.458 10C1.732 5.943 5.522 3 10 3s8.268 2.943 9.542 7c-1.274 4.057-5.064 7-9.542 7S1.732 14.057.458 10zM14 10a4 4 0 11-8 0 4 4 0 018 0z"
                      clip-rule="evenodd" />
                  </svg>
                  <span>{{ post.views }}</span>
                </div>
                <div class="stat-item" @click="handleFavorite(post)">
                  <svg width="16" height="16" fill="currentColor" viewBox="0 0 20 20">
                    <path
                      d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z" />
                  </svg>
                  <span>收藏</span>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<style scoped>
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

/* ==================== 主要内容区域样式 ==================== */
.main-content {
  grid-column: 2;
}

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
  grid-column: 3;
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
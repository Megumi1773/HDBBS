<template>
  <!-- 登录页面主容器 -->
  <div class="login-container">
    <!-- 背景装饰元素 -->
    <div class="bg-decoration">
      <div class="floating-shape shape-1"></div>
      <div class="floating-shape shape-2"></div>
      <div class="floating-shape shape-3"></div>
      <div class="floating-shape shape-4"></div>
    </div>

    <div class="login-wrapper">
      <!-- 头部区域：Logo和标题 -->
      <div class="header-section">
        <!-- Logo容器 -->
        <div class="logo-container">
          <el-icon class="logo-icon" :size="32">
            <ChatDotRound />
          </el-icon>
          <div class="logo-glow"></div>
        </div>
        <!-- 主标题 -->
        <h1 class="title">欢迎回来</h1>
        <!-- 副标题 -->
        <p class="subtitle">登录您的论坛账户</p>
      </div>

      <!-- 登录表单卡片 -->
      <el-card class="login-card" shadow="always">
        <!-- 卡片顶部装饰 -->
        <div class="card-decoration"></div>

        <!-- 登录表单 -->
        <el-form ref="loginFormRef" :model="form" :rules="rules" @submit.prevent="handleLogin" class="login-form"
          size="large">
          <!-- 邮箱/用户名输入框 -->
          <el-form-item prop="email" class="form-item">
            <template #label>
              <span class="form-label">
                <el-icon class="label-icon">
                  <User />
                </el-icon>
                邮箱或用户名
              </span>
            </template>
            <el-input v-model="form.email" placeholder="请输入邮箱或用户名" class="form-input" clearable>
              <template #prefix>
                <el-icon class="input-icon">

                </el-icon>
              </template>
            </el-input>
          </el-form-item>

          <!-- 密码输入框 -->
          <el-form-item prop="password" class="form-item">
            <template #label>
              <span class="form-label">
                <el-icon class="label-icon">
                  <Lock />
                </el-icon>
                密码
              </span>
            </template>
            <el-input v-model="form.password" :type="showPassword ? 'text' : 'password'" placeholder="请输入密码"
              class="form-input">
              <!-- 密码输入框前缀图标 -->
              <template #prefix>
                <el-icon class="input-icon">
                  <Lock />
                </el-icon>
              </template>
              <!-- 密码显示/隐藏切换按钮 -->
              <template #suffix>
                <el-icon class="password-toggle" @click="showPassword = !showPassword">
                  <View v-if="!showPassword" />
                  <Hide v-else />
                </el-icon>
              </template>
            </el-input>
          </el-form-item>

          <!-- 记住我复选框和忘记密码链接 -->
          <div class="form-options">
            <el-checkbox v-model="form.remember" class="remember-checkbox">
              <span class="checkbox-text">记住我</span>
            </el-checkbox>
            <el-link type="primary" class="forgot-link" :underline="false">
              忘记密码？
            </el-link>
          </div>

          <!-- 登录按钮 -->
          <el-form-item class="login-button-item">
            <el-button type="primary" class="login-button" :loading="loading" @click="handleLogin" size="large">
              <!-- 加载状态图标 -->
              <template #loading>
                <el-icon class="is-loading">
                  <Loading />
                </el-icon>
              </template>
              <span class="button-text">{{ loading ? '登录中...' : '登录' }}</span>
              <div class="button-shine"></div>
            </el-button>
          </el-form-item>
        </el-form>

        <!-- 分割线 -->
        <el-divider class="login-divider">
          <span class="divider-text">或者使用以下方式登录</span>
        </el-divider>

        <!-- 第三方登录按钮组 -->
        <div class="social-login">
          <!-- GitHub登录按钮 -->
          <el-button class="social-button github-button" plain>
            <el-icon class="social-icon">
              <!-- GitHub图标SVG -->
              <svg viewBox="0 0 24 24" width="20" height="20">
                <path fill="currentColor"
                  d="M12 0c-6.626 0-12 5.373-12 12 0 5.302 3.438 9.8 8.207 11.387.599.111.793-.261.793-.577v-2.234c-3.338.726-4.033-1.416-4.033-1.416-.546-1.387-1.333-1.756-1.333-1.756-1.089-.745.083-.729.083-.729 1.205.084 1.839 1.237 1.839 1.237 1.07 1.834 2.807 1.304 3.492.997.107-.775.418-1.305.762-1.604-2.665-.305-5.467-1.334-5.467-5.931 0-1.311.469-2.381 1.236-3.221-.124-.303-.535-1.524.117-3.176 0 0 1.008-.322 3.301 1.23.957-.266 1.983-.399 3.003-.404 1.02.005 2.047.138 3.006.404 2.291-1.552 3.297-1.23 3.297-1.23.653 1.653.242 2.874.118 3.176.77.84 1.235 1.911 1.235 3.221 0 4.609-2.807 5.624-5.479 5.921.43.372.823 1.102.823 2.222v3.293c0 .319.192.694.801.576 4.765-1.589 8.199-6.086 8.199-11.386 0-6.627-5.373-12-12-12z" />
              </svg>
            </el-icon>
            <span>GitHub</span>
          </el-button>

          <!-- Google登录按钮 -->
          <el-button class="social-button google-button" plain>
            <el-icon class="social-icon">
              <!-- Google图标SVG -->
              <svg viewBox="0 0 24 24" width="20" height="20">
                <path fill="currentColor"
                  d="M22.56 12.25c0-.78-.07-1.53-.2-2.25H12v4.26h5.92c-.26 1.37-1.04 2.53-2.21 3.31v2.77h3.57c2.08-1.92 3.28-4.74 3.28-8.09z" />
                <path fill="currentColor"
                  d="M12 23c2.97 0 5.46-.98 7.28-2.66l-3.57-2.77c-.98.66-2.23 1.06-3.71 1.06-2.86 0-5.29-1.93-6.16-4.53H2.18v2.84C3.99 20.53 7.7 23 12 23z" />
                <path fill="currentColor"
                  d="M5.84 14.09c-.22-.66-.35-1.36-.35-2.09s.13-1.43.35-2.09V7.07H2.18C1.43 8.55 1 10.22 1 12s.43 3.45 1.18 4.93l2.85-2.22.81-.62z" />
                <path fill="currentColor"
                  d="M12 5.38c1.62 0 3.06.56 4.21 1.64l3.15-3.15C17.45 2.09 14.97 1 12 1 7.7 1 3.99 3.47 2.18 7.07l3.66 2.84c.87-2.6 3.3-4.53 6.16-4.53z" />
              </svg>
            </el-icon>
            <span>Google</span>
          </el-button>
        </div>

        <!-- 注册链接区域 -->
        <div class="signup-section">
          <span class="signup-text">还没有账户？</span>
          <el-link type="primary" class="signup-link" :underline="false">
            立即注册
            <el-icon class="signup-arrow">

            </el-icon>
          </el-link>
        </div>
      </el-card>

      <!-- 页脚版权信息 -->
      <div class="footer">
        <p>© 2024 论坛名称. 保留所有权利.</p>
        <div class="footer-links">
          <el-link class="footer-link" :underline="false">隐私政策</el-link>
          <span class="footer-divider">·</span>
          <el-link class="footer-link" :underline="false">服务条款</el-link>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'


// 表单引用，用于表单验证
const loginFormRef = ref()

// 表单数据对象
const form = reactive({
  email: '',      // 邮箱或用户名
  password: '',   // 密码
  remember: false // 记住我选项
})

// 表单验证规则
const rules = {
  email: [
    { required: true, message: '请输入邮箱或用户名', trigger: 'blur' },
    { min: 3, message: '长度至少为3个字符', trigger: 'blur' }
  ],
  password: [
    { required: true, message: '请输入密码', trigger: 'blur' },
    { min: 6, message: '密码长度至少为6个字符', trigger: 'blur' }
  ]
}

// UI状态管理
const showPassword = ref(false) // 控制密码显示/隐藏
const loading = ref(false)      // 控制登录按钮加载状态

/**
 * 处理登录逻辑
 * 包含表单验证、API调用和错误处理
 */
const handleLogin = async () => {
  // 检查表单引用是否存在
  if (!loginFormRef.value) return

  try {
    // 执行表单验证
    const valid = await loginFormRef.value.validate()
    if (!valid) return

    // 设置加载状态
    loading.value = true

    // 模拟API调用（实际项目中替换为真实的登录API）
    await new Promise(resolve => setTimeout(resolve, 2000))

    // 打印登录尝试信息（调试用）
    console.log('Login attempt:', form)

    // 显示成功消息
    ElMessage.success({
      message: '登录成功！欢迎回来',
      type: 'success',
      duration: 3000
    })

    // 登录成功后的处理逻辑
    // 例如：跳转到仪表板页面
    // router.push('/dashboard')

  } catch (error) {
    // 错误处理
    console.error('Login error:', error)
    ElMessage.error({
      message: '登录失败，请检查您的凭据',
      type: 'error',
      duration: 3000
    })
  } finally {
    // 无论成功或失败都要重置加载状态
    loading.value = false
  }
}
</script>

<style scoped>
/* 登录页面主容器样式 */
.login-container {
  min-height: 100vh;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 50%, #f093fb 100%);
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 20px;
  position: relative;
  overflow: hidden;
  animation: fadeIn 0.8s ease-out;
}

/* 背景装饰元素 */
.bg-decoration {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  pointer-events: none;
  z-index: 0;
}

.floating-shape {
  position: absolute;
  border-radius: 50%;
  background: rgba(255, 255, 255, 0.1);
  backdrop-filter: blur(10px);
  animation: float 6s ease-in-out infinite;
}

.shape-1 {
  width: 80px;
  height: 80px;
  top: 10%;
  left: 10%;
  animation-delay: 0s;
}

.shape-2 {
  width: 120px;
  height: 120px;
  top: 20%;
  right: 15%;
  animation-delay: 2s;
}

.shape-3 {
  width: 60px;
  height: 60px;
  bottom: 30%;
  left: 20%;
  animation-delay: 4s;
}

.shape-4 {
  width: 100px;
  height: 100px;
  bottom: 20%;
  right: 10%;
  animation-delay: 1s;
}

/* 登录表单包装器 */
.login-wrapper {
  width: 100%;
  max-width: 440px;
  position: relative;
  z-index: 1;
}

/* 头部区域样式 */
.header-section {
  text-align: center;
  margin-bottom: 40px;
  animation: slideDown 0.8s ease-out 0.2s both;
}

/* Logo容器样式 */
.logo-container {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  width: 80px;
  height: 80px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border-radius: 50%;
  margin-bottom: 20px;
  box-shadow:
    0 10px 40px rgba(102, 126, 234, 0.4),
    0 0 0 10px rgba(255, 255, 255, 0.1);
  position: relative;
  transition: all 0.3s ease;
}

.logo-container:hover {
  transform: translateY(-5px) scale(1.05);
  box-shadow:
    0 15px 50px rgba(102, 126, 234, 0.5),
    0 0 0 15px rgba(255, 255, 255, 0.15);
}

.logo-glow {
  position: absolute;
  top: -10px;
  left: -10px;
  right: -10px;
  bottom: -10px;
  background: linear-gradient(135deg, #667eea, #764ba2);
  border-radius: 50%;
  opacity: 0.3;
  filter: blur(20px);
  z-index: -1;
  animation: pulse 2s ease-in-out infinite;
}

/* Logo图标样式 */
.logo-icon {
  color: white;
  filter: drop-shadow(0 2px 4px rgba(0, 0, 0, 0.2));
}

/* 主标题样式 */
.title {
  font-size: 36px;
  font-weight: 800;
  color: white;
  margin: 0 0 12px 0;
  text-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
  background: linear-gradient(135deg, #ffffff 0%, #f0f0f0 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

/* 副标题样式 */
.subtitle {
  color: rgba(255, 255, 255, 0.9);
  font-size: 18px;
  margin: 0;
  font-weight: 300;
  letter-spacing: 0.5px;
}

/* 登录卡片样式 */
.login-card {
  border-radius: 24px;
  border: none;
  box-shadow:
    0 25px 80px rgba(0, 0, 0, 0.15),
    0 0 0 1px rgba(255, 255, 255, 0.1);
  overflow: hidden;
  backdrop-filter: blur(20px);
  background: rgba(255, 255, 255, 0.95);
  animation: slideUp 0.8s ease-out 0.4s both;
  position: relative;
}

.card-decoration {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 4px;
  background: linear-gradient(90deg, #667eea 0%, #764ba2 50%, #f093fb 100%);
}

/* 卡片内容区域样式 */
:deep(.el-card__body) {
  padding: 48px;
}

/* 登录表单样式 */
.login-form {
  width: 100%;
}

/* 表单项样式 */
.form-item {
  margin-bottom: 28px;
}

/* 表单标签样式 */
.form-label {
  font-weight: 600;
  color: #2d3748;
  font-size: 15px;
  display: flex;
  align-items: center;
  gap: 8px;
}

.label-icon {
  color: #667eea;
  font-size: 16px;
}

/* Element Plus表单标签样式覆盖 */
:deep(.el-form-item__label) {
  padding-bottom: 12px;
}

/* 表单输入框样式 */
.form-input {
  height: 52px;
}

/* Element Plus输入框包装器样式覆盖 */
:deep(.el-input__wrapper) {
  border-radius: 16px;
  background-color: #f8fafc;
  border: 2px solid #e2e8f0;
  transition: all 0.3s ease;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.04);
}

/* 输入框悬停状态 */
:deep(.el-input__wrapper:hover) {
  border-color: #667eea;
  background-color: white;
  box-shadow: 0 4px 12px rgba(102, 126, 234, 0.1);
}

/* 输入框聚焦状态 */
:deep(.el-input__wrapper.is-focus) {
  border-color: #667eea;
  background-color: white;
  box-shadow:
    0 0 0 4px rgba(102, 126, 234, 0.1),
    0 4px 12px rgba(102, 126, 234, 0.15);
}

/* 输入框图标样式 */
.input-icon {
  color: #a0aec0;
  transition: color 0.3s ease;
}

:deep(.el-input__wrapper.is-focus) .input-icon {
  color: #667eea;
}

/* 密码显示/隐藏切换按钮样式 */
.password-toggle {
  color: #a0aec0;
  cursor: pointer;
  transition: all 0.3s ease;
  padding: 4px;
  border-radius: 6px;
}

.password-toggle:hover {
  color: #667eea;
  background-color: rgba(102, 126, 234, 0.1);
}

/* 表单选项区域样式（记住我和忘记密码） */
.form-options {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 36px;
}

/* 记住我复选框样式 */
.remember-checkbox {
  font-size: 14px;
  color: #4a5568;
}

.checkbox-text {
  margin-left: 8px;
  font-weight: 500;
}

/* Element Plus复选框标签样式覆盖 */
:deep(.el-checkbox__label) {
  color: #4a5568;
  font-weight: 500;
}

:deep(.el-checkbox__input.is-checked .el-checkbox__inner) {
  background-color: #667eea;
  border-color: #667eea;
}

/* 忘记密码链接样式 */
.forgot-link {
  font-size: 14px;
  font-weight: 600;
  color: #667eea;
  transition: all 0.3s ease;
}

.forgot-link:hover {
  color: #5a67d8;
  transform: translateX(2px);
}

/* 登录按钮容器样式 */
.login-button-item {
  margin-bottom: 0;
}

/* 登录按钮样式 */
.login-button {
  width: 100%;
  height: 52px;
  border-radius: 16px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border: none;
  font-weight: 700;
  font-size: 16px;
  box-shadow:
    0 6px 20px rgba(102, 126, 234, 0.4),
    0 2px 4px rgba(0, 0, 0, 0.1);
  transition: all 0.3s ease;
  position: relative;
  overflow: hidden;
}

.button-text {
  position: relative;
  z-index: 2;
}

.button-shine {
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.3), transparent);
  transition: left 0.5s ease;
}

/* 登录按钮悬停效果 */
.login-button:hover {
  transform: translateY(-2px);
  box-shadow:
    0 8px 25px rgba(102, 126, 234, 0.5),
    0 4px 8px rgba(0, 0, 0, 0.15);
}

.login-button:hover .button-shine {
  left: 100%;
}

.login-button:active {
  transform: translateY(0);
}

/* 分割线样式 */
.login-divider {
  margin: 40px 0;
}

/* 分割线文字样式 */
.divider-text {
  color: #718096;
  font-size: 14px;
  background: white;
  padding: 0 20px;
  font-weight: 500;
}

/* 第三方登录按钮组样式 */
.social-login {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 16px;
  margin-bottom: 36px;
}

/* 第三方登录按钮样式 */
.social-button {
  height: 48px;
  border-radius: 16px;
  border: 2px solid #e2e8f0;
  background: white;
  color: #4a5568;
  font-weight: 600;
  transition: all 0.3s ease;
  position: relative;
  overflow: hidden;
}

.social-button::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(0, 0, 0, 0.05), transparent);
  transition: left 0.5s ease;
}

.social-button:hover::before {
  left: 100%;
}

/* GitHub按钮特殊样式 */
.github-button:hover {
  border-color: #24292e;
  color: #24292e;
  background: #f6f8fa;
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(36, 41, 46, 0.15);
}

/* Google按钮特殊样式 */
.google-button:hover {
  border-color: #4285f4;
  color: #4285f4;
  background: #f8fbff;
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(66, 133, 244, 0.15);
}

/* 第三方登录按钮图标样式 */
.social-icon {
  margin-right: 10px;
  transition: transform 0.3s ease;
}

.social-button:hover .social-icon {
  transform: scale(1.1);
}

/* 注册链接区域样式 */
.signup-section {
  text-align: center;
  padding: 20px 0;
  background: linear-gradient(135deg, #f8fafc 0%, #e2e8f0 100%);
  margin: 0 -48px -48px -48px;
  border-radius: 0 0 24px 24px;
}

/* 注册提示文字样式 */
.signup-text {
  color: #718096;
  font-size: 15px;
  margin-right: 8px;
  font-weight: 500;
}

/* 注册链接样式 */
.signup-link {
  font-weight: 700;
  font-size: 15px;
  color: #667eea;
  transition: all 0.3s ease;
  display: inline-flex;
  align-items: center;
  gap: 4px;
}

.signup-link:hover {
  color: #5a67d8;
  transform: translateX(3px);
}

.signup-arrow {
  font-size: 14px;
  transition: transform 0.3s ease;
}

.signup-link:hover .signup-arrow {
  transform: translateX(2px);
}

/* 页脚样式 */
.footer {
  text-align: center;
  margin-top: 40px;
  animation: fadeIn 1s ease-out 0.8s both;
}

.footer p {
  color: rgba(255, 255, 255, 0.8);
  font-size: 13px;
  margin: 0 0 12px 0;
  font-weight: 400;
}

.footer-links {
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 8px;
}

.footer-link {
  color: rgba(255, 255, 255, 0.7);
  font-size: 12px;
  font-weight: 500;
  transition: all 0.3s ease;
}

.footer-link:hover {
  color: white;
  transform: translateY(-1px);
}

.footer-divider {
  color: rgba(255, 255, 255, 0.5);
  font-size: 12px;
}

/* 动画定义 */
@keyframes fadeIn {
  from {
    opacity: 0;
    transform: translateY(30px);
  }

  to {
    opacity: 1;
    transform: translateY(0);
  }
}

@keyframes slideDown {
  from {
    opacity: 0;
    transform: translateY(-30px);
  }

  to {
    opacity: 1;
    transform: translateY(0);
  }
}

@keyframes slideUp {
  from {
    opacity: 0;
    transform: translateY(30px);
  }

  to {
    opacity: 1;
    transform: translateY(0);
  }
}

@keyframes float {

  0%,
  100% {
    transform: translateY(0px) rotate(0deg);
  }

  50% {
    transform: translateY(-20px) rotate(180deg);
  }
}

@keyframes pulse {

  0%,
  100% {
    opacity: 0.3;
    transform: scale(1);
  }

  50% {
    opacity: 0.5;
    transform: scale(1.1);
  }
}

/* 响应式设计 - 移动端适配 */
@media (max-width: 480px) {
  .login-container {
    padding: 16px;
  }

  :deep(.el-card__body) {
    padding: 32px 24px;
  }

  .title {
    font-size: 28px;
  }

  .subtitle {
    font-size: 16px;
  }

  .logo-container {
    width: 70px;
    height: 70px;
  }

  /* 移动端第三方登录按钮改为单列布局 */
  .social-login {
    grid-template-columns: 1fr;
    gap: 12px;
  }

  .signup-section {
    margin: 0 -24px -32px -24px;
  }

  .floating-shape {
    display: none;
  }
}

/* 深色模式适配 */
@media (prefers-color-scheme: dark) {
  .login-card {
    background: rgba(26, 32, 44, 0.95);
  }

  .form-label {
    color: #e2e8f0;
  }

  :deep(.el-input__wrapper) {
    background-color: #2d3748;
    border-color: #4a5568;
    color: #e2e8f0;
  }

  .divider-text {
    background: #2d3748;
    color: #a0aec0;
  }

  .signup-section {
    background: linear-gradient(135deg, #2d3748 0%, #4a5568 100%);
  }
}
</style>

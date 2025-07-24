<template>
  <div class="min-h-screen bg-gray-50">
    <!-- Header -->
    <header class="bg-gradient-to-r from-purple-500 to-blue-500 text-white shadow-lg">
      <div class="container mx-auto px-4 py-4">
        <div class="flex items-center justify-between">
          <!-- Logo and Title -->
          <div class="flex items-center space-x-3">
            <div class="w-8 h-8 bg-white/20 rounded-lg flex items-center justify-center">
              <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 20 20">
                <path
                  d="M10.394 2.08a1 1 0 00-.788 0l-7 3a1 1 0 000 1.84L5.25 8.051a.999.999 0 01.356-.257l4-1.714a1 1 0 11.788 1.84L7.667 9.088l1.94.831a1 1 0 00.787 0l7-3a1 1 0 000-1.84l-7-3zM3.31 9.397L5 10.12v4.102a8.969 8.969 0 00-1.05-.174 1 1 0 01-.89-.89 11.115 11.115 0 01.25-3.762zM9.3 16.573A9.026 9.026 0 007 14.935v-3.957l1.818.78a3 3 0 002.364 0l5.508-2.361a11.026 11.026 0 01.25 3.762 1 1 0 01-.89.89 8.968 8.968 0 00-5.35 2.524 1 1 0 01-1.4 0zM6 18a1 1 0 001-1v-2.065a8.935 8.935 0 00-2-.712V17a1 1 0 001 1z" />
              </svg>
            </div>
            <h1 class="text-xl font-bold">校园论坛</h1>
          </div>

          <!-- Search Bar -->
          <div class="flex-1 max-w-md mx-8">
            <div class="relative">
              <input type="text" placeholder="搜索帖子、用户或话题..."
                class="w-full px-4 py-2 pl-10 rounded-full bg-white/10 backdrop-blur-sm border border-white/20 text-white placeholder-white/70 focus:outline-none focus:ring-2 focus:ring-white/30">
              <svg class="absolute left-3 top-2.5 w-5 h-5 text-white/70" fill="none" stroke="currentColor"
                viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                  d="m21 21-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
              </svg>
            </div>
          </div>

          <!-- User Actions -->
          <div class="flex items-center space-x-4">
            <div class="relative">
              <button class="p-2 rounded-full bg-white/10 hover:bg-white/20 transition-colors">
                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                    d="M15 17h5l-5 5v-5zM9 7H4l5-5v5zm6 10V7a1 1 0 00-1-1H5a1 1 0 00-1 1v10a1 1 0 001 1h9a1 1 0 001-1z" />
                </svg>
              </button>
              <span
                class="absolute -top-1 -right-1 bg-red-500 text-xs rounded-full w-5 h-5 flex items-center justify-center">3</span>
            </div>
            <div class="w-8 h-8 bg-orange-400 rounded-full flex items-center justify-center font-bold text-sm">
              张
            </div>
          </div>
        </div>
      </div>
    </header>

    <div class="container mx-auto px-4 py-6">
      <div class="grid grid-cols-1 lg:grid-cols-4 gap-6">
        <!-- Left Sidebar -->
        <div class="lg:col-span-1">
          <div class="bg-white rounded-xl shadow-sm p-6">
            <div class="flex items-center space-x-2 mb-6">
              <svg class="w-5 h-5 text-purple-500" fill="currentColor" viewBox="0 0 20 20">
                <path
                  d="M7 3a1 1 0 000 2h6a1 1 0 100-2H7zM4 7a1 1 0 011-1h10a1 1 0 110 2H5a1 1 0 01-1-1zM2 11a2 2 0 012-2h12a2 2 0 012 2v4a2 2 0 01-2 2H4a2 2 0 01-2-2v-4z" />
              </svg>
              <h2 class="font-semibold text-gray-800">论坛版块</h2>
            </div>

            <nav class="space-y-2">
              <a href="#" class="flex items-center space-x-3 px-4 py-3 rounded-lg bg-purple-500 text-white font-medium">
                <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 20 20">
                  <path
                    d="M10.707 2.293a1 1 0 00-1.414 0l-7 7a1 1 0 001.414 1.414L4 10.414V17a1 1 0 001 1h2a1 1 0 001-1v-2a1 1 0 011-1h2a1 1 0 011 1v2a1 1 0 001 1h2a1 1 0 001-1v-6.586l.293.293a1 1 0 001.414-1.414l-7-7z" />
                </svg>
                <span>首页</span>
              </a>

              <a href="#"
                class="flex items-center space-x-3 px-4 py-3 rounded-lg text-gray-600 hover:bg-gray-50 transition-colors">
                <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 20 20">
                  <path d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
                </svg>
                <span>学习交流</span>
              </a>

              <a href="#"
                class="flex items-center space-x-3 px-4 py-3 rounded-lg text-gray-600 hover:bg-gray-50 transition-colors">
                <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 20 20">
                  <path d="M13 6a3 3 0 11-6 0 3 3 0 016 0zM18 8a2 2 0 11-4 0 2 2 0 014 0zM14 15a4 4 0 00-8 0v3h8v-3z" />
                </svg>
                <span>求职招聘</span>
              </a>

              <a href="#"
                class="flex items-center space-x-3 px-4 py-3 rounded-lg text-gray-600 hover:bg-gray-50 transition-colors">
                <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 20 20">
                  <path fill-rule="evenodd"
                    d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7-4a1 1 0 11-2 0 1 1 0 012 0zM9 9a1 1 0 000 2v3a1 1 0 001 1h1a1 1 0 100-2v-3a1 1 0 00-1-1H9z"
                    clip-rule="evenodd" />
                </svg>
                <span>娱乐休闲</span>
              </a>

              <a href="#"
                class="flex items-center space-x-3 px-4 py-3 rounded-lg text-gray-600 hover:bg-gray-50 transition-colors">
                <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 20 20">
                  <path fill-rule="evenodd"
                    d="M3 6a3 3 0 013-3h10a1 1 0 01.8 1.6L14.25 8l2.55 3.4A1 1 0 0116 13H6a1 1 0 00-1 1v3a1 1 0 11-2 0V6z"
                    clip-rule="evenodd" />
                </svg>
                <span>运动健身</span>
              </a>

              <a href="#"
                class="flex items-center space-x-3 px-4 py-3 rounded-lg text-gray-600 hover:bg-gray-50 transition-colors">
                <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 20 20">
                  <path fill-rule="evenodd"
                    d="M6 2a2 2 0 00-2 2v12a2 2 0 002 2h8a2 2 0 002-2V4a2 2 0 00-2-2H6zm1 2a1 1 0 000 2h6a1 1 0 100-2H7zm6 7a1 1 0 011 1v3a1 1 0 11-2 0v-3a1 1 0 011-1zm-3 3a1 1 0 100 2h.01a1 1 0 100-2H10zm-4 1a1 1 0 011-1h.01a1 1 0 110 2H7a1 1 0 01-1-1zm1-4a1 1 0 100 2h.01a1 1 0 100-2H7zm2 1a1 1 0 011-1h.01a1 1 0 110 2H10a1 1 0 01-1-1zm4-4a1 1 0 100 2h.01a1 1 0 100-2H13zM9 9a1 1 0 011-1h.01a1 1 0 110 2H10a1 1 0 01-1-1zM7 8a1 1 0 000 2h.01a1 1 0 000-2H7z"
                    clip-rule="evenodd" />
                </svg>
                <span>美食分享</span>
              </a>

              <a href="#"
                class="flex items-center space-x-3 px-4 py-3 rounded-lg text-gray-600 hover:bg-gray-50 transition-colors">
                <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 20 20">
                  <path fill-rule="evenodd"
                    d="M11.49 3.17c-.38-1.56-2.6-1.56-2.98 0a1.532 1.532 0 01-2.286.948c-1.372-.836-2.942.734-2.106 2.106.54.886.061 2.042-.947 2.287-1.561.379-1.561 2.6 0 2.978a1.532 1.532 0 01.947 2.287c-.836 1.372.734 2.942 2.106 2.106a1.532 1.532 0 012.287.947c.379 1.561 2.6 1.561 2.978 0a1.533 1.533 0 012.287-.947c1.372.836 2.942-.734 2.106-2.106a1.533 1.533 0 01.947-2.287c1.561-.379 1.561-2.6 0-2.978a1.532 1.532 0 01-.947-2.287c.836-1.372-.734-2.942-2.106-2.106a1.532 1.532 0 01-2.287-.947zM10 13a3 3 0 100-6 3 3 0 000 6z"
                    clip-rule="evenodd" />
                </svg>
                <span>创业创新</span>
              </a>

              <a href="#"
                class="flex items-center space-x-3 px-4 py-3 rounded-lg text-gray-600 hover:bg-gray-50 transition-colors">
                <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 20 20">
                  <path fill-rule="evenodd"
                    d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-8-3a1 1 0 00-.867.5 1 1 0 11-1.731-1A3 3 0 0113 8a3.001 3.001 0 01-2 2.83V11a1 1 0 11-2 0v-1a1 1 0 011-1 1 1 0 100-2zm0 8a1 1 0 100-2 1 1 0 000 2z"
                    clip-rule="evenodd" />
                </svg>
                <span>问答求助</span>
              </a>
            </nav>
          </div>
        </div>

        <!-- Main Content -->
        <div class="lg:col-span-2">
          <div class="bg-white rounded-xl shadow-sm">
            <!-- Header -->
            <div class="flex items-center justify-between p-6 border-b border-gray-100">
              <div class="flex items-center space-x-2">
                <svg class="w-5 h-5 text-blue-500" fill="currentColor" viewBox="0 0 20 20">
                  <path fill-rule="evenodd"
                    d="M10 18a8 8 0 100-16 8 8 0 000 16zm1-12a1 1 0 10-2 0v4a1 1 0 00.293.707l2.828 2.829a1 1 0 101.415-1.415L11 9.586V6z"
                    clip-rule="evenodd" />
                </svg>
                <h2 class="text-lg font-semibold text-gray-800">最新帖子</h2>
              </div>
              <button
                class="px-4 py-2 bg-green-500 text-white rounded-lg hover:bg-green-600 transition-colors font-medium">
                + 发布新帖
              </button>
            </div>

            <!-- Posts -->
            <div class="divide-y divide-gray-100">
              <!-- Post 1 -->
              <div class="p-6">
                <div class="flex items-start space-x-4">
                  <div
                    class="w-10 h-10 bg-purple-500 rounded-full flex items-center justify-center text-white font-medium">
                    李
                  </div>
                  <div class="flex-1">
                    <div class="flex items-center justify-between mb-2">
                      <div>
                        <h3 class="font-medium text-gray-900">李明同学</h3>
                        <p class="text-sm text-gray-500">2小时前</p>
                      </div>
                      <span class="px-3 py-1 bg-purple-100 text-purple-700 rounded-full text-sm">学习交流</span>
                    </div>
                    <h4 class="text-lg font-semibold text-gray-900 mb-2">期末复习资料分享 - 高等数学重点知识点整理</h4>
                    <p class="text-gray-600 mb-4">
                      马上就要期末考试了，整理了一份高等数学的重点知识点，包括极限、导数、积分等核心内容，希望能帮助大家，一起加油！附件中有详细的思维导图和练习题...</p>
                    <div class="flex items-center space-x-6 text-sm text-gray-500">
                      <div class="flex items-center space-x-1">
                        <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 20 20">
                          <path
                            d="M2 10.5a1.5 1.5 0 113 0v6a1.5 1.5 0 01-3 0v-6zM6 10.333v5.43a2 2 0 001.106 1.79l.05.025A4 4 0 008.943 18h5.416a2 2 0 001.962-1.608l1.2-6A2 2 0 0015.56 8H12V4a2 2 0 00-2-2 1 1 0 00-1 1v.667a4 4 0 01-.8 2.4L6.8 7.933a4 4 0 00-.8 2.4z" />
                        </svg>
                        <span>128</span>
                      </div>
                      <div class="flex items-center space-x-1">
                        <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 20 20">
                          <path fill-rule="evenodd"
                            d="M18 13V5a2 2 0 00-2-2H4a2 2 0 00-2 2v8a2 2 0 002 2h3l3 3 3-3h3a2 2 0 002-2zM5 7a1 1 0 011-1h8a1 1 0 110 2H6a1 1 0 01-1-1zm1 3a1 1 0 100 2h3a1 1 0 100-2H6z"
                            clip-rule="evenodd" />
                        </svg>
                        <span>45</span>
                      </div>
                      <div class="flex items-center space-x-1">
                        <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 20 20">
                          <path d="M10 12a2 2 0 100-4 2 2 0 000 4z" />
                          <path fill-rule="evenodd"
                            d="M.458 10C1.732 5.943 5.522 3 10 3s8.268 2.943 9.542 7c-1.274 4.057-5.064 7-9.542 7S1.732 14.057.458 10zM14 10a4 4 0 11-8 0 4 4 0 018 0z"
                            clip-rule="evenodd" />
                        </svg>
                        <span>1.2k</span>
                      </div>
                      <div class="flex items-center space-x-1">
                        <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 20 20">
                          <path
                            d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z" />
                        </svg>
                        <span>收藏</span>
                      </div>
                    </div>
                  </div>
                </div>
              </div>

              <!-- Post 2 -->
              <div class="p-6">
                <div class="flex items-start space-x-4">
                  <div
                    class="w-10 h-10 bg-blue-500 rounded-full flex items-center justify-center text-white font-medium">
                    王
                  </div>
                  <div class="flex-1">
                    <div class="flex items-center justify-between mb-2">
                      <div>
                        <h3 class="font-medium text-gray-900">王小美</h3>
                        <p class="text-sm text-gray-500">4小时前</p>
                      </div>
                      <span class="px-3 py-1 bg-blue-100 text-blue-700 rounded-full text-sm">求职招聘</span>
                    </div>
                    <h4 class="text-lg font-semibold text-gray-900 mb-2">【实习机会】某知名互联网公司前端开发实习生招聘</h4>
                    <p class="text-gray-600 mb-4">
                      公司正在招聘前端开发实习生，要求熟悉React、Vue等框架，有项目经验优先。实习期间会有导师指导，表现优秀可转正，感兴趣的同学可以私信我...</p>
                    <div class="flex items-center space-x-6 text-sm text-gray-500">
                      <div class="flex items-center space-x-1">
                        <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 20 20">
                          <path
                            d="M2 10.5a1.5 1.5 0 113 0v6a1.5 1.5 0 01-3 0v-6zM6 10.333v5.43a2 2 0 001.106 1.79l.05.025A4 4 0 008.943 18h5.416a2 2 0 001.962-1.608l1.2-6A2 2 0 0015.56 8H12V4a2 2 0 00-2-2 1 1 0 00-1 1v.667a4 4 0 01-.8 2.4L6.8 7.933a4 4 0 00-.8 2.4z" />
                        </svg>
                        <span>89</span>
                      </div>
                      <div class="flex items-center space-x-1">
                        <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 20 20">
                          <path fill-rule="evenodd"
                            d="M18 13V5a2 2 0 00-2-2H4a2 2 0 00-2 2v8a2 2 0 002 2h3l3 3 3-3h3a2 2 0 002-2zM5 7a1 1 0 011-1h8a1 1 0 110 2H6a1 1 0 01-1-1zm1 3a1 1 0 100 2h3a1 1 0 100-2H6z"
                            clip-rule="evenodd" />
                        </svg>
                        <span>67</span>
                      </div>
                      <div class="flex items-center space-x-1">
                        <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 20 20">
                          <path d="M10 12a2 2 0 100-4 2 2 0 000 4z" />
                          <path fill-rule="evenodd"
                            d="M.458 10C1.732 5.943 5.522 3 10 3s8.268 2.943 9.542 7c-1.274 4.057-5.064 7-9.542 7S1.732 14.057.458 10zM14 10a4 4 0 11-8 0 4 4 0 018 0z"
                            clip-rule="evenodd" />
                        </svg>
                        <span>2.1k</span>
                      </div>
                      <div class="flex items-center space-x-1">
                        <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 20 20">
                          <path
                            d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z" />
                        </svg>
                        <span>收藏</span>
                      </div>
                    </div>
                  </div>
                </div>
              </div>

              <!-- Post 3 -->
              <div class="p-6">
                <div class="flex items-start space-x-4">
                  <div
                    class="w-10 h-10 bg-green-500 rounded-full flex items-center justify-center text-white font-medium">
                    陈
                  </div>
                  <div class="flex-1">
                    <div class="flex items-center justify-between mb-2">
                      <div>
                        <h3 class="font-medium text-gray-900">陈大厨</h3>
                        <p class="text-sm text-gray-500">6小时前</p>
                      </div>
                      <span class="px-3 py-1 bg-green-100 text-green-700 rounded-full text-sm">美食分享</span>
                    </div>
                    <h4 class="text-lg font-semibold text-gray-900 mb-2">宿舍简易版麻辣香锅制作教程，零厨艺也能做！</h4>
                    <p class="text-gray-600 mb-4">作为一个资深吃货，终于研究出了适合宿舍制作的简易麻辣香锅，用料简单，步骤清晰，保证你们都能学会。图片教程已上传，快来试试吧...</p>
                    <div class="flex items-center space-x-6 text-sm text-gray-500">
                      <div class="flex items-center space-x-1">
                        <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 20 20">
                          <path
                            d="M2 10.5a1.5 1.5 0 113 0v6a1.5 1.5 0 01-3 0v-6zM6 10.333v5.43a2 2 0 001.106 1.79l.05.025A4 4 0 008.943 18h5.416a2 2 0 001.962-1.608l1.2-6A2 2 0 0015.56 8H12V4a2 2 0 00-2-2 1 1 0 00-1 1v.667a4 4 0 01-.8 2.4L6.8 7.933a4 4 0 00-.8 2.4z" />
                        </svg>
                        <span>156</span>
                      </div>
                      <div class="flex items-center space-x-1">
                        <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 20 20">
                          <path fill-rule="evenodd"
                            d="M18 13V5a2 2 0 00-2-2H4a2 2 0 00-2 2v8a2 2 0 002 2h3l3 3 3-3h3a2 2 0 002-2zM5 7a1 1 0 011-1h8a1 1 0 110 2H6a1 1 0 01-1-1zm1 3a1 1 0 100 2h3a1 1 0 100-2H6z"
                            clip-rule="evenodd" />
                        </svg>
                        <span>32</span>
                      </div>
                      <div class="flex items-center space-x-1">
                        <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 20 20">
                          <path d="M10 12a2 2 0 100-4 2 2 0 000 4z" />
                          <path fill-rule="evenodd"
                            d="M.458 10C1.732 5.943 5.522 3 10 3s8.268 2.943 9.542 7c-1.274 4.057-5.064 7-9.542 7S1.732 14.057.458 10zM14 10a4 4 0 11-8 0 4 4 0 018 0z"
                            clip-rule="evenodd" />
                        </svg>
                        <span>892</span>
                      </div>
                      <div class="flex items-center space-x-1">
                        <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 20 20">
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

        <!-- Right Sidebar -->
        <div class="lg:col-span-1 space-y-6">
          <!-- Hot Topics -->
          <div class="bg-white rounded-xl shadow-sm p-6">
            <div class="flex items-center space-x-2 mb-4">
              <svg class="w-5 h-5 text-orange-500" fill="currentColor" viewBox="0 0 20 20">
                <path fill-rule="evenodd"
                  d="M12.395 2.553a1 1 0 00-1.45-.385c-.345.23-.614.558-.822.88-.214.33-.403.713-.57 1.116-.334.804-.614 1.768-.84 2.734a31.365 31.365 0 00-.613 3.58 2.64 2.64 0 01-.945-1.067c-.328-.68-.398-1.534-.398-2.654A1 1 0 005.05 6.05 6.981 6.981 0 003 11a7 7 0 1011.95-4.95c-.592-.591-.98-.985-1.348-1.467-.363-.476-.724-1.063-1.207-2.03zM12.12 15.12A3 3 0 017 13s.879.5 2.5.5c0-1 .5-4 1.25-4.5.5 1 .786 1.293 1.371 1.879A2.99 2.99 0 0113 13a2.99 2.99 0 01-.879 2.121z"
                  clip-rule="evenodd" />
              </svg>
              <h3 class="font-semibold text-gray-800">热门话题</h3>
            </div>

            <div class="space-y-3">
              <div class="flex items-center justify-between">
                <span class="text-gray-700">#期末复习攻略#</span>
                <div class="text-right text-sm text-gray-500">
                  <div>1.2万讨论</div>
                  <div>3.5万阅读</div>
                </div>
              </div>

              <div class="flex items-center justify-between">
                <span class="text-gray-700">#校园美食推荐#</span>
                <div class="text-right text-sm text-gray-500">
                  <div>8.9k讨论</div>
                  <div>2.1万阅读</div>
                </div>
              </div>

              <div class="flex items-center justify-between">
                <span class="text-gray-700">#实习求职经验#</span>
                <div class="text-right text-sm text-gray-500">
                  <div>6.7k讨论</div>
                  <div>1.8万阅读</div>
                </div>
              </div>

              <div class="flex items-center justify-between">
                <span class="text-gray-700">#宿舍生活日常#</span>
                <div class="text-right text-sm text-gray-500">
                  <div>5.4k讨论</div>
                  <div>1.5万阅读</div>
                </div>
              </div>

              <div class="flex items-center justify-between">
                <span class="text-gray-700">#社团活动分享#</span>
                <div class="text-right text-sm text-gray-500">
                  <div>4.2k讨论</div>
                  <div>1.2万阅读</div>
                </div>
              </div>
            </div>
          </div>

          <!-- Community Stats -->
          <div class="bg-white rounded-xl shadow-sm p-6">
            <div class="flex items-center space-x-2 mb-4">
              <svg class="w-5 h-5 text-blue-500" fill="currentColor" viewBox="0 0 20 20">
                <path
                  d="M2 11a1 1 0 011-1h2a1 1 0 011 1v5a1 1 0 01-1 1H3a1 1 0 01-1-1v-5zM8 7a1 1 0 011-1h2a1 1 0 011 1v9a1 1 0 01-1 1H9a1 1 0 01-1-1V7zM14 4a1 1 0 011-1h2a1 1 0 011 1v12a1 1 0 01-1 1h-2a1 1 0 01-1-1V4z" />
              </svg>
              <h3 class="font-semibold text-gray-800">社区统计</h3>
            </div>

            <div class="text-center mb-6">
              <div class="text-3xl font-bold text-green-500 mb-1">1,247</div>
              <div class="text-gray-600">在线用户</div>
            </div>

            <div class="grid grid-cols-2 gap-4 text-center">
              <div>
                <div class="text-lg font-semibold text-blue-500">156</div>
                <div class="text-sm text-gray-600">今日新帖</div>
              </div>
              <div>
                <div class="text-lg font-semibold text-purple-500">25.6k</div>
                <div class="text-sm text-gray-600">总用户</div>
              </div>
            </div>
          </div>

          <!-- Active Users -->
          <div class="bg-white rounded-xl shadow-sm p-6">
            <div class="flex items-center space-x-2 mb-4">
              <svg class="w-5 h-5 text-green-500" fill="currentColor" viewBox="0 0 20 20">
                <path d="M13 6a3 3 0 11-6 0 3 3 0 016 0zM18 8a2 2 0 11-4 0 2 2 0 014 0zM14 15a4 4 0 00-8 0v3h8v-3z" />
              </svg>
              <h3 class="font-semibold text-gray-800">活跃用户</h3>
            </div>

            <div class="space-y-3">
              <div class="flex items-center space-x-3">
                <div
                  class="w-8 h-8 bg-purple-500 rounded-full flex items-center justify-center text-white text-sm font-medium">
                  李</div>
                <div class="flex-1">
                  <div class="font-medium text-gray-900">李明同学</div>
                  <div class="text-sm text-gray-500">刚刚活跃</div>
                </div>
                <div class="w-2 h-2 bg-green-400 rounded-full"></div>
              </div>

              <div class="flex items-center space-x-3">
                <div
                  class="w-8 h-8 bg-blue-500 rounded-full flex items-center justify-center text-white text-sm font-medium">
                  王</div>
                <div class="flex-1">
                  <div class="font-medium text-gray-900">王小美</div>
                  <div class="text-sm text-gray-500">5分钟前</div>
                </div>
                <div class="w-2 h-2 bg-green-400 rounded-full"></div>
              </div>

              <div class="flex items-center space-x-3">
                <div
                  class="w-8 h-8 bg-green-500 rounded-full flex items-center justify-center text-white text-sm font-medium">
                  陈</div>
                <div class="flex-1">
                  <div class="font-medium text-gray-900">陈大厨</div>
                  <div class="text-sm text-gray-500">10分钟前</div>
                </div>
                <div class="w-2 h-2 bg-yellow-400 rounded-full"></div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
// Vue 3 Composition API setup
// Add any reactive data or methods here if needed
</script>

<style scoped>
/* Additional custom styles if needed */
</style>
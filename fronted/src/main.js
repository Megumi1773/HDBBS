import {
  createApp
} from 'vue'
import {
  createPinia
} from 'pinia'
import App from './App.vue'
import './moke/user.js'
import router from './router'

import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
const app = createApp(App)
app.use(ElementPlus)
app.use(createPinia())
app.use(router)
app.mount('#app')
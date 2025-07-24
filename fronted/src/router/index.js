import {
  c
} from 'naive-ui'
import {
  createRouter,
  createWebHistory
} from 'vue-router'

const router = createRouter({
  history: createWebHistory(
    import.meta.env.BASE_URL),
  routes: [{
      path: '/',
      name: 'home',
      component: () => import('../components/home.vue')
    },
    {
      path: '/about',
      name: 'about',

    },
  ],
})

export default router
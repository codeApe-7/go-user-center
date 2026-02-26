import { createRouter, createWebHistory } from 'vue-router'
import { useAuthStore } from '../store/auth'

import Login from '../views/Login.vue'
import Register from '../views/Register.vue'
import Profile from '../views/Profile.vue'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: '/', redirect: '/profile' },
    { path: '/login', component: Login },
    { path: '/register', component: Register },
    { path: '/profile', component: Profile }
  ]
})

router.beforeEach((to) => {
  const auth = useAuthStore()
  const publicPaths = ['/login', '/register']
  if (!publicPaths.includes(to.path) && !auth.token) {
    return '/login'
  }
  return true
})

export default router

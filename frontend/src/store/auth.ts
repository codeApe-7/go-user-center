import { defineStore } from 'pinia'
import { api } from '../utils/api'

export interface User {
  id: number
  uuid: string
  email: string
  nickname: string
  avatarUrl: string
  bio: string
  createdAt: string
  updatedAt: string
}

export const useAuthStore = defineStore('auth', {
  state: () => ({
    token: localStorage.getItem('token') as string | null,
    user: null as User | null
  }),
  actions: {
    setToken(token: string) {
      this.token = token
      localStorage.setItem('token', token)
    },
    clear() {
      this.token = null
      this.user = null
      localStorage.removeItem('token')
    },
    async login(email: string, password: string) {
      const r = await api.post('/auth/login', { email, password })
      this.setToken(r.data.accessToken)
      this.user = r.data.user
    },
    async register(email: string, password: string, nickname: string) {
      const r = await api.post('/auth/register', { email, password, nickname })
      this.user = r.data.user
    },
    async fetchMe() {
      const r = await api.get('/me')
      this.user = r.data.user
    }
  }
})

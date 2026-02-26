import axios from 'axios'

const baseURL = import.meta.env.VITE_API_BASE || 'http://127.0.0.1:18601/api'

export const api = axios.create({
  baseURL,
  timeout: 10000
})

api.interceptors.request.use((config) => {
  const token = localStorage.getItem('token')
  if (token) {
    config.headers = config.headers || {}
    config.headers.Authorization = `Bearer ${token}`
  }
  return config
})

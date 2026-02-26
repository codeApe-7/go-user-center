<template>
  <el-container class="app">
    <el-header height="56px" class="header">
      <div class="brand">Go User Center</div>
      <div class="spacer" />
      <template v-if="auth.token">
        <el-button text @click="go('/profile')">个人中心</el-button>
        <el-button type="danger" plain @click="logout">退出</el-button>
      </template>
      <template v-else>
        <el-button text @click="go('/login')">登录</el-button>
        <el-button text @click="go('/register')">注册</el-button>
      </template>
    </el-header>

    <el-main class="main">
      <router-view />
    </el-main>
  </el-container>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router'
import { useAuthStore } from './store/auth'

const router = useRouter()
const auth = useAuthStore()

function go(path: string) {
  router.push(path)
}

function logout() {
  auth.clear()
  router.push('/login')
}
</script>

<style scoped>
.app { min-height: 100vh; }
.header {
  display: flex;
  align-items: center;
  border-bottom: 1px solid var(--el-border-color);
}
.brand { font-weight: 700; letter-spacing: 0.5px; }
.spacer { flex: 1; }
.main { max-width: 980px; margin: 0 auto; width: 100%; }
</style>

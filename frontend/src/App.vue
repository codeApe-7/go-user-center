<template>
  <el-container class="app-shell">
    <el-header height="64px" class="topbar">
      <div class="brand" @click="go('/')">
        <div class="brand-dot" />
        <div>
          <div class="brand-name">Go User Center</div>
          <div class="brand-sub">Learning Edition</div>
        </div>
      </div>

      <div class="spacer" />

      <template v-if="auth.token">
        <el-button text @click="go('/profile')">个人中心</el-button>
        <el-button type="danger" plain @click="logout">退出</el-button>
      </template>
      <template v-else>
        <el-button text @click="go('/login')">登录</el-button>
        <el-button type="primary" plain @click="go('/register')">注册</el-button>
      </template>
    </el-header>

    <el-main class="main-wrap">
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
.app-shell {
  min-height: 100vh;
}

.topbar {
  position: sticky;
  top: 0;
  z-index: 20;
  display: flex;
  align-items: center;
  padding: 0 20px;
  border-bottom: 1px solid var(--line);
  backdrop-filter: blur(10px);
  background: rgba(255, 255, 255, 0.78);
}

.brand {
  display: flex;
  gap: 10px;
  align-items: center;
  cursor: pointer;
}

.brand-dot {
  width: 12px;
  height: 12px;
  border-radius: 999px;
  background: linear-gradient(135deg, var(--brand), var(--brand-2));
  box-shadow: 0 0 0 6px rgba(79, 70, 229, 0.12);
}

.brand-name {
  font-size: 15px;
  font-weight: 700;
}

.brand-sub {
  font-size: 12px;
  color: var(--text-sub);
  line-height: 1;
}

.spacer {
  flex: 1;
}

.main-wrap {
  max-width: 1080px;
  margin: 0 auto;
  width: 100%;
  padding: 24px 16px 36px;
}
</style>

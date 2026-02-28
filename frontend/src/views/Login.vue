<template>
  <el-row justify="center" class="auth-page">
    <el-col :xs="24" :sm="18" :md="12" :lg="10" :xl="8">
      <el-card class="auth-card" shadow="never">
        <template #header>
          <div class="title-wrap">
            <div class="title">欢迎回来</div>
            <div class="subtitle">登录你的账号继续使用用户中心</div>
          </div>
        </template>

        <el-form :model="form" label-position="top" @submit.prevent>
          <el-form-item label="邮箱">
            <el-input v-model="form.email" placeholder="you@example.com" size="large" />
          </el-form-item>
          <el-form-item label="密码">
            <el-input v-model="form.password" type="password" show-password size="large" />
          </el-form-item>

          <el-button type="primary" size="large" style="width: 100%" :loading="loading" @click="onLogin">
            登录
          </el-button>

          <div class="hint">
            没有账号？<el-link type="primary" @click="router.push('/register')">去注册</el-link>
          </div>
        </el-form>
      </el-card>
    </el-col>
  </el-row>
</template>

<script setup lang="ts">
import { reactive, ref } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { useAuthStore } from '../store/auth'

const router = useRouter()
const auth = useAuthStore()
const loading = ref(false)

const form = reactive({
  email: '',
  password: ''
})

async function onLogin() {
  loading.value = true
  try {
    await auth.login(form.email, form.password)
    ElMessage.success('登录成功')
    router.push('/profile')
  } catch (e: any) {
    ElMessage.error(e?.response?.data?.error || '登录失败')
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
.auth-page {
  padding-top: 4vh;
}

.auth-card {
  border: 1px solid rgba(99, 102, 241, 0.18);
  border-radius: 16px;
  background: var(--card);
  box-shadow: 0 16px 40px rgba(15, 23, 42, 0.08);
}

.title-wrap {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.title {
  font-weight: 700;
  font-size: 20px;
}

.subtitle {
  color: var(--text-sub);
  font-size: 13px;
}

.hint {
  margin-top: 12px;
  text-align: center;
  color: var(--text-sub);
  font-size: 13px;
}
</style>

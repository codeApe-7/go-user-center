<template>
  <el-row justify="center" class="auth-page">
    <el-col :xs="24" :sm="18" :md="12" :lg="10" :xl="8">
      <el-card class="auth-card" shadow="never">
        <template #header>
          <div class="title-wrap">
            <div class="title">创建账号</div>
            <div class="subtitle">开启你的 Go 用户中心学习之旅</div>
          </div>
        </template>

        <el-form :model="form" label-position="top" @submit.prevent>
          <el-form-item label="邮箱">
            <el-input v-model="form.email" placeholder="you@example.com" size="large" />
          </el-form-item>
          <el-form-item label="昵称">
            <el-input v-model="form.nickname" placeholder="比如：新之助" size="large" />
          </el-form-item>
          <el-form-item label="密码">
            <el-input v-model="form.password" type="password" show-password size="large" />
          </el-form-item>

          <el-button type="primary" size="large" style="width: 100%" :loading="loading" @click="onRegister">
            创建账号
          </el-button>

          <div class="hint">
            已有账号？<el-link type="primary" @click="router.push('/login')">去登录</el-link>
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
  password: '',
  nickname: ''
})

async function onRegister() {
  loading.value = true
  try {
    await auth.register(form.email, form.password, form.nickname)
    ElMessage.success('注册成功，请登录')
    router.push('/login')
  } catch (e: any) {
    ElMessage.error(e?.response?.data?.error || '注册失败')
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

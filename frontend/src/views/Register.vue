<template>
  <el-row justify="center">
    <el-col :xs="24" :sm="16" :md="10" :lg="8">
      <el-card>
        <template #header>
          <div class="title">注册</div>
        </template>

        <el-form :model="form" label-position="top" @submit.prevent>
          <el-form-item label="邮箱">
            <el-input v-model="form.email" placeholder="you@example.com" />
          </el-form-item>
          <el-form-item label="昵称">
            <el-input v-model="form.nickname" placeholder="比如：新之助" />
          </el-form-item>
          <el-form-item label="密码">
            <el-input v-model="form.password" type="password" show-password />
          </el-form-item>

          <el-button type="primary" style="width: 100%" :loading="loading" @click="onRegister">
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
.title { font-weight: 700; }
.hint { margin-top: 12px; color: var(--el-text-color-secondary); font-size: 13px; }
</style>

<template>
  <el-row justify="center">
    <el-col :xs="24" :sm="22" :md="16" :lg="12">
      <el-card>
        <template #header>
          <div class="title">个人中心</div>
        </template>

        <div v-if="auth.user" class="profile">
          <el-avatar :size="64" :src="auth.user.avatarUrl || undefined">
            {{ auth.user.nickname?.slice(0, 1) || 'U' }}
          </el-avatar>

          <div class="info">
            <div class="line"><b>邮箱：</b>{{ auth.user.email }}</div>
            <div class="line"><b>UUID：</b>{{ auth.user.uuid }}</div>
          </div>
        </div>

        <el-divider />

        <el-form :model="form" label-position="top">
          <el-form-item label="昵称">
            <el-input v-model="form.nickname" />
          </el-form-item>
          <el-form-item label="头像 URL">
            <el-input v-model="form.avatarUrl" placeholder="https://..." />
          </el-form-item>
          <el-form-item label="简介">
            <el-input v-model="form.bio" type="textarea" :rows="3" />
          </el-form-item>

          <el-button type="primary" :loading="saving" @click="save">保存资料</el-button>
          <el-button @click="refresh">刷新</el-button>
        </el-form>

        <el-divider />

        <el-form :model="pw" label-position="top">
          <el-form-item label="旧密码">
            <el-input v-model="pw.oldPassword" type="password" show-password />
          </el-form-item>
          <el-form-item label="新密码">
            <el-input v-model="pw.newPassword" type="password" show-password />
          </el-form-item>
          <el-button type="warning" :loading="pwSaving" @click="changePassword">修改密码</el-button>
        </el-form>
      </el-card>
    </el-col>
  </el-row>
</template>

<script setup lang="ts">
import { onMounted, reactive, ref } from 'vue'
import { ElMessage } from 'element-plus'
import { useAuthStore } from '../store/auth'
import { api } from '../utils/api'

const auth = useAuthStore()
const saving = ref(false)
const pwSaving = ref(false)

const form = reactive({
  nickname: '',
  avatarUrl: '',
  bio: ''
})

const pw = reactive({
  oldPassword: '',
  newPassword: ''
})

async function refresh() {
  try {
    await auth.fetchMe()
    if (auth.user) {
      form.nickname = auth.user.nickname || ''
      form.avatarUrl = auth.user.avatarUrl || ''
      form.bio = auth.user.bio || ''
    }
  } catch (e: any) {
    ElMessage.error(e?.response?.data?.error || '加载失败')
  }
}

async function save() {
  saving.value = true
  try {
    const r = await api.put('/me', form)
    auth.user = r.data.user
    ElMessage.success('保存成功')
  } catch (e: any) {
    ElMessage.error(e?.response?.data?.error || '保存失败')
  } finally {
    saving.value = false
  }
}

async function changePassword() {
  pwSaving.value = true
  try {
    await api.put('/me/password', pw)
    pw.oldPassword = ''
    pw.newPassword = ''
    ElMessage.success('密码已更新')
  } catch (e: any) {
    ElMessage.error(e?.response?.data?.error || '修改失败')
  } finally {
    pwSaving.value = false
  }
}

onMounted(refresh)
</script>

<style scoped>
.title { font-weight: 700; }
.profile { display: flex; gap: 16px; align-items: center; }
.info .line { margin: 4px 0; color: var(--el-text-color-regular); }
</style>

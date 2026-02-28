<template>
  <el-row justify="center">
    <el-col :xs="24" :sm="22" :md="18" :lg="14">
      <el-card class="profile-card" shadow="never">
        <template #header>
          <div class="header-wrap">
            <div>
              <div class="title">个人中心</div>
              <div class="subtitle">管理你的个人资料与账号安全</div>
            </div>
            <el-tag type="primary" effect="plain">已登录</el-tag>
          </div>
        </template>

        <div v-if="auth.user" class="profile-overview">
          <el-avatar :size="72" :src="auth.user.avatarUrl || undefined" class="avatar">
            {{ auth.user.nickname?.slice(0, 1) || 'U' }}
          </el-avatar>

          <div class="meta">
            <div class="name">{{ auth.user.nickname || '未设置昵称' }}</div>
            <div class="line"><b>邮箱：</b>{{ auth.user.email }}</div>
            <div class="line"><b>UUID：</b>{{ auth.user.uuid }}</div>
          </div>
        </div>

        <el-divider content-position="left">资料设置</el-divider>

        <el-form :model="form" label-position="top">
          <el-row :gutter="16">
            <el-col :xs="24" :md="12">
              <el-form-item label="昵称">
                <el-input v-model="form.nickname" size="large" />
              </el-form-item>
            </el-col>
            <el-col :xs="24" :md="12">
              <el-form-item label="头像 URL">
                <el-input v-model="form.avatarUrl" placeholder="https://..." size="large" />
              </el-form-item>
            </el-col>
          </el-row>

          <el-form-item label="简介">
            <el-input v-model="form.bio" type="textarea" :rows="3" />
          </el-form-item>

          <div class="actions">
            <el-button type="primary" :loading="saving" @click="save">保存资料</el-button>
            <el-button @click="refresh">刷新</el-button>
          </div>
        </el-form>

        <el-divider content-position="left">账号安全</el-divider>

        <el-form :model="pw" label-position="top">
          <el-row :gutter="16">
            <el-col :xs="24" :md="12">
              <el-form-item label="旧密码">
                <el-input v-model="pw.oldPassword" type="password" show-password size="large" />
              </el-form-item>
            </el-col>
            <el-col :xs="24" :md="12">
              <el-form-item label="新密码">
                <el-input v-model="pw.newPassword" type="password" show-password size="large" />
              </el-form-item>
            </el-col>
          </el-row>

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
.profile-card {
  border: 1px solid rgba(99, 102, 241, 0.18);
  border-radius: 16px;
  background: var(--card);
  box-shadow: 0 16px 40px rgba(15, 23, 42, 0.08);
}

.header-wrap {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 12px;
}

.title {
  font-size: 20px;
  font-weight: 700;
}

.subtitle {
  font-size: 13px;
  color: var(--text-sub);
}

.profile-overview {
  display: flex;
  align-items: center;
  gap: 16px;
}

.avatar {
  border: 2px solid rgba(79, 70, 229, 0.25);
}

.meta .name {
  font-size: 18px;
  font-weight: 700;
  margin-bottom: 4px;
}

.meta .line {
  margin: 3px 0;
  color: var(--text-sub);
  font-size: 14px;
}

.actions {
  display: flex;
  gap: 10px;
}
</style>

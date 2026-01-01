<template>
  <div class="settings-card">
    <h2>设置（Settings）</h2>

    <!-- Config -->
    <section>
      <h3>配置文件</h3>
      <div class="row">
        <span class="label">当前配置</span>
        <span class="value">{{ configPath || '未选择' }}</span>
        <button @click="chooseConfigFile">选择文件</button>
      </div>
    </section>

    <!-- Appearance -->
    <section>
      <h3>外观</h3>

      <div class="row">
        <span class="label">主题</span>
        <select v-model="theme">
          <option value="light">浅色</option>
          <option value="dark">深色</option>
          <option value="system">跟随系统</option>
        </select>
      </div>

      <div class="row">
        <span class="label">字号</span>
        <select v-model="fontSize">
          <option value="small">小</option>
          <option value="normal">中</option>
          <option value="large">大</option>
        </select>
      </div>
    </section>
  </div>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue'
import {
  GetCurrentConfigPath,
  SelectAndLoadConfig
} from '@wailsjs/go/main/App'


const configPath = ref('')
const theme = ref('system')
const fontSize = ref('normal')

onMounted(async () => {
  configPath.value = await GetCurrentConfigPath()
})

async function chooseConfigFile() {
  const result = await SelectAndLoadConfig()
  if (result) {
    configPath.value = result
  }
}

/* 外观设置后面可以统一写入 localStorage */
watch([theme, fontSize], () => {
  localStorage.setItem('theme', theme.value)
  localStorage.setItem('fontSize', fontSize.value)
})
</script>

<style scoped>
.settings-card {
  background: #fff;
  border-radius: 8px;
  padding: 16px 20px;
  box-shadow: 0 4px 12px rgba(0,0,0,.06);
}

section {
  margin-top: 16px;
}

.row {
  display: flex;
  align-items: center;
  margin: 8px 0;
}

.label {
  width: 100px;
  color: #475569;
}

.value {
  flex: 1;
  font-size: 13px;
  color: #334155;
}

button {
  padding: 4px 10px;
  font-size: 12px;
}
</style>

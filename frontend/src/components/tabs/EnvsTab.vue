<template>
  <div class="envs-card">
    <h2>环境变量（Envs）</h2>

    <div
      v-for="env in envs"
      :key="env.name"
      class="env-row"
    >
      <!-- 状态 -->
      <span
        class="status"
        :class="statusClass(env)"
      >
        {{ statusIcon(env) }}
      </span>

      <!-- 名称 -->
      <div class="env-main">
        <div class="env-name">
          {{ env.name }}
          <span class="env-scope">({{ env.scope }})</span>
        </div>

        <!-- 描述 -->
        <div class="env-desc">
          <template v-if="env.correct">
            {{ env.value || '已正确配置' }}
          </template>

          <template v-else-if="env.exists && env.missing?.length">
            缺少：
            <span
              v-for="m in env.missing"
              :key="m"
              class="missing-item"
            >
              {{ m }}
            </span>
          </template>

          <template v-else>
            未设置
          </template>
        </div>
      </div>

      <!-- 操作 -->
      <div class="env-action">
        <button
          v-if="!env.exists"
          @click="onSet(env)"
        >
          设置
        </button>

        <button
          v-else-if="env.missing?.length"
          @click="onAppend(env)"
        >
          补全
        </button>

        <span
          v-else
          class="done-text"
        >
          已完成
        </span>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import {
  GetEnvStatus,
  SetEnv,
  AppendEnv
} from '@wailsjs/go/main/App'

const envs = ref([])

async function load() {
  envs.value = await GetEnvStatus()
}

function statusIcon(env) {
  if (!env.exists) return '✕'
  if (!env.correct) return '!'
  return '✓'
}

function statusClass(env) {
  if (!env.exists) return 'status-missing'
  if (!env.correct) return 'status-warn'
  return 'status-ok'
}

async function onSet(env) {
  await SetEnv(env.name, env.value, env.scope)
  await load()
}

async function onAppend(env) {
  await AppendEnv(env.name, env.missing, env.scope)
  await load()
}

onMounted(load)
</script>

<style scoped>
.envs-card {
  background: #ffffff;
  border-radius: 8px;
  padding: 16px 20px;
  box-shadow: 0 4px 12px rgba(0,0,0,.06);
}

.env-row {
  display: flex;
  align-items: center;
  padding: 10px 6px;
  border-radius: 6px;
}

.env-row:hover {
  background: #f0f4ff;
}

.status {
  width: 24px;
  text-align: center;
  font-size: 16px;
  margin-right: 8px;
}

.status-ok {
  color: #22c55e;
}

.status-warn {
  color: #eab308;
}

.status-missing {
  color: #94a3b8;
}

.env-main {
  flex: 1;
}

.env-name {
  font-weight: 600;
}

.env-scope {
  font-size: 12px;
  color: #64748b;
  margin-left: 6px;
}

.env-desc {
  font-size: 13px;
  color: #475569;
  margin-top: 2px;
}

.missing-item {
  background: #e0e7ff;
  border-radius: 4px;
  padding: 2px 6px;
  margin-right: 4px;
  font-size: 12px;
}

.env-action {
  margin-left: 12px;
}

button {
  padding: 4px 10px;
  font-size: 12px;
  border-radius: 4px;
  border: 1px solid #c7d2fe;
  background: #eef2ff;
  cursor: pointer;
}

button:hover {
  background: #e0e7ff;
}

.done-text {
  font-size: 12px;
  color: #94a3b8;
}
</style>

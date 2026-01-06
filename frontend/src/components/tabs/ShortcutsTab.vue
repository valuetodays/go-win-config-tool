<template>
    <div class="tab-content">
        <div v-for="s in shortcuts" :key="s.name" class="env-row">
            <!-- 状态 -->
            <span class="status" :class="statusClass(s)">
                {{ statusIcon(s) }}
            </span>

            <!-- 名称 -->
            <div class="env-main">
                <div class="env-name">
                    {{ s.name }}
                    <!-- <span class="env-scope">({{ s.scope }})</span> -->
                    <!-- {{ s }} -->
                </div>

                <!-- 描述 -->
                <div class="env-desc">
                    <template v-if="s.correct">
                        {{ s.source }} -> {{ s.target }}
                    </template>
                    <template v-else>
                        未设置
                    </template>
                </div>
            </div>

            <!-- 操作 -->
            <div class="env-action">
                <button v-if="!s.exists" @click="onSet(s)">
                    设置
                </button>
                <span v-else class="done-text">
                    已完成
                </span>
            </div>
        </div>

    </div>
</template>

<script setup>
import { ref, onMounted } from "vue"
import { GetShortcuts, GenerateShortcut } from "@wailsjs/go/main/App"

const shortcuts = ref([])

const load = async () => {
    shortcuts.value = await GetShortcuts()
}

const fix = async (name) => {
    await GenerateShortcut(name)
    await load()
}

onMounted(load)

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
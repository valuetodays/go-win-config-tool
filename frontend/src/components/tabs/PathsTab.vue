<template>
  <div class="paths-card">
    <h3>目录（Paths）</h3>

    <ul class="list">
      <li v-for="item in paths" :key="item.path">
        <StatusIcon :exists="item.exists" />

        <span class="path">{{ item.path }}</span>

        <button
          v-if="!item.exists"
          @click="create(item.path)"
        >
          创建
        </button>
      </li>
    </ul>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted  } from 'vue'
import StatusIcon from '@/components/common/StatusIcon.vue'
import {
  GetPathsStatus,
  CreatePath
} from '@wailsjs/go/main/App'
import { registerReload, unregisterReload } from '@/services/configBus'

const paths = ref([])

const load = async () => {
  paths.value = await GetPathsStatus()
}

const create = async (path) => {
  await CreatePath(path)
}

onMounted(() => {
  load()
  registerReload(load)
})

onUnmounted(() => {
  unregisterReload(load)
})
</script>

<style scoped>
.paths-card {
  background: #ffffff;
  border-radius: 8px;
  padding: 16px 20px;
  box-shadow: 0 4px 12px rgba(0,0,0,.06);
}

.list {
  list-style: none;
  padding: 0;
}

li {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-bottom: 8px;
}

.path {
  flex: 1;
  font-family: monospace;
}
</style>

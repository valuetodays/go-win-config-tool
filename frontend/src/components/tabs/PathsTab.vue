<template>
  <div class="card">
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
import { ref, onMounted } from 'vue'
import StatusIcon from '../common/StatusIcon.vue'
import {
  GetPathsStatus,
  CreatePath
} from '@wailsjs/go/main/App'

const paths = ref([])

const load = async () => {
  paths.value = await GetPathsStatus()
}

const create = async (path) => {
  await CreatePath(path)
  await load()
}

onMounted(load)
</script>

<style scoped>
.card {
  padding: 16px;
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

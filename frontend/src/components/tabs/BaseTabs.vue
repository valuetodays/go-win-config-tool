<template>
  <div>
    <div class="tabs">
      <button
        v-for="t in tabs"
        :key="t.name"
        :class="{ active: t.name === active }"
        @click="active = t.name"
      >
        {{ t.label }}
      </button>
    </div>

    <component :is="active" />
  </div>
</template>

<script setup>
import { ref } from 'vue'
import PathsTab from './PathsTab.vue'

const tabs = [
  { name: PathsTab, label: '📁 Paths' },
  // 以后直接加：
  // { name: FilesTab, label: '📄 Files' }
]

const active = ref(PathsTab)
</script>

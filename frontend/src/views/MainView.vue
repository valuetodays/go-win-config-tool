<template>
  <div class="main-view app-root">
    <!-- Tabs Header -->
    <div class="tabs">
      <div
        v-for="tab in tabs"
        :key="tab.key"
        class="tab"
        :class="{ active: currentTab === tab.key }"
        @click="currentTab = tab.key"
      >
        {{ tab.label }}
      </div>
    </div>

    <!-- Tabs Content -->
    <div class="tab-content">
      <component :is="currentComponent" />
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'

import PathsTab from '@/components/tabs/PathsTab.vue'
import ShortcutsTab from '@/components/tabs/ShortcutsTab.vue'
import EnvsTab from '@/components/tabs/EnvsTab.vue'
import SoftwaresTab from '@/components/tabs/SoftwaresTab.vue'
import SettingsTab from '@/components/tabs/SettingsTab.vue'

const tabs = [
  { key: 'paths', label: 'Paths', component: PathsTab },
  { key: "shortcuts", label: "Shortcuts", component: ShortcutsTab },
  { key: 'envs', label: 'Envs', component: EnvsTab },
  { key: 'softwares', label: 'Softwares', component: SoftwaresTab },
  { key: 'settings', label: 'Settings', component: SettingsTab },
]

const currentTab = ref('paths')

const currentComponent = computed(() => {
  return tabs.find(t => t.key === currentTab.value)?.component
})
</script>

<style scoped>
.main-view {
  padding: 20px;
}

.tabs {
  display: flex;
  border-bottom: 1px solid #e5e7eb;
  margin-bottom: 16px;
}

.tab {
  padding: 8px 14px;
  cursor: pointer;
  font-size: 14px;
  color: #64748b;
}

.tab.active {
  color: #1e293b;
  font-weight: 600;
  border-bottom: 2px solid #6366f1;
}



.app-root {
  min-height: 100vh;
  background: #f7f9fc;
}
</style>
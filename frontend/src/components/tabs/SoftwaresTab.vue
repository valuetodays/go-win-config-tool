<template>
  <div class="softwares-card">
    <h2>软件（Softwares）</h2>

    <div
      v-for="sw in softwares"
      :key="sw.name"
      class="software-block"
    >
      <div class="software-header">
        <span
          class="status"
          :class="sw.exists ? 'ok' : 'missing'"
        >
          {{ sw.exists ? '✓' : '✕' }}
        </span>

        <div class="software-title">
          {{ sw.name }}
          <div class="software-root">
            {{ sw.rootDir }}
          </div>
        </div>
      </div>

      <div class="software-items">
        <div
          v-for="item in sw.items"
          :key="item.path"
          class="software-item"
        >
          <span
            class="item-status"
            :class="item.exists ? 'ok' : 'missing'"
          >
            {{ item.exists ? '✓' : '✕' }}
          </span>

          <span class="item-type">
            {{ item.type }}
          </span>

          <span class="item-path">
            {{ item.path }}
          </span>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { GetSoftwareStatus } from '@wailsjs/go/main/App'

const softwares = ref([])

async function load() {
  softwares.value = await GetSoftwareStatus()
}

onMounted(load)
</script>

<style scoped>
.softwares-card {
  background: #fff;
  border-radius: 8px;
  padding: 16px 20px;
  box-shadow: 0 4px 12px rgba(0,0,0,.06);
}

.software-block {
  margin-bottom: 20px;
  padding-bottom: 12px;
  border-bottom: 1px dashed #e5e7eb;
}

.software-header {
  display: flex;
  align-items: center;
  margin-bottom: 8px;
}

.status {
  width: 24px;
  text-align: center;
  font-size: 16px;
}

.software-title {
  font-weight: 600;
}

.software-root {
  font-size: 12px;
  color: #64748b;
}

.software-items {
  padding-left: 24px;
}

.software-item {
  display: flex;
  align-items: center;
  font-size: 13px;
  margin: 4px 0;
}

.item-status {
  width: 18px;
  text-align: center;
}

.item-type {
  width: 40px;
  color: #6366f1;
}

.item-path {
  color: #334155;
}

.ok {
  color: #22c55e;
}

.missing {
  color: #94a3b8;
}
</style>

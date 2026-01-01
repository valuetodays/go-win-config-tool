import { createApp } from 'vue'
import App from './App.vue'
import { EventsOn } from '@wailsjs/runtime/runtime'
import { reloadAllTabs } from '@/services/configBus'

createApp(App).mount('#app')

EventsOn('config:changed', () => {
  console.log('[config] changed') // 👈 一定加这一行
  reloadAllTabs()
})
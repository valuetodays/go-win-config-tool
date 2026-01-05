<template>
  <span :class="exists ? 'ok' : 'missing'">
    {{ exists ? '✔' : '✖' }}
  </span>
</template>

<script setup>
defineProps({
  exists: Boolean
})
</script>

<style scoped>
.ok {
  color: #3cb371;
}
.missing {
  color: #a52929;
}
</style>

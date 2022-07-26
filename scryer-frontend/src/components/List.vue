<script setup>
  import { storeToRefs } from 'pinia'

  import { useDeviceStore } from '../stores/device'

  const { devices, loading, error } = storeToRefs(useDeviceStore())
</script>

<template>
  <q-list bordered padding separator v-if="devices" v-for="device in devices" :key="device.device_id">
    <q-item clickable v-ripple>
      <q-item-section>{{ device.display_name }}</q-item-section>
    </q-item>
  </q-list>
</template>

<style scoped>
</style>

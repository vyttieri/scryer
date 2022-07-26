<script setup>
  import { storeToRefs } from 'pinia'

  import { useDeviceStore } from '../stores/device'

  const { devices, loading, error } = storeToRefs(useDeviceStore())
</script>

<template>
  <q-list bordered padding separator v-if="devices" v-for="device in devices" :key="device.device_id">
    <q-item>
      <q-item-section>{{ device.display_name }}</q-item-section>
      <q-item-section>{{ device.active_state }}</q-item-section>
      <q-item-section>{{ device.latest_device_point.device_state.drive_status }}</q-item-section>
    </q-item>
  </q-list>
</template>

<style scoped>
</style>

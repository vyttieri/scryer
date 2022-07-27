<script setup>
  import { storeToRefs } from 'pinia'

  import { useDeviceStore } from '../stores/device'
  import { usePreferencesStore } from '../stores/preferences'

  const { devices, loading, error } = storeToRefs(useDeviceStore())
  const { getDeviceVisibility, setDeviceVisibility, setCenter } = usePreferencesStore()


</script>

<template>
  <q-list bordered padding separator v-if="devices" v-for="device in devices" :key="device.device_id">
    <q-item>
      <q-item-section avatar clickable @click="setDeviceVisibility(device)">
        <q-icon clickable v-if="getDeviceVisibility(device)" color="primary" name="visibility" />
        <q-icon clickable v-if="!getDeviceVisibility(device)" color="primary" name="visibility_off" />
      </q-item-section>
      <q-item-section avatar>
        <q-icon clickable color="primary" name="my_location" @click="setCenter(device)"/>
      </q-item-section>
      <q-item-section>{{ device.display_name }}</q-item-section>
      <q-item-section>{{ device.active_state }}</q-item-section>
      <q-item-section>{{ device.latest_accurate_device_point.device_state.drive_status }}</q-item-section>
    </q-item>
  </q-list>
</template>

<style scoped>
</style>

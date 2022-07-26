<script setup>
import { storeToRefs } from 'pinia'

import { useDeviceStore } from './stores/device'

import List from './components/List.vue'


const { devices, loading, error } = storeToRefs(useDeviceStore())
const { fetchDevices }  = useDeviceStore()

fetchDevices()
</script>

<template>
  <header>
  </header>
  <main>
    <p v-if="loading">Loading device data...</p>
    <p v-if="error">{{ error.message }}</p>
    <p v-if="devices" v-for="device in devices" :key="device.id">
      <p>{{ device }}</p>
    </p>
    <div id="q-app">
      <div class="q-pa-md">
        <List />
      </div>
    </div>
  </main>
</template>

<style scoped>
</style>

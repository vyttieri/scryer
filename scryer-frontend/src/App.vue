<script setup>
import { storeToRefs } from 'pinia'

import { useDeviceStore } from './stores/device'

import List from './components/List.vue'
import Map from './components/Map.vue'

const { loading, error } = storeToRefs(useDeviceStore())
const { fetchDevices }  = useDeviceStore()

fetchDevices()
</script>

<template>
  <header>
  </header>
  <main>
    <p v-if="loading">Loading device data...</p>
    <p v-if="error">{{ error.message }}</p>
    <div id="q-app">
      <div class="q-pa-md">
        <div class="row">
          <div class="col-4">
            <List />
          </div>
          <div class="col-8">
            <Map />
          </div>
        </div>
      </div>
    </div>
  </main>
</template>

<style scoped>
</style>

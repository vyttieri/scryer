<script setup>
import { storeToRefs } from 'pinia'
import { onMounted } from 'vue'

import { useDeviceStore } from './stores/device'

import Table from './components/Table.vue'
import Map from './components/Map.vue'

const { loading, error } = storeToRefs(useDeviceStore())
const { fetchDevices }  = useDeviceStore()


onMounted(() => {
  fetchDevices()

  // 60,000 ms = 1 minute; set interval slightly longer than server-side cache timeout ()
  setInterval(fetchDevices, 61000)
})
</script>

<template>
  <header>
  </header>
  <main>
    <p v-if="loading">Loading device data...</p>
    <p v-if="error">{{ error.message }}</p>
    <div id="q-app">
      <q-card>
        <div class="row main">
          <div class="col-4">
            <Table />
          </div>
          <div class="col-8">
            <Map />
          </div>
        </div>
      </q-card>
    </div>
  </main>
</template>

<style scoped>
  .q-card { height: 100vh; }
  .main { height: 100%; }
</style>

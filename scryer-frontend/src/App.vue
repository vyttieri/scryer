<script setup>
import { storeToRefs } from 'pinia'
import { onMounted } from 'vue'

import { useDeviceStore } from './stores/device'

import Header from './components/Header.vue'
import Map from './components/Map.vue'
import Table from './components/Table.vue'

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
    <q-layout view="hHh lpR fFf">
      <Header />
      <q-drawer show-if-above v-model="leftDrawerOpen" side="left" bordered>
        <Table />
      </q-drawer>
      <q-page-container style="width:  100vw; height: 100vh;">
        <Map />
      </q-page-container>
    </q-layout>
  </main>
</template>

<style scoped>
  .q-card { height: 100vh; }
</style>

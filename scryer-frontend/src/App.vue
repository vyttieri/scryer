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
      <q-drawer show-if-above v-model="leftDrawerOpen" side="left" bordered :width="350">
        <Table :center="center" @set-center="setCenter" />
      </q-drawer>
      <q-page-container style="width:  100vw; height: 100vh;">
        <Map :center="center"  @set-center="setCenter" />
      </q-page-container>
    </q-layout>
  </main>
</template>

<script>
const californiaCenter = { lat: 36.778300, lng: -119.417900 }
/* I opted to keep "center" out of the data stores, as it doesn't deal with
 API or backend data. Also, seemed like a good way to learn a little bit
 about props and data management using plain Vue.

If the app were to grow, it'd probably make sense to eventually move such
"client-side settings" to a data store as well.
*/
export default {
  data() {
    return {
      center: californiaCenter,
    }
  },
  methods: {
    setCenter(position) {
      this.center = position
    }
  },
}
</script>

<script setup>
import { storeToRefs } from 'pinia'
import { onMounted, ref } from 'vue'

import { useDeviceStore } from '@/stores/device'

import Header from '@/components/Header.vue'
import Map from '@/components/Map.vue'
import Table from '@/components/Table.vue'

const { fetchDevices }  = useDeviceStore()

onMounted(() => {
  fetchDevices()

  // 60,000 ms = 1 minute; set interval slightly longer than server-side cache timeout ()
  setInterval(fetchDevices, 61000)
})

/* I opted to keep "center" and "openedMarkerId" out of the data stores, as it doesn't deal with
 API or backend data. Also, seemed like a good way to learn a little bit
 about props and data management using plain Vue.

If the app were to grow, it'd probably make sense to eventually move such
"client-side settings" to a data store as well.
*/
const californiaCenter = { lat: 36.778300, lng: -119.417900 }
let center = ref(californiaCenter)
const setCenter = position => center.value = position

let openedMarkerId = ref(null)
const setOpenMarker = deviceId => openedMarkerId.value = deviceId
</script>

<template>
  <header>
  </header>
  <main>
    <q-layout view="hHh lpR fFf">
      <Header />
      <q-drawer show-if-above v-model="leftDrawerOpen" side="left" bordered :width="350">
        <Table
          :center="center"
          @set-center="setCenter"
          :openedMarkerId="openedMarkerId"
          @set-open-marker="setOpenMarker"
        />
      </q-drawer>
      <q-page-container style="width:  100vw; height: 100vh;">
        <Map
          :center="center"
          @set-center="setCenter"
          :openedMarkerId="openedMarkerId"
          @set-open-marker="setOpenMarker"
        />
      </q-page-container>
    </q-layout>
  </main>
</template>

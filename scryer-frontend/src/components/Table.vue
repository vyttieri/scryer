<script setup>
import draggable from 'vuedraggable'

import { storeToRefs } from 'pinia'

import { useDeviceStore } from '@/stores/device'
import { useDevicePreferenceStore } from '@/stores/devicePreferences'

const deviceStore = useDeviceStore()
const { sortedDevices, loading, error } = storeToRefs(deviceStore)

const devicePreferenceStore = useDevicePreferenceStore()
const { getDeviceVisibility, setDeviceVisibility, getDeviceIcon, setDeviceIcon, setSortPosition } = devicePreferenceStore

function onListDrag(e) {
  // The only drag event we care about right now is if something is moved to a new place
  if (e.moved) {
    setSortPosition(e.moved.element.device_id, e.moved.newIndex)
  }
}
</script>

<template>
  <q-list bordered padding separator q-hoverable v-if="sortedDevices">
    <q-item class="list-header">
      <!-- a bit janky, but basically adding empty headers so formatting works for q-list -->
      <q-item-section></q-item-section>
      <q-item-section></q-item-section>
      <q-item-section>Device Name</q-item-section>
      <q-item-section>Active State</q-item-section>
      <q-item-section>Drive Status</q-item-section>
    </q-item>
    <draggable
      group="devices"
      :list="sortedDevices"
      @start="dragging = true"
      @end="dragging = false"
      item-key="device_id"
      @change="onListDrag"
    >
      <template #item="{ element }">
         <q-item class="list-device">
            <q-item-section >
              <q-icon v-if="getDeviceVisibility(element)" color="primary" name="visibility" @click="setDeviceVisibility(element)" class="cursor-pointer list-icon" />
              <q-icon v-else color="negative" name="visibility_off" @click="setDeviceVisibility(element)" class="cursor-pointer list-icon" />
              <q-icon color="primary" name="my_location" @click="$emit('set-center', element.latest_accurate_device_point.device_point_detail.lat_lng)" class="cursor-pointer list-icon" />
            </q-item-section>
            <q-item-section class="list-fab">
              <q-fab color="red" text-color="white" :icon="getDeviceIcon(element)">
                <q-fab-action color="red" text-color="white" icon="directions_car" @click="setDeviceIcon(element.device_id, 'directions_car')" />
                <q-fab-action color="red" text-color="white" icon="two_wheeler" @click="setDeviceIcon(element.device_id, 'two_wheeler')" />
                <q-fab-action color="red" text-color="white" icon="agriculture" @click="setDeviceIcon(element.device_id, 'agriculture')" />
                <q-fab-action color="red" text-color="white" icon="child_friendly" @click="setDeviceIcon(element.device_id, 'child_friendly')" />
              </q-fab>
            </q-item-section>
            <q-item-section>{{ element.display_name }}</q-item-section>
            <q-item-section>{{ element.active_state }}</q-item-section>
            <q-item-section>{{ element.latest_accurate_device_point.device_state.drive_status }}</q-item-section>
          </q-item>
      </template>
    </draggable>
  </q-list>
</template>

<style scoped>
.list-header {
  font-weight:  bold;
}
.list-device {
  border:  1px solid #eee;
}
.list-icon {
  font-size:  20px;
}
.list-fab {
  position:  relative;
  left:  -20px;
}

/* Applies when list element is being dragged */
.sortable-chosen {
  background:  #eee;
}
</style>

<script>
export default {
  props: {
    center: {
      type: String,
      required: true,
      // validator(value) {
      //   return 'lat' in value &&'lng' in value && typeof(value[lat]) === 'string' && typeof(value[lng)]) === 'string'
      // }
    },
  },
}
</script>

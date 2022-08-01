<script setup>
import draggable from 'vuedraggable'

import { storeToRefs } from 'pinia'

import { useDeviceStore } from '@/stores/device'
import { useDevicePreferencesStore } from '@/stores/devicePreferences'

const { devices, loading, error } = storeToRefs(useDeviceStore())
const { getDeviceVisibility, setDeviceVisibility, getDeviceIcon, setDeviceIcon } = useDevicePreferencesStore()
</script>

<template>
  <q-list bordered padding separator q-hoverable v-if="devices">
    <draggable
      group="devices"
      :list="devices"
      @start="dragging = true"
      @end="dragging = false"
      item-key="device_id"
    >
      <template #item="{ element }">
         <q-item bordered separator class="list-device">
            <q-item-section>
              <q-icon clickable v-if="getDeviceVisibility(element)" color="primary" name="visibility" @click="setDeviceVisibility(element)" />
              <q-icon clickable v-else color="negative" name="visibility_off" @click="setDeviceVisibility(element)" />
              <q-icon clickable color="primary" name="my_location" @click="$emit('set-center', element.latest_accurate_device_point.device_point_detail.lat_lng)"/>
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
  .list-device {
    border:  1px solid #eee;
  }
  .table { height: 500px; width:  500px;}
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

<script setup>
import { storeToRefs } from 'pinia'

import { useDeviceStore } from '@/stores/device'
import { useDevicePreferencesStore } from '@/stores/devicePreferences'

const { devices, loading, error } = storeToRefs(useDeviceStore())
const { getDeviceVisibility, setDeviceVisibility, getDeviceIcon } = useDevicePreferencesStore()
console.log('got some deviceIcon', getDeviceIcon)

const columns = [
  { name: 'visible', required: false },
  { name: 'center' , required: false },
  { name: 'icon', required: true },
  {
    name: 'name',
    field: row => row.display_name,
    label: 'Device Name',
  },
  {
    name: 'active_state',
    field: row => row.active_state,
    label: 'Active State',
  },
  {
    name: 'drive_status',
    field: row => row.latest_accurate_device_point.device_state.drive_status,
    label: 'Drive Status',
  },
]

const rows = devices

// This method is used to call setSort on the preferences store
// in the event of a click on a column in the header row
const onRowClick = (event, row) => {
  console.log(event, row)
}
</script>

<template>
  <q-table
    title="GPS Devices"
    :rows="rows"
    :columns="columns"
    row-key="name"
    class="table"
  >
    <template v-slot:body="props">
      <q-tr :props="props">
        <q-td key="icon" clickable @click="setDeviceVisibility(props.row)">
          <q-icon v-if="getDeviceVisibility(props.row)" color="primary" name="visibility" />
          <q-icon v-else color="negative" name="visibility_off" />
        </q-td>
        <q-td key="icon">
          <q-icon clickable color="primary" name="my_location" @click="$emit('set-center', props.row.latest_accurate_device_point.device_point_detail.lat_lng)" />
        </q-td>
        <q-td key="icon">
          <q-icon clickable color="primary" :name="getDeviceIcon(props.row)" />
        </q-td>
        <q-td key="name" :props="props">
          {{ props.row.display_name }}
        </q-td>
        <q-td key="name" :props="props">
          {{ props.row.active_state }}
        </q-td>
        <q-td key="name" :props="props">
          {{ props.row.latest_accurate_device_point.device_state.drive_status }}
        </q-td>
      </q-tr>

    </template>
  </q-table>

<!--   <q-list bordered padding separator v-if="devices" v-for="device in devices" :key="device.device_id">
    <q-item>
      <q-item-section avatar clickable @click="setDeviceVisibility(device)">
        <q-icon clickable v-if="getDeviceVisibility(device)" color="primary" name="visibility" />
        <q-icon clickable v-else color="negative" name="visibility_off" />
      </q-item-section>
      <q-item-section avatar>
        <q-icon clickable color="primary" name="my_location" @click="setCenter(device)"/>
      </q-item-section>
      <q-item-section>{{ device.display_name }}</q-item-section>
      <q-item-section>{{ device.active_state }}</q-item-section>
      <q-item-section>{{ device.latest_accurate_device_point.device_state.drive_status }}</q-item-section>
    </q-item>
  </q-list>
 --></template>

<style scoped>
  .table { height: 500px; width:  500px;}
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
  }
}
</script>
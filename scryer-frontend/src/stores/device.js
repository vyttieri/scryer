import { defineStore } from 'pinia'

import { useDevicePreferencesStore } from '@/stores/devicePreferences'

export const useDeviceStore = defineStore({
  id: 'device',
  state: () => ({
    devices: [],
    loading: false,
    error: null,
  }),
  getters: {
    // deviceIds: state => state.devices.reduce(device => device.device_id),
    // sortedDevices: state => state.devices.sortBy((a, b) => a > b ? 1 : -1)
    visibleDevices: state => {
      const { getDeviceVisibility } = useDevicePreferencesStore()
      console.log('in visibleDevices')
      console.log('devices', state.devices)
      state.devices.forEach(device => console.log(device.latest_accurate_device_point.lat_lng))
      console.log('visibleDevices', state.devices.filter(device => getDeviceVisibility(device)))
      return state.devices.filter(device => getDeviceVisibility(device))
    },
  },
  actions: {
    async fetchDevices() {
      console.log('fetching devices')
      this.loading = true

      try {
        const devices = await fetch('http://localhost:5173/ping')
          .then(response => response.json())
          .then(data => data.response.result_list)
        console.log('logging devices from fetchDevices', devices)
        // TODO: This probably isn't working as planned
        // this.$patch({ devices: devices })
        this.devices = devices

        const { initOrPatchDevicePreferences } = useDevicePreferencesStore()
        initOrPatchDevicePreferences(devices)
      } catch (error) {
        console.log('error fetching devices', error)
        this.error = error
      }

      this.loading = false
    },
  }
})

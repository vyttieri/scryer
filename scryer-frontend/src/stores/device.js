import { defineStore, storeToRefs } from 'pinia'

import { useDevicePreferencesStore } from '@/stores/devicePreferences'

export const useDeviceStore = defineStore({
  id: 'device',
  state: () => ({
    devices: [],
    loading: false,
    error: null,
  }),
  getters: {
    sortedDevices: state => {
      const { sortPositions }  = storeToRefs(useDevicePreferencesStore())
      let deviceMap = {}
      sortPositions.value.forEach((deviceId, i) => deviceMap[deviceId] = i)
      console.log('dedvice map', deviceMap)
      console.log('what is devices', state.devices)
      return state.devices.sort((a, b) => deviceMap[a.device_id] > deviceMap[b.device_id] ? 1 : -1)
    },
    visibleDevices: state => {
      const { getDeviceVisibility } = useDevicePreferencesStore()

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

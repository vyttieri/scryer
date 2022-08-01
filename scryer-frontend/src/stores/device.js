import { defineStore, storeToRefs } from 'pinia'

import { useDevicePreferenceStore } from '@/stores/devicePreferences'

export const useDeviceStore = defineStore({
  id: 'device',
  state: () => ({
    devices: [],
    loading: false,
    error: null,
  }),
  getters: {
    sortedDevices: state => {
      const { sortPositions }  = storeToRefs(useDevicePreferenceStore())
      let deviceMap = {}
      sortPositions.value.forEach((deviceId, i) => deviceMap[deviceId] = i)
      return state.devices.sort((a, b) => deviceMap[a.device_id] > deviceMap[b.device_id] ? 1 : -1)
    },
    visibleDevices: state => {
      const { getDeviceVisibility } = useDevicePreferenceStore()

      return state.devices.filter(device => getDeviceVisibility(device))
    },
  },
  actions: {
    async fetchDevices() {
       this.loading = true

      try {
        const devices = await fetch('http://localhost:5173/ping')
          .then(response => response.json())
          .then(data => data.response.result_list)
        this.devices = devices

        const { initOrPatchDevicePreferences } = useDevicePreferenceStore()
        initOrPatchDevicePreferences(devices)
      } catch (error) {
        this.error = error
      }

      this.loading = false
    },
  }
})

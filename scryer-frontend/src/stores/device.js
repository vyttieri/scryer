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
       this.error = null

      try {
        await fetch('http://localhost:5173/ping')
          .then(response => response.json())
          .then(data => data.response.result_list)
          .then(devices => {
            this.devices = devices

            const { initOrPatchDevicePreferences } = useDevicePreferenceStore()
            initOrPatchDevicePreferences(devices)

            // The fetch was honestly happening so quickly you couldn't see it in the UI.
            // Adding a 1s delay here ensures the user sees that something is happening.
            setTimeout(() => {
              this.loading = false
              this.error = null
             }, 1000) // 1000ms = 1second
          })
      } catch (error) {
        this.error = error
        this.loading = false
      }
    },
  }
})

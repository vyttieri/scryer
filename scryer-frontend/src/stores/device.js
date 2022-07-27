import { defineStore } from 'pinia'

import { usePreferencesStore } from './preferences'

export const useDeviceStore = defineStore({
  id: 'device',
  state: () => ({
    devices: [],
    loading: false,
    error: null,
  }),
  getters: {
    deviceIds: state => state.devices.reduce(device => device.device_id),
    // sortedDevices: state => state.devices.sortBy((a, b) => a > b ? 1 : -1)
    visibleDevices: state => {
      const { getDeviceVisibility } = usePreferencesStore()

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

        // TODO: This probably isn't working as planned
        this.$patch({ devices: devices })

        const { initOrPatchDevicePreferences } = usePreferencesStore()
        initOrPatchDevicePreferences(devices)
      }
      finally {
        this.loading = false
      }
    },
  }
})

import { defineStore } from 'pinia'

export const useDeviceStore = defineStore({
  id: 'device',
  state: () => ({
    devices: [],
    loading: false,
    error: null,
  }),
  getters: {
    getDevices: state => state.devices,
  },
  actions: {
    async fetchDevices() {
      this.loading = true

      try {
        this.devices = await fetch('http://localhost:5173/ping')
          .then(response => response.json())
      } catch (error) {
        this.error = error
      } finally {
        this.loading = false
      }
    }
  }
})

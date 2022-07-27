import { defineStore } from 'pinia'

export const usePreferencesStore = defineStore({
  id: 'preferences',
  state: () => ({
    devices: {},
  }),
  getters: {
    // TODO: Compose stores instead?
    getDeviceVisibility: state => {
      return deviceId => {
        var device = state.devices[deviceId]

        return device === undefined ? true : device.visible
      }
    },
    visibleDevices(state) {
      return state.devices.filter(device => {
        // TODO: Better way to default visible
        let visible = devices[device.device_id].visible
        visible = visible == undefined ? true : visible
        return device.visible
      })
    }
  },
  actions: {
    setDeviceVisibility(device) {
      console.log('we clicked the device icon', device)
      // TODO: Better way to default visible
      var devicePreferences = this.devices[device.device_id]
      var visible = devicePreferences === undefined ? true : !devicePreferencesdevice === undefined ? true : !device.visible.visible
      this.devices[device.device_id] = visible
    },
  },
})

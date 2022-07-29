import { defineStore } from 'pinia'

export const useDevicePreferencesStore = defineStore({
  id: 'devicePreferences',
  state: () => {
    devicePreferences: {}
  },
  getters: {
    getDeviceVisibility: state => {
      return device => {
        console.log('getting device visibility', device)
        console.log('fucking state', state)
        return state.devicePreferences[device.device_id].visible
      }
    },
  },
  actions: {
    initOrPatchDevicePreferences(devices) {
      console.log('logging devices from preferences', devices)

      let devicePreferences = devices.reduce((acc, device) => {
        return { ...acc, [device.device_id]: { visible: true } }
      }, {})

      // spread will overwrite with second object, so we're effectively only adding new devices
      this.devicePreferences = { ...devicePreferences, ...this.devicePreferences}
      console.log('updated devicePreferences', devicePreferences)
    },
    setDeviceVisibility(device) {
      console.log('we clicked the device icon', device, device.device_id)

      var visible = this.devicePreferences[device.device_id].visible
      this.devicePreferences[device.device_id] = { 'visible': !visible }
    },
  },
})

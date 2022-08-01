import { defineStore } from 'pinia'

export const useDevicePreferencesStore = defineStore({
  id: 'devicePreferences',
  state: () => ({
    devicePreferences: {
      // visible: true, icon: 'car', sortPosition: 1
    }
  }),
  getters: {
    getDeviceVisibility: state => {
      return device => {
        return state.devicePreferences[device.device_id].visible
      }
    },
    getDeviceIcon: state => {
      return device => state.devicePreferences[device.device_id].icon
    },
    // for backend consumption:
    jsonDevicePreferences: state => {
      return Object.keys(state.devicePreferences).map(deviceId => {
        return { ...state.devicePreferences[deviceId], deviceId: deviceId }
      })
    }
  },
  actions: {
    initOrPatchDevicePreferences(devices) {
      console.log('logging devices from preferences', devices)

      let devicePreferences = devices.reduce((acc, device) => {
        return { ...acc, [device.device_id]: { visible: true, icon: 'directions_car', sortPosition: 0 } }
      }, {})

      // spread will overwrite with second object, so we're effectively only adding new devices
      this.devicePreferences = { ...devicePreferences, ...this.devicePreferences}
      console.log('updated devicePreferences', devicePreferences)
    },
    setDeviceVisibility(device) {
      console.log('we clicked the device icon', device, device.device_id)

      var visible = this.devicePreferences[device.device_id].visible
      this.devicePreferences[device.device_id].visible = !visible
    },
    setDevicePreferences(devicePreferences) {
      this.devicePreferences = devicePreferences
    },
    async updateDevicePreferences() {
      console.log('updateDevicePreferences hit')
      await fetch(`http://localhost:5173/user/preferences`, {
        method: 'PUT',
        headers: {
          'Accept': 'application/json',
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ devicePreferences: this.jsonDevicePreferences }),
      })
    },

  },
})

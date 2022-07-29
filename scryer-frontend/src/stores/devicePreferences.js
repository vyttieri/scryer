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
    // this maps from simple names to material ui icon names
    getDeviceIcon: state => {
      return device => {
        const icon = state.devicePreferences[device.device_id].icon

        if (icon === 'car') return 'directions_car'
        else if (icon === 'truck') return 'local_shipping'
        else if (icon === 'scooter') return 'delivery_dining'
        else if (icon === 'motorcycle') return 'two_wheeler'
      }
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
        return { ...acc, [device.device_id]: { visible: true, icon: 'car' } }
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
    }
  },
})

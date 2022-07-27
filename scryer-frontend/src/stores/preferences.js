import { defineStore } from 'pinia'

export const usePreferencesStore = defineStore({
  id: 'preferences',
  state: () => ({
    devicePreferences: {},
    sortColumn: 'display_name',
  }),
  getters: {
    // TODO: Compose stores instead?
    getDeviceVisibility: state => {
      return deviceId => {
        var device = state.devicePreferences[deviceId]

        return device === undefined ? true : device.visible
      }
    },
  },
  actions: {
    initOrPatchDevicePreferences(devices) {
      console.log('logging devices', devices)
      let reducedDevices = {}
      // TODO: use Reduce or Map or something
      devices.forEach(device => reducedDevices[device.device_id] = { visible: true })
      console.log('logging reducedDevices', reducedDevices)
      this.devicePreferences = reducedDevices
    },
    setDeviceVisibility(device) {
      console.log('we clicked the device icon', device, device.device_id)

      // TODO: Better way to default visible
      var visible = this.devicePreferences[device.device_id].visible
      this.devicePreferences[device.device_id] = { 'visible': !visible }
    },
  },
})

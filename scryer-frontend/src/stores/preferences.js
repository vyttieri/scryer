import { defineStore } from 'pinia'

export const usePreferencesStore = defineStore({
  id: 'preferences',
  state: () => ({
    devicePreferences: {},
    sortColumn: 'display_name',
  }),
  getters: {
    getDeviceVisibility: state => {
      return device => state.devicePreferences[device.device_id].visible
    },
  },
  actions: {
    initOrPatchDevicePreferences(devices) {
      console.log('logging devices', devices)
      let devicePreferences = {}
      // TODO: use Reduce or Map or something
      devices.forEach(device => devicePreferences[device.device_id] = { visible: true })
      console.log('logging reducedDevices', devicePreferences)
      // TODO: Figure out how $patch works
      // spread will overwrite with second object, so we're effectively only adding new devices
      this.devicePreferences = { ...devicePreferences, ...this.devicePreferences}
    },
    setDeviceVisibility(device) {
      console.log('we clicked the device icon', device, device.device_id)

      var visible = this.devicePreferences[device.device_id].visible
      this.devicePreferences[device.device_id] = { 'visible': !visible }
    },
  },
})

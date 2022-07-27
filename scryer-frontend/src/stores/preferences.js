import { defineStore } from 'pinia'

export const usePreferencesStore = defineStore({
  id: 'preferences',
  state: () => ({
    center: { lat: 35, lng: -110 },
    devicePreferences: {},
    sortColumn: 'display_name',
    sortAscending: false,
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
    setCenter(device) {
      console.log('centering', device.latest_accurate_device_point.device_point_detail.lat_lng)
      this.center = device.latest_accurate_device_point.device_point_detail.lat_lng
    },
    setSort(column) {
      console.log('clicked sort', column)
    }
  },
})

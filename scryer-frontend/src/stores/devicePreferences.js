import { defineStore } from 'pinia'

export const useDevicePreferencesStore = defineStore({
  id: 'devicePreferences',
  state: () => ({
    devicePreferences: {
    },
    // Storing sortPositions as an ordered list so we can do reordering with VueDraggable.
    sortPositions: [],
  }),
  persist: true,
  getters: {
    getDeviceVisibility: state => {
      return device => {
        return state.devicePreferences[device.device_id].visible
      }
    },
    getDeviceIcon: state => {
      return device => state.devicePreferences[device.device_id].icon
    },
    getDeviceSortPosition: state => {
      return device => state.devicePreferences[device.device_id].sortPosition
    },
    // for backend consumption:
    jsonDevicePreferences: state => {
      // Add sortPositions back to individual devicePreference records
      let sortMap = {}
      state.sortPositions.forEach((deviceId, i) => {
        sortMap[deviceId] = i
      })

      return Object.keys(state.devicePreferences).map(deviceId => {
        return { ...state.devicePreferences[deviceId], deviceId: deviceId, sortPosition: sortMap[deviceId] }
      })
    }
  },
  actions: {
    initOrPatchDevicePreferences(devices) {
      console.log('logging devices from preferences', devices)

      let devicePreferences = devices.reduce((acc, device) => {
        return { ...acc, [device.device_id]: { visible: true, icon: 'directions_car' } }
      }, {})

      // spread will overwrite with second object, so we're effectively only adding new devices here
      this.devicePreferences = { ...devicePreferences, ...this.devicePreferences}

      let sortPositions = devices
        .sort((deviceA, deviceB) => deviceA.display_name > deviceB.display_name ? 1 : -1)
        .map(device => device.device_id)
      if (this.sortPositions === []) {
        this.sortPositions = sortPositions
      }

      console.log('updated devicePreferences', devicePreferences)
      console.log('updated sortPositions', sortPositions)
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
    setDeviceIcon(deviceId, icon) {
      this.devicePreferences[deviceId].icon = icon
    },
    setSortPosition(deviceId, newPosition) {
      console.log(deviceId, newPosition)

      // remove old position
      this.sortPositions = this.sortPositions.filter(listDeviceId => deviceId !== listDeviceId)
      console.log('intermediate', this.sortPositions)
      // insert at new position
      this.sortPositions.splice(newPosition, 0, deviceId)
      console.log('final', this.sortPositions)
    },
  },
})

import { defineStore, storeToRefs } from 'pinia'

import { useDeviceStore } from '@/stores/device'
import { useDevicePreferenceStore } from '@/stores/devicePreferences'
import { useUserStore } from '@/stores/user'

export const useAuthStore = defineStore({
  id: 'auth',
  state: () => ({
    error: null,
    loading: false,
  }),
  getters: {
    loggedIn: state => {
      const { username } = storeToRefs(useUserStore())

      return username.value !== null
    }
  },
  actions: {
    async login(username, password) {
      this.loading = true
      this.error = null
      try {
        await fetch('http://localhost:5173/login', {
          method: 'POST',
          headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({ username: username, password: password })
        })
          .then(response => {
            if (response.ok) {
              return response.json()
            } else {
              throw new Error()
            }
          })
          .then(data => {
            useUserStore().setUser(data.user.id, data.user.username)

            // convert from backend [] format to frontend {} format
            const devicePreferences = data.user.devicePreferences.reduce((acc, devicePreference) => {
                    return {
                      ...acc,
                      [devicePreference.deviceId]: {
                        id: devicePreference.id,
                        icon: devicePreference.icon,
                        sortOrder: devicePreference.sortOrder,
                        visible: devicePreference.visible,
                      }
                    }
                  },
            {})
            useDevicePreferenceStore().setDevicePreferences(devicePreferences)
          })
      } catch (error) {
        this.error = 'Failed to login'
      }
    },
    async logout() {
      await fetch('http://localhost:5173/logout')
        .then(response => {
          if (response.status === 200) {
            // reset user store
            useUserStore().$reset()

            // reset devicePreferences store. Basically reset it to null and then init
            // from devices list again
            const { devices } = storeToRefs(useDeviceStore())
            const devicePreferenceStore = useDevicePreferenceStore()
            devicePreferenceStore.$reset()
            devicePreferenceStore.initOrPatchDevicePreferences(devices.value)
          } else {
            throw new Error("failed to logout, ", response.status)
          }
        })
    },
  }
})

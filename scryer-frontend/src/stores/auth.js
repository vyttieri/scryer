import { defineStore } from 'pinia'

import { useUserStore } from './user'

export const useAuthStore = defineStore({
  id: 'auth',
  state: () => ({
    error: null,
    loading: false,
    accessToken: localStorage.getItem('accessToken'),
  }),
  getters: {
    loggedIn: state => state.accessToken !== null
  },
  actions: {
    async login(username, password) {
      this.loading = true
      this.error = null
      try {
        console.log('logging in', username, password)
        await fetch('http://localhost:5173/login', {
          method: 'POST',
          headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({ username: username, password: password })
        })
          .then(response => response.json())
          .then(result => {
            this.accessToken = result.token

            localStorage.setItem('accessToken', result.token)
            console.log(result.token)

            useUserStore().getUser()
          })

      } catch (error) {
        console.log('got error logging in', error)
        this.error = error
      }

    },
    logout() {
      this.accessToken = null
      localStorage.removeItem('accessToken')
    },
    async refreshToken() {
      try {
        const accessToken = await fetch('http://localhost:5173/auth/refresh_token')
      } catch (error) {

      }
    }
  }
})

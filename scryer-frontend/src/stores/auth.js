import { defineStore, storeToRefs } from 'pinia'

import { useUserStore } from './user'

export const useAuthStore = defineStore({
  id: 'auth',
  state: () => ({
    error: null,
    loading: false,
  }),
  getters: {
    // Data should have a single source of truth,
    // and I'm trying to follow the single responsibility principle...
    // so that's why this is a bit weird.
    loggedIn: state => {
      const { userId } = storeToRefs(useUserStore())

      return userId === '' ? false : true
    }
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
          .then(user => {
            useUserStore().setUser(user.userId, user.username)
          })

      } catch (error) {
        console.log('got error logging in', error)
        this.error = error
      }
    },
    async logout() {
      try {
          console.log('logging out')
          await fetch('http://localhost:5173/logout')
            .then(response => {
              if (response.status === 200) {
                useUserStore().setUser(null, null)
                console.log('successfully logged out')
              } else {
                // panic at the disco
              }
            })
        } catch (error) {
          console.log('got error logging in', error)
          this.error = error
        }
    },
  }
})

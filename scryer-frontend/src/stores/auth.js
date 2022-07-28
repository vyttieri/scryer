import { defineStore } from 'pinia'

export const useAuthStore = defineStore({
  id: 'auth',
  state: () => ({
    user: JSON.parse(localStorage.getItem('user')),
    error: null,
    loading: false,
    accessToken: JSON.parse(localStorage.getItem('accessToken'))
  }),
  actions: {
    async login(username, password) {
      this.loading = true
      try {
        console.log('logging in', username, password)
        const result = await fetch('http://localhost:5173/login', {
          method: 'POST',
          headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
          },
          body: JSON.stringify({ username: username, password: password })
        })

        console.log('result', result)
        this.accessToken = result.token
        console.log(result.token)

      } catch (error) {
        console.log('got error logging in', error)
        this.error = error
      }

    },
    logout() {
      this.user = null
      localStorage.removeItem('user')
    },
    async refreshToken() {
      try {
        const accessToken = await fetch('http://localhost:5173/auth/refresh_token')
      } catch (error) {

      }
    }
  }
})

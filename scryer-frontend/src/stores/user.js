import { defineStore } from 'pinia'

import { useAuthStore } from './auth'

export const useUserStore = defineStore({
	id: 'user',
	state: () => ({
		user: {}
	}),
	actions: {
		async register(username, password) {
		 const result = await fetch('http://localhost:5173/login', {
					method: 'POST',
					headers: {
						'Accept': 'application/json',
						'Content-Type': 'application/json',
					},
					body: JSON.stringify({ username: username, password: password }),
			})
		 	.then(response => response.json())

		 console.log('result', result)

		 const authResult = useAuthStore().login(username, password)
		},
		async getPreferences(user) {
			await fetch.post('blah', user)
		},
		async updatePreferences(user) {
			'asdf'
		},
	}
})

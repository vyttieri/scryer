import { defineStore } from 'pinia'

import { useAuthStore } from './auth'

export const useUserStore = defineStore({
	id: 'user',
	state: () => ({
		user: {}
	}),
	actions: {
		async register(username, password, passwordConfirmation) {
		 const result = await fetch('http://localhost:5173/users', {
					method: 'POST',
					headers: {
						'Accept': 'application/json',
						'Content-Type': 'application/json',
					},
					body: JSON.stringify({ username: username, password: password, passwordConfirmation: passwordConfirmation }),
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

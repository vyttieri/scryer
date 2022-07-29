import { defineStore } from 'pinia'

import { useAuthStore } from './auth'
import { usePreferencesStore } from './preferences'

export const useUserStore = defineStore({
	id: 'user',
	state: () => ({
		userId: '',
		username: ''
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
		 this.userId = result.userId
		 this.username = result.username

		 const authResult = useAuthStore().login(username, password)
		},
		async getUser() {
			console.log('getting user')
			var user = await fetch(`http://localhost:5173/users/${userId}`, {
				method: 'GET',
				headers: {
					'Authorization': `Bearer ${useAuthStore().accessToken}`
				},
			})
				.then(response => response.json())

			this.userId = result.userId
			this.username = result.username

			console.log('got user', user)
		},
		async updatePreferences(user) {
			'asdf'
		},
	}
})

import { defineStore } from 'pinia'

import { useAuthStore } from '@/stores/auth'
import { useDevicePreferencesStore } from '@/stores/devicePreferences'

export const useUserStore = defineStore({
	id: 'user',
	state: () => ({
		userId: null,
		username: null
	}),
	actions: {
		async register(username, password, passwordConfirmation) {
			const jsonUser = {
				username,
				password,
				passwordConfirmation,
				devicePreferences: useDevicePreferencesStore().jsonDevicePreferences,
			}
			console.log('stringify user:', JSON.stringify(jsonUser))
		 await fetch('http://localhost:5173/users', {
				method: 'POST',
				headers: {
					'Accept': 'application/json',
					'Content-Type': 'application/json',
				},
				body: JSON.stringify(jsonUser),
			})
		 	.then(response => response.json())
		 	.then(result => {
   			  this.userId = result.user.userId
			  this.username = result.user.username
		 	})
		},
		setUser(userId, username) {
			this.userId = userId
			this.username = username
		},
		async updateDevicePreferences() {
			await fetch(`http://localhost:5173/users/${this.userId}/preferences`, {
				method: 'POST',
				headers: {
					'Accept': 'application/json',
					'Content-Type': 'application/json',
				},
				body: JSON.stringify(useDevicePreferencesStore().jsonDevicePreferences),
			})
		},
	}
})

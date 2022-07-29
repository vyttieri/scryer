import { defineStore, storeToRefs } from 'pinia'

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
			let devicePreferences = storeToRefs(useDevicePreferencesStore()).devicePreferences
			const jsonDevicePreferences = Object.keys(devicePreferences).map(deviceId => {
				return { ...devicePreferences.deviceId, deviceId: deviceId }
			})
			console.log('got them devicePreferences in register', jsonDevicePreferences)

			const jsonUser = {
				username,
				password,
				passwordConfirmation,
				devicePreferences: jsonDevicePreferences,
			}
			console.log('jsonUser', jsonUser)
		 await fetch('http://localhost:5173/users', {
				method: 'POST',
				headers: {
					'Accept': 'application/json',
					'Content-Type': 'application/json',
				},
				body: JSON.stringify(jsonUser),
			})
		 	.then(response => response.json())
		 	.then(user => {
		 		console.log('sup user')
   			  this.userId = user.userId
			  this.username = user.username
		 	})
		},
		async getUser() {
			console.log('getting user')
			var user = await fetch(`http://localhost:5173/users/${userId}`, {
				method: 'GET',
				headers: {
				},
			})
				.then(response => response.json())

			this.userId = result.userId
			this.username = result.username

			console.log('got user', user)
		},
		setUser(userId, username) {
			this.userId = userId
			this.username = username
		},
		async updatePreferences(user) {
			'asdf'
		},
	}
})

import { defineStore } from 'pinia'

import { useAuthStore } from '@/stores/auth'
import { useDevicePreferenceStore } from '@/stores/devicePreferences'

export const useUserStore = defineStore({
	id: 'user',
	state: () => ({
		userId: null,
		username: null,
	}),
	persist: true,
	actions: {
		async register(username, password, passwordConfirmation) {
			const jsonUser = {
				username,
				password,
				passwordConfirmation,
				devicePreferences: useDevicePreferenceStore().jsonDevicePreferences,
			}
		 await fetch('http://localhost:5173/register', {
				method: 'POST',
				headers: {
					'Accept': 'application/json',
					'Content-Type': 'application/json',
				},
				body: JSON.stringify(jsonUser),
			})
		 	.then(response => response.json())
		 	.then(result => {
   			  this.userId = result.user.id
			  this.username = result.user.username
		 	})
		},
		setUser(userId, username) {
			this.userId = userId
			this.username = username
		},
	}
})

import { defineStore } from 'pinia'

export const useUserStore = defineStore({
	id: 'user',
	state: () => ({
		user: {}
	}),
	actions: {
		async register(user) {
			await fetch.post('blah', user)
		},
		async getPreferences(user) {
			await fetch.post('blah', user)
		},
		async updatePreferences(user) {
			'asdf'
		},
	}
})
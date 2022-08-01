<script setup>
import { storeToRefs } from 'pinia'

import { useAuthStore } from '@/stores/auth'
import { useUserStore } from '@/stores/user'
import { useDevicePreferencesStore } from '@/stores/devicePreferences'
import LoginForm from '@/components/LoginForm.vue'


const authStore = useAuthStore()
const devicePreferencesStore = useDevicePreferencesStore()
const userStore = useUserStore()
const { username } = storeToRefs(userStore)
</script>

<template>
  <q-header elevated class="bg-primary text-white">
    <q-toolbar>
      <q-toolbar-title>
        <q-avatar>
          <img src="../../../scryer.png" />
        </q-avatar>
        <span v-if="username !== null">Welcome to Scryer, {{username}}!</span>
        <span v-else>Welcome to Scryer!</span>
      </q-toolbar-title>

      <div>
        <q-btn v-if="username !== null" clickable><q-icon name="save" @click="devicePreferencesStore.updateDevicePreferences()" /></q-btn>
        <q-btn v-if="username === null" label="Login" color="primary" @click="login = true" />
          <q-menu>
            <LoginForm v-if="register === false" @toggle-forms="toggleForms" />
            <RegistrationForm v-else @toggle-forms="toggleForms" />
          </q-menu>
        <q-btn v-if="username !== null" label="Logout" color="primary" @click="authStore.logout()"/>
      </div>
    </q-toolbar>

    <q-ajax-bar />
  </q-header>
</template>

<script>
export default {
  data() {
    return {
      register: false,
    }
  },
  methods: {
    toggleForms() {
      this.register = !this.register
    },
  },
}
</script>

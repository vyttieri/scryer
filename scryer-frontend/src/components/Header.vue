<script setup>
import { storeToRefs } from 'pinia'

import { useAuthStore } from '@/stores/auth'
import { useDeviceStore } from '@/stores/device'
import { useUserStore } from '@/stores/user'
import { useDevicePreferenceStore } from '@/stores/devicePreferences'

import LoginForm from '@/components/LoginForm.vue'
import RegistrationForm from '@/components/RegistrationForm.vue'

const authStore = useAuthStore()
const { loggedIn } = storeToRefs(authStore)
const { logout } = authStore

const devicePreferenceStore = useDevicePreferenceStore()

const { loading } = storeToRefs(useDeviceStore())

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
        <span v-if="loggedIn">Welcome to Scryer, {{ username }}!</span>
        <span v-else>Welcome to Scryer!</span>
      </q-toolbar-title>

      <div>
        <q-spinner v-if="loading" style="position: relative; left: -10px;" />
        <q-btn v-if="loggedIn" clickable><q-icon name="save" @click="devicePreferenceStore.updateDevicePreferences()" /></q-btn>
        <q-btn v-if="!loggedIn" label="Login" color="primary" @click="login = true" />
        <q-menu v-if="!loggedIn" style="padding: 4px;">
            <LoginForm v-if="register === false" @toggle-forms="toggleForms" />
            <RegistrationForm v-else @toggle-forms="toggleForms" />
        </q-menu>
        <q-btn v-if="loggedIn" label="Logout" color="primary" @click="logout()"/>
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

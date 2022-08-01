<script setup>
import { storeToRefs } from 'pinia'

import { useAuthStore } from '@/stores/auth'
import { useUserStore } from '@/stores/user'
import { useDevicePreferencesStore } from '@/stores/devicePreferences'

import Popup from '@/components/Popup.vue'

const authStore = useAuthStore()

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
          <p v-if="username !== null">Welcome to Scryer, {{username}}!</p>
          <p v-else>Welcome to Scryer!</p>
        </q-toolbar-title>

<!--         <q-btn icon="login" color="secondary" @click="prompt = true" />

      <q-dialog v-model="prompt" persistent>
        <q-card style="min-width: 350px; MIN-HEIGHT: 350px;">
          fuckera
            <q-btn flat label="Cancel" v-close-popup />
        </q-card>
      </q-dialog> -->a

      <q-btn v-if="username !== null" clickable><q-icon name="save" @click="useDevicePreferencesStore().updateDevicePreferences()" /></q-btn>
      <q-btn v-if="username === null" label="Login" color="primary" @click="$refs.dialogRef.show(); " />
      <q-btn v-else label="Logout" color="primary" @click="authStore.logout()"/>
      <Popup />

    </q-toolbar>
    <q-ajax-bar :hijack-filter="filterAjaxBar" />
  </q-header>
</template>

<script>
export default {
  setup() {
    return {
      filterAjaxBar(url) {
        // We only want to display the ajax bar when we go to save user preferences.
        return /^localhost:5173\/users\/[\d]+\/preferences/.test(url)
      }
    }
  }
}
</script>

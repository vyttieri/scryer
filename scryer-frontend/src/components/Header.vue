<script setup>
import { storeToRefs } from 'pinia'

import { useAuthStore } from '../stores/auth'

import Popup from './Popup.vue'

const authStore = useAuthStore()
const { loggedIn } = storeToRefs(authStore)
</script>

<template>
    <q-header elevated class="bg-primary text-white">
      <q-toolbar>
        <q-toolbar-title>
          <q-avatar>
            <img src="../../../scryer.png" />
          </q-avatar>
          Scryer: Your portal to the future
        </q-toolbar-title>

<!--         <q-btn icon="login" color="secondary" @click="prompt = true" />

      <q-dialog v-model="prompt" persistent>
        <q-card style="min-width: 350px; MIN-HEIGHT: 350px;">
          fuckera
            <q-btn flat label="Cancel" v-close-popup />
        </q-card>
      </q-dialog> -->a

      <q-btn v-if="!loggedIn" label="Login" color="primary" @click="$refs.dialogRef.show(); " />
      <q-btn v-if="loggedIn" label="Logout" color="primary" @click="authStore.logout()"/>
      <Popup />
    </q-toolbar>
  </q-header>
</template>

<script setup>
import { useAuthStore } from '../stores/auth.js'

function onSubmit(username, password) {
  const authStore = useAuthStore()
  console.log('subbmitted, values:', username, password)
  authStore.login(username, password)
}
</script>

<template>
  <q-form @submit="onSubmit(username, password)">
    <q-input v-model="username" filled type="text" hint="Username" />
    <q-input v-model="password" filled type="password" hint="Password" />
    <q-btn label="Login" type="submit" />
    <div v-if="errors">{{errors.ApiError}}</div>
  </q-form>
</template>

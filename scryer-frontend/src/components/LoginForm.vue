<script setup>
import { useAuthStore } from '../stores/auth.js'

function onSubmit(values) {
  const authStore = useAuthStore()
  const { username, password } = values

  console.log(username, password)
  authStore.login(username, password)
}
</script>

<template>
  <q-form @submit="onSubmit">
    <q-input v-model="username" filled type="text" hint="Username" />
    <q-input v-model="password" filled type="password" hint="Password" />
    <q-btn label="Login" type="submit" />
    <div v-if="errors">{{errors.ApiError}}</div>
  </q-form>
</template>

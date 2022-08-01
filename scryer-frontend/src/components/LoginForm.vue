<script setup>
import { useAuthStore } from '../stores/auth.js'
const authStore = useAuthStore()
</script>

<template>
  <q-form @submit="authStore.login(username, password)">
    <q-input v-model="username" filled type="text" hint="Username" />
    <q-input v-model="password" filled type="password" hint="Password" style="margin-bottom: 10px;" />
    <q-btn label="Login" type="submit" />
    <a @click="$emit('toggle-forms')" class="cursor-pointer" style="padding: 5px; margin-left: 20px; display: inline-block;">Register</a>
    <div v-if="errors">{{errors.ApiError}}</div>
  </q-form>
</template>

<script>
  export default {
    data() {
      return {
        errors: null,
        username: '',
        password: '',
      }
    }
  }
</script>
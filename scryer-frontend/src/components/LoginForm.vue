<script setup>
import { storeToRefs } from 'pinia'
import { ref } from 'vue'

import { useAuthStore } from '@/stores/auth'
const authStore = useAuthStore()
const { login } = authStore
const { error } = storeToRefs(authStore)

let username = ref('')
let password = ref('')
</script>

<template>
  <q-form @submit="login(username, password)">
    <q-input
      v-model="username"
      filled
      type="text"
      hint="Username"
      :rules="[ val => val.length > 0 || 'Please enter a username' ]"
    />
    <q-input
      v-model="password"
      filled
      type="password"
      hint="Password"
      style="margin-bottom: 10px;"
      :rules="[ val => val.length > 0 || 'Please enter a password' ]"
    />
    <div class="text-negative" v-if="error" style="margin-left: 10px;">{{ error }}</div>
    <q-btn label="Login" type="submit" />
    <a @click="$emit('toggle-forms')" class="cursor-pointer" style="padding: 5px; margin-left: 20px; display: inline-block;">Register</a>
  </q-form>
</template>

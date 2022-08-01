<script setup>
import { useAuthStore } from '@/stores/auth'
const { login } = useAuthStore()
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

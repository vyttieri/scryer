<script setup>
import { useUserStore } from '@/stores/user'
const { register } = useUserStore()
</script>

<template>
  <q-form @submit="register(username, password, passwordConfirmation)">
    <q-input
      v-model="username"
      filled
      type="text"
      hint="Username"
      :rules="[ val => val.length <= 24 || 'Please use a maximum of 32 characters',
                val => val.length > 0 || 'Please enter a username' ]"
    />
    <q-input
      v-model="password"
      filled
      type="password"
      hint="Password"
      :rules="[ val => val.length <= 32 || 'Please use a maximum of 64 characters',
                val => val.length > 0 || 'Please enter a password']"
    />
    <q-input
      v-model="passwordConfirmation"
      filled
      type="password"
      hint="Confirm Password"
      style="margin-bottom: 10px;"
      :rules="[ val => val.length <= 32 || 'Please use a maximum of 64 characters',
          val => val.length > 0 || 'Please confirm password']"
 />
    <q-btn label="Register" type="submit" />
    <a @click="$emit('toggle-forms')" class="cursor-pointer" style="padding: 5px; margin-left: 20px; display: inline-block;">Login</a>
    <!-- <div v-if="errors">{{errors}}</div> -->
  </q-form>
</template>

<script>
  export default {
    data() {
      return {
        errors: null,
        username: '',
        password: '',
        passwordConfirmation: ''
      }
    }
  }
</script>
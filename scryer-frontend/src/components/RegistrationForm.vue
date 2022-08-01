<script setup>
import { useUserStore } from '@/stores/user'

const userStore = useUserStore()
</script>

<template>
  <q-form @submit="userStore.register(username, password, passwordConfirmation)">
      <q-input v-model="username" filled type="text" hint="Username" />
      <q-input v-model="password" filled type="password" hint="Password" />
      <q-input v-model="passwordConfirmation" filled type="password" hint="Confirm Password" style="margin-bottom: 10px;" />
      <q-btn label="Register" type="submit" />
      <a @click="$emit('toggle-forms')" class="cursor-pointer" style="padding: 5px; margin-left: 20px; display: inline-block;">Login</a>
    <div v-if="errors">{{errors.ApiError}}</div>
  </q-form>
</template>


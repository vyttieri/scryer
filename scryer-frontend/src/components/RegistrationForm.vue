<script setup>
import { useUserStore } from '@/stores/user'

const userStore = useUserStore()
</script>

<template>
  <q-form @submit="userStore.register(username, password, passwordConfirmation)">
    <q-input v-model="username" filled type="text" hint="Username" />
    <q-input v-model="password" filled type="password" hint="Password" />
    <q-input v-model="passwordConfirmation" filled type="password" hint="Confirm Password" />
    <q-btn label="Register" type="submit" />
    <a @click="$emit('toggle-forms')" class="cursor-pointer">Login</a>
    <div v-if="errors">{{errors.ApiError}}</div>
  </q-form>
</template>


<script setup>
import Popup from './Popup.vue'


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

      <q-btn label="Login" color="primary" @click="$refs.dialogRef.show(); " />
      <Popup />
    </q-toolbar>
  </q-header>
</template>

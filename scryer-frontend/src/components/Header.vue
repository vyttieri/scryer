<template>

    <q-header elevated class="bg-primary text-white">
      <q-toolbar>
        <q-toolbar-title>
          <q-avatar>
          	<img src="../../../scryer.png" />
          </q-avatar>
          Scryer: Your portal to the future
        </q-toolbar-title>

        <q-btn icon="login" color="secondary" @click="prompt = true" />

	    <q-dialog v-model="prompt" persistent>
	      <q-card style="min-width: 350px; MIN-HEIGHT: 350PX;">
	        <q-card-section>
	          <div class="text-h6">Your address</div>
	        </q-card-section>

	        <q-card-section class="q-pt-none">
	          <q-input dense v-model="address" autofocus @keyup.enter="prompt = false" />
	        </q-card-section>

	        <q-card-actions align="right" class="text-primary">
	          <q-btn flat label="Cancel" v-close-popup />
	          <q-btn flat label="Add address" v-close-popup />
	        </q-card-actions>
	      </q-card>
	    </q-dialog>
      </q-toolbar>
    </q-header>
</template>

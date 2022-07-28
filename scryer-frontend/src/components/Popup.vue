<script setup>
	const $q = useQuasar()

	$q.dialog({

	})
</script>

<template>
  <div class="q-pa-md q-gutter-sm">
    <q-btn label="Login" color="primary" @click="login = true" />

    <q-dialog v-model="login" persistent>
      <q-card style="min-width: 350px">
        <q-card-section>
          <div class="text-h6">Your address</div>
        </q-card-section>

        <q-card-section class="q-pt-none">
          <q-input dense v-model="address" autofocus @keyup.enter="login = false" />
        </q-card-section>

        <q-card-actions align="right" class="text-primary">
          <q-btn flat label="Cancel" v-close-popup />
          <q-btn flat label="Add address" v-close-popup />
        </q-card-actions>
      </q-card>
    </q-dialog>
  </div>
</template>

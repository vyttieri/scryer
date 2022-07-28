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

<q-btn label="Alert" color="primary" @click="alert = true" />

    <q-dialog ref="dialogRef" @hide="onDialogHide">
      <q-card class="q-dialog-plugin">
        <q-card-section>
        	<div>WHAT THE FUCK</div>
        </q-card-section>

        <q-card-actions align="right">
        	<q-btn color="primary" label="OK" @click="onOKClick" />
	        <q-btn color="primary" label="Cancel" @click="onCancelClick" />
        </q-card-actions>
      </q-card>
    </q-dialog>
      </q-toolbar>
    </q-header>
</template>

<script>
	import { useDialogPluginComponent } from 'quasar'

	export default {
		emits: [
			...useDialogPluginComponent.emits
		],
		setup() {
		    // REQUIRED; must be called inside of setup()
		    const { dialogRef, onDialogHide, onDialogOK, onDialogCancel } = useDialogPluginComponent()
		    // dialogRef      - Vue ref to be applied to QDialog
		    // onDialogHide   - Function to be used as handler for @hide on QDialog
		    // onDialogOK     - Function to call to settle dialog with "ok" outcome
		    //                    example: onDialogOK() - no payload
		    //                    example: onDialogOK({ /*.../* }) - with payload
		    // onDialogCancel - Function to call to settle dialog with "cancel" outcome

		    return {
		      // This is REQUIRED;
		      // Need to inject these (from useDialogPluginComponent() call)
		      // into the vue scope for the vue html template
		      dialogRef,
		      onDialogHide,

		      // other methods that we used in our vue html template;
		      // these are part of our example (so not required)
		      onOKClick () {
		      	console.log('sup')
		        // on OK, it is REQUIRED to
		        // call onDialogOK (with optional payload)
		        onDialogOK()
		        // or with payload: onDialogOK({ ... })
		        // ...and it will also hide the dialog automatically
		      },

		      // we can passthrough onDialogCancel directly
		      onCancelClick: onDialogCancel
		    }
		}
	}
</script>
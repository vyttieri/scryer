import { createApp } from 'vue'
import { createPinia } from 'pinia'
import { Quasar, Dialog } from 'quasar'
import VueGoogleMaps from '@fawmi/vue-google-maps'
import VueDraggable from 'vue-draggable'

import '@quasar/extras/material-icons/material-icons.css'
import 'quasar/src/css/index.sass'

import App from './App.vue'

const app = createApp(App)

app.use(VueGoogleMaps, {
	load: {
		key: import.meta.env.VITE_GOOGLE_MAPS_API_KEY
	}
})
app.use(createPinia())

app.use(Quasar, {
	plugins: {
		Dialog
	},
})

app.use(VueDraggable)

app.mount('#app')

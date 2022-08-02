import { createApp } from 'vue'
import { createPinia } from 'pinia'
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate'
import { Quasar } from 'quasar'
import VueGoogleMaps from '@fawmi/vue-google-maps'

import '@quasar/extras/material-icons/material-icons.css'
import 'quasar/src/css/index.sass'

import App from '@/App.vue'

const app = createApp(App)

app.use(VueGoogleMaps, {
	load: {
		key: import.meta.env.VITE_GOOGLE_MAPS_API_KEY
	}
})

const pinia = createPinia()
pinia.use(piniaPluginPersistedstate)
app.use(pinia)

app.use(Quasar)

app.mount('#app')

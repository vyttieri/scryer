import { createApp } from 'vue'
import { createPinia } from 'pinia'
import { Quasar } from 'quasar'

import App from './App.vue'
import './assets/main.css'

const app = createApp(App)

app.use(createPinia())

app.use(Quasar, {
	plugins: {},
})

app.mount('#app')

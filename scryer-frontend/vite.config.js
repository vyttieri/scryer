import { fileURLToPath, URL } from 'node:url'
import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import { quasar, transformAssetUrls } from '@quasar/vite-plugin'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue({
      template: { transformAssetUrls },
    }),
    quasar({
      sassVariables: 'src/quasar-variables.sass',
    })
  ],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    }
  },
  server: {
    proxy: {
      '/ping': {
        target: 'http://localhost:8080/',
        changeOrigin: true,
        secure: false,
      },
      '/login': {
        target: 'http://localhost:8080/',
        changeOrigin: true,
        secure: false,
      },
      '/logout': {
        target: 'http://localhost:8080/',
        changeOrigin: true,
        secure: false,
      },
      '/register': {
        target: 'http://localhost:8080/',
        changeOrigin: true,
        secure: false,
      },
      '/user/preferences': {
        target: 'http://localhost:8080/',
        changeOrigin: true,
        secure: false,
      },
    }
  }
})

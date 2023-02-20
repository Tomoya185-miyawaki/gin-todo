export default {
  // Global page headers: https://go.nuxtjs.dev/config-head
  head: {
    title: 'gin-todo-front',
    htmlAttrs: {
      lang: 'ja'
    },
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { hid: 'description', name: 'description', content: '' },
      { name: 'format-detection', content: 'telephone=no' }
    ],
    link: [
      { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }
    ]
  },

  // Global CSS: https://go.nuxtjs.dev/config-css
  css: [
  ],

  // Plugins to run before rendering page: https://go.nuxtjs.dev/config-plugins
  plugins: [
  ],

  // Auto import components: https://go.nuxtjs.dev/config-components
  components: true,

  // Modules for dev and build (recommended): https://go.nuxtjs.dev/config-modules
  buildModules: [
  ],

  // Modules: https://go.nuxtjs.dev/config-modules
  modules: [
  ],

  // Build Configuration: https://go.nuxtjs.dev/config-build
  build: {
  },

  // Server Configuration
  server: {
    port: 8000,
    host: '0.0.0.0'
  },

  // Env Configuration
  env: {
    baseUrl: process.env.BASE_URL || 'http://localhost:3000',
    baseAPIUrl: process.env.BASE_API_URL || 'http://localhost:3000/api/v1',
  }
}

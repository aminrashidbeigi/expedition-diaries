export default {
  // Global page headers: https://go.nuxtjs.dev/config-head
  head: {
    title: 'history-travelers',
    htmlAttrs: {
      lang: 'en'
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

  env: {
    baseUrl:
      process.env.NODE_ENV === 'dev'
        ? 'http://localhost:3000'
        : 'http://37.32.25.134:3000',
    baseAPI:
      process.env.NODE_ENV === 'dev'
        ? 'http://localhost:8080'
        : 'http://37.32.25.134:8080',
    compact: true
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
    // https://go.nuxtjs.dev/tailwindcss
    '@nuxtjs/tailwindcss',
  ],

  // Modules: https://go.nuxtjs.dev/config-modules
  modules: [
    // https://go.nuxtjs.dev/axios
    '@nuxtjs/axios',
    '@nuxtjs/vuetify',
  ],

  // Axios module configuration: https://go.nuxtjs.dev/config-axios
  axios: {
    // Workaround to avoid enforcing hard-coded localhost:3000: https://github.com/nuxt-community/axios-module/issues/308
    baseURL: '/',
  },

  // Build Configuration: https://go.nuxtjs.dev/config-build
  build: {
    transpile: ['@svg-maps/world']
  },

  server: {
    host: process.env.NODE_ENV === 'dev'
        ? '127.0.0.1'
        : '0.0.0.0'
  },

  loaders: [
    { test: /\.js$/, loader: 'babel', query: {compact: false} }
  ]
}
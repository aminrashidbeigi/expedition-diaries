import axios from 'axios';

export default {
  // Global page headers: https://go.nuxtjs.dev/config-head
  head: {
    title: 'Expedition Diaries',
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
        ? 'http://127.0.0.1:3000'
        : 'http://127.0.0.1:3000',
    baseAPI:
      process.env.NODE_ENV === 'dev'
        ? 'http://127.0.0.1:8080'
        : 'https://expedition-diaries.com/api',
    baseImageURL:
      process.env.NODE_ENV === 'dev'
        ? 'http://127.0.0.1:3000/images/'
        : 'https://expedition-diaries.com/images/',
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
    '@nuxtjs/sitemap',
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
        : '127.0.0.1'
  },

  loaders: [
    { test: /\.js$/, loader: 'babel', query: {compact: false} }
  ],
 
  publicRuntimeConfig: {
    baseImageURL:
      process.env.NODE_ENV === 'dev'
        ? 'http://127.0.0.1:3000/images/'
        : 'https://expedition-diaries.com/static/images/',
  },

  // sitemap: {
  //   hostname: 'https://expedition-diaries.com',
  //   gzip: true,
  //   generate: false,
  //   exclude: [],
  //   routes: async () => {
  //       const { data } = await axios.get('https://expedition-diaries.com/api/countries')
  //       return data.map((country) => `/countries/${country.Code}`)
  //   }
  // },
}

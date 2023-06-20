export default {
  target: 'static',

  head: {
    title: 'App Test',
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { hid: 'description', name: 'description', content: '' },
      { name: 'format-detection', content: 'telephone=no' }
    ],
    link: [
      { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }
    ],
    __dangerouslyDisableSanitizers: ['script'],
    script: [
      {
        hid: 'NewRelic',
        src: '/nr.js',
        defer: true,
        type: 'text/javascript',
      }
    ]
  },

  css: [
  ],

  plugins: [
  ],

  components: true,

  buildModules: [
    '@nuxtjs/eslint-module',
  ],

  modules: [
    'bootstrap-vue/nuxt',
    '@nuxtjs/axios',
    '@nuxtjs/auth-next',
    '@nuxtjs/pwa',
  ],

  axios: {
    baseURL: 'http://localhost:8080/api/',
  },

  auth: {
    strategies: {
      local: {
        token: {
          property: 'data.token',
        },
        user: {
          property: 'data',
          // autoFetch: true
        },
        endpoints: {
          login: { url: '/auth/login', method: 'post' },
          logout: { url: '/auth/logout', method: 'post' },
          user: { url: '/auth/profile', method: 'get' },
        },
      },
    },

    redirect: {
      login: '/login',
      logout: '/login',
      home: '/',
    },

    plugins: ['~/plugins/axios'],
  },

  ssr: false,

  pwa: {
    manifest: {
      lang: 'id'
    }
  },

  build: {
    crawler: false,
    fallback: '404.html'
  }
}

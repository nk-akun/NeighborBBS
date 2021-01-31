export default {
  // Global page headers: https://go.nuxtjs.dev/config-head
  head: {
    title: "site",
    htmlAttrs: {
      lang: "en"
    },
    meta: [
      { charset: "utf-8" },
      { name: "viewport", content: "width=device-width, initial-scale=1" },
      { hid: "description", name: "description", content: "" }
    ],
    link: [
      { rel: "icon", type: "image/x-icon", href: "/favicon.ico" },
      {
        rel: "alternate",
        type: "application/atom+xml",
        title: "文章",
        href: "/atom.xml"
      },
      {
        rel: "alternate",
        type: "application/atom+xml",
        title: "话题",
        href: "/topic_atom.xml"
      },
      {
        rel: "alternate",
        type: "application/atom+xml",
        title: "开源项目",
        href: "/project_atom.xml"
      },
      {
        rel: "stylesheet",
        href: "//cdn.staticfile.org/bulma/0.8.0/css/bulma.min.css"
      },
      {
        rel: "stylesheet",
        href: "//at.alicdn.com/t/font_1142441_1or22jfsge3.css"
      }
    ]
  },

  // Global CSS
  css: [
    "element-ui/lib/theme-chalk/index.css",
    { src: "~/assets/styles/main.scss", lang: "scss" }
  ],

  // Plugins to run before rendering page: https://go.nuxtjs.dev/config-plugins
  /*
   ** Plugins to load before mounting the App
   */
  plugins: [
    "@/plugins/element-ui",
    "@/plugins/filters",
    "@/plugins/axios",
    "@/plugins/bbs-go",
    { src: "@/plugins/infinite-scroll", ssr: false },
    { src: "@/plugins/vue-lazyload", ssr: false }
  ],

  // Auto import components: https://go.nuxtjs.dev/config-components
  components: true,

  // Modules for dev and build (recommended): https://go.nuxtjs.dev/config-modules
  buildModules: [],

  // Modules: https://go.nuxtjs.dev/config-modules
  modules: [
    // https://go.nuxtjs.dev/axios
    "@nuxtjs/axios"
  ],

  // Axios module configuration: https://go.nuxtjs.dev/config-axios
  axios: { proxy: true, credentials: false },
  proxy: {
    "/api/": "http://127.0.0.1:8080"
  },
  // Build Configuration: https://go.nuxtjs.dev/config-build
  build: {
    optimizeCSS: true,
    extractCSS: true,
    splitChunks: {
      layouts: true,
      pages: true,
      commons: true
    },
    postcss: {
      preset: {
        features: {
          customProperties: false
        }
      }
    },
    /*
     ** You can extend webpack config here
     */
    extend(config, ctx) {},
    transpile: [/^element-ui/]
  }
};

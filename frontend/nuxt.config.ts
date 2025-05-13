export default({
  compatibilityDate: "2024-11-01",
  devtools: { enabled: true },
  modules: ["@nuxtjs/tailwindcss", "@nuxtjs/leaflet"],
  css: ["~/assets/css/tailwind.css"],
  plugins: [
    { src: "~/plugins/leaflet.client.js", mode: "client"}, 
    '@/plugins/axios.js',
    '@/plugins/theme.js'
  ],
  ssr: false,
  runtimeConfig: {
    public: {
      apiBase: "http://localhost:8080",
    },
  },
  app: {
    head: {
      link: [
        { rel: 'stylesheet', href: 'https://fonts.googleapis.com/css2?family=Inter:wght@400;500;600&family=Orbitron:wght@400;600&display=swap' }
      ]
    }
  }
});

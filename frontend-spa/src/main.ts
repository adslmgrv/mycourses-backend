import { createApp } from "vue"
import "./style.css"
import App from "./App.vue"
import router from "./router"
import { createPinia } from "pinia"
import piniaPluginPersistedstate from "pinia-plugin-persistedstate"
import "./api/base"
import "./api/auth"

const pinia = createPinia()
pinia.use(piniaPluginPersistedstate)

const app = createApp(App)
app.use(pinia)
app.use(router)
app.mount("#app")

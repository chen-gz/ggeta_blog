import './assets/main.css'

import { createApp } from 'vue'
import App from './App.vue'
import router from "../router.js";

let app = createApp(App).use(router).mount('#app')
// app.mount('#app')

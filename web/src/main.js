import Vue from 'vue'
import axios from 'axios'

import 'normalize.css/normalize.css'

import '@/plugins/element'

import '@/styles/index.styl'

import App from './App.vue'
import router from './router'
import store from './store'

import '@/icons'
import '@/permission'

import './registerServiceWorker'

Vue.config.productionTip = false

// Set base URL to backend API service
const backendAddr = process.env.OFRIENDS_BACKEND_ADDRS || 'http://localhost:8080'
console.log(`OFRIENDS_BACKEND_ADDRS: ${backendAddr}`)
axios.defaults.baseURL = backendAddr
Vue.prototype.$http = axios

new Vue({
  router,
  store,
  el: '#app',
  render: h => h(App)
}).$mount('#app')

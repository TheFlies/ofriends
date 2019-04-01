import Vue from 'vue'

import 'normalize.css/normalize.css'

import '@/plugins/flag-icon'
import '@/plugins/element'

import '@/styles/index.styl'

import App from './App.vue'
import router from './router'
import store from './store'

import '@/icons'
import '@/permission'

import '@/plugins/axios'

import './registerServiceWorker'

Vue.config.productionTip = false

new Vue({
  router,
  store,
  el: '#app',
  render: h => h(App)
}).$mount('#app')

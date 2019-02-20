import Vue from 'vue'
import App from './App.vue'
import axios from 'axios'

import router from './router'
import store from './store'

import Vuetify from 'vuetify'
import 'vuetify/dist/vuetify.min.css'
import Moment from 'vue-moment'

Vue.use(Moment);

Vue.use(Vuetify)

axios.defaults.baseURL = '/'
axios.defaults.headers.get['Accepts'] = 'application/json'

const reqInterceptor = axios.interceptors.request.use(config => {
  console.log('Request Interceptor', config)
  return config
})
const resInterceptor = axios.interceptors.response.use(res => {
  console.log('Response Interceptor', res)
  return res
})

axios.interceptors.request.eject(reqInterceptor)
axios.interceptors.response.eject(resInterceptor)

Vue.filter('short', function (value, limit) {
  if (!value) return ''
  if (!limit) limit = 20
  value = value.toString()
  if (value.length > limit) {
    return value.substring(0, limit) + "...";
  }
})

new Vue({
  el: '#app',
  router,
  store,
  render: h => h(App)
})

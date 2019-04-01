import Vue from 'vue'
import axios from 'axios'
import router from '../router'
import store from '../store'

import { getToken } from '../utils/auth'

import { Message } from 'element-ui'

// Set base URL to backend API service
const backendAddr = process.env.OFRIENDS_BACKEND_ADDRS || 'http://localhost:8080'
console.log(`OFRIENDS_BACKEND_ADDRS: ${backendAddr}`)
axios.defaults.baseURL = backendAddr
// axios.defaults.withCredentials = true

// axios default
axios.interceptors.request.use(
  config => {
    if (store.getters.token) {
      config.headers['Authorization'] = `Bearer ${getToken()}`
    }
    return config
  },
  error => {
    console.log(error)
    Promise.reject(error)
  }
)

axios.interceptors.response.use(
  (response) => {
    return response
  },
  (error) => {
    if (error.response && error.response.status === 401) {
      Message({
        message: 'Unauthorize Access!!!',
        duration: 5000,
        showClose: true,
        type: 'error',
        onClose: () => {
          store.dispatch('FedLogOut').then(() => {
            router.replace(`/login?redirect=${router.currentRoute.path}`)
          })
        }
      })
    }
    if (error.response && error.response.status === 500) {
      Message({
        message: 'Server Error!!!. We don\'t expect this kind of error. Please report this to system administrator.',
        duration: 5000,
        showClose: true,
        type: 'error',
        onClose: () => {
          // store.dispatch('FedLogOut').then(() => {
          //   if (router.currentRoute.path !== '/login') {
          //     router.replace(`/login?redirect=${router.currentRoute.path}`)
          //   }
          // })
        }
      })
    }
    return Promise.reject(error)
  }
)

Vue.prototype.$http = axios

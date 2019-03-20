import Vue from 'vue'
import axios from 'axios'
import router from '../router'

import { Message } from 'element-ui'

// Set base URL to backend API service
const backendAddr = process.env.OFRIENDS_BACKEND_ADDRS || 'http://localhost:8080'
console.log(`OFRIENDS_BACKEND_ADDRS: ${backendAddr}`)
axios.defaults.baseURL = backendAddr

// axios default
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
          router.replace(`/login?redirect=${router.currentRoute.path}`)
        }
      })
    }
    return Promise.reject(error)
  }
)

Vue.prototype.$http = axios

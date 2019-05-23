import axios from 'axios'
import querystring from 'querystring'

export function login(username, password) {
  return axios.post('/login', querystring.stringify({ username, password }), {
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded'
    }
  })
}

export function register(user) {
  return axios.post('/register', user)
}

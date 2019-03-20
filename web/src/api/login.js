import axios from 'axios'
import querystring from 'querystring'

export function login(username, password) {
  return axios.post('/api/v1/login', querystring.stringify({ username, password }), {
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded'
    }
  })
}

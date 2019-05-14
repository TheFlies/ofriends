import axios from 'axios'

export function getAllUsers() {
  return axios.get('/api/v1/user/getAll')
}

export function updateUser(user) {
  return axios.put('/api/v1/user', user)
}

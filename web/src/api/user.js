import axios from 'axios'

export function getAllUsers() {
  return axios.get('/api/v1/user')
}

export function updateUser(user) {
  return axios.put('/api/v1/user/', user)
}

export function deleteUser(id) {
  return axios.delete('/api/v1/user/' + id)
}
import axios from 'axios'

export function getAllUsers() {
  return axios.get('/users')
}

export function updateUser(user) {
  return axios.put('/users', user)
}

export function deleteUser(id) {
  return axios.delete('/users/' + id)
}
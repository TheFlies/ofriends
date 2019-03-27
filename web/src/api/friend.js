import axios from 'axios'

export function getAllFriends() {
  return axios.get('/friends')
}

export function getFriendById(id) {
  return axios.get('/friends/' + id)
}

export function createFriend(friend) {
  return axios.post('/friends', friend)
}

export function updateFriend(friend) {
  return axios.put('/friends', friend)
}

export function deleteFriendById(id) {
  return axios.delete('/friends/' + id)
}


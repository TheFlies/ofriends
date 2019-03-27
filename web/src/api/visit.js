import axios from 'axios'

export function getAllVisitsByFriendID(friendId) {
  return axios.get('/visits/friend/' + friendId)
}

export function createVisit(visit) {
  return axios.post('/visits', visit)
}

export function updateVisit(visit) {
  return axios.put('/visits', visit)
}

export function deleteVisitById(id) {
  return axios.delete('/visits/' + id)
}


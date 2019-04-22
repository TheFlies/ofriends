import axios from 'axios'

export function getAllVisits() {
  return axios.get('/visits')
}

export function getAllVisitsByCustomerID(customerId) {
  return axios.get('/visits/customer/' + customerId)
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


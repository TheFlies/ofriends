import axios from 'axios'

export function getAllCusVisitAssocs() {
  return axios.get('/customers/associates')
}

export function getCusVisitAssocsByVisitID(visitID) {
  return axios.get('/customers/associates/visits/' + visitID)
}

export function createCusVisitAssoc(associate) {
  return axios.post('/customers/associates', associate)
}

export function modifyCusVisitAssocs(associate) {
  return axios.put('/customers/associates', associate)
}

export function deleteCusVisitAssocByID(associateID) {
  return axios.delete('/customers/associates/' + associateID)
}

export function deleteCusVisitAssocsByVisitID(visitID) {
  return axios.delete('/customers/associates/visits/' + visitID)
}


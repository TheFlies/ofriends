import axios from 'axios'

export function getAllAcVisitAssocs() {
  return axios.get('/activities/associates')
}

export function getActVisitAssocsByVisitID(visitID) {
  return axios.get('/activities/associates/visits/' + visitID)
}

export function createActVisitAssoc(associate) {
  return axios.post('/activities/associates', associate)
}

export function modifyActVisitAssocs(associate) {
  return axios.put('/activities/associates', associate)
}

export function deleteActVisitAssocByID(associateID) {
  return axios.delete('/activities/associates/' + associateID)
}

export function deleteActVisitAssocsByVisitID(visitID) {
  return axios.delete('/activities/associates/visits/' + visitID)
}


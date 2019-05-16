import axios from 'axios'

export function getAllAcitivityVisitAssociates() {
  return axios.get('/activities/associates')
}

export function getAcitivityVisitAssociatesByVisitID(visitID) {
  return axios.get('/activities/associates/visits/' + visitID)
}

export function createAcitivityVisitAssociate(associate) {
  return axios.post('/activities/associates', associate)
}

export function modifyAcitivityVisitAssociates(associate) {
  return axios.put('/activities/associates', associate)
}

export function deleteAcitivityVisitAssociateByID(associateID) {
  return axios.delete('/activities/associates/' + associateID)
}

export function deleteAcitivityVisitAssociatesByVisitID(visitID) {
  return axios.delete('/activities/associates/visits/' + visitID)
}


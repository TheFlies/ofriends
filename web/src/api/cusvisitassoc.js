import axios from 'axios'

export function getAllCusVisitAssocs() {
  return axios.get('/cusvisitassocs')
}

export function getCusVisitAssocsByVisitID(visitID) {
  return axios.get('/cusvisitassocs?visitid=' + visitID)
}

export function createCusVisitAssoc(associate) {
  return axios.post('/cusvisitassocs', associate)
}

export function modifyCusVisitAssocs(associate) {
  return axios.put('/cusvisitassocs', associate)
}

export function deleteCusVisitAssocByID(associateID) {
  return axios.delete('/cusvisitassocs/' + associateID)
}

export function deleteCusVisitAssocsByVisitID(visitID) {
  return axios.delete('/cusvisitassocs?visitid=' + visitID)
}


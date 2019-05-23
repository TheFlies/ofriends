import axios from 'axios'

export function getAllAcVisitAssocs() {
  return axios.get('/actvisitassocs')
}

export function getActVisitAssocsByVisitID(visitID) {
  return axios.get('/actvisitassocs?visitid=' + visitID)
}

export function createActVisitAssoc(associate) {
  return axios.post('/actvisitassocs', associate)
}

export function modifyActVisitAssocs(associate) {
  return axios.put('/actvisitassocs', associate)
}

export function deleteActVisitAssocByID(associateID) {
  return axios.delete('/actvisitassocs/' + associateID)
}

export function deleteActVisitAssocsByVisitID(visitID) {
  return axios.delete('/actvisitassocs?visitid=' + visitID)
}


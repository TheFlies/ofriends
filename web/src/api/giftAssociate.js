import axios from 'axios'

export function getAllGiftAssociates() {
  return axios.get('/gifts/associates')
}

export function getGiftAssociatesByVisitID(visitID) {
  return axios.get('/gifts/associates/visits/' + visitID)
}

export function createGiftAssociate(gift) {
  return axios.post('/gifts/associates', gift)
}

export function modifyGiftAssociates(gift) {
  return axios.put('/gifts/associates', gift)
}

export function deleteGiftAssociateById(giftID) {
  return axios.delete('/gifts/associates/' + giftID)
}


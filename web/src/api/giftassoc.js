import axios from 'axios'

export function getAllGiftAssocs() {
  return axios.get('/giftassociates')
}

export function getGiftAssocsByCusVistAssocID(assignID) {
  return axios.get('/giftassociates/' + assignID)
}

export function createGiftAssociate(gift) {
  return axios.post('/giftassociates', gift)
}

export function modifyGiftAssociates(gift) {
  return axios.put('/giftassociates', gift)
}

export function deleteGiftAssociateById(giftID) {
  return axios.delete('/giftassociates/' + giftID)
}

export function deleteGiftAssocsByCusVistAssocID(assignID) {
  return axios.delete('/giftassociates?cusvisitassocid=' + assignID)
}


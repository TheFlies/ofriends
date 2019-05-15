import axios from 'axios'

export function getAllGiftAssociates() {
  return axios.get('/gifts/associates')
}

export function getGiftAssociatesByVisitIDNCustomerID(visitID, customerID) {
  return axios.get('/gifts/associates/visits/' + visitID + '/' + customerID)
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

export function deleteGiftAssociatesByVisitIDNCustomerID(visitID, customerID) {
  return axios.delete('/gifts/associates/visits/' + visitID + '/' + customerID)
}


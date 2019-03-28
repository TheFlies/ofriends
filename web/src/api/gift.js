import axios from 'axios'

export function getAllGiftsByVisitID(visitId) {
  return axios.get('/gifts/visit/' + visitId)
}

export function getAllGifts() {
  return axios.get('/gifts')
}

export function getGiftById() {
  return axios.get('/gifts')
}

export function createGifts(gift) {
  return axios.post('/gifts', gift)
}

export function modifyGifts(gift) {
  return axios.put('/gifts', gift)
}

export function deleteGiftById(giftID) {
  return axios.delete('/gifts/' + giftID)
}


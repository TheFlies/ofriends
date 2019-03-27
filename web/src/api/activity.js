import axios from 'axios'

export function getAllActivitiesByVisitID(visitId) {
  return axios.get('/activity/visit/' + visitId)
}

export function createActivity(activity) {
  return axios.post('/activity', activity)
}

export function updateActivity(activity) {
  return axios.put('/activity', activity)
}

export function deleteActivityById(id) {
  return axios.delete('/activity/' + id)
}


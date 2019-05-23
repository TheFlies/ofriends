import axios from 'axios'

export function getAllActivities() {
  return axios.get('/activities')
}

export function getActivityByID(id) {
  return axios.get('/activities/' + id)
}
export function createActivity(activity) {
  return axios.post('/activities', activity)
}

export function updateActivity(activity) {
  return axios.put('/activities', activity)
}

export function deleteActivityByID(id) {
  return axios.delete('/activities/' + id)
}


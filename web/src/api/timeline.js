import axios from 'axios'

export function getTimelineByDay(dayTime) {
  return axios.get('/timeline/' + dayTime)
}
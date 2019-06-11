import axios from 'axios'

export function getTimelineByDay(dayTime) {
  return axios.get('/timeline?daytime=' + dayTime)
}

export function getTimelineInRange(starttime, endtime) {
  return axios.get('/timeline?starttime=' + starttime + '&endtime=' + endtime)
}

import moment from 'moment'

export function getHumanDate(timestamp) {
  return moment(timestamp).format('MMMM Do YYYY, h:mm a')
}

export function calculateTillNow(timestamp) {
  return moment(timestamp).fromNow()
}

export function todayIsAfter(timestamp) {
  return moment().isAfter(moment(timestamp))
}

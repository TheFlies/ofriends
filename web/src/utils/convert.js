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

export function getRoleName(priority) {
  let rolName
  if (priority === 2) {
    rolName = 'User'
  } else if (priority === 3) {
    rolName = 'Admin'
  } else {
    rolName = 'None'
  }
  return rolName
}

export function indexMethod(index) {
  return index + 1
}

import axios from 'axios'

function colorGeneration(dc) {
  switch (dc) {
    case '14':
      return '#f4b299'
    default:
      return '#b2b2b2'
  }
}

export function getAllFriends() {
  return axios.get('/friends')
}

export function getFriendById(id) {
  return axios.get('/friends/' + id)
}

export function createFriend(friend) {
  return axios.post('/friends', friend)
}

export function updateFriend(friend) {
  return axios.put('/friends', friend)
}

export function deleteFriendById(id) {
  return axios.delete('/friends/' + id)
}

export function transform(f) {
  return {
    title: `${f.title}. ${f.name}`,
    subtitle: f.position,
    color: colorGeneration(f.dc),
    dc: f.dc || 'TMA',
    project: f.project,
    company: f.company,
    country: f.country
  }
}

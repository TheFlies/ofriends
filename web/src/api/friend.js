import axios from 'axios'

function colorGeneration(dc) {
  switch (dc) {
    case '14':
      return '#f4b299'
    default:
      return '#b2b2b2'
  }
}

export function friends() {
  return axios.get('/friends')
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

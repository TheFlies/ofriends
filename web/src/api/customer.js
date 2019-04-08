import axios from 'axios'

function colorGeneration(dc) {
  switch (dc) {
    case '14':
      return '#f4b299'
    default:
      return '#b2b2b2'
  }
}

export function getAllCustomers() {
  return axios.get('/customers')
}

export function getCustomerById(id) {
  return axios.get('/customers/' + id)
}

export function createCustomer(customer) {
  return axios.post('/customers', customer)
}

export function updateCustomer(customer) {
  return axios.put('/customers', customer)
}

export function deleteCustomerById(id) {
  return axios.delete('/customers/' + id)
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

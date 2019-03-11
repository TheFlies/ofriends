import axios from 'axios'
import { API_FRIEND } from '@/const'

const state = {
  customers: [{
    name: 'dummy',
    title: 'Mr',
    position: 'architect lead',
    project: 'dummy project',
    arrive: '01-01-0001',
    depart: '31-12-9999'
  }],
  filter: {
    name: '',
    time: 'Current',
    arrive: new Date(),
    depart: new Date()
  }
}

const getters = {
  customers (state, getters, rootState) {
    return state.customers.filter((customer) => {
      return (customer.name.toLowerCase().indexOf(state.filter.name.toLowerCase()) !== -1)
    })
  }
}

const mutations = {
  setFilter (state, payload) {
    state.filter = payload
  },
  setCustomers (state, payload) {
    state.customers = payload
  }
}

const actions = {
  getCustomers (context) {
    axios({
      url: API_FRIEND
    }).then(r => {
      context.commit('setCustomers', r.data)
    })
  }
}

export default {
  namespaced: true,
  state,
  getters,
  mutations,
  actions
}

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
  mutations,
  actions
}

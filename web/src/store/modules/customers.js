const state = {
  customers: [],
  filter: {
    name: 'Johnnnnn',
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

export default {
  namespaced: true,
  state,
  mutations
}

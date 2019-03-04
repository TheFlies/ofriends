const state = {
  customers: [],
  filter: {
    name: 'John',
    time: 'Current',
    arrive: new Date(),
    depart: new Date()
  }
}

const mutations = {
  setFilter (state) {

  }
}

export default {
  namespaced: true,
  state,
  mutations
}

import Vue from 'vue'
import Vuex from 'vuex'
import customers from './modules/customers.js'
import gift from './modules/gift.js'

Vue.use(Vuex)

export default new Vuex.Store({
  // state: {
  //   user: {
  //     userId: null,
  //     user: null,
  //     userImage: null
  //   }
  // },
  // mutations: {

  // },
  // actions: {

  // }
  modules: {
    customers,
    gift
  }
})

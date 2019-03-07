import Vue from 'vue'
import Vuex from 'vuex'
import customers from './modules/customers.js'
import gift from './modules/gift.js'

Vue.use(Vuex)

export default new Vuex.Store({
  // strict: process.env.NODE_ENV !== 'production',
  modules: {
    customers,
    gift
  }
})

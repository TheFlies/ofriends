import Vue from 'vue'
import VueRouter from 'vue-router'

import HomePage from './components/home/home.vue'


Vue.use(VueRouter)

const routes = [
  { path: '/', component: HomePage },
]

export default new VueRouter({mode: 'history', routes})
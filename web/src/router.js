import Vue from 'vue'
import Router from 'vue-router'
import Gifts from './views/Gifts.vue'
import Hello from '@/components/Hello'

Vue.use(Router)

export default new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/',
      name: 'home',
      component: Hello
    },
    {
      path: '/friend',
      name: 'Friend',
      component: () => import('./views/friends/Add.vue')
    },
    {
      path: '/gifts',
      name: 'gifts',
      component: Gifts
    },
    {
      path: '/friends',
      name: 'friends',
      // route level code-splitting
      // this generates a separate chunk (about.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import(/* webpackChunkName: "about" */ './views/friends/Show.vue')
    }
  ]
})

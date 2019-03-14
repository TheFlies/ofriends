import Vue from 'vue'
import Router from 'vue-router'
import Gifts from './views/Gifts.vue'
import Hello from '@/components/Hello'

import Layout from '@/views/layout/Layout'

Vue.use(Router)

export default new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/',
      redirect: 'home',
      component: Layout,
      children: [
        {
          path: 'home',
          component: Hello,
          meta: {
            title: 'Home',
            icon: 'form'
          }
        }
      ]
    },
    {
      path: '/friend',
      component: Layout,
      children: [
        {
          path: '',
          name: 'Friend',
          component: () => import('./views/friends/Add.vue'),
          meta: {
            title: 'Friends',
            icon: 'user'
          }
        }
      ]
    },
    {
      path: '/gifts',
      component: Layout,
      children: [
        {
          path: '',
          name: 'gifts',
          component: Gifts,
          meta: {
            title: 'Gifts',
            icon: 'eye'
          }
        }
      ]
    // },
    // {
    //   path: '/about',
    //   name: 'about',
    //   // route level code-splitting
    //   // this generates a separate chunk (about.[hash].js) for this route
    //   // which is lazy-loaded when the route is visited.
    //   component: () => import(/* webpackChunkName: "about" */ './views/About.vue')
    }
  ]
})

import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

import Layout from '@/views/layout/Layout'

/**
* hidden: true                   if `hidden:true` will not show in the sidebar(default is false)
* alwaysShow: true               if set true, will always show the root menu, whatever its child routes length
*                                if not set alwaysShow, only more than one route under the children
*                                it will becomes nested mode, otherwise not show the root menu
* redirect: noredirect           if `redirect:noredirect` will no redirect in the breadcrumb
* name:'router-name'             the name is used by <keep-alive> (must set!!!)
* meta : {
    title: 'title'               the name show in subMenu and breadcrumb (recommend set)
    icon: 'svg-name'             the icon show in the sidebar
    breadcrumb: false            if false, the item will hidden in breadcrumb(default is true)
  }
**/
export const constantRouterMap = [
  { path: '/login', component: () => import('@/views/login/index'), hidden: true },
  { path: '/404', component: () => import('@/views/404'), hidden: true },

  {
    path: '/',
    redirect: 'home',
    component: Layout,
    meta: {
      title: 'Home',
      icon: 'dashboard'
    },
    children: [
      // {
      //   path: 'home',
      //   component: Hello,
      //   meta: {
      //     title: 'Hello',
      //     icon: 'form'
      //   }
      // },
      {
        path: 'home',
        component: () => import('./views/home/index.vue'),
        meta: {
          title: 'Timeline',
          icon: 'dashboard'
        }
      }
    ]
  },

  {
    path: '/customers',
    component: Layout,
    meta: {
      title: 'Customers',
      icon: 'user-setting'
    },
    children: [
      {
        path: '',
        name: 'Customers',
        component: () => import('./views/customers/Show.vue'),
        meta: {
          title: 'Customers',
          icon: 'user-setting'
        }
      },
      {
        path: '/customers/:id',
        name: 'Customer',
        component: () => import('./views/customers/ShowDetail.vue'),
        hidden: true,
        meta: {
          title: 'Customer Profile'
        }
      }
    ]
  },
  {
    path: '/visits',
    component: Layout,
    children: [
      {
        path: '',
        name: 'visit',
        component: () => import('./views/visits/Show.vue'),
        meta: {
          title: 'Visit',
          icon: 'visit'
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
        component: () => import('./views/Gifts.vue'),
        meta: {
          title: 'Gifts',
          icon: 'gift'
        }
      }
    ]
  },
  {
    path: '/activites',
    component: Layout,
    children: [
      {
        path: '',
        name: 'activity',
        component: () => import('./views/activity/Show.vue'),
        meta: {
          title: 'Activity',
          icon: 'hiking'
        }
      }
    ]
  },
  { path: '*', redirect: '/404', hidden: true }
]

export default new Router({
  mode: 'history',
  scrollBehavior: () => ({ y: 0 }),
  base: process.env.BASE_URL,
  routes: constantRouterMap
})

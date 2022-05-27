import Vue from 'vue'
import VueRouter, { RouteConfig } from 'vue-router'
import Catelog from '../views/Catelog.vue'

Vue.use(VueRouter)

  const routes: Array<RouteConfig> = [
  {
    path: '/',
    name: 'index',
    redirect: '/catelog'
  },
  {
    path: '/catelog',
    name: 'catelog',
    component: Catelog
  },
  {
    path: '/userAgent',
    name: 'userAgent',
    component: () => import("../views/useragent/index.vue")
  },
  {
    path: '/websocket',
    name: 'websocket',
    component: () => import("../views/websocket/index.vue")
  },
  {
    path: '/1',
    name: 'About',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/About.vue')
  }
]

const router = new VueRouter({
  base: process.env.BASE_URL,
  routes
})

export default router

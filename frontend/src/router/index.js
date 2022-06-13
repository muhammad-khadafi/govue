/**
 * Vue Router
 *
 * @library
 *
 * https://router.vuejs.org/en/
 */

// Lib imports
import Vue from 'vue'
import VueAnalytics from 'vue-analytics'
import Router from 'vue-router'
import Meta from 'vue-meta'

// Routes
import paths from './paths'

Vue.use(Router)

// Create a new router
const router = new Router({
  mode: 'history',
  // routes: paths.map(path => route(path.path, path.view, path.name)).concat([
  //   { path: '*', redirect: '/' }
  // ]),
  routes: paths,
  scrollBehavior (to, from, savedPosition) {
    if (savedPosition) {
      return savedPosition
    }
    if (to.hash) {
      return { selector: to.hash }
    }
    return { x: 0, y: 0 }
  },
})

Vue.use(Meta)

// Bootstrap Analytics
// Set in .env
// https://github.com/MatteoGabriele/vue-analytics
if (process.env.GOOGLE_ANALYTICS) {
  Vue.use(VueAnalytics, {
    id: process.env.GOOGLE_ANALYTICS,
    router,
    autoTracking: {
      page: process.env.NODE_ENV !== 'development',
    },
  })
}

router.beforeEach((to, from, next) => {
  console.log('URL guard called...')

  console.log(`to : ${to.path}, from ${from.path}`)
  const loggedIn = localStorage.getItem('user')
  console.log('loggedIn ~', loggedIn)
  const userData = JSON.parse(loggedIn)

  console.log('user data ~', userData)

  if (to.matched.some(record => record.meta.requiresAuth) && !loggedIn) {
    console.log('1')
    next('/')
  }

  if (to.matched.some(record => record.meta.loginForm) && loggedIn) {
    console.log('2')
    next('/home')
  }

  if (!to.matched.length) {
    console.log('3')
    next('/404')
  }

  console.log('4')

  // if (to.matched.length && loggedIn && to.fullPath !== '/403' && userData.menus[0].menuCode && !isPublic) {
  //   console.log("test5")
  //   console.log('Checking privilege menu...')
  //   let isAllow = false
  //   for (var i in userData.menus) {
  //     // console.log(userData.menus[i].menuCode + " - " + to.name)
  //     if (to.name === userData.menus[i].menuCode) {
  //       isAllow = true
  //     }
  //   }
  //   console.log('Is menu allowed? ' + isAllow)
  //   if (!isAllow) {
  //     console.log('Access rejected!', next())
  //     next('/403')
  //   }
  //   console.log('Is menu allowed? ' + isAllow)
  // }
  // console.log("test6")
  next()
})

export default router

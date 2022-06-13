import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import './plugins'
import vuetify from './plugins/vuetify'
import { sync } from 'vuex-router-sync'
import Vuelidate from 'vuelidate'
import { apiServer } from '../src/setup-axios'

sync(store, router)

Vue.config.productionTip = false
Vue.use(Vuelidate)
Vue.use(require('vue-shortkey'))

new Vue({
  router,
  store,
  created () {
    const userString = localStorage.getItem('user')
      console.log('user string ; ', userString)
    if (userString) {
      const userData = JSON.parse(userString)
      console.log('user data ; ', userData)
      this.$store.commit('SET_USER_DATA', userData)
    }
    apiServer.interceptors.response.use(
      response => response,
      error => {
        console.log('logging-error')
        if (error.response.status === 401) {
          this.$store.dispatch('logout')
        }
        return Promise.reject(error)
      }
    )
  },
  vuetify,
  render: h => h(App)
}).$mount('#app')

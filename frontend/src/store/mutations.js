// https://vuex.vuejs.org/en/mutations.html

import { apiServer } from '../setup-axios'
import router from '../router';

export default {
  SET_USER_DATA (state, user_response) {
    state.user = user_response
    localStorage.setItem('user', JSON.stringify(user_response))
    apiServer.defaults.headers.common.Authorization = `Bearer ${
      user_response.token
      }`
  },
  LOGOUT () {
    localStorage.removeItem('user')
    router.push('/logout')
  },
  SET_STATUS (state, payload) {
    state.status.snackbarColor = payload.snackbarColor
    state.status.snackbarText = payload.snackbarText
  },
}

import Vue from 'vue'
import Vuetify from 'vuetify/lib'

Vue.use(Vuetify)

export default new Vuetify({
  theme: {
    themes: {
      light: {
        primary: '#0E0E4B',
        secondary: '#D7F7F5',
        success: '#75CAC3',
        tertiary: '#2A6171',
        orange: '#F3D516',

        // for toolbar, card, and button color template

        red: '#861B20',
        green: '#02DC8A',
        blue: '#16B3AC',
        brown: '#E8A77D',
        header: '#FFFFFF',
        footer: '#FFFFFF',
        drawer: '#fff',
        icon_drawer: '#51A2A2',
        menu_drawer: '#16B3AC',
        submenu_drawer: '#D2E920',


      },
      dark: {
        primary: '#022C43',
        secondary: '#053F5E',
        tertiary: '#115173',
        orange: '#F6D637',


        // for toolbar and card color template

        red: '#115173',
        green: '#115173',
        blue: '#115173',
        brown: '#115173',
        // header:'#000000',
        // footer:'#000000',
        // drawer: '#000000',
        drawer: '#363636',
        icon_drawer: '#51A2A2',
      }
    }
  },
  icons: {
    iconfont: 'mdi',
    iconfont: 'mdiSvg'
  }
})

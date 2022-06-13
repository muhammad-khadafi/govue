<template>
  <v-navigation-drawer
    id="app-drawer"
    v-model="inputValue"
    app
    permanent
    floating
    color="drawer"
    :width="260"
  >
    <v-list>
      <router-link to="/home">
        <v-list-item
          class="logo-drawer"
        >
          <v-list-item-avatar style="border-radius: 0 !important;">
            <v-img
              src="../../assets/logo-GOVUE-drawer.png"
              contain
            />
          </v-list-item-avatar>
          <!-- <v-list-item-title class="headline text-uppercase font-weight-black">
            GOVUE
          </v-list-item-title> -->
        </v-list-item>
      </router-link>
      <!-- <template>
        <v-toolbar
          flat
          color="drawer"
        >
          <v-list-item-content class="user-drawer mt-4 text-center">
            <v-list-item-title class="mt-4 font-weight-bold">
              anda login sebagai {{ $store.state.user.user.first_name }} {{ $store.state.user.user.last_name }}
            </v-list-item-title>
            <v-list-item-subtitle class="mb-2 subtitle">
              <v-select
                v-model="roleSelected"
                :items="roles"
                item-text="role_name"
                item-value="id"
                @change="onChangeRole()"
              />
            </v-list-item-subtitle>
          </v-list-item-content>
        </v-toolbar>
      </template> -->

      <!-- <v-list-item class="item-admin mt-2">
              <v-list-item-content>
                <v-list-item-title class="subtitle admin">Dear {{ $store.state.user.name }}, you are login as Administrator</v-list-item-title>
              </v-list-item-content>
      </v-list-item> -->
    </v-list>

    <!-- <v-divider></v-divider> -->

    <v-list nav>
      <!-- Bug in Vuetify for first child of v-list not receiving proper border-radius -->

      <!-- <v-list-item
            v-for="(menu, i) in menus"
            :key="i"
            :to="menu.routerTo"
            active-class="primary white--text"
          >
            <v-list-item-action>
              <v-icon>{{ menu.icon }}</v-icon>
            </v-list-item-action>

            <v-list-item-title v-text="menu.menuName" />
      </v-list-item>-->

      <!-- <v-list-item v-for="(menu, i) in menus" :key="i">
            <v-list style="width:100%">
              <v-list-group>
                <template v-slot:activator>
                  <v-list-item-action>
                    <v-icon>{{menu[i,0].icon}}</v-icon>
                  </v-list-item-action>
                  <v-list-item-title v-text="menu[i,0].label" />
                </template>
                <v-list-item v-for="(sub, j) in menu[1]" :key="j" :to="sub.to">
                  <v-list-item-icon>
                    <v-icon v-text="sub.icon"></v-icon>
                  </v-list-item-icon>
                  <v-list-item-title v-text="sub.label"></v-list-item-title>
                </v-list-item>
              </v-list-group>
            </v-list>
      </v-list-item>-->

      <v-list-item
        v-for="(menu, i) in menus"
        :key="i"
        class="menu"
      >
        <v-list class="menu-drawer">
          <v-list-group active-class="">
            <template v-slot:activator>
              <!-- <v-list-item-action>
                <v-icon active-class="">{{menu.icon}}</v-icon>
              </v-list-item-action> -->
              <v-list-item-title
                v-tooltip.top-center="menu.menu_name"
                class="item-menu-drawer"
                v-text="menu.menu_name"
              />
            </template>
            <v-list-item
              class="pl-4 submenu"
              v-for="(sub, j) in menu.sub_menus"
              :key="j"
              :to="sub.router_to"
            >
              <!-- <v-list-item-icon>
                <v-icon v-text="sub.icon" />
              </v-list-item-icon> -->
              <v-list-item-title
                v-tooltip.top-center="sub.menu_name"
                v-text="sub.menu_name"
              />
            </v-list-item>
          </v-list-group>
        </v-list>
      </v-list-item>
    </v-list>

    <template v-slot:append>
      <div class="footer-drawer">
        <div class="avatar-footer-drawer">
          <v-img
              src="../../assets/avatar-footer-drawer.png"
              contain
            />
        </div>
        <div class="desc-footer-drawer">
          <p class="role">Halo, <span style="font-weight: 700;">{{ $store.state.user.username }}</span></p>
<!--          <v-select-->
<!--              v-model="roleSelected"-->
<!--              :items="roles"-->
<!--              item-text="role_name"-->
<!--              item-value="id"-->
<!--              @change="onChangeRole()"-->
<!--              class="username"-->
<!--            />-->
        </div>
      </div>

    </template>
  </v-navigation-drawer>
</template>

<script>
// Utilities
  import { mapMutations, mapState } from 'vuex'
  // import { authComputed } from '../../store/helpers.js'
  // import axios from 'axios'

  import { apiServer } from '../../setup-axios'

  export default {
    props: {
      opened: {
        type: Boolean,
        default: false,
      },
    },
    data: () => ({
      user: [],
      imageSrc: '',
      fileFoto: null,
      encodedFoto: '',
      organisasi: {
        nama: '',
      },
      institusi: {
        nama: '',
      },
      roles: [],
      roleSelected: null,
      mini: true,
      hidden: false,
      hover: true,
      links: [
        {
          to: '/dashboard',
          icon: 'mdi-view-dashboard',
          text: 'Dashboard',
        },
        {
          to: '/user-profile',
          icon: 'mdi-account',
          text: 'User Profile',
        },
        {
          to: '/crud',
          icon: 'mdi-table',
          text: 'Sample CRUD',
        },
        {
          to: '/file',
          icon: 'mdi-upload-outline',
          text: 'File',
        },
      ],
      menus: [],
      menus_: [
        [
          {
            label: 'Box',
            icon: 'mdi-alarm',
          },
          [
            {
              label: 'Management',
              icon: 'mdi-alarm',
              to: '/dashboard',
            },
            {
              label: 'Setting',
              icon: 'mdi-settings',
              to: '/user-profile',
            },
          ],
        ],
        [
          {
            label: 'Settings',
            icon: 'mdi-settings',
          },
          [
            {
              label: 'Management',
              icon: 'mdi-alarm',
              to: '/crud',
            },
            {
              label: 'Settingz',
              icon: 'mdi-settings',
              to: '/crud',
            },
          ],
        ],
      ],
    }),

    computed: {
      ...mapState('app', ['image', 'color']),
      inputValue: {
        get () {
          return this.$store.state.app.drawer
        },
        set (val) {
          this.setDrawer(val)
        },
      },
    },

    mounted () {
      // this.getMenu()
      // this.getRoles()
      // this.getInstitusi()
      // this.getFotoUser()
    },

    methods: {
      ...mapMutations('app', ['setDrawer', 'toggleDrawer']),

      getMenu () {
        apiServer
          .get('/user-management/menu/' + this.$store.state.user.roleActive.role_code)
          .then(response => {
            this.menus = response.data.result
            console.log('Load Menu Debug')
            console.log(this.menus)
          })
          .catch(error => {
            console.log('There was an error:', error.response) // Logs out the error
          })
      },
      getRoles () {
        this.roles = this.$store.state.user.roles
        console.log('debug roles : ', this.roles)
        this.roleSelected = this.$store.state.user.roleActive.id
        console.log('debug role selected : ' + this.roleSelected)
      },
      getOrganisasi () {
        if (this.$store.state.user.organizationId != null) {
          apiServer
            .get('/organisasi/' + this.$store.state.user.organizationId)
            .then(response => {
              this.organisasi = response.data
            })
            .catch(error => {
              console.log('there was an error:', error)
              this.showSnackbar('error', 'Error, Harap hubungi Administrator!')
            })
        }
      },
      getInstitusi () {
        this.institusi = this.$store.state.user.institusi
      },
      onChangeRole () {
        console.log('change role..')
        const user = localStorage.getItem('user')
        const userData = JSON.parse(user)
        for (var i in userData.roles) {
          if (userData.roles[i].id === this.roleSelected) {
            this.resetRoleMenuSession(userData.roles[i], userData)
          }
        }
      },
      resetRoleMenuSession (newRole, userData) {
        apiServer
          .get('/user-management/menu/' + newRole.role_code)
          .then(response => {
            userData.menu = response.data
            userData.roleActive = newRole
            console.log('new Role ', newRole)
            this.$store.commit('SET_USER_DATA', userData)
            console.log('new state ', this.$store.state.user)
            this.$router.replace('/home')
            location.reload()
          })
          .catch(error => {
            console.log('There was an error:', error.response) // Logs out the error
          })
      },
      getFotoUser () {
        apiServer
          .get('/userweb-profil/' + this.$store.state.user.userId)
          .then(response => {
            this.user = response.data
            if (this.user.fotoName !== null) {
              this.getFoto()
            } else {
              this.resetFoto()
            }
          })
          .catch(error => {
            console.log('There was an error: ', error.response) // Logs out the error
            this.showSnackbar('error', 'Error, please contact your administrator!')
          })
      },
      getFoto () {
        apiServer.get('/userweb-foto/' + this.$store.state.user.userId)
          .then(response => {
            if (response.data != null) {
              this.encodedFoto = response.data
              this.fileFoto = this.dataURLtoFile(
                this.encodedFoto,
                this.user.fotoName
              )
              const fileReader = new FileReader()
              fileReader.addEventListener('load', () => {
                this.imageSrc = fileReader.result
              })
              fileReader.readAsDataURL(this.fileFoto)
            }
          })
      },
      resetFoto () {
        this.imageSrc = ''
        this.fileFoto = null
      },
      dataURLtoFile (dataurl, filename) {
        var arr = dataurl.split(',')
        var mime = arr[0].match(/:(.*?);/)[1]
        var bstr = atob(arr[1])
        var n = bstr.length
        var u8arr = new Uint8Array(n)

        while (n--) {
          u8arr[n] = bstr.charCodeAt(n)
        }

        return new File([u8arr], filename, { type: mime })
      },
    },
  }
</script>

<style lang="scss">
// .menu {
//   margin-left: -10px;
// }

// .avatar-drawer {
//   width: 50%;
//   margin-left: 25%;
// }
// .admin-drawer {
//   margin-left: -8px;
//   margin-bottom: -30px;
// }

// .admin {
//   font-size: 75%;
// }
</style>

<template>
  <v-app-bar
    absolute
    flat
    color="header"
  >
    <v-spacer />

    <v-toolbar-items>
      <v-row
        align="center"
        class="mx-0"
      >
        
        <!-- <v-btn
          v-tooltip.top-center="userprofile"
          active-class="icon_drawer--text darken-4"
          to="/profil-pengguna"
          icon
        >
          <v-icon>mdi-account</v-icon>
        </v-btn> -->

        <v-btn
          v-tooltip.top-center="exit"
          active-class="icon_drawer--text darken-4"
          icon
          @click="logout()"
        >
          <v-icon>mdi-logout</v-icon>
        </v-btn>
      </v-row>
    </v-toolbar-items>
  </v-app-bar>
</template>

<script>
// Utilities
  import Vue from 'vue'
  import { VTooltip, VPopover, VClosePopover } from 'v-tooltip'
  import { mapMutations } from 'vuex'
  import { apiServer } from "../../setup-axios";

  Vue.directive('tooltip', VTooltip)
  Vue.directive('close-popover', VClosePopover)
  Vue.component('v-popover', VPopover)

  export default {
    data: () => ({
      notifications: [
        'Mike, John responded to your email',
        'You have 5 new tasks',
        "You're now a friend with Andrew",
        'Another Notification',
        'Another One',
      ],
      title: null,
      responsive: false,
      dashboard: 'Dashboard',
      userprofile: 'Profil Pengguna',
      exit: 'Keluar',
      userObject: {
        id: ""
      }
    }),

    watch: {
      $route (val) {
        this.title = val.name
      },
    },

    mounted () {
      this.onResponsiveInverted()
      window.addEventListener('resize', this.onResponsiveInverted)
    },
    beforeDestroy () {
      window.removeEventListener('resize', this.onResponsiveInverted)
    },

    methods: {
      ...mapMutations('app', ['setDrawer', 'toggleDrawer']),
      onClick () {
        this.setDrawer(!this.$store.state.app.drawer)
      },
      onResponsiveInverted () {
        if (window.innerWidth < 991) {
          this.responsive = true
        } else {
          this.responsive = false
        }
      },
      logout () {
        apiServer
            .post('/user-management/logout')
            .then(response => {
              console.log("Logout success");
              this.$store.dispatch('logout')
            })
            .catch(error => {
              console.log("There was an error:", error.response);
            });
      },
    },
  }
</script>

<style lang="scss">
/* Fix coming in v2.0.8 */
#core-app-bar {
  width: auto;
}

#core-app-bar a {
  text-decoration: none;
}

.tooltip {
  // font-family: Roboto !important;
  display: block !important;
  z-index: 10000;
}

.tooltip .tooltip-inner {
  background: rgb(88, 88, 88);
  color: white;
  border-radius: 0.25rem;
  padding: 5px 10px 4px;
}

.tooltip .tooltip-arrow {
  width: 0;
  height: 0;
  border-style: solid;
  position: absolute;
  margin: 5px;
  border-color: rgb(88, 88, 88);
  z-index: 1;
}

.tooltip[x-placement^="top"] {
  margin-bottom: 5px;
}

.tooltip[x-placement^="top"] .tooltip-arrow {
  border-width: 5px 5px 0 5px;
  border-left-color: transparent !important;
  border-right-color: transparent !important;
  border-bottom-color: transparent !important;
  bottom: -5px;
  left: calc(50% - 5px);
  margin-top: 0;
  margin-bottom: 0;
}

.tooltip[x-placement^="bottom"] {
  margin-top: 5px;
}

.tooltip[x-placement^="bottom"] .tooltip-arrow {
  border-width: 0 5px 5px 5px;
  border-left-color: transparent !important;
  border-right-color: transparent !important;
  border-top-color: transparent !important;
  top: -5px;
  left: calc(50% - 5px);
  margin-top: 0;
  margin-bottom: 0;
}

.tooltip[x-placement^="right"] {
  margin-left: 5px;
}

.tooltip[x-placement^="right"] .tooltip-arrow {
  border-width: 5px 5px 5px 0;
  border-left-color: transparent !important;
  border-top-color: transparent !important;
  border-bottom-color: transparent !important;
  left: -5px;
  top: calc(50% - 5px);
  margin-left: 0;
  margin-right: 0;
}

.tooltip[x-placement^="left"] {
  margin-right: 5px;
}

.tooltip[x-placement^="left"] .tooltip-arrow {
  border-width: 5px 0 5px 5px;
  border-top-color: transparent !important;
  border-right-color: transparent !important;
  border-bottom-color: transparent !important;
  right: -5px;
  top: calc(50% - 5px);
  margin-left: 0;
  margin-right: 0;
}

.tooltip.popover .popover-inner {
  background: #f9f9f9;
  color: black;
  padding: 24px;
  border-radius: 5px;
  box-shadow: 0 5px 30px rgba(black, .1);
}

.tooltip.popover .popover-arrow {
  border-color: #f9f9f9;
}

.tooltip[aria-hidden='true'] {
  visibility: hidden;
  opacity: 0;
  transition: opacity .15s, visibility .15s;
}

.tooltip[aria-hidden='false'] {
  visibility: visible;
  opacity: 1;
  transition: opacity .15s;
}

</style>

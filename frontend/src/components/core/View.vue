<template>
  <!-- <v-content> -->
    <div id="core-view">
      <v-fade-transition mode="out-in">
        <router-view />
      </v-fade-transition>
    </div>
  <!-- </v-content> -->
</template>

<script>
  export default {
    name: 'CoreView',

    metaInfo () {
      return {
        title: 'GOVUE'
      }
    }
  }
</script>

<style lang="scss">
  // #core-view{
  //   margin-top: 5%;
  //   margin-bottom: 5%;
  // }
</style>

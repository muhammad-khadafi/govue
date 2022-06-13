<template>
  <v-container
    class="text-center"
    style="padding: 0;"
  >
    <v-row align="center">
      <v-col style="padding: 0;">
      <v-img
              src="../assets/error-404.png"
              contain
              class="error-404"
            />
        <h1 class="display-2 primary--text">Error 403 - Forbidden</h1>

        <p>Anda tidak memiliki akses untuk membuka halaman ini</p>

        <v-btn
          :to="`/`"
          color="primary"
          outlined
        >
          Kembali
        </v-btn>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
  export default {
    name: 'FourOhFour',
    props: {
      lang: {
        type: String,
        default: 'en',
      },
    },
    created () {
      console.log('403 Forbidden')
      this.$ssrContext && this.$ssrContext.res.status(404)
    },
  }
</script>

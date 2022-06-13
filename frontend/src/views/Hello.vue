<template>
  <v-container fluid class="home">
    <p class="desc">
      {{ message }}
    </p>
  </v-container>
</template>

<script>
import { apiServer } from "../setup-axios";

export default {
  data: () => ({
    message: "",
  }),

  created() {
    this.getMessage()
  },

  methods: {
    getMessage() {
      this.isLoading = true;
      apiServer
        .get("/hello-auth")
        .then((response) => {
          // const responseFormatted = JSON.parse(response)
          this.message = response.data.Message;
        })
        .catch((error) => {
          console.log("There was an error:", error.response); // Logs out the error
        });
    }
  },
};
</script>

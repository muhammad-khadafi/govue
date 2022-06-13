<template>
  <div class="login-wrapper">
    <v-row>
      <v-col cols="12" lg="6" md="6">
        <v-content class="login-content">
          <v-img src="../assets/logo-GOVUE.png" class="login-logo" />
          <p class="login-desc">
            Please enter your username and password to login and access GOVUE
          </p>
          <v-form>
            <div class="username">
              <p class="label">Username</p>
              <v-text-field
                id="mxUsername"
                class=""
                ref="username"
                v-model="username"
                :class="{ mx_empty_filed: !username }"
                :rules="[() => !!username || 'This field is required']"
                placeholder="Username"
                required
                outlined
              />
            </div>
            <div class="password">
              <p class="label">Password</p>
              <v-text-field
                id="mxPassword"
                class=""
                ref="password"
                v-model="password"
                :class="{ mx_empty_filed: !password }"
                :rules="[() => !!password || 'This field is required']"
                :append-icon="showPassword ? 'mdi-eye-off' : 'mdi-eye'"
                :type="showPassword ? 'text' : 'password'"
                placeholder="Password"
                counter
                required
                outlined
                @keydown.enter="login"
                @click:append="showPassword = !showPassword"
              />
            </div>
          </v-form>

          <div class="recaptcha">
            <template>
              <div>
                <vue-recaptcha
                  ref="recaptcha"
                  @verify="onVerify"
                  @expired="onExpired"
                  :sitekey="sitekey"
                >
                </vue-recaptcha>
              </div>
            </template>
          </div>
          <div class="text-center">
            <template>
              <!-- <vue-recaptcha sitekey="6LeGVgEVAAAAAI0_xZPUIysLsSsCk3IAPUFkuZP8"> -->
              <v-btn
                class="button-login"
                :loading="loading"
                @click="login()"
                >Login
              </v-btn>
            </template>
          </div>
          <div class="text-center pb-10">
            <template>
              <v-btn class="text-capitalize" text @click="forgotPassword()">Forgot Password? </v-btn>
            </template>
          </div>
          <v-snackbar v-model="snackbar" :color="color" :top="true">
            {{ errorMessages }}
            <v-btn dark @click="snackbar = false">Close</v-btn>
          </v-snackbar>
        </v-content>
      </v-col>
      <v-col class="desktop" cols="12" lg="6" md="6">
        <v-img src="../assets/dashboard.gif" class="img-dashboard" />
      </v-col>
    </v-row>
  </div>
</template>

<script>
import VueRecaptcha from "vue-recaptcha";
import { apiServer } from "../setup-axios";

export default {
  components: {
    VueRecaptcha,
  },
  data: function () {
    return {
      username: "",
      password: "",
      sitekey: "6LdXtKQZAAAAABq6xKaAhPbyYrjQHQnvx9owa0vy",
      errorMessages: "Incorrect login info",
      snackbar: false,
      color: "primary",
      showPassword: false,
      loading: false,
    };
  },

  methods: {
    onVerify: function (response) {
      this.recaptcha = response;
      console.log("Verify: " + response);
    },
    onExpired: function () {
      console.log("Expired");
    },
    resetRecaptcha() {
      this.$refs.recaptcha.reset(); // Direct call reset method
    },
    forgotPassword() {
      this.$router.push("/forgot-password");
    },
    login: function () {
      if (
        this.username &&
        this.password
        // && this.recaptcha
      ) {
        let username = this.username;
        let password = this.password;
        this.loading = true;
        console.log("debug : " + username);
        this.$store
          .dispatch("login", { username, password })
          .then(() => this.$router.push("/home"))
          .catch((err) => {
            this.loading = false;
            console.log(err);
            this.snackbar = true;
          });
      }
    },
  },
  metaInfo() {
    return {
      title: "GOVUE | Login",
    };
  },
};
</script>

<style scoped>

</style>

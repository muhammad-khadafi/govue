<template>
  <v-content class="forgot-password-content">
    <!-- <v-img src="../assets/bg.png" class="background-forgot-password"/> -->
      <v-container
        dark
        fill-height
        fluid
        class="pattern-bg"
      >
      <v-layout
        align-center
        justify-center>
        <v-flex
          xs10
          sm8
          md4>
          <v-card flat
            class="elevation-5 card-forgot-password">
              <p class="title-forgot-password text-uppercase text-center font-weight-bold">
                Reset Password
              </p>

            <v-card-text>
              <v-form ref="form" v-if="isFormValid" >
                <v-text-field
                  class="mt-2 ml-10 mr-10"
                  ref="email"
                  v-model="defaultItem.email"
                  :rules="[() => !!defaultItem.email || 'This field is required']"
                  prepend-inner-icon="mdi-email"
                  label="Email"
                  placeholder="Email"
                  required
                  outlined
                />
              </v-form>
              <p v-if="!isFormValid" class="text-center">
                Silahkan cek email anda, anda akan menerima username beserta password baru
              </p>

              <p v-if="!isFormValid" class="text-center">
                <v-btn color="success"
                    v-if="!isFormValid"
                    @click="back()"
                    >Kembali ke Halaman Login
                </v-btn>
              </p>

            </v-card-text>

            <div class="recaptcha">
              <template>
                  <div>
                    <vue-recaptcha
                      ref="recaptcha"
                      @verify="onVerify"
                      @expired="onExpired"
                      :sitekey="sitekey">
                    </vue-recaptcha>
                  </div>
              </template>
            </div>

            <!-- {{defaultItem}} -->

            <p class="text-center pb-10 pt-2">
              <template>
                <v-btn
                  rounded
                  x-large
                  class="button-forgot-password font-weight-bold"
                  v-if="isFormValid"
                  :loading="loadingForgot"
                  @click="save()"
                  >Reset Password
                </v-btn>
              </template>
            </p>
            <v-snackbar
              v-model="snackbar"
              :color="color"
              :top='true'
            >
              {{ errorMessages }}
              <v-btn
                dark
                @click="snackbar = false"
              >
                Close
              </v-btn>
            </v-snackbar>
          </v-card>
        </v-flex>
      </v-layout>
    </v-container>
  </v-content>
</template>

<script>
  import VueRecaptcha from 'vue-recaptcha'
  import { apiServer } from '../setup-axios'
  import { siteKey } from '../config/env'

  export default {
    components: { VueRecaptcha },
    data: function () {
      return {
        password: '',
        sitekey: siteKey,
        errorMessages: 'Incorrect login info',
        snackbar: false,
        color: 'primary',
        showPassword: false,
        isFormValid: true,
        loadingForgot: false,
        defaultItem: {
          email: '',
          isValid: '',
        },
      }
    },

    // Sends action to Vuex that will log you in and redirect to the dash otherwise, error
    methods: {
      onVerify: function (response) {
        this.recaptcha = response
        console.log('Verify: ' + response)
      },
      onExpired: function () {
        console.log('Expired')
      },
      resetRecaptcha () {
        this.$refs.recaptcha.reset() // Direct call reset method
      },
      login: function () {
        if (this.username && this.password && this.recaptcha) {
          const username = this.username
          const password = this.password
          this.loading = true
          console.log(`[DEBUG] Username: ${username}`)
          this.$store.dispatch('login', { username, password })
            .then(() => this.$router.push('/home'))
            .catch(err => {
              this.loading = false
              console.log(err)
              this.snackbar = true
            })
        }
      },
      back () {
        this.$router.push('/')
      },
      save () {
        apiServer
          .put('/userweb/forgot-password', this.defaultItem)
          .then(response => {
            console.log(response.data)
            if (response.data.isValid) {
              this.loadingForgot = true
              this.isFormValid = false
              this.errorMessages = 'reset password success'
              this.snackbar = true
            } else {
              this.loadingForgot = false
              this.errorMessages = 'email invalid'
              this.snackbar = true
            }
          })
          .catch(error => {
            console.log('There was an error: ', error.response) // Logs out the error
          })
      },
    },
    metaInfo () {
      return {
        title: 'Lupa Password',
      }
    },
  }
</script>

<style scoped>
  .background-forgot-password{
    position: fixed;
    height: 100%;
    left: 0;
    bottom: 0;
    z-index: 1;
  }

  .recaptcha {
    margin-top: -4%;
    margin-left: 25%;
  }

  .pattern-bg {
    background-color: #ffffff;
    background-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 2000 1500'%3E%3Cdefs%3E%3Crect stroke='%23ffffff' stroke-width='0.33' width='1' height='1' id='s'/%3E%3Cpattern id='a' width='3' height='3' patternUnits='userSpaceOnUse' patternTransform='rotate(360 1000 750) scale(14.4) translate(-930.56 -697.92)'%3E%3Cuse fill='%23fafafa' href='%23s' y='2'/%3E%3Cuse fill='%23fafafa' href='%23s' x='1' y='2'/%3E%3Cuse fill='%23f5f5f5' href='%23s' x='2' y='2'/%3E%3Cuse fill='%23f5f5f5' href='%23s'/%3E%3Cuse fill='%23f0f0f0' href='%23s' x='2'/%3E%3Cuse fill='%23f0f0f0' href='%23s' x='1' y='1'/%3E%3C/pattern%3E%3Cpattern id='b' width='7' height='11' patternUnits='userSpaceOnUse' patternTransform='rotate(360 1000 750) scale(14.4) translate(-930.56 -697.92)'%3E%3Cg fill='%23ebebeb'%3E%3Cuse href='%23s'/%3E%3Cuse href='%23s' y='5' /%3E%3Cuse href='%23s' x='1' y='10'/%3E%3Cuse href='%23s' x='2' y='1'/%3E%3Cuse href='%23s' x='2' y='4'/%3E%3Cuse href='%23s' x='3' y='8'/%3E%3Cuse href='%23s' x='4' y='3'/%3E%3Cuse href='%23s' x='4' y='7'/%3E%3Cuse href='%23s' x='5' y='2'/%3E%3Cuse href='%23s' x='5' y='6'/%3E%3Cuse href='%23s' x='6' y='9'/%3E%3C/g%3E%3C/pattern%3E%3Cpattern id='h' width='5' height='13' patternUnits='userSpaceOnUse' patternTransform='rotate(360 1000 750) scale(14.4) translate(-930.56 -697.92)'%3E%3Cg fill='%23ebebeb'%3E%3Cuse href='%23s' y='5'/%3E%3Cuse href='%23s' y='8'/%3E%3Cuse href='%23s' x='1' y='1'/%3E%3Cuse href='%23s' x='1' y='9'/%3E%3Cuse href='%23s' x='1' y='12'/%3E%3Cuse href='%23s' x='2'/%3E%3Cuse href='%23s' x='2' y='4'/%3E%3Cuse href='%23s' x='3' y='2'/%3E%3Cuse href='%23s' x='3' y='6'/%3E%3Cuse href='%23s' x='3' y='11'/%3E%3Cuse href='%23s' x='4' y='3'/%3E%3Cuse href='%23s' x='4' y='7'/%3E%3Cuse href='%23s' x='4' y='10'/%3E%3C/g%3E%3C/pattern%3E%3Cpattern id='c' width='17' height='13' patternUnits='userSpaceOnUse' patternTransform='rotate(360 1000 750) scale(14.4) translate(-930.56 -697.92)'%3E%3Cg fill='%23e5e5e5'%3E%3Cuse href='%23s' y='11'/%3E%3Cuse href='%23s' x='2' y='9'/%3E%3Cuse href='%23s' x='5' y='12'/%3E%3Cuse href='%23s' x='9' y='4'/%3E%3Cuse href='%23s' x='12' y='1'/%3E%3Cuse href='%23s' x='16' y='6'/%3E%3C/g%3E%3C/pattern%3E%3Cpattern id='d' width='19' height='17' patternUnits='userSpaceOnUse' patternTransform='rotate(360 1000 750) scale(14.4) translate(-930.56 -697.92)'%3E%3Cg fill='%23ffffff'%3E%3Cuse href='%23s' y='9'/%3E%3Cuse href='%23s' x='16' y='5'/%3E%3Cuse href='%23s' x='14' y='2'/%3E%3Cuse href='%23s' x='11' y='11'/%3E%3Cuse href='%23s' x='6' y='14'/%3E%3C/g%3E%3Cg fill='%23e0e0e0'%3E%3Cuse href='%23s' x='3' y='13'/%3E%3Cuse href='%23s' x='9' y='7'/%3E%3Cuse href='%23s' x='13' y='10'/%3E%3Cuse href='%23s' x='15' y='4'/%3E%3Cuse href='%23s' x='18' y='1'/%3E%3C/g%3E%3C/pattern%3E%3Cpattern id='e' width='47' height='53' patternUnits='userSpaceOnUse' patternTransform='rotate(360 1000 750) scale(14.4) translate(-930.56 -697.92)'%3E%3Cg fill='%23F60'%3E%3Cuse href='%23s' x='2' y='5'/%3E%3Cuse href='%23s' x='16' y='38'/%3E%3Cuse href='%23s' x='46' y='42'/%3E%3Cuse href='%23s' x='29' y='20'/%3E%3C/g%3E%3C/pattern%3E%3Cpattern id='f' width='59' height='71' patternUnits='userSpaceOnUse' patternTransform='rotate(360 1000 750) scale(14.4) translate(-930.56 -697.92)'%3E%3Cg fill='%23F60'%3E%3Cuse href='%23s' x='33' y='13'/%3E%3Cuse href='%23s' x='27' y='54'/%3E%3Cuse href='%23s' x='55' y='55'/%3E%3C/g%3E%3C/pattern%3E%3Cpattern id='g' width='139' height='97' patternUnits='userSpaceOnUse' patternTransform='rotate(360 1000 750) scale(14.4) translate(-930.56 -697.92)'%3E%3Cg fill='%23F60'%3E%3Cuse href='%23s' x='11' y='8'/%3E%3Cuse href='%23s' x='51' y='13'/%3E%3Cuse href='%23s' x='17' y='73'/%3E%3Cuse href='%23s' x='99' y='57'/%3E%3C/g%3E%3C/pattern%3E%3C/defs%3E%3Crect fill='url(%23a)' width='100%25' height='100%25'/%3E%3Crect fill='url(%23b)' width='100%25' height='100%25'/%3E%3Crect fill='url(%23h)' width='100%25' height='100%25'/%3E%3Crect fill='url(%23c)' width='100%25' height='100%25'/%3E%3Crect fill='url(%23d)' width='100%25' height='100%25'/%3E%3Crect fill='url(%23e)' width='100%25' height='100%25'/%3E%3Crect fill='url(%23f)' width='100%25' height='100%25'/%3E%3Crect fill='url(%23g)' width='100%25' height='100%25'/%3E%3C/svg%3E");
    background-attachment: fixed;
    background-size: cover;
  }

  .container-forgot-password{
    margin-left: 25%;
    margin-top: -3%;
  }

  .card-forgot-password{
    z-index: 1;
  }

  .title-forgot-password{
    padding-top: 5%;
    margin-bottom: -5px;
    color: #393939;
    font-size: 30px;
  }

  .button-forgot-password{
    width: 60%;
    height: 50px;
    border-radius: 25px;
    font-size: 1.0rem;
    outline: none;
    border:none;
    text-transform: uppercase;
    background-image: linear-gradient(to right, #a30000, #a30000,  #a30000);
    cursor: pointer;
    color: #fff;
    font-family: 'Poppins',sans-serif;
    background-size: 200%;
    transition: .5s;
  }

  .button-forgot-password:hover{
    background-position: right;
  }

  @media screen and (max-width: 1200px){
    .title-forgot-password{
      padding-top: 15%;
      margin-bottom: -10px;
      font-size: 35px;
    }
  }

  @media screen and (max-width: 1050px){
    .container-forgot-password{
      margin-left: 0%;
    }
    .background-forgot-password{
      display: none;
    }
    .forgot-password-content{
      background-image: linear-gradient(to right, #50B2AF,#51989B, #537F87);
    }
  }

  @media screen and (max-width: 960px){
    .container-forgot-password{
      margin-left: 0%;
    }
  }

  @media screen and (max-width: 600px){
    .container-forgot-password{
      margin-left: 0%;
    }
    .background-forgot-password{
      display: none;
    }
  }

    @media screen and (max-width: 400px){
    .container-forgot-password{
      margin-left: 0%;
    }
    .background-forgot-password{
      display: none;
      }
  }
</style>

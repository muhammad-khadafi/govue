<template>
  <v-container fluid>
    <material-card
      color="dark"
      title="Daftar Peran per Pengguna"
    >
      <v-data-table
        :headers="headers"
        :items="userRole"
        class="elevation-1"
        :loading="isLoading"
        loading-text="Loading... Please wait"
        :search="search"
      >
        <template v-slot:top>

          <v-toolbar flat class="mt-1">
            <v-text-field
              v-model="search"
              append-icon="mdi-magnify"
              label="Cari"
              single-line
              hide-details
              class="search"
            ></v-text-field>
            <div class="flex-grow-1"></div>
            <v-dialog v-model="dialog" max-width="500px">
              <!--template v-slot:activator="{ on }">
                <v-btn color="blue" dark class="mb-2" v-on="on">Tambah</v-btn>
              </template-->

              <v-card>
                <v-card-title>
                  <span class="headline">{{ formTitle }}</span>
                </v-card-title>

                <v-card-text>
                  <v-container>
                    <v-row>
                      <v-col cols="12">
                        <v-select
                          :items="user"
                          item-text="username"
                          item-value="id"
                          label="Username"
                          clearable
                          v-model="editedItem.idUserweb"
                          required
                          @input="$v.editedItem.idUserweb.$touch()"
                          @blur="$v.editedItem.idUserweb.$touch()"
                          :error-messages="userwebErrors"
                          :disabled="editedIndex > -1"
                        ></v-select>
                      </v-col>

                      <v-col cols="12">
                        <v-select
                          :items="role"
                          multiple
                          chips
                          item-text="roleName"
                          item-value="id"
                          label="Peran"
                          v-model="editedItem.idRoles"
                          required
                          @input="$v.editedItem.idRoles.$touch()"
                          @blur="$v.editedItem.idRoles.$touch()"
                          :error-messages="rolesErrors"
                        ></v-select>
                      </v-col>
                    </v-row>
                  </v-container>
                </v-card-text>

                <v-card-actions>
                  <div class="flex-grow-1"></div>
                  <v-btn color="blue darken-1" text @click="close">Batal</v-btn>
                  <v-btn color="blue darken-1" text @click="save" :disabled="saveLock">Simpan</v-btn>
                </v-card-actions>
              </v-card>
            </v-dialog>
          </v-toolbar>
        </template>

        <template v-slot:item.isActive="{ item }">
          <v-icon>{{ item.isActive ? "mdi-check" : "mdi-minus" }}</v-icon>
        </template>

        <template v-slot:item.action="{ item }">
          <v-icon small class="mr-2" @click="editItem(item)"
            v-tooltip.top-center="editTip">mdi-pencil</v-icon>
          <v-icon small @click="deleteItem(item)">{{ item.rolesStr=="" ? "" : "mdi-delete" }} </v-icon>
        </template>
        <template v-slot:no-data></template>
      </v-data-table>

      <v-snackbar v-model="snackbar" :color="snackbarColor" :timeout="2000" :top="true">
        {{ snackbarText }}
        <v-btn dark text @click="snackbar = false">Close</v-btn>
      </v-snackbar>
    </material-card>
  </v-container>
</template>

<script>
import { apiServer } from "../setup-axios";
import { required } from "vuelidate/lib/validators";

//import { VTooltip } from "v-tooltip";

export default {
  data: () => ({
    saveLock: false,
    editTip: "Edit",
    deleteTip: "Hapus",
    search: "",
    snackbar: false,
    snackbarText: "",
    snackbarColor: "",
    isLoading: true,
    menu: false,
    dialog: false,
    headers: [
      { text: "Username", value: "username" },
      { text: "Peran", value: "rolesStr" },
      { text: "Aksi", value: "action", sortable: false }
    ],
    editedIndex: -1,
    editedItem: {
      idUserweb: "",
      idRoles: [],
      username: ""
    },
    defaultItem: {
      idUserweb: "",
      idRoles: [],
      username: ""
    },
    userRole: [],
    role: [],
    user: []
  }),

  validations: {
    editedItem: {
      idUserweb: {
        required
      },
      idRoles: {
        required
      }
    }
  },

  computed: {
    formTitle() {
      return this.editedIndex === -1 ? "Tambah Peran per Pengguna" : "Edit Peran per Pengguna";
    },
    userwebErrors() {
      const errors = [];
      if (!this.$v.editedItem.idUserweb.$dirty) return errors;
      !this.$v.editedItem.idUserweb.required &&
        errors.push("Username harus di isi");
      return errors;
    },
    rolesErrors() {
      const errors = [];
      if (!this.$v.editedItem.idRoles.$dirty) return errors;
      !this.$v.editedItem.idRoles.required && errors.push("Peran harus di isi");
      return errors;
    }
  },

  watch: {
    dialog(val) {
      val || this.close();
    }
  },

  created() {
    this.getData();
    this.getUser();
    this.getRole();
  },

  methods: {
    getData() {
      this.isLoading = true;
      apiServer
        .get(
          "/user-role-all/" +
            this.$store.state.user.organizationId +
            "/" +
            this.$store.state.user.role.roleCode
        )
        .then(response => {
          this.userRole = response.data;
          console.log("getData() " + response.data); // For now, logs out the response
          this.isLoading = false;
        })
        .catch(error => {
          console.log("There was an error:", error.response); // Logs out the error
          this.isLoading = false;
        });
    },
    getUser() {
    apiServer
        .get(
          "/userweb-all-for-user-role/" +
            this.$store.state.user.organizationId +
            "/" +
            this.$store.state.user.role.roleCode
        )
        .then(response => {
          this.user = response.data;
          console.log("getData() ", response.data);
        })
        .catch(error => {
          console.log("There was an error:", error.response);
        });
    },
    getRole() {
      apiServer
        .get("/role")
        .then(response => {
          this.role = response.data;
          console.log(response.data);
        })
        .catch(error => {
          console.log("There was an error:", error.response); // Logs out the error
        });
    },
    postData() {
      apiServer
        .post("/user-role", this.editedItem)
        .then(response => {
          this.getData();
          this.showSnackbar("success", "Menambahkan data berhasil");
        })
        .catch(error => {
          console.log("There was an error:", error.response);
          this.showSnackbar(
            "error",
            "Error, silakan kontak administrator anda"
          );
        });
    },
    deleteData(idUserweb) {
      apiServer
        .delete("/user-role/" + idUserweb + "/" + this.$store.state.user.name)
        .then(response => {
          this.getData();
          this.showSnackbar("success", "Menghapus data berhasil");
        })
        .catch(error => {
          console.log("There was an error:", error.response);
          if (error.response.data != null) {
            this.showSnackbar("error", error.response.data);
          } else {
            this.showSnackbar(
              "error",
              "Error, silakan kontak administrator anda!"
            );
          }
        });
    },
    updateData() {
      apiServer
        .put("/user-role", this.editedItem)
        .then(response => {
          this.getData();
          this.showSnackbar("success", "Mengedit data berhasil");
        })
        .catch(error => {
          console.log("There was an error:", error.response);
          this.showSnackbar(
            "error",
            "Error, silakan kontak administrator anda!"
          );
        });
    },
    editItem(item) {
      this.editedIndex = this.userRole.indexOf(item);
      this.editedItem = Object.assign({}, item);
      this.dialog = true;
    },

    deleteItem(item) {
      confirm("Are you sure you want to delete this item?") &&
        this.deleteData(item.idUserweb);
    },
    close() {
      this.$v.$reset();
      this.dialog = false;
      this.saveLock = false;
      setTimeout(() => {
        this.editedItem = Object.assign({}, this.defaultItem);
        this.editedIndex = -1;
      }, 300);
    },
    save() {
      this.$v.$touch();
      if (!this.$v.$invalid) {
        this.editedItem.username = this.$store.state.user.name;
        this.saveLock = true;
        if (this.editedIndex > -1) {
          this.updateData();
        } else {
          this.postData();
        }
        this.close();
      }
    },
    showSnackbar(color, message) {
      console.log("snack " + color + " - " + message);
      this.snackbar = true;
      this.snackbarText = message;
      this.snackbarColor = color;
    }
  }
};
</script>

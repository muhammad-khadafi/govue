<template>
  <v-container fluid>
    <material-card
      color="primary"
      title="Daftar Peran"
      class="card-wrapper"
    >
      <v-data-table
        :headers="headers"
        :items="role"
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
              label="Cari..."
              single-line
              hide-details
              class="search"
            ></v-text-field>
            <div class="flex-grow-1"></div>
            <v-dialog v-model="dialog" max-width="500px">
              <template v-slot:activator="{ on }">
                <v-btn color="primary" dark class="mb-2" v-on="on">Tambah</v-btn>
              </template>

              <v-card class="card-modal">
                <v-card-title>
                  <span class="headline">{{ formTitle }}</span>
                </v-card-title>

                <v-card-text>
                  <v-form>
                    <div class="form-input">
                      <p class="label">Kode Peran</p>
                      <v-text-field
                        v-model="editedItem.role_code"
                        placeholder="Kode Peran"
                        required
                        @input="$v.editedItem.role_code.$touch()"
                        @blur="$v.editedItem.role_code.$touch()"
                        :error-messages="roleCodeErrors"
                        outlined
                      />
                    </div>
                    <div class="form-input">
                      <p class="label">Nama Peran</p>
                      <v-text-field
                        v-model="editedItem.role_name"
                        placeholder="Nama Peran"
                        required
                        @input="$v.editedItem.role_name.$touch()"
                        @blur="$v.editedItem.role_name.$touch()"
                        :error-messages="roleNameErrors"
                        outlined
                      />
                    </div>
                  </v-form>
                  <!-- <v-container>
                    <v-row>
                      <v-col cols="12">
                        <v-text-field
                          v-model="editedItem.role_code"
                          label="Kode Peran"
                          required
                          @input="$v.editedItem.role_code.$touch()"
                          @blur="$v.editedItem.role_code.$touch()"
                          :error-messages="roleCodeErrors"
                        ></v-text-field>
                      </v-col>

                      <v-col cols="12">
                        <v-text-field
                          v-model="editedItem.role_name"
                          label="Nama Peran"
                          required
                          @input="$v.editedItem.role_name.$touch()"
                          @blur="$v.editedItem.role_name.$touch()"
                          :error-messages="roleNameErrors"
                        ></v-text-field>
                      </v-col>
                    </v-row>
                  </v-container> -->
                </v-card-text>

                <v-card-actions>
                  <div class="flex-grow-1"></div>
                  <v-btn color="red darken-1" text @click="close">Cancel</v-btn>
                  <v-btn color="blue darken-1" text @click="save">Save</v-btn>
                </v-card-actions>
              </v-card>
            </v-dialog>
          </v-toolbar>
        </template>

        <template v-slot:item.isActive="{ item }">
          <v-icon>{{ item.isActive ? "mdi-check" : "mdi-minus" }}</v-icon>
        </template>

        <template v-slot:item.action="{ item }">
          <v-icon small class="mr-2" @click="editItem(item)">mdi-pencil</v-icon>
          <v-icon small @click="deleteItem(item)">mdi-delete</v-icon>
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
import { required, maxLength } from "vuelidate/lib/validators";

export default {
  data: () => ({
    isTest: true,
    search: "",
    snackbar: false,
    snackbarText: "",
    snackbarColor: "",
    isLoading: true,
    menu: false,
    dialog: false,
    role: [],
    headers: [
      { text: "Kode Peran", value: "role_code" },
      { text: "Nama Peran", value: "role_name" },
      { text: "", value: "action", sortable: false }
    ],
    editedIndex: -1,
    editedItem: {
      role_code: "",
      role_name: "",
    },
    defaultItem: {
      role_code: "",
      role_name: "",
    },
    roleCodeEdited: "",
    roleNameEdited: "",
    roleIdEdited: ""
  }),

  validations: {
    editedItem: {
      role_code: {
        required,
        maxLength: maxLength(50),
        async isUnique(value) {
          if (value === "") return true;

          if (value === this.roleCodeEdited) return true;

          for (var index in this.role) {
            if (
              this.role[index].role_code.toUpperCase().trim() ===
              value.toUpperCase().trim()
            ) {
              return false;
            }
          }
          return true;
        }
      },
      role_name: {
        required,
        maxLength: maxLength(100),
        async isUnique(value) {
          if (value === "") return true;

          if (value === this.roleNameEdited) return true;

          for (var index in this.role) {
            if (
              this.role[index].role_name.toUpperCase().trim() ===
              value.toUpperCase().trim()
            ) {
              return false;
            }
          }
          return true;
        }
      }
    }
  },

  computed: {
    formTitle() {
      return this.editedIndex === -1 ? "Tambah Peran" : "Ubah Peran";
    },
    roleCodeErrors() {
      const errors = [];
      if (!this.$v.editedItem.role_code.$dirty) return errors;
      !this.$v.editedItem.role_code.required &&
        errors.push("Kode peran harus di isi");
      !this.$v.editedItem.role_code.maxLength &&
        errors.push("Kode peran panjang maksimal 50 karakter");
      !this.$v.editedItem.role_code.isUnique &&
        errors.push("Kode peran sudah terdaftar");
      return errors;
    },
    roleNameErrors() {
      const errors = [];
      if (!this.$v.editedItem.role_name.$dirty) return errors;
      !this.$v.editedItem.role_name.required &&
        errors.push("Nama peran harus di isi");
      !this.$v.editedItem.role_name.maxLength &&
        errors.push("Nama peran panjang maksimal 100 karakter");
      !this.$v.editedItem.role_name.isUnique &&
        errors.push("Nama peran sudah terdaftar");
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
  },

  methods: {
    getData() {
      this.isLoading = true;
      console.log(this.isLoading);
      apiServer
        .get("/user-management/roles")
        .then(response => {
          this.role = response.data;
          console.log(response.data);
          this.isLoading = false;
        })
        .catch(error => {
          console.log("There was an error:", error.response); // Logs out the error
          this.isLoading = false;
        });
    },
    postData() {
      apiServer
        .post("/user-management/roles", this.editedItem)
        .then(response => {
          this.getData();
          this.showSnackbar("success", "Add new data successful");
        })
        .catch(error => {
          console.log("There was an error:", error.response);
          this.showSnackbar(
            "error",
            "Error, please contact your administrator!"
          );
        });
    },
    deleteData(id) {
      apiServer
        .delete("/user-management/roles/" + id )
        .then(response => {
          this.getData();
          this.showSnackbar("success", "Delete data successful");
        })
        .catch(error => {
          console.log("There was an error:", error.response);
          this.showSnackbar("error", error.response.data);
        });
    },
    updateData() {
      apiServer
        .put("/user-management/roles/" + this.roleIdEdited, this.editedItem)
        .then(response => {
          this.getData();
          this.showSnackbar("success", "Update data successful");
        })
        .catch(error => {
          console.log("There was an error:", error.response);
          this.showSnackbar(
            "error",
            "Error, please contact your administrator!"
          );
        });
    },
    editItem(item) {
      this.roleCodeEdited = item.role_code;
      this.roleNameEdited = item.role_name;
      this.roleIdEdited = item.id;
      this.editedIndex = this.role.indexOf(item);
      this.editedItem = Object.assign({}, item);
      this.dialog = true;
    },

    deleteItem(item) {
      confirm("Are you sure you want to delete this item?") &&
        this.deleteData(item.id);
    },
    close() {
      this.$v.$reset();
      this.roleCodeEdited = "";
      this.roleNameEdited = "";
      this.dialog = false;
      setTimeout(() => {
        this.editedItem = Object.assign({}, this.defaultItem);
        this.editedIndex = -1;
      }, 300);
    },
    save() {
      this.$v.$touch();
      if (!(this.$v.$anyError)) {
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
    },

    changeBool() {
        this.isTest = !this.isTest;
      }
  }
};
</script>

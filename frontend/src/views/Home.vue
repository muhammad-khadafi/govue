<template>
  <v-container fluid class="home">
    <v-img src="../assets/dashboard.gif" contain />
    <p class="desc">
      Welcome, <span style="font-weight: 700;">{{ $store.state.user.username }}</span>
    </p>
  </v-container>
</template>

<script>
import { apiServer, fileUrl } from "../setup-axios";
import moment from "moment";

export default {
  data: () => ({
    daftarDokumen: [],
    headers: [
      { text: "Judul", width: "35%", value: "judul" },
      { text: "Institusi", value: "namaOrganisasi" },
      { text: "Tahun", align: "center", value: "tahunTerbit" },
      { text: "Aksi", value: "action", align: "center", sortable: false },
    ],
    isLoading: false,
    search: "",
    search2: "",
    snackbar: false,
    snackbarText: "",
    snackbarColor: "",
    showKerjaAktif: false,
    showUserLogin: false,
    kerjasamaAktif: [],
    headersKerjasama: [
      { text: "Judul kerjasama", value: "judul" },
      { text: "Pihak Pertama", value: "pihakPertamaStr" },
      { text: "Pihak Kedua", value: "pihakKeduaStr" },
      { text: "Tanggal Mulai", value: "tanggalMulai" },
      { text: "Tanggal Selesai", value: "tanggalSelesai" },
      { text: "", value: "action", sortable: false },
    ],
    kerjasamaTidakLengkap: [],
    headersKerjasamaTidakLengkap: [
      { text: "Judul kerjasama", value: "judul" },
      { text: "Pihak Pertama", value: "pihakPertamaStr" },
      { text: "Pihak Kedua", value: "pihakKeduaStr" },
      { text: "Tanggal Mulai", value: "tanggalMulai" },
      { text: "Tanggal Selesai", value: "tanggalSelesai" },
      { text: "Penanggung Jawab", value: "jmlPenanggungJawab" },
      { text: "Dokumen", value: "jmlDokumen" },
      { text: "", value: "action", sortable: false },
    ],
    penggunaLogin: [],
    headersPenggunaLogin: [
      { text: "Username", value: "username" },
      { text: "Waktu Login Terakhir", value: "lastLoginStr" },
    ],
  }),

  created() {},

  methods: {
    formatDate(value) {
      return moment(value).format("DD MMM YYYY");
    },
    editItem(item) {
      this.$store.dispatch("setKerjasama", item);
      this.$router.push("/kerjasama-edit");
    },
    getDokumen() {
      this.isLoading = true;
      apiServer
        .get("/dokumen-terkini/" + this.$store.state.user.organizationId)
        .then((response) => {
          this.isLoading = false;
          this.daftarDokumen = response.data;
        })
        .catch((error) => {
          console.log("There was an error:", error.response); // Logs out the error
          this.isLoading = false;
        });
    },
    downloadDokumen(item) {
      apiServer
        .post("/dokumen-organisasi/download/" + item.id)
        .then((response) => {
          const ticket = response.data.ticket;
          window.location =
            fileUrl +
            "/download/dokumen-organisasi/" +
            item.id +
            "?ticket=" +
            ticket;
        })
        .catch((error) => {
          console.log("There was an error:", error.response); // Logs out the error
          if (error.response.status === 500) {
            this.showSnackbar("error", "File tidak ditemukan.");
          } else {
            this.showSnackbar("error", "Error, Harap hubungi Administrator!");
          }
        });
    },
    showSnackbar(color, message) {
      this.snackbar = true;
      this.snackbarColor = color;
      this.snackbarText = message;
    },
    getPenggunaLogin() {
      this.isLoading = true;
      console.log(this.isLoading);
      apiServer
        .get("/user-login-list")
        .then((response) => {
          this.penggunaLogin = response.data;
          console.log("getData() " + JSON.stringify(response.data)); // For now, logs out the response
          this.isLoading = false;
        })
        .catch((error) => {
          console.log("There was an error:", error.response); // Logs out the error
          this.isLoading = false;
        });
    },
    getKerjasamaAktif() {
      this.isLoading = true;
      console.log(this.isLoading);
      apiServer
        .get("/kerjasama-aktif")
        .then((response) => {
          this.kerjasamaAktif = response.data;
          console.log("getData() " + response.data); // For now, logs out the response
          this.isLoading = false;
        })
        .catch((error) => {
          console.log("There was an error:", error.response); // Logs out the error
          this.isLoading = false;
        });
    },
    getKerjasamaTidakLengkap() {
      this.isLoading = true;
      console.log(this.isLoading);
      apiServer
        .get("/kerjasama-tidak-lengkap")
        .then((response) => {
          this.kerjasamaTidakLengkap = response.data;
          this.isLoading = false;
        })
        .catch((error) => {
          console.log("There was an error:", error.response); // Logs out the error
          this.isLoading = false;
        });
    },
  },
};
</script>

<style lang="scss">
</style>

<template>
  <v-container fluid>
    <div class="mt-5 ml-3">
      <h1>Sandbox</h1>
      <h3>area bermain, selamat melakukan eksplorasi</h3>
      <v-form class="mt-3">
        <div class="form-input">
          <p class="label">Nama Peran</p>
          <input
            v-shortkey.focus="['ctrl', 'i']"
            class="custom-input"
            type="text"
            placeholder="Input Biasa"
          />
          <v-text-field
            placeholder="v-text-fiels"
            v-shortkey.focus="['c']"
            required
            outlined
            class="mt-3"
            @focus="onFocus"
          />
        </div>
      </v-form>
    </div>
  </v-container>
</template>

<script>
export default {
  methods: {
    onFocus() {
      this.value = "v-shortkey.focus='['alt', 'b']'";
    },
  },
};
</script>

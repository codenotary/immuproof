<template>
  <v-app>
    <v-main>
      <main-page></main-page>
    </v-main>
  </v-app>
</template>

<script>
import MainPage from "./components/pages/MainPage";

export default {
  components: {
      MainPage,
  }
};
</script>

<style lang="scss">
#app {
   background-color: #FAF9F8;
}
</style>

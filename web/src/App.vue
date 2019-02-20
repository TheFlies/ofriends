<template>
  <v-app id="inspire">
    <app-header></app-header>
    <v-content>
      <router-view></router-view>
    </v-content>
    <app-footer></app-footer>
  </v-app>
</template>

<script>
  import Header from './components/header/header.vue'
  import Footer from './components/footer/footer.vue'
  export default {
    components: {
        'app-header': Header,
        'app-footer': Footer
    },
  }
</script>

<style>
  @import '../assets/css/main.css'
</style>

<template>
  <div id="app">
    <Header/>
    <Customers msg="Welcome to Your Vue.js App"/>
    <Footer/>
  </div>
</template>

<script>
import Customers from './components/customer/Customers.vue'
import Header from './components/common/Header.vue'
import Footer from './components/common/Footer.vue'

export default {
  name: 'app',
  components: {
    Customers,
    Header,
    Footer
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
}
body {
  margin: 0
}
* {
  box-sizing: border-box
}
</style>

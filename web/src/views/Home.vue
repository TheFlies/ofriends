<template>
  <div class="home">
    <img alt="Vue logo" src="../assets/logo.png">
    <AddCustomer msg="Welcome d tosss Your Vue.js App" />
  </div>
</template>

<script>
// @ is an alias to /src
import AddCustomer from '@/components/customers/CustomerAdd.vue'

export default {
  name: 'Home',
  components: {
    AddCustomer
  }
}
</script>

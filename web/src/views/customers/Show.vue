<template>
  <div class="show-customer">
    <CustomerList />
  </div>
</template>

<script>
// @ is an alias to /src
import CustomerList from '@/components/customers/CustomerList.vue'

export default {
  name: 'Show',
  components: {
    CustomerList
  }
}
</script>

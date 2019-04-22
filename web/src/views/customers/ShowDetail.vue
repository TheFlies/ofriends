<template lang="pug">
  el-row
    el-col(:span="24")
      el-row
        customer-show-detail
      el-row
        visit-list-by-customer
</template>

<script>
// @ is an alias to /src
import CustomerShowDetail from '@/components/customers/CustomerShowDetail.vue'
import VisitListByCustomer from '@/components/visits/VisitListByCustomer.vue'

export default {
  name: 'Show',
  components: {
    CustomerShowDetail,
    VisitListByCustomer
  }
}
</script>

<template lang="pug">
el-container
  el-header(style="width: 100%; margin:auto")
    el-button(type="primary" icon="el-icon-plus" plain="" style="float:right" @click="isVisibleAdd = !isVisibleAdd")
      | New customer
  el-main
    el-table(v-loading="loading"
      :data="tableData.filter(data => !search || data.name.toLowerCase().includes(search.toLowerCase()))"
      style="width: 100%; margin:auto")
      el-table-column(type="index" :index="indexMethod")
      el-table-column(label="Name" width="120" sortable="" prop="name")
        template(slot-scope="scope")
          router-link.link-type(:to="{ name: 'Customer', params: { id: scope.row.id }}")
            span {{ scope.row.name }}
      el-table-column(label="Position" width="120" sortable="" prop="position")
      el-table-column(label="Project" width="120" sortable="" prop="project")
      el-table-column(label="Age" width="120" sortable="" prop="age")
      el-table-column(label="Company" width="120" sortable="" prop="company")
      el-table-column(label="Country" width="120" sortable="" prop="country")
        template(slot-scope="scope")
          span {{ ($countries.find(it => it.code === scope.row.country) || { name: null }).name }}
      el-table-column(label="City" width="120" sortable="" prop="city")
      el-table-column(label="Food Note" width="120" prop="foodNote")
      el-table-column(label="Family Note" width="120" prop="familyNote")
      el-table-column(label="Next Visit Note" width="120" prop="nextVisitNote")
      el-table-column(align="right")
        template(slot="header" slot-scope="scope")
          el-input(v-model="search" size="mini" placeholder="Type to search")
        template(slot-scope="scope")
          el-button(size="mini" @click="customer = scope.row; isVisibleUpdate = !isVisibleUpdate")
            | Edit
          el-button(size="mini" type="danger" @click="isVisibleDelete = !isVisibleDelete; scopeCustomer = scope; customerName = scope.row.name")
            | Delete
  edit-customer(:is-visible-update.sync="isVisibleUpdate" :customer.sync="customer" @isUpdateCustomer="handleUpdate")
  delete-customer(:is-visible-delete.sync="isVisibleDelete" :customer-name.sync="customerName" @isDeleteCustomer="handleDelete")
  add-customer(:is-visible-add.sync="isVisibleAdd" @isAddCustomer="handleAdd")
</template>

<script>
import EditCustomer from '@/components/customers/CustomerUpdate.vue'
import DeleteCustomer from '@/components/customers/CustomerDelete.vue'
import AddCustomer from '@/components/customers/CustomerAdd.vue'
import {
  getAllCustomers,
  createCustomer,
  updateCustomer,
  deleteCustomerById
} from '@/api/customer'

export default {
  name: 'ListCustomers',
  components: {
    EditCustomer,
    DeleteCustomer,
    AddCustomer
  },
  data() {
    return {
      tableData: [],
      isVisibleAdd: false,
      isVisibleUpdate: false,
      isVisibleDelete: false,
      search: '',
      loading: true,
      customer: {},
      customerName: '',
      scopeCustomer: {}
    }
  },
  mounted() {
    getAllCustomers()
      .then(resp => {
        if (resp.data != null) {
          this.tableData = resp.data
        }
        this.loading = false
      })
      .catch(err => {
        console.log(err)
      })
  },
  methods: {
    handleAdd: function(isAddCustomer, customer) {
      if (isAddCustomer) {
        this.loading = true
        customer.age = parseInt(customer.age, 10)
        createCustomer(customer)
          .then(resp => {
            this.$notify({
              title: 'Success',
              message: 'Update successfully!',
              type: 'success'
            })
            customer.id = resp.data.id
            this.tableData.splice(0, 0, customer)
          })
          .catch(err => {
            console.log(err)
            this.$notify.error({
              title: 'Error',
              message: err
            })
          })
        this.loading = false
      }
    },
    handleUpdate: function(isUpdateCustomer) {
      if (isUpdateCustomer) {
        this.loading = true
        this.customer.age = parseInt(this.customer.age, 10)
        updateCustomer(this.customer)
          .then(resp => {
            console.log(resp.data)
            this.$notify({
              title: 'Success',
              message: 'Update successfully!',
              type: 'success'
            })
          })
          .catch(err => {
            console.log(err)
            this.$notify.error({
              title: 'Error',
              message: err
            })
          })
        this.loading = false
      }
    },
    handleDelete: function(isDeleteCustomer) {
      if (isDeleteCustomer) {
        this.loading = true
        deleteCustomerById(this.scopeCustomer.row.id)
          .then(resp => {
            this.tableData.splice(this.scopeCustomer.$index, 1)
            this.$notify({
              title: 'Success',
              message: 'Delete successfully!',
              type: 'success'
            })
          })
          .catch(err => {
            console.log(err)
            this.$notify.error({
              title: 'Error',
              message: err
            })
          })
      }
      this.loading = false
    },
    indexMethod(index) {
      return index * 1
    }
  }
}
</script>
<style scoped lang="stylus">
.link-type {
  color: #1989fa;
  text-decoration: underline;
}
</style>

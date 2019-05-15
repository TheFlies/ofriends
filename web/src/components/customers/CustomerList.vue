<template lang="pug">
      <el-card class="box-card">
        <div slot="header" class="clearfix">
          //- <span>Customer</span>
          <el-tooltip class="item" effect="dark" content="Add customer" placement="right-start" >
            <el-button type="primary" icon="el-icon-plus" plain @click="isVisibleAdd = !isVisibleAdd">
              | New customer
            </el-button> 
          </el-tooltip>          
        </div>
        <div class="text item">
          <el-table v-loading="loading" :data="tableData.filter(data => !search || data.name.toLowerCase().includes(search.toLowerCase()))" style="width: 100%; margin:auto">
            <el-table-column label="Name" width="120" sortable="" prop="name"/>
            <el-table-column label="Position" width="120" sortable="" prop="position"/>
            <el-table-column label="Project" width="120" sortable="" prop="project"/>
            <el-table-column label="Age" width="120" sortable="" prop="age"/>
            <el-table-column label="Company" width="120" sortable="" prop="company"/>
            <el-table-column label="Country" width="120" sortable="" prop="country"/>
            <el-table-column label="City" width="120" sortable="" prop="city"/>
            <el-table-column label="Pre-approved visa" width="120">
                <template slot-scope="scope">
                  <el-checkbox v-model="scope.row.preApproveVisa" />
                </template>
            </el-table-column>
            <el-table-column label="Passport Info" width="120" prop="passportInfo"/>
            <el-table-column label="Food Note" width="120" prop="foodNote"/>
            <el-table-column label="Family Note" width="120" prop="familyNote"/>
            <el-table-column label="Next Visit Note" width="120" prop="nextVisitNote"/>
            <el-table-column align="right">
              <template slot="header" slot-scope="scope">
                <el-input v-model="search" size="mini" placeholder="Type to search base on participant" />
              </template>
                <el-button size="mini" @click="customer = scope.row; isVisibleUpdate = !isVisibleUpdate">
                  | Edit
                </el-button> 
                <el-button size="mini" type="danger" @click="isVisibleDelete = !isVisibleDelete; scopeCustomer = scope; customerName = scope.row.name">
                  | Delete
                </el-button>
            </el-table-column>
          </el-table>
          <edit-customer :is-visible-update.sync="isVisibleUpdate" :customer.sync="customer" @isUpdateCustomer="handleUpdate"/>
          <delete-customer :is-visible-delete.sync="isVisibleDelete" :customer-name.sync="customerName" @isDeleteCustomer="handleDelete"/>
          <add-customer :is-visible-add.sync="isVisibleAdd" @isAddCustomer="handleAdd"/>
        </div>
      </el-card>      
</template>

<script>
import EditCustomer from '@/components/customers/CustomerUpdate.vue'
import DeleteCustomer from '@/components/customers/CustomerDelete.vue'
import AddCustomer from '@/components/customers/CustomerAdd.vue'
import {
  getAllCustomers,
  createCustomer,
  updateCustomer,
  deleteCustomerByID
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
              type: 'success',
              position: 'bottom-right'
            })
            customer.id = resp.data.id
            this.tableData.splice(0, 0, customer)
          })
          .catch(err => {
            console.log(err)
            this.$notify.error({
              title: 'Error',
              message: err,
              position: 'bottom-right'
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
              type: 'success',
              position: 'bottom-right'
            })
          })
          .catch(err => {
            console.log(err)
            this.$notify.error({
              title: 'Error',
              message: err,
              position: 'bottom-right'
            })
          })
        this.loading = false
      }
    },
    handleDelete: function(isDeleteCustomer) {
      if (isDeleteCustomer) {
        this.loading = true
        deleteCustomerByID(this.scopeCustomer.row.id)
          .then(resp => {
            this.tableData.splice(this.scopeCustomer.$index, 1)
            this.$notify({
              title: 'Success',
              message: 'Delete successfully!',
              type: 'success',
              position: 'bottom-right'
            })
          })
          .catch(err => {
            console.log(err)
            this.$notify.error({
              title: 'Error',
              message: err,
              position: 'bottom-right'
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

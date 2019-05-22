<template>
  <el-container>
    <el-main>
      <el-card class="box-card">
        <div slot="header" class="clearfix" style="">
          <span style="font-size: 18px;">Customer Associate</span>
          <el-button type="primary" icon="el-icon-plus" plain style="float:right" @click="getCustomerAssociate">
            Assign customer
          </el-button>
        </div>
        <div class="text item">
          <el-table v-loading="loading" :data="tableData.filter(data => !search || data.name.toLowerCase().includes(search.toLowerCase()))" style="width: 100%">
            <el-table-column type="expand">
              <template slot-scope="scope">
                <GiftListByVisit :visit-id="visit.id" :customer-id="scope.row.id" />
              </template>
            </el-table-column>
            <el-table-column label="Customer Name" prop="customerName" sortable />
            <el-table-column label="Pre-approved visa" width="120">
              <template slot-scope="scope">
                <el-checkbox v-model="scope.row.preApproveVisa" disabled />
              </template>
            </el-table-column>
            <el-table-column label="Created By" prop="createdBy" sortable />
            <el-table-column label="Note" prop="note" sortable />
            <el-table-column align="right">
              <CustomerAssociateUpdate :is-visible-update.sync="isVisibleUpdate" :customer.sync="customer" @isUpdateCustomer="handleCusAssocUpdate" />
              <CustomerAssociateDelete :is-visible-delete.sync="isVisibleDelete" :customer-name.sync="customerName" @isDeleteCustomer="handleCusAssocDelete" />
              <CustomerAssociateAdd :assigned-customers.sync="assignedCustomers" :is-visible-assign.sync="isVisibleAssign" @isCustomerAssociateAdd="handleCusAssocAdd" />
              <template slot-scope="scope">
                <el-button size="mini" @click="customer = scope.row; isVisibleUpdate = !isVisibleUpdate; ">
                  Edit
                </el-button>
                <el-button size="mini" type="danger" @click="isVisibleDelete = !isVisibleDelete; scopeCustomer = scope; customerName = scope.row.customerName">
                  Delete
                </el-button>
              </template>
            </el-table-column>
          </el-table>
        </div>
      </el-card>
    </el-main>
  </el-container>
</template>

<script>
import CustomerAssociateUpdate from '@/components/customerAssocicates/CustomerAssociateUpdate.vue'
import CustomerAssociateDelete from '@/components/customerAssocicates/CustomerAssociateDelete.vue'
import CustomerAssociateAdd from '@/components/customerAssocicates/CustomerAssociateAdd.vue'
import GiftListByVisit from '@/components/gifts/GiftListByVisit.vue'

import {
  getCusVisitAssocsByVisitID,
  createCusVisitAssoc,
  modifyCusVisitAssocs,
  deleteCusVisitAssocByID
} from '@/api/cusvisitassoc'

export default {
  name: 'CustomerListByVisit',
  components: {
    CustomerAssociateUpdate,
    CustomerAssociateDelete,
    CustomerAssociateAdd,
    GiftListByVisit
  },
  props: {
    customerID: { type: String, default: '' },
    visit: {
      type: Object,
      default: () => ({
        id: '',
        name: '',
        lab: [],
        arrivedTime: 0,
        departureTime: 0,
        accommodation: '',
        pickup: ''
      })
    }
  },

  data() {
    return {
      tableData: [],
      isVisibleUpdate: false,
      isVisibleDelete: false,
      isVisibleAssign: false,
      search: '',
      giftAssociates: {},
      customerName: '',
      loading: true,
      customer: {},
      scopeCustomer: {},
      gift: {},
      assignedCustomers: []
    }
  },
  mounted() {
    getCusVisitAssocsByVisitID(this.visit.id)
      .then(resp => {
        if (resp.data != null) {
          this.tableData = resp.data
          this.tableData.forEach((customer) => {
            this.assignedCustomers.push(customer.customerID)
          })
        }
      })
      .catch(err => {
        console.log(err)
      })
    this.loading = false
  },

  methods: {
    getCustomerAssociate: function() {
      this.assignedCustomers = []
      this.tableData.forEach((customer, index) => {
        this.assignedCustomers.push(customer.customerID)
      })
      this.isVisibleAssign = !this.isVisibleAssign
    },

    // handleCusAssocAdd: Create if this customer is not exist in current assigned customers.
    // Delete if the current assigned customers does not exist in updated assigned customers.
    handleCusAssocAdd: function(isCustomerAssociateAdd, updatedCustomerAssociates) {
      var currentAssigneds = this.assignedCustomers
      if (isCustomerAssociateAdd) {
        updatedCustomerAssociates.forEach((customer, index) => {
          // If currentAssigneds doesn not include this customer, assign this customer
          var position = currentAssigneds.indexOf(customer.initial)
          if (position < 0) {
            var cusVisitAssocs = {}
            cusVisitAssocs.customerID = customer.initial
            cusVisitAssocs.visitID = this.visit.id
            cusVisitAssocs.customerName = customer.label
            cusVisitAssocs.preApproveVisa = false
            cusVisitAssocs.CreatedBy = ''
            cusVisitAssocs.note = ''
            createCusVisitAssoc(cusVisitAssocs)
              .then(resp => {
                this.$notify({
                  title: 'Success',
                  message: 'Update successfully!',
                  type: 'success',
                  position: 'bottom-right'
                })
                cusVisitAssocs.id = resp.data.id
                this.tableData.splice(0, 0, cusVisitAssocs)
              })
              .catch(err => {
                console.log(err)
                this.$notify.error({
                  title: 'Error',
                  message: err,
                  position: 'bottom-right'
                })
              })
          } else {
            // Remove this gifts are not in currentAssignedGifts
            currentAssigneds.splice(position, 1)
          }
        })
        // Current elements in currentAssigneds are customers that is not in updated Customers
        if (currentAssigneds.length > 0) {
          currentAssigneds.forEach((id) => {
            // Find to get gift associate ID, then delete this gift associate
            var data = this.tableData.filter(data => !id || data.customerID.includes(id))
            deleteCusVisitAssocByID(data[0].id)
              .then(resp => {
                var giftIndex = this.assignedCustomers.indexOf(id)
                this.tableData.splice(giftIndex, 1)
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
          })
        }
        this.loading = false
      }
    },

    handleCusAssocDelete: function(isDeleteCustomer) {
      if (isDeleteCustomer) {
        this.loading = true
        deleteCusVisitAssocByID(this.scopeCustomer.row.id)
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

    handleCusAssocUpdate: function(isUpdateCustomer) {
      if (isUpdateCustomer) {
        this.loading = true
        modifyCusVisitAssocs(this.customer)
          .then(resp => {
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
    }
  }
}
</script>


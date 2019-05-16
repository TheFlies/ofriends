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
            <el-table-column label="Customer Name" prop="name" sortable />
            <el-table-column label="Position" prop="position" sortable />
            <el-table-column label="Age" prop="age" sortable />
            <el-table-column label="Project" prop="project" sortable />
            <el-table-column align="right">
              <!-- <CustomerAssociateUpdate
                :is-visible-update.sync="isVisibleUpdate"
                :gift.sync="gift"
                @isUpdateGift="handleGiftAssociateUpdate"
              /> -->
              <CustomerAssociateDelete
                :is-visible-delete.sync="isVisibleDelete"
                :customer-name.sync="customerName"
                @isDeleteCustomer="handleCustomerAssociateDelete"
              />
              <CustomerAssociateAdd :assigned-customers.sync="assignedCustomers" :is-visible-assign.sync="isVisibleAssign" @isCustomerAssociateAdd="handleCustomerAssociateAdd" />
              <template slot-scope="scope">
                <!--<el-button
                  size="mini"
                  @click="gift = scope.row; isVisibleUpdate = !isVisibleUpdate; "
                >
                  Edit
                </el-button> -->
                <el-button size="mini" type="danger" @click="isVisibleDelete = !isVisibleDelete; scopeCustomer = scope; customerName = scope.row.name">
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
// import CustomerAssociateUpdate from '@/components/customerAssocicates/CustomerAssociateUpdate.vue'
import CustomerAssociateDelete from '@/components/customerAssocicates/CustomerAssociateDelete.vue'
import CustomerAssociateAdd from '@/components/customerAssocicates/CustomerAssociateAdd.vue'
import GiftListByVisit from '@/components/gifts/GiftListByVisit.vue'

import {
  // getGiftAssociatesByVisitID,
  // createGiftAssociate,
  // deleteGiftAssociateById,
  // modifyGiftAssociates,
  deleteGiftAssociatesByVisitIDNCustomerID
} from '@/api/giftAssociate'

import {
  getCustomerByID
} from '@/api/customer'

import {
  updateVisit
} from '@/api/visit'

export default {
  name: 'CustomerListByVisit',
  components: {
    // CustomerAssociateUpdate,
    CustomerAssociateDelete,
    CustomerAssociateAdd,
    GiftListByVisit
  },
  props: {
    customerID: { type: String, default: '' },
    visit: { type: Object }
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
      scopeCustomer: {},
      gift: {},
      assignedCustomers: []
    }
  },
  mounted() {
    this.visit.customerID.forEach((id, index) => {
      getCustomerByID(id).then(resp => {
        if (resp.data != null) {
          this.tableData.push(resp.data)
        }
      })
    })
    this.loading = false
  },

  methods: {
    getCustomerAssociate: function() {
      this.assignedCustomers = []
      this.tableData.forEach((customer, index) => {
        this.assignedCustomers.push(customer.id)
      })
      this.isVisibleAssign = !this.isVisibleAssign
    },

    // handleGiftAssociateAdd: Create if this customer is not exist in current assigned customers.
    // Delete if the current assigned customers does not exist in updated assigned gifts.
    handleCustomerAssociateAdd: function(isCustomerAssociateAdd, updatedCustomerAssociates) {
      this.tableData = []
      var customerIDList = []
      if (isCustomerAssociateAdd) {
        updatedCustomerAssociates.forEach((customer, index) => {
          customerIDList.push(customer.initial)
        })
      }
      this.visit.customerID = customerIDList
      updateVisit(this.visit)
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

      // Load list again
      this.visit.customerID.forEach((id, index) => {
        getCustomerByID(id).then(resp => {
          if (resp.data != null) {
            this.tableData.push(resp.data)
          }
        })
      })
      this.loading = false
    },

    handleCustomerAssociateDelete: function(isDeleteCustomer) {
      if (isDeleteCustomer) {
        this.loading = true
        var position = this.visit.customerID.indexOf(this.scopeCustomer.row.id)
        if (position >= 0) {
          this.visit.customerID.splice(position, 1)
          var positionCustomer = this.tableData.indexOf(this.scopeCustomer.row)
          if (positionCustomer >= 0) {
            deleteGiftAssociatesByVisitIDNCustomerID(this.visit.id, this.scopeCustomer.row.id)
              .then(resp => {
                // console.log(resp)
              })
              .catch(err => {
                console.log(err)
                this.$notify.error({
                  title: 'Error',
                  message: err,
                  position: 'bottom-right'
                })
              })
            this.tableData.splice(positionCustomer, 1)
          }
          updateVisit(this.visit)
            .then(resp => {
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
      }
      this.loading = false
    }
  }
}
</script>


<template>
  <el-dialog
    title="Assign Customer"
    :visible.sync="isVisibleAssign"
    width="30%"
    append-to-body
    :before-close="handleBackdropClick"
  >
    <div style="text-align: center">
      <el-transfer
        v-model="value"
        style="text-align: left; display: inline-block"
        filterable
        :filter-method="filterMethod"
        filter-placeholder="State Abbreviations"
        :titles="['Customers', 'Assigned Customers']"
        :data="data"
      />
      <div style="margin-top: 20px">
        <el-button
          type="primary"
          @click="submit()"
        >
          Save
        </el-button>
        <el-button @click="resetForm('customerAssociate')">
          Cancel
        </el-button>
      </div>
    </div>
  </el-dialog>
</template>

<script>
import { getAllCustomers } from '@/api/customer'

export default {
  name: 'CustomerAssociateAdd',

  props: {
    isVisibleAssign: { type: Boolean, default: false },
    assignedCustomers: { type: Array, default: () => [] }
  },
  data() {
    return {
      data: [],
      value: []
    }
  },
  watch: {
    isVisibleAssign: function(val) {
      getAllCustomers().then(resp => {
        this.value = []
        this.data = []
        if (resp.data != null) {
          resp.data.forEach((customer, index) => {
            this.data.push({
              label: customer.name,
              key: index,
              initial: customer.id
            })
            var position = this.assignedCustomers.indexOf(customer.id)
            if (position >= 0) {
              this.value.push(index)
            }
          })
        }
      })
        .catch(err => {
          console.log(err)
        })
    }
  },
  methods: {
    filterMethod(query, item) {
      return item.initial.toLowerCase().indexOf(query.toLowerCase()) > -1
    },

    handleBackdropClick() {
      this.$emit('update:isVisibleAssign', false)
    },
    submit() {
      var customerAssociate = []
      this.value.forEach((val) => {
        customerAssociate.push(this.data[val])
      })
      this.$emit('update:isVisibleAssign', false)
      this.$emit('isCustomerAssociateAdd', true, customerAssociate)
    },
    resetForm(formName) {
      this.$emit('update:isVisibleAssign', false)
    }
  }
}
</script>

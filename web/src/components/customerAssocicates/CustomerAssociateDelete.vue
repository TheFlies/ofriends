<template>
  <el-dialog
    title="Delete Customer"
    :visible.sync="isVisibleDelete"
    width="30%"
    append-to-body
    :before-close="handleBackdropClick"
  >
    <span>Deleting {{ customerName }} will lead to delete Gifts in that Customer. <br>Are you sure? </span>
    <span
      slot="footer"
      class="dialog-footer"
    >
      <el-button @click="handleBackdropClick">Cancel</el-button>
      <el-button
        type="primary"
        @click="deleteCustomer"
      >Confirm</el-button>
    </span>
  </el-dialog>
</template>

<script>
export default {
  name: 'CustomerAssociateDelete',
  props: {
    isVisibleDelete: { type: Boolean, default: false },
    customerName: { type: String, default: '' },
    isDeleteCustomer: { type: Boolean, default: false }
  },
  methods: {
    handleBackdropClick() {
      this.$emit('update:isVisibleDelete', false)
    },
    deleteCustomer() {
      this.$emit('update:isVisibleDelete', false)
      this.$emit('isDeleteCustomer', true)
    }
  }
}
</script>

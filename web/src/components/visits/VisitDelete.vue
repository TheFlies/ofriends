<template>
  <el-dialog
    title="Delete Visit"
    :visible.sync="isVisibleDelete"
    width="30%"
    append-to-body
    :before-close="handleBackdropClick"
  >
    <span>Deleting {{ visitName }} will lead to delete Activities, Customers and Gifts in that Visit. <br>Are you sure? </span>
    <span slot="footer" class="dialog-footer">
      <el-button @click="handleBackdropClick">Cancel</el-button>
      <el-button type="primary" @click="deleteVisit">Confirm</el-button>
    </span>
  </el-dialog>
</template>

<script>
export default {
  name: 'VisitDelete',
  props: {
    isVisibleDelete: { type: Boolean, default: false },
    visitInfo: { type: String, default: '' },
    isDeleteVisit: { type: Boolean, default: false },
    visitName: { type: String, default: '' }
  },
  methods: {
    handleBackdropClick() {
      this.$emit('update:isVisibleDelete', false)
    },
    deleteVisit() {
      this.$emit('update:isVisibleDelete', false)
      this.$emit('isDeleteVisit', true)
    }
  }
}
</script>

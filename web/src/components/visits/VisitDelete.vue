<template>
    <el-dialog
    title=""
    :visible.sync="isVisibleDelete"
    width="30%"
    append-to-body
    :before-close="handleBackdropClick">
    <span>Do you want to delete: {{ visitInfo }}?</span>
    <span slot="footer" class="dialog-footer">
        <el-button @click="handleBackdropClick">Cancel</el-button>
        <el-button type="primary" @click="deleteVisit">Confirm</el-button>
    </span>
    </el-dialog>
</template>

<script>
export default {
  name: 'visitDelete',
  props: {
    isVisibleDelete: { type: Boolean, default: false },
    visitInfo: { type: String, default: '' },
    isDeleteVisit: { type: Boolean, default: false }
  },
  methods: {
    handleBackdropClick () {
      this.$emit('update:isVisibleDelete', false)
    },
    deleteVisit () {
      this.$emit('update:isVisibleDelete', false)
      this.$emit('isDeleteVisit', true)
    }
  }
}
</script>

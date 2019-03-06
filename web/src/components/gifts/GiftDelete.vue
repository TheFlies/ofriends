<template>
    <el-dialog
    title="Tips"
    :visible.sync="isVisibleDelete"
    width="30%"
    append-to-body
    :before-close="handleBackdropClick">
    <span>Do you want to delete: {{ giftName }}?</span>
    <span slot="footer" class="dialog-footer">
        <el-button @click="handleBackdropClick">Cancel</el-button>
        <el-button type="primary" @click="deleteGift">Confirm</el-button>
    </span>
    </el-dialog>
</template>

<script>
export default {
  name: 'giftDelete',
  props: {
    isVisibleDelete: { type: Boolean, default: false },
    giftName: { type: String, default: '' },
    isDeleteGift: { type: Boolean, default: false }
  },
  methods: {
    handleBackdropClick () {
      this.$emit('update:isVisibleDelete', false)
    },
    deleteGift () {
      this.$emit('update:isVisibleDelete', false)
      this.$emit('isDeleteGift', true)
    }
  }
}
</script>

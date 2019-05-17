<template>
  <el-dialog
    title=""
    :visible.sync="isVisibleDelete"
    width="30%"
    append-to-body
    :before-close="handleBackdropClick"
  >
    <span>Do you want to delete: {{ userName }}?</span>
    <span slot="footer" class="dialog-footer">
      <el-button @click="handleBackdropClick">Cancel</el-button>
      <el-button type="primary" @click="deleteUser">Confirm</el-button>
    </span>
  </el-dialog>
</template>

<script>
export default {
  name: 'UserDelete',
  props: {
    isVisibleDelete: { type: Boolean, default: false },
    userName: { type: String, default: '' },
    isDeleteUser: { type: Boolean, default: false }
  },
  methods: {
    handleBackdropClick() {
      this.$emit('update:isVisibleDelete', false)
    },
    deleteUser() {
      this.$emit('update:isVisibleDelete', false)
      this.$emit('isDeleteUser', true)
    }
  }
}
</script>

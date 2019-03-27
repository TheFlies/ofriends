<template>
    <el-dialog
    title=""
    :visible.sync="isVisibleDelete"
    width="30%"
    append-to-body
    :before-close="handleBackdropClick">
    <span>Do you want to delete: {{ friendName }}?</span>
    <span slot="footer" class="dialog-footer">
        <el-button @click="handleBackdropClick">Cancel</el-button>
        <el-button type="primary" @click="deleteFriend">Confirm</el-button>
    </span>
    </el-dialog>
</template>

<script>
export default {
  name: 'friendDelete',
  props: {
    isVisibleDelete: { type: Boolean, default: false },
    friendName: { type: String, default: '' },
    isDeleteFriend: { type: Boolean, default: false }
  },
  methods: {
    handleBackdropClick () {
      this.$emit('update:isVisibleDelete', false)
    },
    deleteFriend () {
      this.$emit('update:isVisibleDelete', false)
      this.$emit('isDeleteFriend', true)
    }
  }
}
</script>

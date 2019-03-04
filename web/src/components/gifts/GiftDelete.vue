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
    giftId: { type: String, default: '' },
    giftName: { type: String, default: '' },
    rowIndex: { type: Number, default: 0 }
  },
  methods: {
    handleBackdropClick () {
      this.$emit('update:isVisibleDelete', false)
    },
    deleteGift () {
      this.$http.delete('http://localhost:8080/api/v1/gift/' + this.giftId)
        .then(resp => {
          console.log(resp.data)
        })
        .catch(err => {
          console.log(err)
        })
      this.$emit('update:isVisibleDelete', false)
    }
  }
}
</script>

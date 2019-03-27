<template>
  <el-dialog
    title
    :visible.sync="isVisibleDelete"
    width="30%"
    append-to-body
    :before-close="handleBackdropClick"
  >
    <span>Do you want to delete?</span>
    <span slot="footer" class="dialog-footer">
      <el-button @click="handleBackdropClick">Cancel</el-button>
      <el-button type="primary" @click="deleteAct">Confirm</el-button>
    </span>
  </el-dialog>
</template>

<script>
export default {
  name: "ActivityDelete",
  props: {
    isVisibleDelete: { type: Boolean, default: false },
    isDeleteAct: { type: Boolean, default: false }
  },
  methods: {
    handleBackdropClick() {
      this.$emit("update:isVisibleDelete", false);
    },
    deleteAct() {
      this.$emit("update:isVisibleDelete", false);
      this.$emit("isDeleteAct", true);
    }
  }
};
</script>

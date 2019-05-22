<template>
  <el-dialog title="update Activity" :visible.sync="isVisibleUpdate" width="30%" append-to-body :before-close="handleBackdropClick">
    <el-form ref="activity" :model="activity" :rules="rules" label-width="130px" class="activity-form">
      <el-form-item label="Name" prop="name">
        <el-input v-model="activity.name" placeholder="What is the activity name?" />
      </el-form-item>
      <el-form-item label="Detail Plan" prop="detail">
        <el-input v-model="activity.detail" type="textarea" placeholder="Planning..." />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="submitForm('activity')">
          Save
        </el-button>
        <el-button @click="resetForm('activity')">
          Cancel
        </el-button>
      </el-form-item>
    </el-form>
  </el-dialog>
</template>

<script>
export default {
  name: 'ActivityUpdate',
  props: {
    isVisibleUpdate: { type: Boolean, default: false },
    activity: {
      type: Object,
      default: () => ({
        id: '',
        detail: ''
      })
    }
  },
  data() {
    return {
      rules: {}
    }
  },
  methods: {
    handleBackdropClick() {
      this.$refs['activity'].resetFields()
      this.$emit('update:isVisibleUpdate', false)
    },
    submitForm(formName) {
      this.$refs[formName].validate(valid => {
        if (valid) {
          this.$emit('update:isVisibleUpdate', false)
          this.$emit('isUpdateAct', true)
        } else {
          console.log('error submit!!')
          return false
        }
      })
    },
    resetForm(formName) {
      this.$refs[formName].resetFields()
      this.$emit('update:isVisibleUpdate', false)
    }
  }
}
</script>

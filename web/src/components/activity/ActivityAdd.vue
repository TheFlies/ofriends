<template>
  <el-dialog
    title="Add Activity"
    :visible.sync="isVisibleAdd"
    width="30%"
    append-to-body
    :before-close="handleBackdropClick"
  >
    <el-form ref="activity" :model="activity" :rules="rules" label-width="130px" class="activity-form">
      <el-form-item label="Name" prop="name">
        <el-input v-model="activity.name" placeholder="What is the activity name?" />
      </el-form-item>
      <el-form-item label="Detail Plan" prop="detail">
        <el-input
          v-model="activity.detail"
          type="textarea"
          placeholder="Planning..."
        />
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
const activityDefault = {
  name: '',
  detail: ''
}
export default {
  name: 'ActivityAdd',
  props: {
    isVisibleAdd: { type: Boolean, default: false }
  },
  data() {
    return {
      activity: activityDefault,
      rules: {
        name: [
          {
            required: true,
            message: 'Please input activity name',
            trigger: 'blur'
          }
        ]
      }
    }
  },
  watch: {
    isVisibleAdd: function(val) {
      if (val) {
        this.activity = {
          name: '',
          detail: ''
        }
      }
    }
  },
  mounted() {
    this.activity = activityDefault
  },
  methods: {
    handleBackdropClick() {
      this.$emit('update:isVisibleAdd', false)
    },
    submitForm(formName) {
      this.$refs[formName].validate(valid => {
        if (valid) {
          this.$emit('update:isVisibleAdd', false)
          this.$emit('isAddAct', true, this.activity)
        } else {
          console.log('error submit!!')
          return false
        }
        this.activity = activityDefault
      })
    },
    resetForm(formName) {
      this.$refs[formName].resetFields()
      this.$emit('update:isVisibleAdd', false)
    }
  }
}
</script>

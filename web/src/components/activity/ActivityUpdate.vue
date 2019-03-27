<template>
  <el-dialog
    title="update Activity"
    :visible.sync="isVisibleUpdate"
    width="30%"
    append-to-body
    :before-close="handleBackdropClick"
  >
    <el-form
      ref="activity"
      :model="activity"
      :rules="rules"
      label-width="130px"
      class="activity-form"
    >
      <el-form-item label="Start time" required>
        <el-col :span="22">
          <el-form-item prop="startTime">
            <el-date-picker
              v-model="activity.startTime"
              type="datetime"
              placeholder="Pick start time"
              format="yyyy/MM/dd HH:mm"
              style="width: 100%;"
            />
          </el-form-item>
        </el-col>
      </el-form-item>
      <el-form-item label="End time" required>
        <el-col :span="22">
          <el-form-item prop="endTime">
            <el-date-picker
              v-model="activity.endTime"
              type="datetime"
              placeholder="Pick End Time"
              format="yyyy/MM/dd HH:mm"
              style="width: 100%;"
            />
          </el-form-item>
        </el-col>
      </el-form-item>
      <el-form-item label="Detail Plan" prop="detail">
        <el-input v-model="activity.detail" type="textarea" placeholder="Planning..." />
      </el-form-item>
      <el-form-item label="Participant" prop="participant">
        <el-input v-model="activity.participant" placeholder="Who joint this activity?" />
      </el-form-item>
      <el-form-item label="Hotel Stayed" prop="hotel">
        <el-input v-model="activity.hotel" placeholder="Where hotel friend stayed?" />
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
        startTime: '',
        endTime: '',
        detail: '',
        participant: '',
        hotel: ''
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

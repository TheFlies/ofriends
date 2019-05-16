<template>
  <el-dialog title="Update Activity" :visible.sync="isVisibleUpdate" width="30%" append-to-body :before-close="handleBackdropClick">
    <el-form ref="activity" :model="activity" label-width="120px">
      <el-form-item label="Name" prop="name">
        <el-input v-model="activity.name" :disabled="true" />
      </el-form-item>
      <el-form-item label="Start time" required>
        <el-col :span="22">
          <el-form-item prop="startTime">
            <el-date-picker
              v-model="activity.startTime"
              type="datetime"
              placeholder="Pick start time"
              format="yyyy/MM/dd HH:mm"
              value-format="timestamp"
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
              value-format="timestamp"
              style="width: 100%;"
            />
          </el-form-item>
        </el-col>
      </el-form-item>
      <el-form-item label="Participant" prop="participant">
        <el-input v-model="activity.participant" />
      </el-form-item>
      <el-form-item label="Note" prop="note">
        <el-input v-model="activity.note" type="textarea" />
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
  name: 'GiftUpdate',
  props: {
    isVisibleUpdate: { type: Boolean, default: false },
    activity: {}
  },
  methods: {
    handleBackdropClick() {
      this.$emit('update:isVisibleUpdate', false)
    },
    submitForm(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          this.$emit('update:isVisibleUpdate', false)
          this.$emit('isUpdateActivity', true)
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

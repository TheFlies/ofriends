<template>
  <el-dialog
    title="Add Activity"
    :visible.sync="isVisibleAdd"
    width="30%"
    append-to-body
    :before-close="handleBackdropClick"
  >
    <el-form ref="activity" :model="activity" :rules="rules" label-width="130px" class="activity-form">
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
      <el-form-item label="Detail Plan" prop="detail">
        <el-input
          v-model="activity.detail"
          type="textarea"
          placeholder="Planning..."
        />
      </el-form-item>
      <el-form-item label="Participant" prop="participant">
        <el-input v-model="activity.participant" placeholder="Who joint this activity?" />
      </el-form-item>
      <el-form-item label="Hotel Stayed" prop="hotel">
        <el-input v-model="activity.hotel" placeholder="Where hotel customer stayed?" />
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
  name: 'ActivityAdd',
  props: {
    isVisibleAdd: { type: Boolean, default: false }
  },
  data() {
    return {
      activity: {
        startTime: new Date().getTime(),
        endTime: new Date().getTime(),
        detail: '',
        participant: '',
        hotel: '',
        visitID: ''
      },
      rules: {
        startTime: [
          {
            required: true,
            message: 'Please input start time',
            trigger: 'change'
          }
        ],
        endTime: [
          {
            required: true,
            message: 'Please input end time',
            trigger: 'change'
          }
        ]
      }
    }
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
      })
    },
    resetForm(formName) {
      this.$refs[formName].resetFields()
      this.$emit('update:isVisibleAdd', false)
    }
  }
}
</script>

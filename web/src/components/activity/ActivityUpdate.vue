<template>
  <el-dialog
    title="update Activity"
    :visible.sync="isVisibleUpdate"
    width="30%"
    append-to-body
    :before-close="handleBackdropClick"
  >
    <el-form
      :model="activity"
      :rules="rules"
      ref="activity"
      label-width="130px"
      class="activity-form"
    >
      <el-form-item label="Start time" required>
        <el-col :span="22">
          <el-form-item prop="startTime">
            <el-date-picker
              type="datetime"
              placeholder="Pick start time"
              v-model="activity.startTime"
              format="yyyy/MM/dd HH:mm"
              style="width: 100%;"
            ></el-date-picker>
          </el-form-item>
        </el-col>
      </el-form-item>
      <el-form-item label="End time" required>
        <el-col :span="22">
          <el-form-item prop="endTime">
            <el-date-picker
              type="datetime"
              placeholder="Pick End Time"
              v-model="activity.endTime"
              format="yyyy/MM/dd HH:mm"
              style="width: 100%;"
            ></el-date-picker>
          </el-form-item>
        </el-col>
      </el-form-item>
      <el-form-item label="Detail Plan" prop="detail">
        <el-input type="textarea" placeholder="Planning..." v-model="activity.detail"></el-input>
      </el-form-item>
      <el-form-item label="Participant" prop="participant">
        <el-input placeholder="Who joint this activity?" v-model="activity.participant"></el-input>
      </el-form-item>
      <el-form-item label="Hotel Stayed" prop="hotel">
        <el-input placeholder="Where hotel friend stayed?" v-model="activity.hotel"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="submitForm('activity')">Save</el-button>
        <el-button @click="resetForm('activity')">Cancel</el-button>
      </el-form-item>
    </el-form>
  </el-dialog>
</template>


<script>
export default {
  name: "ActivityUpdate",
  props: {
    isVisibleUpdate: { type: Boolean, default: false },
    activity: {
      startTime: "",
      endTime: "",
      detail: "",
      participant: "",
      hotel: ""
    }
  },
  data() {
    return {
      rules: {}
    };
  },
  methods: {
    handleBackdropClick() {
      this.$emit("update:isVisibleUpdate", false);
    },
    submitForm(formName) {
      this.$refs[formName].validate(valid => {
        if (valid) {
          this.$emit("update:isVisibleUpdate", false);
          this.$emit("isUpdateAct", true);
        } else {
          console.log("error submit!!");
          return false;
        }
      });
    },
    resetForm(formName) {
      this.$refs[formName].resetFields();
      this.$emit("update:isVisibleUpdate", false);
    }
  }
};
</script>

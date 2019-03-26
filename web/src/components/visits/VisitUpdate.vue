<template>
  <el-dialog
    title="update Visit"
    :visible.sync="isVisibleUpdate"
    width="30%"
    append-to-body
    :before-close="handleBackdropClick"
  >
    <el-form :model="visit" :rules="rules" ref="visit" label-width="120px" class="visit-form">
      <el-form-item label="lab" prop="lab">
        <el-select style="width: 100%;" v-model="visit.lab" placeholder="please select lab visit">
          <el-option label="lab 1" value="lab1"></el-option>
          <el-option label="lab 2" value="lab2"></el-option>
          <el-option label="lab 3" value="lab3"></el-option>
          <el-option label="lab 4" value="lab4"></el-option>
          <el-option label="lab 5" value="lab5"></el-option>
          <el-option label="lab 6" value="lab6"></el-option>
          <el-option label="lab 8" value="lab8"></el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="Visit time" required>
        <el-col :span="11">
          <el-form-item prop="arrivedTime">
            <el-date-picker
              type="date"
              placeholder="Pick arrived time"
              v-model="visit.arrivedTime"
              style="width: 100%;"
            ></el-date-picker>
          </el-form-item>
        </el-col>
        <el-col class="line" :span="2">-</el-col>
        <el-col :span="11">
          <el-form-item prop="departureTime">
            <el-time-picker
              placeholder="Pick Departure Time"
              v-model="visit.departureTime"
              style="width: 100%;"
            ></el-time-picker>
          </el-form-item>
        </el-col>
      </el-form-item>
      <el-form-item label="Pre-approve Visa" prop="preApproveVisa">
        <el-checkbox-group v-model="visit.preApproveVisa">
          <el-checkbox name="preApproveVisa"></el-checkbox>
        </el-checkbox-group>
      </el-form-item>
      <el-form-item label="Passport info" prop="passportInfo">
        <el-input
          type="textarea"
          placeholder="Passport info of friend"
          v-model="visit.passportInfo"
        ></el-input>
      </el-form-item>
      <el-form-item label="Created By" prop="createdBy">
        <el-input placeholder="Who in HR created the pre-approved visa?" v-model="visit.createdBy"></el-input>
      </el-form-item>

      <el-form-item label="Hotel Stayed" prop="hotelStayed">
        <el-input placeholder="Where hotel friend stayed?" v-model="visit.hotelStayed"></el-input>
      </el-form-item>

      <el-form-item label="Pickup" prop="pickup">
        <el-input placeholder="Who pickup friend?" v-model="visit.pickup"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="submitForm('visit')">Save</el-button>
        <el-button @click="resetForm('visit')">Cancel</el-button>
      </el-form-item>
    </el-form>
  </el-dialog>
</template>


<script>
export default {
  name: "visitUpdate",
  props: {
    isVisibleUpdate: { type: Boolean, default: false },
    visit: {
      lab: "",
      arrivedTime: "",
      departureTime: "",
      preApproveVisa: "",
      passportInfo: "",
      createdBy: "",
      hotelStayed: "",
      pickup: 50,
      activities: {
        startDate: "",
        endDate: "",
        detail: "",
        participant: "",
        hotel: ""
      }
    }
  },
  data() {
    return {
      rules: {

      }
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
          this.$emit("isUpdateVisit", true);
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

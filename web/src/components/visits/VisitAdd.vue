<template>
  <el-dialog
    title="Add Visit"
    :visible.sync="isVisibleAdd"
    width="30%"
    append-to-body
    :before-close="handleBackdropClick"
  >
    <el-form ref="visit" :model="visit" :rules="rules" label-width="130px" class="visit-form">
      <el-form-item label="Lab" prop="lab" required>
        <el-select v-model="visit.lab" style="width: 100%;" placeholder="please select lab visit">
          <el-option label="lab 1" value="lab1" />
          <el-option label="lab 2" value="lab2" />
          <el-option label="lab 3" value="lab3" />
          <el-option label="lab 4" value="lab4" />
          <el-option label="lab 5" value="lab5" />
          <el-option label="lab 6" value="lab6" />
          <el-option label="lab 8" value="lab8" />
        </el-select>
      </el-form-item>
      <el-form-item label="Visit time" required>
        <el-col :span="22">
          <el-form-item prop="arrivedTime">
            <el-date-picker
              v-model="visit.arrivedTime"
              type="datetime"
              placeholder="Pick arrived time"
              format="yyyy/MM/dd HH:mm"
              value-format="timestamp"
              style="width: 100%;"
            />
          </el-form-item>
        </el-col>
      </el-form-item>
      <el-form-item label="Departure time" required>
        <el-col :span="22">
          <el-form-item prop="departureTime">
            <el-date-picker
              v-model="visit.departureTime"
              type="datetime"
              placeholder="Pick Departure Time"
              format="yyyy/MM/dd HH:mm"
              value-format="timestamp"
              style="width: 100%;"
            />
          </el-form-item>
        </el-col>
      </el-form-item>
      <el-form-item label="Pre-approve Visa" prop="preApproveVisa">
        <el-checkbox-group v-model="visit.preApproveVisa">
          <el-checkbox name="preApproveVisa" />
        </el-checkbox-group>
      </el-form-item>
      <el-form-item label="Passport info" prop="passportInfo">
        <el-input
          v-model="visit.passportInfo"
          type="textarea"
          placeholder="Passport info of customer"
        />
      </el-form-item>
      <el-form-item label="Created By" prop="createdBy">
        <el-input v-model="visit.createdBy" placeholder="Who in HR created the pre-approved visa?" />
      </el-form-item>

      <el-form-item label="Hotel Stayed" prop="hotelStayed">
        <el-input v-model="visit.hotelStayed" placeholder="Where hotel customer stayed?" />
      </el-form-item>

      <el-form-item label="Pickup" prop="pickup">
        <el-input v-model="visit.pickup" placeholder="Who pickup customer?" />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="submitForm('visit')">
          Save
        </el-button>
        <el-button @click="resetForm('visit')">
          Cancel
        </el-button>
      </el-form-item>
    </el-form>
  </el-dialog>
</template>

<script>
export default {
  name: 'VisitAdd',
  props: {
    isVisibleAdd: { type: Boolean, default: false }
  },
  data() {
    return {
      visit: {
        lab: '',
        arrivedTime: new Date().getTime(),
        departureTime: new Date().getTime(),
        preApproveVisa: false,
        passportInfo: '',
        createdBy: '',
        hotelStayed: '',
        pickup: '',
        customerID: ''
      },
      rules: {
        lab: [
          {
            required: true,
            message: 'Please input lab',
            trigger: 'change'
          }
        ],
        arrivedTime: [
          {
            required: true,
            message: 'Please input arrived time',
            trigger: 'change'
          }
        ],
        departureTime: [
          {
            required: true,
            message: 'Please input departure time',
            trigger: 'change'
          }
        ]
      }
    }
  },
  created() {
    this.visit.customerID = this.$route.params.id
  },
  methods: {
    handleBackdropClick() {
      this.$emit('update:isVisibleAdd', false)
    },
    submitForm(formName) {
      this.$refs[formName].validate(valid => {
        if (valid) {
          this.$emit('update:isVisibleAdd', false)
          this.$emit('isAddVisit', true, this.visit)
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

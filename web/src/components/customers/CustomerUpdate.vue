<template lang="pug">
  <el-dialog title="Update Customer" :visible.sync="isVisibleUpdate" width="60%" append-to-body="" :before-close="handleBackdropClick">
    <el-form class="demo-ruleForm" ref="customer" :model="customer" :rules="rules" label-width="130px">
      <el-row :gutter="20">
        <el-col :span="12">
          <div class="grid-content bg-purple">
            <el-form-item label="Name" prop="name">
              <el-input v-model="customer.name" type="success" placeholder="Fistname and last name" />
            </el-form-item>
            <el-form-item label="Title" prop="title">
              <el-select v-model="customer.title" style="width: 100%;" placeholder="please select customer title" >
                <el-option label="Mr" value="Mr"/>
                <el-option label="Ms" value="Ms"/>
                <el-option label="Mrs" value="Mrs"/>
              </el-select>
            </el-form-item>
            <el-form-item label="Position" prop="position">
              <el-input v-model="customer.position" placeholder="e.g Project manager" />
            </el-form-item>
            <el-form-item label="Project" prop="project">
              <el-input v-model="customer.project" placeholder="Project name..." />
            </el-form-item>
            <el-form-item label="Age" prop="age">
              <el-input v-model="customer.age" placeholder="Age of customer" :min="20" />
            </el-form-item>
            <el-form-item label="Company" prop="company">
              <el-input v-model="customer.company" placeholder="Company name..." />
            </el-form-item>
            <el-form-item label="Country" prop="country">
              <el-select v-model="customer.country" placeholder="Country where customer live..." filterable style="width: 100%;">
                <el-option v-for="item in $countries" :key="item.code" :label="item.name" :value="item.code">
                  <flag :iso="item.code"/>
                  <span style="margin-left: 20px"> {{ item.name }}</span>
                </el-option>
              </el-select>
            </el-form-item>
            <el-form-item label="City" prop="city">
              <el-input v-model="customer.city" placeholder="City where customer live..."/>
            </el-form-item>
          </div>
        </el-col>
        <el-col :span="12">
          <div class="grid-content bg-purple">
            <el-form-item label="Pre-approve Visa" prop="preApproveVisa">
              <el-checkbox-group v-model="customer.preApproveVisa">
                <el-checkbox name="preApproveVisa"/>
              </el-checkbox-group>
            </el-form-item>
            <el-form-item label="Passport info" prop="passportInfo">
                <el-input v-model="customer.passportInfo" type="textarea" placeholder="Passport info of customer" />
            </el-form-item>
            <el-form-item label="Food Note" prop="foodNote">
              <el-input v-model="customer.foodNote" type="textarea" placeholder="is a vegetarian..."/>
            </el-form-item>
            <el-form-item label="Family Note" prop="familyNote">
              <el-input v-model="customer.familyNote" type="textarea" placeholder="Wife 34 years old, have two boys in 10 and 15 ..."/>
            </el-form-item>
            <el-form-item label="Next visit Note" prop="nextVisitNote">
              <el-input v-model="customer.nextVisitNote" type="textarea" placeholder="..."/>
            </el-form-item>
            <el-form-item>
              <el-button type="primary" @click="submitForm('customer')">
                | Save
              </el-button>
              <el-button @click="resetForm('customer')">
                | Cancel
              </el-button>
            </el-form-item>
          </div>
        </el-col>
      </el-row>
    </el-form>
  </el-dialog>
</template>

<script>
export default {
  name: 'CustomerUpdate',
  props: {
    isVisibleUpdate: { type: Boolean, default: false },
    customer: {
      type: Object,
      default: () => ({
        id: '',
        name: '',
        title: 'Mr',
        position: '',
        project: '',
        company: '',
        country: '',
        city: '',
        age: 50,
        foodNote: '',
        familyNote: '',
        nextVisitNote: ''
      })
    }
  },
  data() {
    return {
      rules: {
        name: [
          {
            required: true,
            message: 'Please input customer name',
            trigger: 'change'
          }
        ],
        title: [
          {
            required: true,
            message: 'Please input customer title',
            trigger: 'change'
          }
        ],
        position: [
          {
            required: true,
            message: 'Please input customer position',
            trigger: 'change'
          }
        ],
        project: [
          {
            required: true,
            message: 'Please input project',
            trigger: 'change'
          }
        ]
      }
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
          this.$emit('isUpdateCustomer', true)
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

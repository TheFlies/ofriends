<template>
  <el-dialog
    title="Update Assigned Customer"
    :visible.sync="isVisibleUpdate"
    width="30%"
    append-to-body
    :before-close="handleBackdropClick"
  >
    <el-form
      ref="customer"
      :model="customer"
      :rules="rules"
      label-width="150px"
      class="demo-ruleForm"
    >
      <el-form-item
        label="Name"
        prop="customerName"
      >
        <el-input v-model="customer.customerName" :disabled="true" />
      </el-form-item>
      <el-form-item
        label="Pre-Approve Visa"
        prop="preApproveVisa"
      >
        <el-checkbox v-model="customer.preApproveVisa" />
      </el-form-item>
      <el-form-item label="Created by" prop="createdBy">
        <el-input v-model="customer.createdBy" placeholder="Who in HR created the pre-approved visa?" />
      </el-form-item>
      <el-form-item
        label="Note"
        prop="note"
      >
        <el-input
          v-model="customer.note"
          type="textarea"
        />
      </el-form-item>
      <el-form-item>
        <el-button
          type="primary"
          @click="submitForm('customer')"
        >
          Save
        </el-button>
        <el-button @click="resetForm('customer')">
          Cancel
        </el-button>
      </el-form-item>
    </el-form>
  </el-dialog>
</template>

<script>
export default {
  name: 'CustomerAssociateUpdate',
  props: {
    isVisibleUpdate: { type: Boolean, default: false },
    customer: {
      type: Object,
      default: () => ({
        id: 0,
        name: '',
        idea: '',
        size: '',
        price: 0,
        link: '',
        description: ''
      })
    }
  },
  data() {
    return {
      rules: {
        name: [
          { required: true, message: 'Please input customer name', trigger: 'blur' }
        ]
      }
    }
  },
  methods: {
    handleBackdropClick() {
      this.$refs['customer'].resetFields()
      this.$emit('update:isVisibleUpdate', false)
    },
    submitForm(formName) {
      this.$refs[formName].validate((valid) => {
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

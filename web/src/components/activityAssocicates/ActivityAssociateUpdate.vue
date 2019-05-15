<template>
  <el-dialog
    title="Update Gift"
    :visible.sync="isVisibleUpdate"
    width="30%"
    append-to-body
    :before-close="handleBackdropClick"
  >
    <el-form
      ref="gift"
      :model="gift"
      :rules="rules"
      label-width="120px"
      class="demo-ruleForm"
    >
      <el-form-item
        label="Name"
        prop="giftName"
      >
        <el-input v-model="gift.giftName" :disabled="true" />
      </el-form-item>
      <el-form-item
        label="Quantity"
        prop="quantity"
      >
        <el-input-number v-model="gift.quantity" :min="1" />
      </el-form-item>
      <el-form-item
        label="Note"
        prop="note"
      >
        <el-input
          v-model="gift.note"
          type="textarea"
        />
      </el-form-item>
      <el-form-item>
        <el-button
          type="primary"
          @click="submitForm('gift')"
        >
          Save
        </el-button>
        <el-button @click="resetForm('gift')">
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
    gift: {
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
          { required: true, message: 'Please input Gift name', trigger: 'blur' }
        ]
      }
    }
  },
  methods: {
    handleBackdropClick() {
      this.$emit('update:isVisibleUpdate', false)
    },
    submitForm(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          this.$emit('update:isVisibleUpdate', false)
          this.$emit('isUpdateGift', true)
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

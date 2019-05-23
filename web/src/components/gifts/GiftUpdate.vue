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
        prop="name"
      >
        <el-input v-model="gift.name" />
      </el-form-item>
      <el-form-item
        label="Idea"
        prop="idea"
      >
        <el-input v-model="gift.idea" />
      </el-form-item>
      <el-form-item
        label="Size"
        prop="size"
      >
        <el-input v-model="gift.size" />
      </el-form-item>
      <el-form-item
        label="Price"
        prop="price"
      >
        <el-input-number
          v-model="gift.price"
          :precision="2"
          :step="0.1"
        />
      </el-form-item>
      <el-form-item
        label="Link"
        prop="link"
      >
        <el-input v-model="gift.link" />
      </el-form-item>
      <el-form-item
        label="Description"
        prop="description"
      >
        <el-input
          v-model="gift.description"
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
      this.$refs['gift'].resetFields()
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

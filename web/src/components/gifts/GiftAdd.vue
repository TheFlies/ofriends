<template>
  <el-dialog
    title="Add Gift"
    :visible.sync="isVisibleAdd"
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
        label="Quantity"
        prop="quantity"
      >
        <el-input-number
          v-model="gift.quantity"
          :min="0"
        />
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
const giftDefault = {
        name: '',
        idea: '',
        size: '',
        quantity: 0,
        price: 0,
        link: '',
        description: ''
      }
export default {
  name: 'GiftAdd',
  
  props: {
    isVisibleAdd: { type: Boolean, default: false }
  },
  data() {
    return {
      gift: giftDefault,
      rules: {
        name: [
          { required: true, message: 'Please input Gift name', trigger: 'blur' }
        ]
      }
    }
  },
  mounted() {
    this.gift =  {
      name: '',
      idea: '',
      size: '',
      quantity: 0,
      price: 0,
      link: '',
      description: ''
    }
  },
  methods: {
    handleBackdropClick() {
      this.$emit('update:isVisibleAdd', false)
    },
    submitForm(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          this.$emit('update:isVisibleAdd', false)
          this.$emit('isAddGift', true, this.gift)
        } else {
          console.log('error submit!!')
          return false
        }
      })
      this.gift = giftDefault
    },
    resetForm(formName) {
      this.$refs[formName].resetFields()
      this.$emit('update:isVisibleAdd', false)
    }
  }
}
</script>

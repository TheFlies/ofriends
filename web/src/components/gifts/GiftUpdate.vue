<template>
    <el-dialog
        title="Update Gift"
        :visible.sync="isVisibleUpdate"
        width="30%"
        append-to-body :before-close="handleBackdropClick">
        <el-form :model="gift" :rules="rules" ref="gift" label-width="120px" class="demo-ruleForm">
          <el-form-item label="Name" prop="name">
            <el-input v-model="gift.name"></el-input>
          </el-form-item>
          <el-form-item label="Idea" prop="idea">
            <el-input v-model="gift.idea"></el-input>
          </el-form-item>
          <el-form-item label="Size" prop="size">
            <el-input v-model="gift.idea"></el-input>
          </el-form-item>
          <el-form-item label="Quantity" prop="quantity">
            <el-input-number v-model="gift.quantity" :min="0"/>
          </el-form-item>
          <el-form-item label="Price" prop="price">
            <el-input-number v-model="gift.price" :precision="2" :step="0.1"/>
          </el-form-item>
          <el-form-item label="Link" prop="link">
            <el-input v-model="gift.description"></el-input>
          </el-form-item>
          <el-form-item label="Description" prop="desc">
            <el-input type="textarea" v-model="gift.desc"></el-input>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="submitForm('gift')">Save</el-button>
            <el-button @click="resetForm('gift')">Cancel</el-button>
          </el-form-item>
        </el-form>
    </el-dialog>
</template>

<script>
export default {
  name: 'giftUpdate',
  props: {
    isVisibleUpdate: { type: Boolean, default: false },
    gift: {
      id: 0,
      name: '',
      idea: '',
      size: '',
      quantity: 0,
      price: 0,
      link: '',
      description: ''
    }
  },
  data () {
    return {
      rules: {
        name: [
          { required: true, message: 'Please input Activity name', trigger: 'blur' },
          { min: 3, max: 5, message: 'Length should be 3 to 5', trigger: 'blur' }
        ],
        idea: [
          { required: true, message: 'Please input Activity name', trigger: 'blur' },
          { min: 3, max: 5, message: 'Length should be 3 to 5', trigger: 'blur' }
        ],
        size: [
          { required: true, message: 'Please input Activity name', trigger: 'blur' },
          { min: 3, max: 5, message: 'Length should be 3 to 5', trigger: 'blur' }
        ],
        link: [
          { required: true, message: 'Please input Activity name', trigger: 'blur' },
          { min: 3, max: 5, message: 'Length should be 3 to 5', trigger: 'blur' }
        ],
        desc: [
          { required: true, message: 'Please input activity form', trigger: 'blur' }
        ]
      }
    }
  },
  methods: {
    handleBackdropClick () {
      this.$emit('update:isVisibleUpdate', false)
    },
    submitForm (formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          this.$http.put('http://localhost:8080/api/v1/gift')
            .then(resp => {
              console.log(resp.data)
            })
            .catch(err => {
              console.log(err)
            })
        } else {
          console.log('error submit!!')
          return false
        }
      })
    },
    resetForm (formName) {
      this.$refs[formName].resetFields()
    }
  }
}
</script>

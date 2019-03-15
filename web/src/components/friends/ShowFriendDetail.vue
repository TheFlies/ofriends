<template>
    <el-dialog
        title="Friend Detail"
        :visible.sync="isVisibleAdd"
        width="70%"
        append-to-body :before-close="handleBackdropClick">
        <el-form :model="friend" :rules="rules" ref="friend" label-width="120px" class="demo-ruleForm">
          <el-form-item label="Name" prop="name">
        <el-input type="success" placeholder="Fistname and last name" v-model="friend.name"></el-input>
      </el-form-item>
      <el-form-item label="Title" prop="title">
        <el-select style="width: 100%;" v-model="friend.title" placeholder="please select customer title">
          <el-option label="Mr" value="Mr"></el-option>
          <el-option label="Mrs" value="Mrs"></el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="Position" prop="position">
        <el-input placeholder="e.g Project manager" v-model="friend.position"></el-input>
      </el-form-item>
      <el-form-item label="Project" prop="project">
        <el-input placeholder="Project name..." v-model="friend.project"></el-input>
      </el-form-item>
       <el-form-item label="Age" prop="age">
        <el-input placeholder="Age of customer" v-model="friend.age" :min="20"></el-input>
      </el-form-item>
      <el-form-item label="Company" prop="company">
        <el-input placeholder="Company name..." v-model="friend.company"></el-input>
      </el-form-item>
      <el-form-item label="Country" prop="country">
        <el-input placeholder="Country where customer live..." v-model="friend.country"></el-input>
      </el-form-item>
       <el-form-item label="City" prop="city">
        <el-input placeholder="City where customer live..." v-model="friend.city"></el-input>
      </el-form-item>
      <el-form-item label="Food Note" prop="foodNote">
        <el-input
          type="textarea"
          placeholder="is a vegetarian..."
          v-model="friend.foodNote"
        ></el-input>
      </el-form-item>
      <el-form-item label="Family Note" prop="familyNote">
        <el-input
          type="textarea"
          placeholder="Wife 34 years old, have two boys in 10 and 15 ..."
          v-model="friend.familyNote"
        ></el-input>
      </el-form-item>
      <el-form-item label="Next visit Note" prop="nextVisitNote">
        <el-input
          type="textarea"
          placeholder="..."
          v-model="friend.nextVisitNote"
        ></el-input>
      </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="submitForm('friend')">Save</el-button>
            <el-button @click="resetForm('friend')">Cancel</el-button>
          </el-form-item>
        </el-form>
    </el-dialog>
</template>

<script>
export default {
  name: 'friendAdd',
  props: {
    isVisibleAdd: { type: Boolean, default: false }
  },
  data () {
    return {
      friend: {
        id: 0,
        name: '',
        title: 'Mr',
        position: '',
        project: '',
        age: 50,
        company: '',
        country: '',
        city: '',
        foodNote: '',
        familyNote: '',
        nextVisitNote: ''
      },
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
            message: 'Please input customer title',
            trigger: 'change'
          }
        ],
        project: [
          {
            required: true,
            message: 'Please input project name',
            trigger: 'change'
          }
        ]
      }
    }
  },
  methods: {
    handleBackdropClick () {
      this.$emit('update:isVisibleAdd', false)
    },
    submitForm (formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          this.$emit('update:isVisibleAdd', false)
          this.$emit('isAddFriend', true, this.friend)
        } else {
          console.log('error submit!!')
          return false
        }
      })
    },
    resetForm (formName) {
      this.$refs[formName].resetFields()
      this.$emit('update:isVisibleAdd', false)
    }
  }
}
</script>

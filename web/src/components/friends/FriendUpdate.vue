<template>
  <el-dialog
    title="Update Friend"
    :visible.sync="isVisibleUpdate"
    width="30%"
    append-to-body
    :before-close="handleBackdropClick"
  >
    <el-form ref="friend" :model="friend" :rules="rules" label-width="120px" class="demo-ruleForm">
      <el-form-item label="Name" prop="name">
        <el-input v-model="friend.name" type="success" placeholder="Fistname and last name" />
      </el-form-item>
      <el-form-item label="Title" prop="title">
        <el-select
          v-model="friend.title"
          style="width: 100%;"
          placeholder="please select customer title"
        >
          <el-option label="Mr" value="Mr" />
          <el-option label="Mrs" value="Mrs" />
        </el-select>
      </el-form-item>
      <el-form-item label="Position" prop="position">
        <el-input v-model="friend.position" placeholder="e.g Project manager" />
      </el-form-item>
      <el-form-item label="Project" prop="project">
        <el-input v-model="friend.project" placeholder="Project name..." />
      </el-form-item>
      <el-form-item label="Age" prop="age">
        <el-input v-model="friend.age" placeholder="Age of customer" :min="20" />
      </el-form-item>
      <el-form-item label="Company" prop="company">
        <el-input v-model="friend.company" placeholder="Company name..." />
      </el-form-item>
      <el-form-item label="Country" prop="country">
        <el-input v-model="friend.country" placeholder="Country where customer live..." />
      </el-form-item>
      <el-form-item label="City" prop="city">
        <el-input v-model="friend.city" placeholder="City where customer live..." />
      </el-form-item>
      <el-form-item label="Food Note" prop="foodNote">
        <el-input v-model="friend.foodNote" type="textarea" placeholder="is a vegetarian..." />
      </el-form-item>
      <el-form-item label="Family Note" prop="familyNote">
        <el-input
          v-model="friend.familyNote"
          type="textarea"
          placeholder="Wife 34 years old, have two boys in 10 and 15 ..."
        />
      </el-form-item>
      <el-form-item label="Next visit Note" prop="nextVisitNote">
        <el-input v-model="friend.nextVisitNote" type="textarea" placeholder="..." />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="submitForm('friend')">
          Save
        </el-button>
        <el-button @click="resetForm('friend')">
          Cancel
        </el-button>
      </el-form-item>
    </el-form>
  </el-dialog>
</template>

<script>
export default {
  name: 'FriendUpdate',
  props: {
    isVisibleUpdate: { type: Boolean, default: false },
    friend: {
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
            message: 'Please input friend name',
            trigger: 'change'
          }
        ],
        title: [
          {
            required: true,
            message: 'Please input friend title',
            trigger: 'change'
          }
        ],
        position: [
          {
            required: true,
            message: 'Please input friend position',
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
          this.$emit('isUpdateFriend', true)
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

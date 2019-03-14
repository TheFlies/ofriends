<template>
  <div class="addFriendForm">
    <el-form
      ref="form"
      :model="form"
      :rules="rules"
      label-width="110px"
    >
      <h3>Add New Friend</h3>
      <el-form-item
        label="Name"
        prop="name"
      >
        <el-input
          v-model="form.name"
          type="success"
          placeholder="Fistname and last name"
        />
      </el-form-item>
      <el-form-item
        label="Title"
        prop="title"
      >
        <el-select
          v-model="form.title"
          style="width: 100%;"
          placeholder="please select customer title"
        >
          <el-option
            label="Mr"
            value="Mr"
          />
          <el-option
            label="Mrs"
            value="Mrs"
          />
        </el-select>
      </el-form-item>
      <el-form-item
        label="Position"
        prop="position"
      >
        <el-input
          v-model="form.position"
          placeholder="e.g Project manager"
        />
      </el-form-item>
      <el-form-item
        label="Project"
        prop="project"
      >
        <el-input
          v-model="form.project"
          placeholder="Project name..."
        />
      </el-form-item>
      <el-form-item
        label="Age"
        prop="age"
      >
        <el-input
          v-model="form.age"
          placeholder="Age of customer"
          :min="20"
        />
      </el-form-item>
      <el-form-item
        label="Company"
        prop="company"
      >
        <el-input
          v-model="form.company"
          placeholder="Company name..."
        />
      </el-form-item>
      <el-form-item
        label="Country"
        prop="country"
      >
        <el-input
          v-model="form.country"
          placeholder="Country where customer live..."
        />
      </el-form-item>
      <el-form-item
        label="City"
        prop="city"
      >
        <el-input
          v-model="form.city"
          placeholder="City where customer live..."
        />
      </el-form-item>
      <el-form-item
        label="Food Note"
        prop="foodNote"
      >
        <el-input
          v-model="form.foodNote"
          type="textarea"
          placeholder="is a vegetarian..."
        />
      </el-form-item>
      <el-form-item
        label="Family Note"
        prop="familyNote"
      >
        <el-input
          v-model="form.familyNote"
          type="textarea"
          placeholder="Wife 34 years old, have two boys in 10 and 15 ..."
        />
      </el-form-item>
      <el-form-item
        label="Next visit Note"
        prop="nextVisitNote"
      >
        <el-input
          v-model="form.nextVisitNote"
          type="textarea"
          placeholder="..."
        />
      </el-form-item>
      <el-form-item>
        <el-button
          type="primary"
          @click="onSubmit('form',$event)"
        >
          Add
        </el-button>
        <el-button @click="resetForm('form',$event)">
          Cancel
        </el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
export default {
  name: 'AddFriend',
  props: {
    msg: String
  },
  data() {
    return {
      form: {
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
        nextVisitNote: '',
        response: null
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
    showMessage(status, mesg) {
      var mesgType = 'error'
      if (status === 201) {
        mesgType = 'success'
        mesg = 'Create friend successfully!'
      }
      this.$message({
        message: mesg,
        type: mesgType
      })
    },
    onSubmit(form, event) {
      this.$refs[form].validate(valid => {
        if (valid) {
          var value = this.form
          this.$http
            .post('/friends', {
              Name: value.name,
              Title: value.title,
              Position: value.position,
              Project: value.project,
              Age: parseInt(value.age, 10),
              Company: value.company,
              Country: value.country,
              City: value.city,
              FoodNote: value.foodNote,
              FamilyNote: value.familyNote,
              NextVisitNote: value.nextVisitNote
            })
            .then(
              response => {
                this.form.response = response.data
                this.showMessage(response.status, response.data)
                if (response.status === 201) {
                  this.resetForm(form, event)
                }
              })
            .catch(error => {
              this.showMessage(404, error)
            })
        } else {
          console.log('error submit!!')
          return false
        }
      })
    },
    resetForm(form, event) {
      this.$refs[form].resetFields()
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="stylus">
.addFriendForm {
  max-width: 500px;
  margin: auto;
  border: 1px solid #ebebeb;
  border-radius: 3px;
  transition: 0.2s;
}

.addFriendForm .el-form {
  padding: 24px;
}
</style>

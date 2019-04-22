<template lang="pug">
.showCustomerDetail
  el-form(ref="form" :model="form" :rules="rules" label-width="110px")
    h3(style="align:center;")
      | Customer Detail Infomation
    el-col(:span="11")
      el-form-item(label="Name" prop="name")
        el-input(v-model="form.name" type="success" placeholder="Fistname and last name")
    el-col(:span="11")
      el-form-item(label="Title" prop="title")
        el-select(v-model="form.title" style="width: 100%;" placeholder="please select customer title")
          el-option(label="Mr" value="Mr")
            el-option(label="Mrs" value="Mrs")
    el-col(:span="11")
      el-form-item(label="Position" prop="position")
        el-input(v-model="form.position")
    el-col(:span="11")
      el-form-item(label="Project" prop="project")
        el-input(v-model="form.project")
    el-col(:span="11")
      el-form-item(label="Age" prop="age")
        el-input(v-model="form.age" :min="20")
    el-col(:span="11")
      el-form-item(label="Company" prop="company")
        el-input(v-model="form.company")
    el-col(:span="11")
      el-form-item(label="Country" prop="country")
        el-select(v-model="form.country" placeholder="Country where customer live..." filterable="" style="width: 100%;")
          el-option(v-for="item in $countries" :key="item.code" :label="item.name" :value="item.code")
            flag(:iso="item.code")
            span(style="margin-left: 20px")
              | {{ item.name }}
    el-col(:span="11")
      el-form-item(label="City" prop="city")
        el-input(v-model="form.city")
    el-col(:span="22")
      el-form-item(label="Food Note" prop="foodNote")
        el-input(v-model="form.foodNote" type="textarea")
    el-col(:span="22")
      el-form-item(label="Family Note" prop="familyNote")
        el-input(v-model="form.familyNote" type="textarea")
    el-col(:span="22")
      el-form-item(label="Next visit Note" prop="nextVisitNote")
        el-input(v-model="form.nextVisitNote" type="textarea")
</template>

<script>
export default {
  name: 'CustomerShowDetail',
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
      id: '',
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
  created() {
    this.id = this.$route.params.id
  },
  mounted() {
    // We already set the axios baseURL to the backend service in main.js file.
    this.loading = true
    this.$http
      .get('/customers/' + this.id)
      .then(resp => {
        if (resp.data != null) {
          this.form = resp.data
        }
      })
      .catch(err => {
        console.log(err)
      })
      .finally(() => {
        this.loading = false
      })
  },
  methods: {
    showMessage(status, mesg) {
      var mesgType = 'error'
      if (status === 201) {
        mesgType = 'success'
        mesg = 'Create customer successfully!'
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
            .post('/customers', {
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
            .then(response => {
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
.showCustomerDetail
  max-width 800px
  margin auto
  border 1px solid #ebebeb
  border-radius 3px
  transition 0.2s
  .el-form
    padding: 24px;
</style>

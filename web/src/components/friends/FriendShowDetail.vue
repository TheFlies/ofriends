<template>
  <div class="showFriendDetail">
    <el-form ref="form" :model="form" :rules="rules" label-width="110px">
      <h3 style="align:center;">Friend Detail Infomation</h3>
      <el-col :span="11">
        <el-form-item label="Name" prop="name">
          <el-input type="success" placeholder="Fistname and last name" v-model="form.name"></el-input>
        </el-form-item>
      </el-col>
      <el-col :span="11">
        <el-form-item label="Title" prop="title">
          <el-select
            style="width: 100%;"
            v-model="form.title"
            placeholder="please select customer title"
          >
            <el-option label="Mr" value="Mr"></el-option>
            <el-option label="Mrs" value="Mrs"></el-option>
          </el-select>
        </el-form-item>
      </el-col>
      <el-col :span="11">
        <el-form-item label="Position" prop="position">
          <el-input v-model="form.position"></el-input>
        </el-form-item>
      </el-col>
      <el-col :span="11">
        <el-form-item label="Project" prop="project">
          <el-input v-model="form.project"></el-input>
        </el-form-item>
      </el-col>
      <el-col :span="11">
        <el-form-item label="Age" prop="age">
          <el-input v-model="form.age" :min="20"></el-input>
        </el-form-item>
      </el-col>
      <el-col :span="11">
        <el-form-item label="Company" prop="company">
          <el-input v-model="form.company"></el-input>
        </el-form-item>
      </el-col>
      <el-col :span="11">
        <el-form-item label="Country" prop="country">
          <el-input v-model="form.country"></el-input>
        </el-form-item>
      </el-col>
      <el-col :span="11">
        <el-form-item label="City" prop="city">
          <el-input v-model="form.city"></el-input>
        </el-form-item>
      </el-col>
      <el-col :span="22">
        <el-form-item label="Food Note" prop="foodNote">
          <el-input type="textarea" v-model="form.foodNote"></el-input>
        </el-form-item>
      </el-col>
      <el-col :span="22">
        <el-form-item label="Family Note" prop="familyNote">
          <el-input
            type="textarea"
            v-model="form.familyNote"
          ></el-input>
        </el-form-item>
      </el-col>
      <el-col :span="22">
        <el-form-item label="Next visit Note" prop="nextVisitNote">
          <el-input type="textarea" v-model="form.nextVisitNote"></el-input>
        </el-form-item>
      </el-col>
      <el-form-item>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
export default {
  name: "friendShowDetail",
  data() {
    return {
      form: {
        name: "",
        title: "Mr",
        position: "",
        project: "",
        company: "",
        country: "",
        city: "",
        age: 50,
        foodNote: "",
        familyNote: "",
        nextVisitNote: "",
        response: null
      },
      id:"",
      rules: {
        name: [
          {
            required: true,
            message: "Please input friend name",
            trigger: "change"
          }
        ],
        title: [
          {
            required: true,
            message: "Please input friend title",
            trigger: "change"
          }
        ],
        position: [
          {
            required: true,
            message: "Please input friend position",
            trigger: "change"
          }
        ],
        project: [
          {
            required: true,
            message: "Please input project",
            trigger: "change"
          }
        ]
      }
    };
  },
  created() {
    this.id = this.$route.params.id
  },
  mounted() {
    // We already set the axios baseURL to the backend service in main.js file.
    this.$http
      .get("/friends/" + this.id)
      .then(resp => {
        if (resp.data != null) {
          this.form = resp.data;
        }
        this.loading = false;
      })
      .catch(err => {
        console.log(err);
      });
  },
  methods: {
    showMessage(status, mesg) {
      var mesgType = "error";
      if (status === 201) {
        mesgType = "success";
        mesg = "Create friend successfully!";
      }
      this.$message({
        message: mesg,
        type: mesgType
      });
    },
    onSubmit(form, event) {
      this.$refs[form].validate(valid => {
        if (valid) {
          var value = this.form;
          this.$http
            .post("/friends", {
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
              this.form.response = response.data;
              this.showMessage(response.status, response.data);
              if (response.status === 201) {
                this.resetForm(form, event);
              }
            })
            .catch(error => {
              this.showMessage(404, error);
            });
        } else {
          console.log("error submit!!");
          return false;
        }
      });
    },
    resetForm(form, event) {
      this.$refs[form].resetFields();
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="stylus">
.showFriendDetail {
  max-width: 800px;
  margin: auto;
  border: 1px solid #ebebeb;
  border-radius: 3px;
  transition: 0.2s;
}

.showFriendDetail .el-form {
  padding: 24px;
}
</style>

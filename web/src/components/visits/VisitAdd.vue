<template>
  <el-dialog
    title="Add Visit"
    :visible.sync="isVisibleAdd"
    width="30%"
    append-to-body
    :before-close="handleBackdropClick"
  >
    <el-form ref="visit" :model="visit" :rules="rules" label-width="130px" class="visit-form">
      <el-form-item label="Name" prop="name" required>
        <el-input v-model="visit.name" type="success" placeholder="Visit name" />
      </el-form-item>
      <el-form-item label="Lab" prop="lab" required>
        <el-tag v-for="tag in visit.lab" :key="tag" closable :disable-transitions="false" @close="handleClose(tag)">
          {{ tag }}
        </el-tag>
        <el-input v-if="inputVisible" ref="saveTagInput" v-model="inputValue" class="input-new-tag" size="mini" @keyup.enter.native="handleInputConfirm" @blur="handleInputConfirm" />
        <el-button v-else class="button-new-tag" size="small" @click="showInput">
          + Add Lab
        </el-button>
      </el-form-item>
      <el-form-item label="Arrival time" required>
        <el-col :span="22">
          <el-form-item prop="arrivedTime">
            <el-date-picker
              v-model="visit.arrivedTime"
              type="datetime"
              placeholder="Pick arrived time"
              format="yyyy/MM/dd HH:mm"
              value-format="timestamp"
              style="width: 100%;"
            />
          </el-form-item>
        </el-col>
      </el-form-item>
      <el-form-item label="Departure time" required>
        <el-col :span="22">
          <el-form-item prop="departureTime">
            <el-date-picker
              v-model="visit.departureTime"
              type="datetime"
              placeholder="Pick Departure Time"
              format="yyyy/MM/dd HH:mm"
              value-format="timestamp"
              style="width: 100%;"
            />
          </el-form-item>
        </el-col>
      </el-form-item>
      <el-form-item label="Accommodation" prop="accommodation">
        <el-input v-model="visit.hotelStayed" placeholder="Where hotel customer stayed?" />
      </el-form-item>
      <el-form-item label="Pickup" prop="pickup">
        <el-input v-model="visit.pickup" placeholder="Who pickup customer?" />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="submitForm('visit')">
          Save
        </el-button>
        <el-button @click="resetForm('visit')">
          Cancel
        </el-button>
      </el-form-item>
    </el-form>
  </el-dialog>
</template>

<script>
export default {
  name: 'VisitAdd',
  props: {
    isVisibleAdd: { type: Boolean, default: false }
  },
  data() {
    return {
      inputVisible: false,
      inputValue: '',
      visit: {
        lab: [],
        arrivedTime: new Date().getTime(),
        departureTime: new Date().getTime(),
        createdBy: '',
        hotelStayed: '',
        pickup: '',
        customerID: ''
      },
      rules: {
        arrivedTime: [
          {
            required: true,
            message: 'Please input arrived time',
            trigger: 'blur'
          }
        ],
        departureTime: [
          {
            required: true,
            message: 'Please input departure time',
            trigger: 'blur'
          }
        ]
      }
    }
  },
  watch: {
    isVisibleAdd: function(val) {
      if (val) {
        this.visit = {
          lab: [],
          arrivedTime: new Date().getTime(),
          departureTime: new Date().getTime(),
          createdBy: '',
          hotelStayed: '',
          pickup: '',
          customerID: []
        }
      }
    }
  },
  created() {
    this.visit.customerID = []
    this.visit.activityID = []
  },
  methods: {
    handleBackdropClick() {
      this.$emit('update:isVisibleAdd', false)
    },
    submitForm(formName) {
      this.$refs[formName].validate(valid => {
        if (valid) {
          this.$emit('update:isVisibleAdd', false)
          this.$emit('isAddVisit', true, this.visit)
        } else {
          console.log('error submit!!')
          return false
        }
      })
    },
    resetForm(formName) {
      this.$refs[formName].resetFields()
      this.$emit('update:isVisibleAdd', false)
    },
    handleClose(tag) {
      this.visit.lab.splice(this.visit.lab.indexOf(tag), 1)
    },

    showInput() {
      this.inputVisible = true
      this.$nextTick(_ => {
        this.$refs.saveTagInput.$refs.input.focus()
      })
    },

    handleInputConfirm() {
      if (this.inputValue) {
        if (this.visit.lab.indexOf(this.inputValue) < 0) {
          this.visit.lab.push(this.inputValue)
        }
      }
      this.inputVisible = false
      this.inputValue = ''
    }
  }
}
</script>
<style>
  .el-tag + .el-tag {
    margin-left: 10px;
  }
  .button-new-tag {
    margin-left: 10px;
    height: 32px;
    line-height: 30px;
    padding-top: 0;
    padding-bottom: 0;
  }
  .input-new-tag {
    width: 90px;
    margin-left: 10px;
    vertical-align: bottom;
  }
</style>

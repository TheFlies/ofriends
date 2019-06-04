<template>
  <el-dialog
    title="Update Visit"
    :visible.sync="isVisibleUpdate"
    width="30%"
    append-to-body
    :before-close="handleBackdropClick"
  >
    <el-form ref="visit" :model="visit" :rules="rules" label-width="130px" class="visit-form">
      <el-form-item label="Name" prop="name" required>
        <el-input v-model="visit.name" type="success" placeholder="Visit name" />
      </el-form-item>
      <el-form-item label="Lab" prop="lab" required>
        <el-tag v-for="tag in visit.lab" :key="tag" closable size="small" type="success" :disable-transitions="false" @close="handleClose(tag)">
          {{ tag }}
        </el-tag>
        <template>
          <el-select v-if="inputVisible" ref="saveTagInput" v-model="inputValue" placeholder="Select" class="input-new-tag" size="mini" @change="handleInputConfirm">
            <el-option v-for="item in options" :key="item.value" :label="item.label" :value="item.value" />
          </el-select>
          <el-button v-else class="button-new-tag" size="small" @click="showInput">
            + Add Lab
          </el-button>
        </template>
      </el-form-item>
      <el-form-item label="Arrival time" required>
        <el-col :span="22">
          <el-form-item prop="arrivedTime">
            <el-date-picker
              v-model="visit.arrivedTime"
              type="datetime"
              placeholder="Pick arrived time"
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
              style="width: 100%;"
            />
          </el-form-item>
        </el-col>
      </el-form-item>
      <el-form-item label="Created By" prop="createdBy">
        <el-input v-model="visit.createdBy" placeholder="Who in HR created the pre-approved visa?" />
      </el-form-item>

      <el-form-item label="Accommodation" prop="hotelStayed">
        <el-input v-model="visit.accommodation" placeholder="Where hotel customer stayed?" />
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
  name: 'VisitUpdate',
  props: {
    isVisibleUpdate: { type: Boolean, default: false },
    visit: {
      type: Object,
      default: () => ({
        id: '',
        name: '',
        lab: [],
        arrivedTime: new Date().getTime(),
        departureTime: new Date().getTime(),
        preApproveVisa: false,
        passportInfo: '',
        createdBy: '',
        hotelStayed: '',
        pickup: ''
      })
    }
  },
  data() {
    return {
      options: [{
        value: 'Lab 1',
        label: 'Lab 1'
      }, {
        value: 'Lab 2',
        label: 'Lab 2'
      }, {
        value: 'Lab 3',
        label: 'Lab 3'
      }, {
        value: 'Lab 4',
        label: 'Lab 4'
      }, {
        value: 'Lab 5',
        label: 'Lab 5'
      }, {
        value: 'Lab 6',
        label: 'Lab 6'
      }, {
        value: 'Lab 8',
        label: 'Lab 8'
      }
      ],
      rules: {},
      inputVisible: false,
      inputValue: ''
    }
  },
  methods: {
    handleBackdropClick() {
      this.$refs['visit'].resetFields()
      this.$emit('update:isVisibleUpdate', false)
    },
    submitForm(formName) {
      this.$refs[formName].validate(valid => {
        if (valid) {
          this.$emit('update:isVisibleUpdate', false)
          this.$emit('isUpdateVisit', true)
        } else {
          console.log('error submit!!')
          return false
        }
      })
    },
    resetForm(formName) {
      this.$refs[formName].resetFields()
      this.$emit('update:isVisibleUpdate', false)
    },
    handleClose(tag) {
      this.visit.lab.splice(this.visit.lab.indexOf(tag), 1)
    },

    showInput() {
      this.inputVisible = true
      this.$nextTick(_ => {
        this.$refs.saveTagInput.focus()
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

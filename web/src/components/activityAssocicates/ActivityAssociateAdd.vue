<template>
  <el-dialog
    title="Assign Activity"
    :visible.sync="isVisibleAssign"
    width="30%"
    append-to-body
    :before-close="handleBackdropClick"
  >
    <div style="text-align: center">
      <el-transfer
        v-model="value"
        style="text-align: left; display: inline-block"
        filterable
        :filter-method="filterMethod"
        filter-placeholder="State Abbreviations"
        :titles="['Activities', 'Assigned Activities']"
        :data="data"
      />
      <div style="margin-top: 20px">
        <el-button
          type="primary"
          @click="submit()"
        >
          Save
        </el-button>
        <el-button @click="resetForm('activityAssociate')">
          Cancel
        </el-button>
      </div>
    </div>
  </el-dialog>
</template>

<script>
import { getAllActivities } from '@/api/activity'

export default {
  name: 'ActivityAssociateAdd',

  props: {
    isVisibleAssign: { type: Boolean, default: false },
    assignedActivities: { type: Array, default: () => [] }
  },
  data() {
    return {
      data: [],
      value: []
    }
  },
  watch: {
    isVisibleAssign: function(val) {
        getAllActivities().then(resp => {
          this.value = []
          this.data = []
          if (resp.data != null) {            
            resp.data.forEach((activity, index) => {
            this.data.push({
              starttime: activity.starttime,
              key: index,
              endtime: activity.endtime,
              detail: activity.detail,
              participant: activity.participant,
              hotel: activity.hotel,
              initial: activity.id,
              label:  activity.name,
            })
            var position = this.assignedActivities.indexOf(activity.id)
            if (position >= 0) {
              this.value.push(index)
            }
          })
          }
        })
          .catch(err => {
            console.log(err)
          })
    }
  },
  methods: {
    filterMethod(query, item) {
      return item.initial.toLowerCase().indexOf(query.toLowerCase()) > -1
    },

    handleBackdropClick() {
      this.$emit('update:isVisibleAssign', false)
    },
    submit() {
      var activityAssociate = []
      this.value.forEach((val) => {
        activityAssociate.push(this.data[val])
      })
      this.$emit('update:isVisibleAssign', false)
      this.$emit('isActivityAssociateAdd', true, activityAssociate)
    },
    resetForm(formName) {
      this.$emit('update:isVisibleAssign', false)
    }
  }
}
</script>

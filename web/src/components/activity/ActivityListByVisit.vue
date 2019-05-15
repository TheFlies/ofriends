<template>
  <el-container>
    <el-main>
      <el-card class="box-card">
        <div slot="header" class="clearfix">
          <span style="font-size: 18px;">Activity Associate</span>
          <el-button type="primary" icon="el-icon-plus" plain style="float:right" @click="getActivityAssociate">
            Assign activity
          </el-button>
        </div>
        
        <div class="text item">
          <el-table v-loading="loading" :data="tableData.filter(data => !search || data.name.toLowerCase().includes(search.toLowerCase()))" style="width: 100%; margin:auto">
            <el-table-column type="index" :index="indexMethod" />
            <el-table-column label="Name" prop="name" width="150" />
            <el-table-column label="Start Time" width="130" sortable prop="startTime">
              <template slot-scope="scope">
                {{ getHumanDate(scope.row.startTime) }}
              </template>
            </el-table-column>
            <el-table-column label="End Time" width="130" sortable prop="endTime">
              <template slot-scope="scope">
                {{ getHumanDate(scope.row.endTime) }}
              </template>
            </el-table-column>
            <el-table-column label="Detail" width="300" prop="detail" />
            <el-table-column label="Participant" prop="participant" width="150" />
            <el-table-column label="Hotel" width="280" prop="hotel" />
            <el-table-column align="right">
              <ActivityAssociateAdd :assigned-activities.sync="assignedActivities" :is-visible-assign.sync="isVisibleAssign" @isActivityAssociateAdd="handleActivityAssociateAdd" />
              <!-- <ActivityUpdate
                :is-visible-update.sync="isVisibleUpdate"
                :activity.sync="activity"
                @isUpdateAct="handleUpdate"
              />
              <ActivityDelete :is-visible-delete.sync="isVisibleDelete" @isDeleteAct="handleDelete" />
              <ActivityAdd :is-visible-add.sync="isVisibleAdd" @isAddAct="handleAdd" /> -->
              <template slot-scope="scope">
                <el-button
                  size="mini"
                  @click="activity = scope.row; isVisibleUpdate = !isVisibleUpdate"
                >
                  Edit
                </el-button>
                <el-button
                  size="mini"
                  type="danger"
                  @click="isVisibleDelete = !isVisibleDelete; scopeActivity= scope"
                >
                  Delete
                </el-button>
              </template>
            </el-table-column>
          </el-table>
        </div>
      </el-card>      
    </el-main>
  </el-container>
</template>

<script>
import ActivityUpdate from '@/components/activity/ActivityUpdate.vue'
import ActivityDelete from '@/components/activity/ActivityDelete.vue'
import ActivityAdd from '@/components/activity/ActivityAdd.vue'
import ActivityAssociateAdd from '@/components/activityAssocicates/ActivityAssociateAdd.vue'
import { getHumanDate } from '@/utils/convert'
import { updateVisit } from '@/api/visit'

import {
  getActivityByID,
  createActivity,
  updateActivity,
  deleteActivityById
} from '@/api/activity'

export default {
  name: 'ActivityListByVisit',
  components: {
    ActivityUpdate,
    ActivityDelete,
    ActivityAdd,
    ActivityAssociateAdd
  },
  props: {
    visit: { type: Object}
  },
  data() {
    return {
      tableData: [],
      isVisibleAdd: false,
      isVisibleUpdate: false,
      isVisibleDelete: false,
      search: '',
      loading: false, // need to be true need fix
      activity: {},
      scopeActivity: {},
      assignedActivities: [],
      isVisibleAssign: false,
    }
  },
  mounted() {
    this.visit.activityID.forEach((id, index) => {
      getActivityByID(id).then(resp => {
        if (resp.data != null) {
          this.tableData.push(resp.data)
        }
      })
    })
    this.loading = false
  },
  methods: {
    getActivityAssociate: function() {
      this.assignedActivities = []
      this.tableData.forEach((activity, index) => {
        this.assignedActivities.push(activity.id)
      })
      this.isVisibleAssign = !this.isVisibleAssign
    },
    handleActivityAssociateAdd: function(isActivityAssociateAdd, updatedActivityAssociates) {
      this.tableData = []
      var activityIDList = []
      if(isActivityAssociateAdd) {
        updatedActivityAssociates.forEach((activity, index) => {
          activityIDList.push(activity.initial)
        })
      }
      this.visit.activityID = activityIDList
      updateVisit(this.visit)
        .then(resp => {
          this.$notify({
            title: 'Success',
            message: 'Update successfully!',
            type: 'success',
            position: 'bottom-right'
          })
        })
        .catch(err => {
          console.log(err)
          this.$notify.error({
            title: 'Error',
            message: err
          })
        })

      //Load list again
      this.visit.activityID.forEach((id, index) => {
        getActivityByID(id).then(resp => {
          if (resp.data != null) {
            this.tableData.push(resp.data)
          }
        })
      })
      this.loading = false
    },
    handleAdd: function(isAddAct, activity) {
      activity.visitID = this.activity.visitID
      if (isAddAct) {
        this.loading = true
        createActivity(activity)
          .then(resp => {
            this.$notify({
              title: 'Success',
              message: 'Update successfully!',
              type: 'success',
              position: 'bottom-right'
            })
            activity.id = resp.data.id
            this.tableData.splice(0, 0, activity)
          })
          .catch(err => {
            console.log(err)
            this.$notify.error({
              title: 'Error',
              message: err,
              position: 'bottom-right'
            })
          })
        this.loading = false
      }
    },
    handleUpdate: function(isUpdateAct) {
      if (isUpdateAct) {
        this.loading = true
        updateActivity(this.activity)
          .then(resp => {
            console.log(resp.data)
            this.$notify({
              title: 'Success',
              message: 'Update successfully!',
              type: 'success',
              position: 'bottom-right'
            })
          })
          .catch(err => {
            console.log(err)
            this.$notify.error({
              title: 'Error',
              message: err,
              position: 'bottom-right'
            })
          })
        this.loading = false
      }
    },
    handleDelete: function(isDeleteAct) {
      if (isDeleteAct) {
        this.loading = true
        deleteActivityById(this.scopeActivity.row.id)
          .then(resp => {
            this.tableData.splice(this.scopeActivity.$index, 1)
            this.$notify({
              title: 'Success',
              message: 'Delete successfully!',
              type: 'success',
              position: 'bottom-right'
            })
          })
          .catch(err => {
            console.log(err)
            this.$notify.error({
              title: 'Error',
              message: err,
              position: 'bottom-right'
            })
          })
      }
      this.loading = false
    },
    indexMethod(index) {
      return index * 1
    },
    getHumanDate
  }
}
</script>

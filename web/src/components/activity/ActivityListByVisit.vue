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
            <el-table-column label="Participant" prop="participant" width="150" />
            <el-table-column label="Note" width="300" prop="note" />
            <el-table-column align="right">
              <ActivityAssociateAdd :assigned-activities.sync="assignedActivities" :is-visible-assign.sync="isVisibleAssign" @isActivityAssociateAdd="handleActivityAssociateAdd" />
              <!-- <ActivityUpdate
                :is-visible-update.sync="isVisibleUpdate"
                :activity.sync="activity"
                @isUpdateAct="handleUpdate"
              />
               -->
              <ActivityAssociateUpdate :is-visible-update.sync="isVisibleUpdate" :activity.sync="activity" @isUpdateActivity="handleActivityAssociateUpdate" />
              <ActivityAssociateDelete :is-visible-delete.sync="isVisibleDelete" :activity-name.sync="activityName" @isDeleteActivity="handleAcitivityAssociateDelete" />
              <!-- <ActivityAdd :is-visible-add.sync="isVisibleAdd" @isAddAct="handleAdd" /> -->
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
                  @click="isVisibleDelete = !isVisibleDelete; scopeActivity= scope; activityName = scope.row.activityName"
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
// import ActivityUpdate from '@/components/activity/ActivityUpdate.vue'
// import ActivityDelete from '@/components/activity/ActivityDelete.vue'
// import ActivityAdd from '@/components/activity/ActivityAdd.vue'
import ActivityAssociateAdd from '@/components/activityAssocicates/ActivityAssociateAdd.vue'
import ActivityAssociateDelete from '@/components/activityAssocicates/ActivityAssociateDelete.vue'
import ActivityAssociateUpdate from '@/components/activityAssocicates/ActivityAssociateUpdate.vue'

import { getHumanDate } from '@/utils/convert'
import { updateVisit } from '@/api/visit'
import {
  getAcitivityVisitAssociatesByVisitID,
  createAcitivityVisitAssociate,
  modifyAcitivityVisitAssociates,
  deleteAcitivityVisitAssociateByID
} from '@/api/activityVisitAssociate'

import {
  getActivityByID,
  createActivity,
  updateActivity,
  deleteActivityById
} from '@/api/activity'

export default {
  name: 'ActivityListByVisit',
  components: {
    // ActivityUpdate,
    // ActivityDelete,
    // ActivityAdd,
    ActivityAssociateAdd,
    ActivityAssociateDelete,
    ActivityAssociateUpdate
  },
  props: {
    visit: { type: Object }
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
      activityName: ''
    }
  },
  mounted() {
    getAcitivityVisitAssociatesByVisitID(this.visit.id)
        .then(resp => {
          if (resp.data != null) {
            this.tableData = resp.data
            this.tableData.forEach((activity, index) => {
              this.assignedActivities.push(activity.id)
            })
          }
        })
        .catch(err => {
          console.log(err)          
        })
    this.loading = false
  },
  methods: {
    getActivityAssociate: function() {
      this.assignedActivities = []
      this.tableData.forEach((activity, index) => {
        this.assignedActivities.push(activity.activityID)
      })
      this.isVisibleAssign = !this.isVisibleAssign
    },
    handleActivityAssociateAdd: function(isActivityAssociateAdd, updatedActivityAssociates) {
      var currentAssignedActivityVisitAssocs = this.assignedActivities
      if (isActivityAssociateAdd) {
        updatedActivityAssociates.forEach((activity, index) => {
          // If currentAssignedGifts doesn not include this gift, assign this gift
          var position = currentAssignedActivityVisitAssocs.indexOf(activity.initial)
          if (position < 0) {
            var activityVisitAssocs = {}
            activityVisitAssocs.activityID = activity.initial
            activityVisitAssocs.visitID = this.visit.id
            activityVisitAssocs.customerID = this.customerId
            activityVisitAssocs.name = activity.label
            activityVisitAssocs.startTime = new Date().getTime()
            activityVisitAssocs.endTime = new Date().getTime()
            activityVisitAssocs.participant = ''
            activityVisitAssocs.note = activity.note
            createAcitivityVisitAssociate(activityVisitAssocs)
              .then(resp => {
                this.$notify({
                  title: 'Success',
                  message: 'Update successfully!',
                  type: 'success',
                  position: 'bottom-right'
                })
                activityVisitAssocs.id = resp.data.id
                this.tableData.splice(0, 0, activityVisitAssocs)
              })
              .catch(err => {
                console.log(err)
                this.$notify.error({
                  title: 'Error',
                  message: err,
                  position: 'bottom-right'
                })
              })
          } else {
            // Remove this gifts are not in currentAssignedGifts
            currentAssignedActivityVisitAssocs.splice(position, 1)
          }
        })
        // Current elements in currentAssignedActivityVisitAssocs are gifts that is not in updated Gift
        if (currentAssignedActivityVisitAssocs.length > 0) {
          currentAssignedActivityVisitAssocs.forEach((id) => {
            // Find to get gift associate ID, then delete this gift associate
            var data = this.tableData.filter(data => !id || data.activityID.includes(id))
            deleteAcitivityVisitAssociateByID(data[0].id)
              .then(resp => {
                var giftIndex = this.assignedActivities.indexOf(id)
                this.tableData.splice(giftIndex, 1)
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
          })
        }
        this.loading = false
      }
    },
     handleAcitivityAssociateDelete: function(isDeleteActivity) {
      if (isDeleteActivity) {
        this.loading = true
        deleteAcitivityVisitAssociateByID(this.scopeActivity.row.id)
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
    handleActivityAssociateUpdate: function(isUpdateActivity) {
      if (isUpdateActivity) {
        this.loading = true
        modifyAcitivityVisitAssociates(this.activity)
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
              message: err,
              position: 'bottom-right'
            })
          })
        this.loading = false
      }
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

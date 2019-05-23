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
            <el-table-column label="Name" width="250" prop="activityName" sortable />
            <el-table-column label="Start Time" width="200" sortable prop="startTime">
              <template slot-scope="scope">
                {{ getHumanDate(scope.row.startTime) }}
              </template>
            </el-table-column>
            <el-table-column label="End Time" width="200" sortable prop="endTime">
              <template slot-scope="scope">
                {{ getHumanDate(scope.row.endTime) }}
              </template>
            </el-table-column>
            <el-table-column label="Participant" prop="participant" width="300" />
            <el-table-column label="Note" prop="note" />
            <el-table-column align="right">
              <ActivityAssociateAdd :assigned-activities.sync="assignedActivities" :is-visible-assign.sync="isVisibleAssign" @isActivityAssociateAdd="handleActAssocAdd" />
              <ActivityAssociateUpdate :is-visible-update.sync="isVisibleUpdate" :activity.sync="activity" @isUpdateActivity="handleActAssocUpdate" />
              <ActivityAssociateDelete :is-visible-delete.sync="isVisibleDelete" :act-name.sync="activityName" @isDeleteActivity="handleActAssocDelete" />
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
import ActivityAssociateAdd from '@/components/activityAssocicates/ActivityAssociateAdd.vue'
import ActivityAssociateDelete from '@/components/activityAssocicates/ActivityAssociateDelete.vue'
import ActivityAssociateUpdate from '@/components/activityAssocicates/ActivityAssociateUpdate.vue'

import { getHumanDate } from '@/utils/convert'
import {
  getActVisitAssocsByVisitID,
  createActVisitAssoc,
  modifyActVisitAssocs,
  deleteActVisitAssocByID
} from '@/api/actvisitassoc'

export default {
  name: 'ActivityListByVisit',
  components: {
    ActivityAssociateAdd,
    ActivityAssociateDelete,
    ActivityAssociateUpdate
  },
  props: {
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
    getActVisitAssocsByVisitID(this.visit.id)
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
    handleActAssocAdd: function(isActivityAssociateAdd, updatedActAssocs) {
      var currentAssigns = this.assignedActivities
      if (isActivityAssociateAdd) {
        updatedActAssocs.forEach((activity, index) => {
          // If currentAssigns doesn not include this activity, assign this activity
          var position = currentAssigns.indexOf(activity.initial)
          if (position < 0) {
            var actVisitAssocs = {}
            actVisitAssocs.activityID = activity.initial
            actVisitAssocs.visitID = this.visit.id
            actVisitAssocs.customerID = this.customerId
            actVisitAssocs.activityName = activity.label
            actVisitAssocs.startTime = new Date().getTime()
            actVisitAssocs.endTime = new Date().getTime()
            actVisitAssocs.participant = ''
            actVisitAssocs.note = activity.note
            createActVisitAssoc(actVisitAssocs)
              .then(resp => {
                this.$notify({
                  title: 'Success',
                  message: 'Update successfully!',
                  type: 'success',
                  position: 'bottom-right'
                })
                actVisitAssocs.id = resp.data.id
                this.tableData.splice(0, 0, actVisitAssocs)
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
            // Remove this activities are not in currentAssigns
            currentAssigns.splice(position, 1)
          }
        })
        // Current elements in currentAssigns are activities that is not in updated activities
        if (currentAssigns.length > 0) {
          currentAssigns.forEach((id) => {
            // Find to get activity associate ID, then delete this activity associate
            var data = this.tableData.filter(data => !id || data.activityID.includes(id))
            deleteActVisitAssocByID(data[0].id)
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
    handleActAssocDelete: function(isDeleteActivity) {
      if (isDeleteActivity) {
        this.loading = true
        deleteActVisitAssocByID(this.scopeActivity.row.id)
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
              message: err.response.data,
              position: 'bottom-right'
            })
          })
      }
      this.loading = false
    },

    handleActAssocUpdate: function(isUpdateActivity) {
      if (isUpdateActivity) {
        this.loading = true
        modifyActVisitAssocs(this.activity)
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
    indexMethod(index) {
      return index * 1
    },
    getHumanDate
  }
}
</script>

<template>
  <el-card class="box-card">
    <div slot="header">
      <!-- <span>Activity</span> -->
      <el-tooltip class="item" effect="dark" content="Add activity" placement="right-start">
        <el-button type="primary" icon="el-icon-plus" plain @click="isVisibleAdd = !isVisibleAdd">
          New Activity
        </el-button>
      </el-tooltip>
    </div>
    <div class="text item">
      <el-table v-loading="loading" :data="tableData.filter(data => !search || data.name.toLowerCase().includes(search.toLowerCase()))" style="width: 100%; margin:auto">
        <el-table-column label="Name" width="250" prop="name" />
        <el-table-column label="Detail" prop="detail" />
        <el-table-column align="right">
          <template slot="header" slot-scope="scope">
            <el-input
              v-model="search"
              size="mini"
              placeholder="Type to search by name"
            />
          </template>
          <ActivityUpdate
            :is-visible-update.sync="isVisibleUpdate"
            :activity.sync="activity"
            @isUpdateAct="handleUpdate"
          />
          <ActivityDelete :is-visible-delete.sync="isVisibleDelete" :act-name.sync="actName" @isDeleteAct="handleDelete" />
          <ActivityAdd :is-visible-add.sync="isVisibleAdd" @isAddAct="handleAdd" />
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
              @click="isVisibleDelete = !isVisibleDelete; scopeActivity= scope; actName = scope.row.name"
            >
              Delete
            </el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>
  </el-card>
</template>

<script>
import ActivityUpdate from '@/components/activity/ActivityUpdate.vue'
import ActivityDelete from '@/components/activity/ActivityDelete.vue'
import ActivityAdd from '@/components/activity/ActivityAdd.vue'
import { getHumanDate } from '@/utils/convert'

import {
  getAllActivities,
  createActivity,
  updateActivity,
  deleteActivityByID
} from '@/api/activity'

export default {
  name: 'ActivityList',
  components: {
    ActivityUpdate,
    ActivityDelete,
    ActivityAdd
  },
  data() {
    return {
      tableData: [],
      isVisibleAdd: false,
      isVisibleUpdate: false,
      isVisibleDelete: false,
      search: '',
      loading: true,
      activity: {},
      actName: '',
      scopeActivity: {}
    }
  },
  mounted() {
    getAllActivities()
      .then(resp => {
        if (resp.data != null) {
          this.tableData = resp.data
        }
        this.loading = false
      })
      .catch(err => {
        console.log(err)
      })
  },
  methods: {
    handleAdd: function(isAddAct, activity) {
      // activity.visitID = this.activity.visitID
      if (isAddAct) {
        this.loading = true
        // activity.startTime = activity.startTime.toString()
        // activity.endTime = activity.endTime.toString()
        createActivity(activity)
          .then(resp => {
            this.$notify({
              title: 'Success',
              message: 'Update successfully!',
              type: 'success',
              position: 'bottom-right'
            })
            activity.id = resp.data.id
            // activity.startTime = activity.startTime.toString()
            // activity.endTime = activity.endTime.toString()
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
        deleteActivityByID(this.scopeActivity.row.id)
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
    indexMethod(index) {
      return index * 1
    },
    getHumanDate
  }
}
</script>

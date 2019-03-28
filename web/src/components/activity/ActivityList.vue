<template>
  <el-container style="background: #ecf5ff;">
    <el-main>
      <el-table
        v-loading="loading"
        :data="tableData.filter(data => !search || data.participant.toLowerCase().includes(search.toLowerCase()))"
        style="width: 100%; margin:auto"
      >
        <el-table-column type="index" :index="indexMethod" />
        <el-table-column label="Start Time" width="130" sortable prop="startTime" >
          <template slot-scope="scope">{{getHumanDate(scope.row.startTime)}}</template>
        </el-table-column>
        <el-table-column label="End Time" width="130" sortable prop="endTime" >
          <template slot-scope="scope">{{getHumanDate(scope.row.endTime)}}</template>
        </el-table-column>
        <el-table-column label="Detail" width="300" prop="detail" />
        <el-table-column label="Participant" prop="participant" width="150" />
        <el-table-column label="Hotel" width="280" prop="hotel" />
        <el-table-column align="right">
          <template slot="header" slot-scope="scope">
            <el-input
              v-model="search"
              size="mini"
              placeholder="Type to search base on participant"
            />
          </template>
          <ActivityUpdate
            :is-visible-update.sync="isVisibleUpdate"
            :activity.sync="activity"
            @isUpdateAct="handleUpdate"
          />
          <ActivityDelete :is-visible-delete.sync="isVisibleDelete" @isDeleteAct="handleDelete" />
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
              @click="isVisibleDelete = !isVisibleDelete; scopeActivity= scope"
            >
              Delete
            </el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-main>
  </el-container>
</template>

<script>
import ActivityUpdate from '@/components/activity/ActivityUpdate.vue'
import ActivityDelete from '@/components/activity/ActivityDelete.vue'
import ActivityAdd from '@/components/activity/ActivityAdd.vue'
import {getHumanDate} from "@/utils/convert";

import {
  getAllActivites,
  createActivity,
  updateActivity,
  deleteActivityById
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
      scopeActivity: {}
    }
  },
  mounted() {
    getAllActivites()
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
      activity.visitID = this.activity.visitID
      if (isAddAct) {
        this.loading = true
        activity.startTime = activity.startTime.toString()
        activity.endTime = activity.endTime.toString()
        createActivity(activity)
          .then(resp => {
            this.$notify({
              title: 'Success',
              message: 'Update successfully!',
              type: 'success'
            })
            activity.id = resp.data.id
            activity.startTime = activity.startTime.toString()
            activity.endTime = activity.endTime.toString()
            this.tableData.splice(0, 0, activity)
          })
          .catch(err => {
            console.log(err)
            this.$notify.error({
              title: 'Error',
              message: err
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
              type: 'success'
            })
          })
          .catch(err => {
            console.log(err)
            this.$notify.error({
              title: 'Error',
              message: err
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
              type: 'success'
            })
          })
          .catch(err => {
            console.log(err)
            this.$notify.error({
              title: 'Error',
              message: err
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

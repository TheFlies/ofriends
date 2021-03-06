<template>
  <el-main>
    <el-card class="box-card">
      <div slot="header" class="clearfix">
        <el-tooltip class="item" effect="dark" content="Add visit" placement="right-start">
          <el-button type="primary" icon="el-icon-plus" plain @click="isVisibleAdd = !isVisibleAdd">
            New Visit
          </el-button>
        </el-tooltip>
      </div>
      <div class="text item">
        <el-table stripe v-loading="loading" :data="tableData.filter(data => !search || data.name.toLowerCase().includes(search.toLowerCase()))" style="width: 100%; margin:auto">
          <el-table-column type="index" :index="indexMethod" />
          <el-table-column type="expand">
            <template slot-scope="scope">
              <el-tabs type="card">
                <el-tab-pane label="Activity">
                  <ActivityListByVisit :visit="scope.row" />
                </el-tab-pane>
                <el-tab-pane label="Customer">
                  <CustomerListByVisit :customer-id="scope.row.customerID" :visit="scope.row" />
                </el-tab-pane>
              </el-tabs>
            </template>
          </el-table-column>
          <el-table-column label="Name" width="250" sortable="" prop="name" />
          <el-table-column label="Lab" width="200" sortable prop="lab">
            <template slot-scope="scope">
              <el-tag v-for="tag in scope.row.lab" :key="tag" size="small" type="success" :disable-transitions="false">
                {{ tag }}
              </el-tag>
            </template>
          </el-table-column>
          <el-table-column label="Arrival time" width="200" prop="arrivedTime">
            <template slot-scope="scope">
              {{ getHumanDate(scope.row.arrivedTime) }}
            </template>
          </el-table-column>
          <el-table-column label="Departure time" width="200" prop="departureTime">
            <template slot-scope="scope">
              {{ getHumanDate(scope.row.departureTime) }}
            </template>
          </el-table-column>
          <el-table-column label="Accommodation" width="180" prop="accommodation" />
          <el-table-column label="Pickup" width="120" prop="pickup" />
          <el-table-column align="right">
            <template slot="header" slot-scope="scope">
              <el-input v-model="search" size="mini" placeholder="Type to search by name" />
            </template>
            <VisitUpdate :is-visible-update.sync="isVisibleUpdate" :visit.sync="visit" @isUpdateVisit="handleUpdate" />
            <VisitDelete :is-visible-delete.sync="isVisibleDelete" :visit-name.sync="visitName" @isDeleteVisit="handleDelete" />
            <VisitAdd :is-visible-add.sync="isVisibleAdd" @isAddVisit="handleAdd" />
            <template slot-scope="scope">
              <el-button size="mini" @click="visit = scope.row; isVisibleUpdate = !isVisibleUpdate">
                Edit
              </el-button>
              <el-button size="mini" type="danger" @click="isVisibleDelete = !isVisibleDelete; scopeVisit= scope; visitName = scope.row.name">
                Delete
              </el-button>
            </template>
          </el-table-column>
        </el-table>
      </div>
    </el-card>
  </el-main>
</template>

<script>
import VisitUpdate from '@/components/visits/VisitUpdate.vue'
import VisitDelete from '@/components/visits/VisitDelete.vue'
import VisitAdd from '@/components/visits/VisitAdd.vue'
import CustomerListByVisit from '@/components/customers/CustomerListByVisit.vue'
import ActivityListByVisit from '@/components/activity/ActivityListByVisit.vue'
import { indexMethod, getHumanDate } from '@/utils/convert'
import {
  getAllVisits,
  createVisit,
  updateVisit,
  deleteVisitById
} from '@/api/visit'

export default {
  name: 'VisitList',
  components: {
    VisitUpdate,
    VisitDelete,
    VisitAdd,
    ActivityListByVisit,
    CustomerListByVisit
  },
  data() {
    return {
      tableData: [],
      isVisibleAdd: false,
      isVisibleUpdate: false,
      isVisibleDelete: false,
      search: '',
      loading: true,
      visit: {},
      activities: {},
      scopeVisit: {},
      visitName: ''
    }
  },
  mounted() {
    getAllVisits()
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
    handleAdd: function(isAddVisit, visit) {
      if (isAddVisit) {
        this.loading = true
        createVisit(visit)
          .then(resp => {
            this.$notify({
              title: 'Success',
              message: 'Update successfully!',
              type: 'success',
              position: 'bottom-right'
            })
            visit.id = resp.data.id
            this.tableData.splice(0, 0, visit)
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
    handleUpdate: function(isUpdateVisit) {
      if (isUpdateVisit) {
        this.loading = true
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
        this.loading = false
      }
    },
    handleDelete: function(isDeleteVisit) {
      if (isDeleteVisit) {
        this.loading = true
        deleteVisitById(this.scopeVisit.row.id)
          .then(resp => {
            this.tableData.splice(this.scopeVisit.$index, 1)
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
    indexMethod,
    getHumanDate
  }
}
</script>

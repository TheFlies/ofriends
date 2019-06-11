<template>
  <el-container>
    <el-header style="width: 100%; margin:auto">
      <el-button
        type="primary"
        icon="el-icon-plus"
        plain
        style="float:right"
        @click="isVisibleAdd = !isVisibleAdd"
      >
        New visit
      </el-button>
    </el-header>
    <el-main>
      <el-table
        v-loading="loading"
        :data="tableData.filter(data => !search || data.name.toLowerCase().includes(search.toLowerCase()))"
        style="width: 100%; margin:auto"
      >
        <el-table-column type="expand">
          <template slot-scope="scope">
            <GiftListByVisit :visit-id="scope.row.id" :friend-id="visit.customerID" :friend-name="visit.customername" />
            <ActivityListByVisit :visit-id="scope.row.id" />
          </template>
        </el-table-column>
        <el-table-column type="index" :index="indexMethod" />
        <el-table-column label="Lab" width="70" sortable prop="lab" />
        <el-table-column label="Arrival time" width="130" prop="arrivedTime">
          <template slot-scope="scope">
            {{ getHumanDate(scope.row.arrivedTime) }}
          </template>
        </el-table-column>
        <el-table-column label="Departure time" width="150" prop="departureTime">
          <template slot-scope="scope">
            {{ getHumanDate(scope.row.departureTime) }}
          </template>
        </el-table-column>
        <el-table-column label="Created by" width="120" sortable prop="createdBy" />
        <el-table-column label="Accommodation" width="180" prop="hotelStayed" />
        <el-table-column label="Pickup" width="120" prop="pickup" />
        <el-table-column align="right">
          <VisitUpdate
            :is-visible-update.sync="isVisibleUpdate"
            :visit.sync="visit"
            @isUpdateVisit="handleUpdate"
          />
          <VisitDelete :is-visible-delete.sync="isVisibleDelete" @isDeleteVisit="handleDelete" />
          <VisitAdd :is-visible-add.sync="isVisibleAdd" @isAddVisit="handleAdd" />
          <template slot-scope="scope">
            <el-button
              size="mini"
              @click="visit = scope.row; isVisibleUpdate = !isVisibleUpdate"
            >
              Edit
            </el-button>
            <el-button
              size="mini"
              type="danger"
              @click="isVisibleDelete = !isVisibleDelete; scopeVisit= scope"
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
import VisitUpdate from '@/components/visits/VisitUpdate.vue'
import VisitDelete from '@/components/visits/VisitDelete.vue'
import VisitAdd from '@/components/visits/VisitAdd.vue'
import ActivityListByVisit from '@/components/activity/ActivityListByVisit.vue'
import GiftListByVisit from '@/components/gifts/GiftListByVisit.vue'
import { indexMethod, getHumanDate } from '@/utils/convert'
import {
  getAllVisitsByCustomerID,
  createVisit,
  updateVisit,
  deleteVisitById
} from '@/api/visit'

export default {
  name: 'VisitListByVisit',
  components: {
    VisitUpdate,
    VisitDelete,
    VisitAdd,
    ActivityListByVisit,
    GiftListByVisit
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
      scopeVisit: {}
    }
  },
  created() {
    this.visit.customerID = this.$route.params.id
    this.visit.customerName = this.$route.params.customername
  },
  mounted() {
    getAllVisitsByCustomerID(this.visit.customerID)
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
              message: err,
              position: 'bottom-right'
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
              message: err,
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

<template>
  <el-container style="background: #ecf5ff;">
    <el-main>
      <el-table
        v-loading="loading"
        :data="tableData.filter(data => !search || data.Customer.name.toLowerCase().includes(search.toLowerCase()))"
        style="width: 100%; margin:auto"
      >
        <el-table-column type="index" :index="indexMethod" />
        <el-table-column label="Arrived Time" width="180" sortable prop="Visit.arrivedTime">
          <template slot-scope="scope">
            {{ getHumanDate(scope.row.Visit.arrivedTime) }}
          </template>
        </el-table-column>
        <el-table-column label="Departure Time" width="180" sortable prop="Visit.departureTime">
           <template slot-scope="scope">
            {{ getHumanDate(scope.row.Visit.departureTime) }}
          </template>
        </el-table-column>
        <el-table-column label="Customer Name" width="180" sortable prop="Customer.name">
        </el-table-column>
        <el-table-column label="Title" prop="Customer.title" width="80" />
        <el-table-column label="Position" prop="Customer.position" width="150" />
        <el-table-column label="Project" sortable width="180" prop="Customer.project" />
        <el-table-column label="Gift" width="180" sortable prop="priority">
         <template slot-scope="scope">
          <el-popover
            v-for="gift in scope.row.Gifts" :key="gift.id"
            placement="left"
            width="400"
            trigger="click">
            <el-form 
            ref="gift"
            :model="gift"
            label-width="80px">
              <h3 style="text-align: center; background-color: #FFC0CB; color: #fff; padding: 10px">Gift Info</h3>
              <el-form-item label="Name" style="color: #259dd8;">
                <span>{{gift.name}}</span>
              </el-form-item>
               <el-form-item label="Idea" style="color: #259dd8;">
                <span>{{gift.idea}}</span>
              </el-form-item>
               <el-form-item label="Size" style="color: #259dd8;">
                <span>{{gift.size}}</span>
              </el-form-item>
               <el-form-item label="Price" style="color: #259dd8;">
                <span>{{gift.price}}</span>
              </el-form-item>
               <el-form-item label="Link" style="color: #259dd8;">
                 <a :href="`${gift.link}`" target="_blank">
                  {{gift.link}}
                  </a>
              </el-form-item>
              <el-form-item label="Description" style="color: #259dd8;">
                <span>{{gift.description}}</span>
              </el-form-item>
            </el-form>
            <el-button slot="reference">{{ gift.name }}</el-button>
          </el-popover>
          </template>
        </el-table-column>
        <el-table-column label="Activity" sortable prop="priority">
           <template slot-scope="scope">
            <el-popover
            v-for="activity in scope.row.Activities" :key="activity.id"
            placement="left"
            width="400"
            trigger="click">
            <el-form 
            ref="activity"
            :model="activity"
            label-width="80px">
             <h3 style="text-align: center; background-color: #259dd8; color: #fff; padding: 10px">Activity Info</h3>
              <el-form-item label="Name" style="color: #259dd8;">
                <span>{{activity.name}}</span>
              </el-form-item>
               <el-form-item label="Detail" style="color: #259dd8;">
                <span>{{activity.detail}}</span>
              </el-form-item>
            </el-form>
            <el-button slot="reference">{{ activity.name }}</el-button>
          </el-popover>
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column align="right">
          <template slot="header" slot-scope="scope">
            <el-input
              v-model="search"
              size="mini"
              placeholder="Type to search base on user name"
            />
          </template>
        </el-table-column>
      </el-table>
    </el-main>
  </el-container>
</template>
<script>
import DeleteUser from "@/components/user/UserDelete.vue";
import { getAllUsers, updateUser, deleteUser } from "@/api/user";
import { getTimelineByDay } from "@/api/timeline"
import { getHumanDate } from '@/utils/convert'

export default {
  name: "ListUsers",
  components: {
    DeleteUser
  },
  data() {
    return {
      tableData: [],
      isVisibleAdd: false,
      isVisibleDelete: false,
      search: "",
      loading: true,
      user: {},
      userName: "",
      scopeUser: {}
    };
  },
  mounted() {
    getTimelineByDay(Date.now())
      .then(resp => {
        console.log(resp);
        if (resp.data != null) {
          this.tableData = resp.data
          console.log(resp.data);
        }
        this.loading = false;
      })
      .catch(err => {
        console.log(err);
      });
  },
  methods: {
    indexMethod(index) {
      return index * 1;
    },
    getHumanDate,
  }
};
</script>
<style scoped lang="stylus">
.link-type {
  color: #1989fa;
  text-decoration: underline;
}
</style>

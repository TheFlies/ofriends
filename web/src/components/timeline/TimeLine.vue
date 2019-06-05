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
        <el-table-column label="Title" prop="Customer.title" width="120" />
        <el-table-column label="Position" prop="Customer.position" width="150" />
        <el-table-column label="Project" sortable width="180" prop="Customer.project" />
        <el-table-column label="Gift" width="180" sortable prop="priority">
         <template slot-scope="scope">
            <el-tag v-for="gift in scope.row.Gifts" :key="gift" size="small" type="success" :disable-transitions="false">
              {{ gift.name }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="Activity" sortable prop="priority">
           <template slot-scope="scope">
            <el-tag v-for="activity in scope.row.Activities" :key="activity" size="small" type="success" :disable-transitions="false">
              {{ activity.name }}
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
    handleUpdate: function(priority) {
      this.loading = true;
      this.user.priority=priority;
      console.log(this.user);
      updateUser(this.user)
        .then(resp => {
          console.log(resp);
          this.$notify({
            title: "Success",
            message: "Update successfully!",
            type: "success"
          });
        })
        .catch(err => {
          console.log(err);
          this.$notify.error({
            title: "Error",
            message: err
          });
        });
      this.loading = false;
  },
    handleDelete: function(isDeleteUser) {
      if (isDeleteUser) {
        this.loading = true;
        deleteUser(this.scopeUser.row.id)
          .then(resp => {
            this.tableData.splice(this.scopeUser.$index, 1);
            this.$notify({
              title: "Success",
              message: "Delete successfully!",
              type: "success"
            });
          })
          .catch(err => {
            console.log(err);
            this.$notify.error({
              title: "Error",
              message: err
            });
          });
      }
      this.loading = false;
    },
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

<template>
  <el-container style="background: #ecf5ff;">
    <el-main>
      <el-table
        v-loading="loading"
        :data="tableData.filter(data => !search || data.username.toLowerCase().includes(search.toLowerCase()))"
        style="width: 100%; margin:auto"
      >
        <el-table-column type="index" :index="indexMethod" />
        <el-table-column label="User Name" width="130" sortable prop="username">
        </el-table-column>
        <el-table-column label="Full Name" width="130" sortable prop="fullname">
        </el-table-column>
        <el-table-column label="Email" prop="email" width="150" />
        <el-table-column label="Delivery Center" width="280" prop="delivery_center" />
        <el-table-column label="Priority" width="280" prop="priority">
          <template slot-scope="scope">
            {{ getRoleName(scope.row.priority) }}
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
          <DeleteUser :is-visible-delete.sync="isVisibleDelete" :user-name.sync="userName" @isDeleteUser="handleDelete" />
          <template slot-scope="scope">
            <el-button
              size="mini"
              @click="user=scope.row; handleUpdate(2)"
            >
              User Role
            </el-button>
            <el-button
              size="mini"
              @click="user=scope.row; handleUpdate(3)"
            >
              Admin Role
            </el-button>
            <el-button
              size="mini"
              @click="user=scope.row; handleUpdate(1)"
            >
              None
            </el-button>
            <el-button
              size="mini"
              type="danger"
              @click="isVisibleDelete = !isVisibleDelete; scopeUser= scope; userName=scope.row.username"
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
import DeleteUser from "@/components/user/UserDelete.vue";
import { getAllUsers, updateUser, deleteUser } from "@/api/user";
import { getRoleName } from '@/utils/convert'

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
    getAllUsers()
      .then(resp => {
        console.log(resp);
        if (resp.data != null) {
          this.tableData = resp.data;
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
    getRoleName
  }
};
</script>
<style scoped lang="stylus">
.link-type {
  color: #1989fa;
  text-decoration: underline;
}
</style>
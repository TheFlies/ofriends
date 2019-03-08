<template>
  <el-container>
    <el-header style="width: 80%; margin:auto">
      <el-button
        type="primary"
        icon="el-icon-plus"
        plain
        style="float:right"
        v-on:click="isVisibleAdd = !isVisibleAdd"
      >New friend</el-button>
    </el-header>
    <el-main>
      <el-table
        v-loading="loading"
        :data="tableData.filter(data => !search || data.name.toLowerCase().includes(search.toLowerCase()))"
        style="width: 80%; margin:auto"
      >
        <el-table-column type="index" :index="indexMethod"></el-table-column>
        <el-table-column label="Name" width="120" sortable prop="name"></el-table-column>
        <el-table-column label="Position" width="120" sortable prop="position"></el-table-column>
        <el-table-column label="Project" width="120" sortable prop="project"></el-table-column>
        <el-table-column label="Age" width="120" sortable prop="age"></el-table-column>
        <el-table-column label="Company" width="120" sortable prop="company"></el-table-column>
        <el-table-column label="Country" width="120" sortable prop="country"></el-table-column>
        <el-table-column label="City" width="120" sortable prop="city"></el-table-column>
        <el-table-column label="Food Note" width="120" prop="foodNote"></el-table-column>
        <el-table-column label="Family Note" width="120" prop="familyNote"></el-table-column>
        <el-table-column label="Next Visit Note" width="120" prop="nextVisitNote"></el-table-column>
        <el-table-column align="right">
          <template slot="header" slot-scope="scope">
            <el-input v-model="search" size="mini" placeholder="Type to search"/>
          </template>
          <EditFriend
            :isVisibleUpdate.sync="isVisibleUpdate"
            @isUpdateFriend="handleUpdate"
            :friend.sync="friend"
          />
          <DeleteFriend
            :isVisibleDelete.sync="isVisibleDelete"
            @isDeleteFriend="handleDelete"
            :friendName.sync="friendName"
          />
          <AddFriend :isVisibleAdd.sync="isVisibleAdd" @isAddFriend="handleAdd"/>
          <template slot-scope="scope">
            <el-button
              size="mini"
              v-on:click="friend = scope.row; isVisibleUpdate = !isVisibleUpdate"
            >Edit</el-button>
            <el-button
              size="mini"
              type="danger"
              v-on:click="isVisibleDelete = !isVisibleDelete; scopeFriend = scope; friendName = scope.row.name"
            >Delete</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-main>
  </el-container>
</template>

<script>
import EditFriend from "@/components/friends/FriendUpdate.vue";
import DeleteFriend from "@/components/friends/FriendDelete.vue";
import AddFriend from "@/components/friends/FriendAdd.vue";
export default {
  name: "ListFriends",
  components: {
    EditFriend,
    DeleteFriend,
    AddFriend
  },
  data() {
    return {
      tableData: [],
      isVisibleAdd: false,
      isVisibleUpdate: false,
      isVisibleDelete: false,
      search: "",
      loading: true,
      friend: {},
      friendName: "",
      scopeFriend: {}
    };
  },
  mounted() {
    // We already set the axios baseURL to the backend service in main.js file.
    this.$http
      .get("/friends")
      .then(resp => {
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
    handleAdd: function(isAddFriend, friend) {
      if (isAddFriend) {
        this.loading = true;
        this.$http
          .post("http://localhost:8080/friends", {
            Name: friend.name,
            Title: friend.title,
            Position: friend.position,
            Project: friend.project,
            Age: parseInt(friend.age, 10),
            Company: friend.company,
            Country: friend.country,
            City: friend.city,
            FoodNote: friend.foodNote,
            FamilyNote: friend.familyNote,
            NextVisitNote: friend.nextVisitNote
          })
          .then(resp => {
            this.$notify({
              title: "Success",
              message: "Update successfully!",
              type: "success"
            });
            friend.id = resp.data.id;
            this.tableData.splice(0, 0, friend);
          })
          .catch(err => {
            console.log(err);
            this.$notify.error({
              title: "Error",
              message: err
            });
          });
        this.loading = false;
      }
    },
    handleUpdate: function(isUpdateFriend) {
      if (isUpdateFriend) {
        this.loading = true;
        this.$http
          .put("http://localhost:8080/friends", this.friend)
          .then(resp => {
            console.log(resp.data);
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
      }
    },
    handleDelete: function(isDeleteFriend) {
      if (isDeleteFriend) {
        this.loading = true;
        console.log(this.scopeFriend.row);
        this.$http
          .delete("http://localhost:8080/friends/" + this.scopeFriend.row.id)
          .then(resp => {
            this.tableData.splice(this.scopeFriend.$index, 1);
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
    }
  }
};
</script>

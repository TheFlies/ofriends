<template>
  <el-container>
    <el-header style="width: 80%; margin:auto">
      <el-button
        type="primary"
        icon="el-icon-plus"
        plain
        style="float:right"
        v-on:click="isVisibleAdd = !isVisibleAdd"
      >New visit</el-button>
    </el-header>
    <el-main>
      <el-table
        v-loading="loading"
        :data="tableData.filter(data => !search || data.name.toLowerCase().includes(search.toLowerCase()))"
        style="width: 80%; margin:auto"
      >
        <el-table-column type="index" :index="indexMethod"></el-table-column>
        <el-table-column label="Lab" width="120" sortable prop="name"></el-table-column>
        <el-table-column label="Arrived Date" width="120" sortable prop="position"></el-table-column>
        <el-table-column label="Departure Date" width="120" sortable prop="project"></el-table-column>
        <el-table-column label="Pre-approved visa" width="120" sortable prop="age"></el-table-column>
        <el-table-column label="Passport Info" width="120" sortable prop="country"></el-table-column>
        <el-table-column label="Created By" width="120" sortable prop="company"></el-table-column>
        <el-table-column label="Hotel Stayed" width="120" sortable prop="city"></el-table-column>
        <el-table-column label="Pickup" width="120" prop="foodNote"></el-table-column>
        <el-table-column align="right">
          <VisitUpdate
            :isVisibleUpdate.sync="isVisibleUpdate"
            @isUpdateVisit="handleUpdate"
            :visit.sync="visit"
          />
          <VisitDelete
            :isVisibleDelete.sync="isVisibleDelete"
            @isDeleteVisit="handleDelete"
            :visitInfo.sync="visitInfo"
          />
          <VisitAdd :isVisibleAdd.sync="isVisibleAdd" @isAddVisit="handleAdd"/>
          <template slot-scope="scope">
            <el-button
              size="mini"
              v-on:click="friend = scope.row; isVisibleUpdate = !isVisibleUpdate"
            >Edit</el-button>
            <el-button
              size="mini"
              type="danger"
              v-on:click="isVisibleDelete = !isVisibleDelete; scopeVisit= scope; visitInfo = scope.row.name"
            >Delete</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-main>
  </el-container>
</template>

<script>
import VisitUpdate from "@/components/visits/VisitUpdate.vue";
import VisitDelete from "@/components/visits/VisitDelete.vue";
import VisitAdd from "@/components/visits/VisitAdd.vue";
export default {
  name: "VisitList",
   components: {
    VisitUpdate,
    VisitDelete,
    VisitAdd
  },
  data() {
    return {
      tableData: [],
      isVisibleAdd: false,
      isVisibleUpdate: false,
      isVisibleDelete: false,
      search: "",
      loading: false, // need to be true need fix
      visit: Object,
      scopeVisit: Function,
      visitInfo:""
    };
  },
  methods: {
    handleAdd: function(isAddVisit, visit) {
      if (isAddVisit) {
        this.loading = true;
        console.log(visit);
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

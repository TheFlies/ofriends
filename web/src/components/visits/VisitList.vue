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
        <el-table-column label="Lab" width="70" sortable prop="lab"></el-table-column>
        <el-table-column label="Arrived Date" width="130" prop="arrivedTime"></el-table-column>
        <el-table-column label="Departure Date" width="150" prop="departureTime"></el-table-column>
        <el-table-column label="Pre-approved visa" width="120">
          <template slot-scope="scope">
            <el-checkbox v-model="scope.row.preApproveVisa"></el-checkbox>
          </template>
        </el-table-column>
        <el-table-column label="Passport Info" width="280" prop="passportInfo"></el-table-column>
        <el-table-column label="Created By" width="120" sortable prop="createdBy"></el-table-column>
        <el-table-column label="Hotel Stayed" width="180" prop="hotelStayed"></el-table-column>
        <el-table-column label="Pickup" width="120" prop="pickup"></el-table-column>
        <el-table-column align="right">
          <VisitUpdate
            :isVisibleUpdate.sync="isVisibleUpdate"
            @isUpdateVisit="handleUpdate"
            :visit.sync="visit"
          />
          <VisitDelete :isVisibleDelete.sync="isVisibleDelete" @isDeleteVisit="handleDelete"/>
          <VisitAdd :isVisibleAdd.sync="isVisibleAdd" @isAddVisit="handleAdd"/>
          <template slot-scope="scope">
            <el-button
              size="mini"
              v-on:click="visit = scope.row; isVisibleUpdate = !isVisibleUpdate"
            >Edit</el-button>
            <el-button
              size="mini"
              type="danger"
              v-on:click="isVisibleDelete = !isVisibleDelete; scopeVisit= scope"
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
      visit: {},
      scopeVisit: {}
    };
  },
  created() {
    this.visit.friendID = this.$route.params.id;
  },
  mounted() {
    // We already set the axios baseURL to the backend service in main.js file.
    this.$http
      .get("/visits/friend/" + this.visit.friendID)
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
    handleAdd: function(isAddVisit, visit) {
      if (isAddVisit) {
        this.loading = true;
        this.$http
          .post("http://localhost:8080/visits", {
            Lab: visit.lab,
            ArrivedTime: visit.arrivedTime.toString(),
            DepartureTime: visit.departureTime.toString(),
            PreApproveVisa: visit.preApproveVisa,
            PassportInfo: visit.passportInfo,
            CreatedBy: visit.createdBy,
            HotelStayed: visit.hotelStayed,
            Pickup: visit.pickup,
            FriendID: visit.friendID
          })
          .then(resp => {
            this.$notify({
              title: "Success",
              message: "Update successfully!",
              type: "success"
            });
            visit.id = resp.data.id;
            visit.arrivedTime = visit.arrivedTime.toString();
            visit.departureTime = visit.departureTime.toString();
            this.tableData.splice(0, 0, visit);
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
    handleUpdate: function(isUpdateVisit) {
      if (isUpdateVisit) {
        this.loading = true;
        this.$http
          .put("http://localhost:8080/visits", this.visit)
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
    handleDelete: function(isDeleteVisit) {
      if (isDeleteVisit) {
        this.loading = true;
        console.log(this.scopeVisit.row);
        this.$http
          .delete("http://localhost:8080/visits/" + this.scopeVisit.row.id)
          .then(resp => {
            this.tableData.splice(this.scopeVisit.$index, 1);
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

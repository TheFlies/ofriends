<template lang="pug">
el-container
  el-main
    el-table(v-loading="loading"
      :data="tableData.filter(data => !search || data.username.toLowerCase().includes(search.toLowerCase()))"
      style="width: 100%; margin:auto")
      el-table-column(type="index" :index="indexMethod")
      el-table-column(label="User Name" width="120" sortable="" prop="username")
      el-table-column(label="Full Name" width="120" sortable="" prop="fullname")
      el-table-column(label="Email" width="120" sortable="" prop="email")
      el-table-column(label="DeliveryCenter" width="120" sortable="" prop="delivery_center")
      el-table-column(label="RoleName" width="120" sortable="" prop="role")
      el-table-column(align="right")
        template(slot="header" slot-scope="scope")
          el-input(v-model="search" size="mini" placeholder="Type to search")
        template(slot-scope="scope")
          el-button(size="mini" @click="user = scope.row; isVisibleUpdate = !isVisibleUpdate")
            | Edit
          el-button(size="mini" type="danger" @click="isVisibleDelete = !isVisibleDelete; scopeUser = scope; username = scope.row.username")
            | Delete
  <!-- edit-user(:is-visible-update.sync="isVisibleUpdate" :user.sync="user" @isUpdateUser="handleUpdate") -->
  <!-- delete-user(:is-visible-delete.sync="isVisibleDelete" :user-name.sync="username" @isDeleteUser="handleDelete") -->
</template>

<script>
import EditCustomer from '@/components/customers/CustomerUpdate.vue'
import DeleteCustomer from '@/components/customers/CustomerDelete.vue'
import AddCustomer from '@/components/customers/CustomerAdd.vue'
import {
  getAllUsers,
  updateUser,
} from '@/api/user'

export default {
  name: 'ListUsers',
  components: {
    EditCustomer,
    DeleteCustomer,
    AddCustomer
  },
  data() {
    return {
      tableData: [],
      isVisibleAdd: false,
      isVisibleUpdate: false,
      isVisibleDelete: false,
      search: '',
      loading: true,
      user: {},
      UserName: '',
      scopeUser: {}
    }
  },
  mounted() {
    getAllUsers()
      .then(resp => {
        console.log(resp)
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
    handleUpdate: function(isUpdateUser) {
      if (isUpdateUser) {
        this.loading = true
        updateUser(this.user)
          .then(resp => {
            console.log(resp)
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
    indexMethod(index) {
      return index * 1
    }
  }
}
</script>
<style scoped lang="stylus">
.link-type {
  color: #1989fa;
  text-decoration: underline;
}
</style>

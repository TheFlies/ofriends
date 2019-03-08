<template>
  <el-container>
    <el-header class="gift-list-header">
      <el-row>
        <el-col :span="12">
          <div style="text-align: left; font-size: 25px">
            <span>Gift</span>
          </div>
        </el-col>
        <el-col :span="12">
          <div style="text-align: right; font-size: 25px">
            <el-tooltip class="item" effect="dark" content="Add gift" placement="right-start">
              <i class="el-icon-plus" style="margin-right: 15px"
                    v-on:click="isVisibleAdd = !isVisibleAdd"></i>
            </el-tooltip>
          </div>
        </el-col>
      </el-row>
    </el-header>
    <el-main>
      <el-table
        v-loading="loading"
        :data="tableData.filter(data => !search || data.name.toLowerCase().includes(search.toLowerCase()))"
        style="width: 100%">
        <el-table-column
          label="Name"
          prop="name"
          sortable>
        </el-table-column>
        <el-table-column
          label="Idea"
          prop="idea"
          sortable>
        </el-table-column>
        <el-table-column
          label="Size"
          prop="size"
          sortable>
        </el-table-column>
        <el-table-column
          label="Quantity"
          prop="quantity"
          sortable>
        </el-table-column>
        <el-table-column
          label="Price"
          prop="price"
          sortable>
        </el-table-column>
        <el-table-column
          label="Link"
          prop="link">
        </el-table-column>
        <el-table-column
          label="Description"
          prop="description">
        </el-table-column>
        <el-table-column
          align="right">
          <template slot="header" slot-scope="scope">
            <el-input
              v-model="search"
              size="mini"
              placeholder="Type to search"/>
          </template>
          <GiftUpdate :isVisibleUpdate.sync="isVisibleUpdate" @isUpdateGift="handleUpdate" :gift.sync="gift"/>
          <GiftDelete :isVisibleDelete.sync="isVisibleDelete" @isDeleteGift="handleDelete" :giftName.sync="giftName"/>
          <GiftAdd :isVisibleAdd.sync="isVisibleAdd" @isAddGift="handleAdd"/>
          <template slot-scope="scope">
            <el-button
              size="mini"
              v-on:click="gift = scope.row; isVisibleUpdate = !isVisibleUpdate">Edit</el-button>
            <el-button
              size="mini"
              type="danger"
              v-on:click="isVisibleDelete = !isVisibleDelete; scopeGift = scope; giftName = scope.row.name">Delete</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-main>
  </el-container>
</template>

<script>
import GiftUpdate from '@/components/gifts/GiftUpdate.vue'
import GiftDelete from '@/components/gifts/GiftDelete.vue'
import GiftAdd from '@/components/gifts/GiftAdd.vue'
export default {
  name: 'ListGift',
  components: {
    GiftUpdate,
    GiftDelete,
    GiftAdd
  },
  mounted () {
    // We already set the axios baseURL to the backend service in main.js file.
    this.$http.get('http://localhost:8080/gifts')
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

  data () {
    return {
      tableData: [],
      isVisibleUpdate: false,
      isVisibleDelete: false,
      isVisibleAdd: false,
      search: '',
      gift: Object,
      giftName: '',
      loading: true,
      scopeGift: Function
    }
  },

  methods: {
    handleAdd: function (isAddGift, gift) {
      if (isAddGift) {
        this.loading = true
        this.$http.post('http://localhost:8080/gifts', gift)
          .then(resp => {
            this.$notify({
              title: 'Success',
              message: 'Update successfully!',
              type: 'success'
            })
            gift.id = resp.data.id
            this.tableData.splice(0, 0, gift)
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
    handleUpdate: function (isUpdateGift) {
      if (isUpdateGift) {
        this.loading = true
        this.$http.put('http://localhost:8080/gifts', this.gift)
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
    handleDelete: function (isDeleteGift) {
      if (isDeleteGift) {
        this.loading = true
        this.$http.delete('http://localhost:8080/gifts/' + this.scopeGift.row.id)
          .then(resp => {
            this.tableData.splice(this.scopeGift.$index, 1)
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
    }
  }
}
</script>

<style>
  .gift-list-header {
    background-color: #E4E7ED;
    color: #333;
    line-height: 60px;
  }
</style>

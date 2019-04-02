<template>
  <el-container>
    <el-main>
      <el-table
        v-loading="loading"
        :data="tableData.filter(data => !search || data.name.toLowerCase().includes(search.toLowerCase()))"
        style="width: 100%"
      >
        <el-table-column
          label="Name"
          prop="name"
          sortable
        />
        <el-table-column
          label="Idea"
          prop="idea"
          sortable
        />
        <el-table-column
          label="Size"
          prop="size"
          sortable
        />
        <el-table-column
          label="Quantity"
          prop="quantity"
          sortable
        />
        <el-table-column
          label="Price"
          prop="price"
          sortable
        />
        <el-table-column
          label="Link"
          prop="link"
        />
        <el-table-column
          label="Description"
          prop="description"
        />
        <el-table-column
          align="right"
        >
          <template slot="header" slot-scope="scope">
            <el-input
              v-model="search"
              size="mini"
              placeholder="Type to search"
            />
          </template>
          <GiftUpdate
            :is-visible-update.sync="isVisibleUpdate"
            :gift.sync="gift"
            @isUpdateGift="handleUpdate"
          />
          <GiftDelete
            :is-visible-delete.sync="isVisibleDelete"
            :gift-name.sync="giftName"
            @isDeleteGift="handleDelete"
          />
          <GiftAdd
            :is-visible-add.sync="isVisibleAdd"
            @isAddGift="handleAdd"
          />
          <template slot-scope="scope">
            <el-button
              size="mini"
              @click="gift = scope.row; isVisibleUpdate = !isVisibleUpdate"
            >
              Edit
            </el-button>
            <el-button
              size="mini"
              type="danger"
              @click="isVisibleDelete = !isVisibleDelete; scopeGift = scope; giftName = scope.row.name"
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
import GiftUpdate from '@/components/gifts/GiftUpdate.vue'
import GiftDelete from '@/components/gifts/GiftDelete.vue'
import GiftAdd from '@/components/gifts/GiftAdd.vue'
import { getAllGifts, createGifts, modifyGifts, deleteGiftById } from '@/api/gift'

export default {
  name: 'ListGift',
  components: {
    GiftUpdate,
    GiftDelete,
    GiftAdd
  },

  data() {
    return {
      tableData: [],
      isVisibleUpdate: false,
      isVisibleDelete: false,
      isVisibleAdd: false,
      search: '',
      gift: {},
      giftName: '',
      loading: true,
      scopeGift: {}
    }
  },
  mounted() {
    getAllGifts().then(resp => {
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
    handleAdd: function(isAddGift, gift) {
      if (isAddGift) {
        this.loading = true
        createGifts(gift).then(resp => {
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
    handleUpdate: function(isUpdateGift) {
      if (isUpdateGift) {
        this.loading = true
        modifyGifts(this.gift).then(resp => {
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
    handleDelete: function(isDeleteGift) {
      if (isDeleteGift) {
        this.loading = true
        deleteGiftById(this.scopeGift.row.id).then(resp => {
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

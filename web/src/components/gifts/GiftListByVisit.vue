<template>
  <el-container style="background: #f0f9eb;">
    <el-header style="width: 100%; margin:auto">
      <el-button
        type="success"
        icon="el-icon-plus"
        plain
        style="float:right"
        @click="isVisibleAdd = !isVisibleAdd"
      >
        New gift
      </el-button>
    </el-header>
    <el-main>
      <el-table
        v-loading="loading"
        :data="tableData.filter(data => !search || data.name.toLowerCase().includes(search.toLowerCase()))"
        style="width: 100%"
      >
        <el-table-column label="Name" prop="name" sortable />
        <el-table-column label="Idea" prop="idea" sortable />
        <el-table-column label="Size" prop="size" sortable />
        <el-table-column label="Quantity" prop="quantity" sortable />
        <el-table-column label="Price" prop="price" sortable />
        <el-table-column label="Link" prop="link" />
        <el-table-column label="Description" prop="description" />
        <el-table-column align="right">
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
          <GiftAdd :is-visible-add.sync="isVisibleAdd" @isAddGift="handleAdd" />
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
import {
  getAllGiftsByVisitID,
  createGifts,
  modifyGifts,
  deleteGiftById
} from '@/api/gift'

export default {
  name: 'GiftListByVisit',
  components: {
    GiftUpdate,
    GiftDelete,
    GiftAdd
  },
  props: {
    visitId: { type: String, default: '' }
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
  created() {
    this.gift.visitID = this.visitId
  },
  mounted() {
    getAllGiftsByVisitID(this.gift.visitID)
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
    handleAdd: function(isAddGift, gift) {
      if (isAddGift) {
        this.loading = true
        createGifts(gift)
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
    handleUpdate: function(isUpdateGift) {
      if (isUpdateGift) {
        this.loading = true
        modifyGifts(this.gift)
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
    handleDelete: function(isDeleteGift) {
      if (isDeleteGift) {
        this.loading = true
        deleteGiftById(this.scopeGift.row.id)
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

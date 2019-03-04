<template>
  <el-table
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
      <GiftUpdate :isVisibleUpdate.sync="isVisibleUpdate" :gift.sync="gift"/>
      <GiftDelete :isVisibleDelete.sync="isVisibleDelete" :rowIndex.sync="rowIndex" :giftName.sync="giftName" :giftId.sync="giftId"/>
      <template slot-scope="scope">
        <el-button
          size="mini"
          v-on:click="handleEdit(scope.$index, scope.row); isVisibleUpdate = !isVisibleUpdate">Edit</el-button>
        <el-button
          size="mini"
          type="danger"
          v-on:click="handleDelete(scope.$index, scope.row.id, scope.row.name); isVisibleDelete = !isVisibleDelete">Delete</el-button>
      </template>
    </el-table-column>
  </el-table>
</template>

<script>
import GiftUpdate from '@/components/gifts/GiftUpdate.vue'
import GiftDelete from '@/components/gifts/GiftDelete.vue'

export default {
  name: 'ListGift',
  components: {
    GiftUpdate,
    GiftDelete
  },
  mounted () {
    // We already set the axios baseURL to the backend service in main.js file.
    this.$http.get('http://localhost:8080/api/v1/gifts')
      .then(resp => {
        this.tableData = resp.data
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
      search: '',
      gift: Object,
      rowIndex: 0,
      giftName: '',
      giftId: ''
    }
  },

  methods: {
    handleEdit: function (index, row) {
      this.gift = row
    },
    handleDelete: function (index, id, name) {
      this.rowIndex = index
      this.giftId = id
      this.giftName = name
    }
  }
}
</script>

<template>
  <div class="customers-container">
    <p class="title-font">Customers</p>
    <el-row :gutter="20" class="space-bottom">
      <el-col :span="12">
        <el-input placeholder="Filter by name" v-model="filterName"></el-input>
      </el-col>
      <el-col :span="6">
        <el-input placeholder="Arrive" v-model="filterArrive"></el-input>
      </el-col>
      <el-col :span="6">
        <el-input placeholder="Depart" v-model="filterDepart"></el-input>
      </el-col>
    </el-row>
    <div class="space-bottom">
      <el-radio-group v-model="filterTime">
        <el-radio-button label="Current"></el-radio-button>
        <el-radio-button label="Upcomming"></el-radio-button>
        <el-radio-button label="Served"></el-radio-button>
        <el-radio-button label="All"></el-radio-button>
      </el-radio-group>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Customers',
  data: () => {
    return {
      filterName: '',
      filterTime: 'Current'
    }
  }
}
</script>
<style lang="scss" scoped>
.customers-container {
  margin: auto;
  padding: 0 20px;
  max-width: 1366px;
}
</style>

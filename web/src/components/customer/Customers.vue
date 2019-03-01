<template>
  <div class="customers-container">
    <p class="title-font">Customers</p>
    <el-row :gutter="20" class="space-bottom">
      <el-col :span="12">
        <el-input placeholder="Filter by name" v-model="filterName"></el-input>
      </el-col>
      <el-col :span="6" class="text-right">
        <el-date-picker v-model="filterArrive" type="date" placeholder="Arrive day"> </el-date-picker>
      </el-col>
      <el-col :span="6" class="text-right">
        <el-date-picker v-model="filterDepart" type="date" placeholder="Depart day"> </el-date-picker>
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
    <div class="customers">
      <Customer/>
      <Customer/>
      <Customer/>
      <Customer/>
      <Customer/>
      <Customer/>
      <Customer/>
      <Customer/>
    </div>
  </div>
</template>

<script>
import Customer from './Customer'

export default {
  name: 'Customers',
  data: () => {
    return {
      filterName: '',
      filterTime: 'Current',
      filterArrive: new Date(),
      filterDepart: new Date()
    }
  },
  components: {
    Customer
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

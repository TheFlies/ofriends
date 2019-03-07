<template>
  <div class="customers-container">
    <p class="title-font">Customers</p>
    <el-row :gutter="20" class="space-bottom">
      <el-col :span="12">
        <el-input placeholder="Filter by name" name="name" v-model="filter.name" @input="updateFilter"></el-input>
      </el-col>
      <el-col :span="6" class="text-right">
        <el-date-picker v-model="filter.arrive" type="date" placeholder="Arrive day"> </el-date-picker>
      </el-col>
      <el-col :span="6" class="text-right">
        <el-date-picker v-model="filter.depart" type="date" placeholder="Depart day"> </el-date-picker>
      </el-col>
    </el-row>
    <div class="space-bottom">
      <el-radio-group v-model="filter.time">
        <el-radio-button label="Current"></el-radio-button>
        <el-radio-button label="Up Comming"></el-radio-button>
        <el-radio-button label="Served"></el-radio-button>
        <el-radio-button label="All"></el-radio-button>
      </el-radio-group>
    </div>
    <div class="customers">
      <Customer/>
    </div>
  </div>
</template>

<script>
import Customer from './Customer'

export default {
  name: 'Customers',
  components: {
    Customer
  },
  computed: {
    filter: {
      get: function () {
        return this.$store.state.customers.filter
      }
    },
    customers: {
      get: function () {
        return this.$store.state.customers.customers
      }
    }
  },
  methods: {
    updateFilter () {
      this.$store.commit('customers/setFilter', this.filter)
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

<template>
  <div class="customer">
    <div class="left">
      <span>{{customer.title}} {{customer.name}} <label>({{customer.position}})</label></span>
      <span class="project">{{customer.project}}</span>
    </div>
    <div class="right">
      <span>{{customer.arrive}}</span>
      <span> => </span>
      <span>{{customer.depart}}</span>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Customer',
  props: {
    customer: {
      name: String,
      title: String,
      position: String,
      project: String,
      arrive: String,
      depart: String
    }
  }
}
</script>
<style lang="scss" scoped>
.customer {
  border: 1px solid $--primary;
  border-radius: 5px;
  margin-bottom: 5px
}
.left {
  width: 70%;
  display: inline-block;
  padding: 1em;
  position: relative;

}
.right {
  width: 30%;
  display: inline-block;
  text-align: center;
  padding: 1em;
  border-left: 1px solid $--primary;
  position: relative;
}
.project {
  position: absolute;
  display: inline-block;
  right: 20px;
}
</style>

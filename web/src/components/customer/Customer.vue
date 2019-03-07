<template>
  <div class="customer">
    <div class="left">
      <span>John <label>(architect lead)</label></span>
      <span class="project">KBTG</span>
    </div>
    <div class="right">
      <span>21-12-2018</span>
      <span> => </span>
      <span>22-12-2019</span>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Customer'
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

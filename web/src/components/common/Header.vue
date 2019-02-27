<template>
  <el-header>
    <div class="flex full-width">
      <div class="home-link flex-item">
         <a href="/">OFriends</a>
      </div>
      <div class="menu-link flex-item">
        <!-- TODO: replace with a Cog icon -->
        <a href="/">Config</a>
        <!-- TODO: replace with a account icon -->
        <a href="/">Account</a>
      </div>
    </div>
  </el-header>
</template>

<script>
export default {
  name: 'Header'
}
</script>

<style lang="scss">
.el-header {
  background-color: $--primary-color;
  min-height: 36px;
  @include flex
}
.home-link {
  min-height: 36px;
  text-align: left;
  a {
    color: $--white;
    font-size: $--font-bg;
    text-decoration: none;
  }
}
.menu-link {
  min-height: 36px;
  text-align: right;
  a {
    color: $--white;
    font-size: $--font-md;
    text-decoration: none;
    margin-left: 20px;
  }
}
</style>

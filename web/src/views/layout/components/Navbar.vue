<template lang="pug">
  .navbar
    hamburger.hamburger-container(:toggle-click="toggleSidebar" :is-active="sidebar.opened")
    breadcrumb
    el-dropdown.avatar-container(trigger="click")
      .avatar-wrapper
        img.user-avatar(src="https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif?imageView2/1/w/80/h/80")
        i.el-icon-caret-bottom
      el-dropdown-menu.user-dropdown(slot="dropdown")
        router-link.inlineBlock(to="/")
          el-dropdown-item Home
        el-dropdown-item(divided)
          span(style="display:block" @click="logout") LogOut
</template>

<script>
import { mapGetters } from 'vuex'
import Breadcrumb from '@/components/Breadcrumb'
import Hamburger from '@/components/Hamburger'

export default {
  components: {
    Breadcrumb,
    Hamburger
  },
  computed: {
    ...mapGetters([
      'sidebar'
    ])
  },
  methods: {
    toggleSidebar() {
      this.$store.dispatch('ToggleSidebar')
    },
    logout() {
      this.$store.dispatch('LogOut').then(() => {
        location.reload()
      })
    }
  }
}
</script>

<style lang="stylus" scoped>
.navbar
  height: 50px
  line-height: 50px
  box-shadow: 0 1px 3px 0 rgba(0, 0, 0, 0.12), 0 0 3px 0 rgba(0, 0, 0, 0.04)
  .hamburger-container
    line-height: 58px
    height: 50px
    float: left
    padding: 0 10px
  .screenfull
    position: absolute
    right: 90px
    top: 16px
    color: red
  .avatar-container
    height: 50px
    display: inline-block
    position: absolute
    right: 35px
    .avatar-wrapper
      cursor: pointer
      margin-top: 5px
      position: relative
      line-height: initial
      .user-avatar
        width: 40px
        height: 40px
        border-radius: 10px
      .el-icon-caret-bottom
        position: absolute
        right: -20px
        top: 25px
        font-size: 12px
</style>

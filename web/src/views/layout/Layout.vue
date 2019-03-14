<template lang="pug">
  div(:class="classObj").app-wrapper
    .drawer-bg(v-if="device === 'mobile' && sidebar.opened" @click="handleClickOutside")
    sidebar.sidebar-container
    .main-container
      navbar
      app-main
</template>

<script>
import { Navbar, Sidebar, AppMain } from './components'
import ResizeMixin from './mixin/ResizeHandler'

export default {
  name: 'Layout',
  components: {
    Navbar,
    Sidebar,
    AppMain
  },
  mixins: [ResizeMixin],
  computed: {
    sidebar() {
      return this.$store.state.app.sidebar
    },
    device() {
      return this.$store.state.app.device
    },
    classObj() {
      return {
        hideSidebar: !this.sidebar.opened,
        openSidebar: this.sidebar.opened,
        withoutAnimation: this.sidebar.withoutAnimation,
        mobile: this.device === 'mobile'
      }
    }
  },
  methods: {
    handleClickOutside() {
      this.$store.dispatch('CloseSideBar', { withoutAnimation: false })
    }
  }
}
</script>

<style lang="stylus" scoped>
  @import '~@/styles/mixin'
  .app-wrapper
    clearfix()
    position: relative
    height: 100%
    width: 100%
    &.mobile.openSidebar
      position: fixed
      top: 0
  .drawer-bg
    background: #000
    opacity: 0.3
    width: 100%
    top: 0
    height: 100%
    position: absolute
    z-index: 999
</style>

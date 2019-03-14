<template lang="pug">
  el-breadcrumb.app-breadcrumb(separator="/")
    transition-group(name="breadcrumb")
      el-breadcrumb-item(v-for="(it, idx) in levelList" :key="it.path")
        span.no-redirect(v-if="it.redirect === 'noredirect' || idx == levelList.length-1") {{ it.meta.title }}
        a(v-else @click.prevent="handleLink(it)") {{ it.meta.title}}
</template>

<script>
import pathToRegexp from 'path-to-regexp'

export default {
  data() {
    return {
      levelList: null
    }
  },
  watch: {
    $route() {
      this.getBreadcrumb()
    }
  },
  created() {
    this.getBreadcrumb()
  },
  methods: {
    getBreadcrumb() {
      let matched = this.$route.matched.filter(item => item.name)

      // const first = matched[0]
      // if (first && first.name !== 'dashboard') {
      //   matched = [{ path: '/dashboard', meta: { title: 'Dashboard' }}].concat(matched)
      // }

      this.levelList = matched.filter(item => item.meta && item.meta.title && item.meta.breadcrumb !== false)
    },
    pathCompile(path) {
      // To solve this problem https://github.com/PanJiaChen/vue-element-admin/issues/561
      const { params } = this.$route
      var toPath = pathToRegexp.compile(path)
      return toPath(params)
    },
    handleLink(item) {
      const { redirect, path } = item
      if (redirect) {
        this.$router.push(redirect)
        return
      }
      this.$router.push(this.pathCompile(path))
    }
  }
}
</script>

<style lang="stylus" scoped>
  .app-breadcrumb.el-breadcrumb
    display inline-block
    font-size 14px
    line-height 50px
    margin-left 10px
    .no-redirect
      color #97a8be
      cursor text
</style>

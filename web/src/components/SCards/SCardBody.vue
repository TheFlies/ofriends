<template lang="pug">
  .body(:style="cssProps")
    slot
</template>

<script>
export default {
  computed: {
    cssProps() {
      return {
        'overflow': 'auto'
      }
    }
  }
}
</script>

<style lang="stylus" scoped>
mq-lg()
  @media (min-width: 1200px)
    {block}

$border = 1px solid rgba(191, 191, 191, 0.4)
.body
  background: #fff
  border: $border
  border-top: 0
  padding: var(--inner-margin)
  +mq-lg()
    height var(--body-height)
  p
    font-size: 14px
    line-height: 18px
    margin-bottom: var(--inner-margin)
  img
    display: block
    width: 100%
</style>

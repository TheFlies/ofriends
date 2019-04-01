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
mq-xs()
  @media (min-width: 320px)
    {block}
mq-sm()
  @media (min-width: 480px)
    {block}
mq-md()
  @media (min-width: 720px)
    {block}
mq-lg()
  @media (min-width: 1200px)
    {block}

$border = 1px solid rgba(191, 191, 191, 0.4)
.body
  background: #fff
  border: $border
  border-top: 0
  padding: var(--inner-margin)
  +mq-xs()
    height var(--body-height)
  +mq-sm()
    height var(--body-height)
  +mq-md()
    height var(--body-height)
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

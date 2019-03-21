<template lang="pug">
  .body(:style="cssProps")
    slot
</template>

<script>
export default {
  props: {
    innerMargin: {
      type: Number,
      default: 15
    },
    numberSize: {
      type: Number,
      default: 35
    },
    cardHeight: {
      type: Number,
      default: 400
    },
    cardWidth: {
      type: Number,
      default: 450,
    }
  },
  computed: {
    cssProps() {
      return {
        '--inner-margin': `${this.innerMargin}px`,
        '--body-height': `${this.bodyHeight()}px`,
        'overflow': 'auto'
      }
    }
  },
  methods: {
    bodyHeight() {
      const headHeight = this.numberSize + 50
      return this.cardHeight - headHeight
    }
  }
}
</script>

<style lang="stylus" scoped>
mq-lg()
  @media (min-width: 1000px)
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

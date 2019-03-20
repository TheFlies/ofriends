<template lang="pug">
  .head(:style="cssProps")
    slot
    h2
      span.small {{ subTitle }}
      |  {{ title }}
</template>

<script>
export default {
  props: {
    title: {
      type: String,
      required: true
    },
    subTitle: {
      type: String
    },
    innerMargin: {
      type: Number,
      default: 15
    },
    numberSize: {
      type: Number,
      default: 35
    }
  },
  computed: {
    cssProps() {
      return {
        '--inner-margin': `${this.innerMargin}px`,
        '--number-size': `${this.numberSize}px`
      }
    }
  }
}
</script>

<style lang="stylus" scoped>
mq-sm()
  @media (min-width: 480px)
    {block}

.head
  position: relative
  display: flex
  align-items: center
  color: #fff
  font-weight: 400
  .number-box
    display: inline
    float: left
    margin var(--inner-margin)
    padding: 10px
    font-size: var(--number-size)
    line-height: var(--number-size)
    font-weight: 600
    background: rgba(0, 0, 0, 0.17)
  h2
    text-transform: uppercase
    font-size: 1.3rem
    font-weight: inherit
    letter-spacing: 2px
    margin: 0
    padding-bottom: 6px
    line-height: 1rem
    +mq-sm()
      font-size 165%
      line-height 1.2rem
    span
      display: block
      font-size: 0.6rem
      margin: 0
      +mq-sm()
        font-size 0.8rem
</style>

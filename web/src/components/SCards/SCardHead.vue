<template lang="pug">
  .head(:style="cssProps")
    slot
    h2
      span.small {{ subTitle }}
      |  {{ title }}
    .left hello
</template>

<script>
export default {
  props: {
    title: {
      type: String,
      required: true
    },
    subTitle: {
      type: String,
      default: null
    },
    innerMargin: {
      type: Number,
      default: 15
    },
    numberSize: {
      type: Number,
      default: 35
    },
    titleSize: {
      type: Number,
      default: 1.3
    },
    subTitleSize: {
      type: Number,
      default: 0.6
    },
    color: {
      type: String,
      default: '#fff'
    }
  },
  computed: {
    cssProps() {
      return {
        '--inner-margin': `${this.innerMargin}px`,
        '--number-size': `${this.numberSize}px`,
        '--title-size': `${this.titleSize}rem`,
        '--sub-title-size': `${this.subTitleSize}rem`,
        '--bg-color': this.color
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
  background-color: var(--bg-color)
  font-weight: 400
  text-align: left
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
    font-size: var(--title-size)
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
      font-size: var(--sub-title-size)
      margin: 0
      +mq-sm()
        font-size 0.8rem
  .left
    text-align right
    display: inline
    position absolute
    width 50px
    height 30px
    right 5px
    top 5px
</style>

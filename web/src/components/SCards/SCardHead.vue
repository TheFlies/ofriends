<template lang="pug">
  .head()
    slot
    a(:href="`/friends/${id}`")
      h2
        span.small {{ subTitle }}
        |  {{ title }}
    .left
      svg-icon(:name="companyIcon(icon)")
    .btm(v-if="attime")
      svg-icon(v-if="arrival" name="arrival")
      svg-icon(v-else name="departure")
      | &nbsp; {{ attime }}
</template>

<script>
export default {
  props: {
    id: {
      type: String,
      required: true
    },
    title: {
      type: String,
      required: true
    },
    subTitle: {
      type: String,
      default: null
    },
    icon: {
      type: String,
      default: null
    },
    arrival: {
      type: Boolean,
      default: true
    },
    attime: {
      type: String,
      default: ''
    }
  },
  methods: {
    companyIcon(company) {
      return (company || 'example').toLowerCase()
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
  .btm
    text-align right
    padding-left 15px
    padding-right 10px
    padding-bottom 10px
    display inline
    position absolute
    background-color var(--bg-color)
    height 28px
    font-size 14px
    right 0px
    bottom -35px
    border-radius 0 0 0 15px
    z-index 1
    font-family 'Courier New', Courier, monospace
    font-weight 600
</style>

<template lang="pug">
.s-card(:class="classProps" :style="cssProps")
  s-card-head(
    :title="item.title"
    :sub-title="item.subtitle"
  )
    .number-box
      span {{ item.dc }}
  s-card-body
    span {{ item.familyNote }}
    br
    span {{ item.footNote }}
    br
    span {{ item.nextVisitNote }}
    br
    span {{ item.country }}
    br
    span {{ item.city }}
    .footer
      svg-icon(name="gift")
      | &nbsp;
      svg-icon(name="places")
      | &nbsp;
      svg-icon(name="hobby")
</template>

<script>
import SCardHead from './SCardHead.vue'
import SCardBody from './SCardBody.vue'

export default {
  components: { SCardHead, SCardBody },
  props: {
    item: {
      type: Object,
      required: true,
    },
    step: {
      type: Number,
      required: true
    },
    color: {
      type: String,
      default: '#332211'
    },
    total: {
      type: Number,
      required: true
    },
    stagger: {
      type: Number,
      default: 180
    },
    titleSize: {
      type: Number,
      default: 1.3
    },
    subTitleSize: {
      type: Number,
      default: 0.6
    }
  },
  computed: {
    cssProps() {
      const props = {
        backgroundColor: this.color,
        '--bg-color': this.color,
        '--margin': `${(this.stagger)}px`,
        order: this.calculateOrder()
      }
      return props
    },
    classProps() {
      const props = []
      if (this.stagger !== 0) {
        props.push('mgt2')
      }
      return props
    }
  },
  methods: {
    calculateOrder() {
      const items = this.total
      const counter = Math.ceil(items / 2)

      let ord = Math.ceil(this.step / 2)
      if (this.step % 2 === 0) {
        ord += counter
      }
      return ord
    }
  }
}
</script>

<style scoped>
@media (min-width: 1200px) {
  .mgt2 {
    margin-top: var(--margin) !important;
  }
}
.head::after {
  border-color: var(--bg-color);
}
</style>

<style lang="stylus" scoped>
.body
  position relative
  margin-top 10px
  padding-top 35px
  .footer
    text-align left
    width 100%
    position absolute
    display inline-block
    bottom 0
    padding 8px 16px 8px 16px
    left 0
    height 34px
    overflow hidden
</style>


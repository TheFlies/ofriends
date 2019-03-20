<template lang="pug">
  .s-card-wrapper(:style="cssProps")
    s-card(v-for="(item, i) in data" :color="item.color" :total="data.length" :step="i + 1" :stagger="i===1?stagger:0")
      s-card-head(:title="item.title" :sub-title="item.subtitle" :inner-margin="innerMargin" :number-size="numberSize")
        .number-box
          span {{ item.dc }}
      s-card-body(:card-height="cardHeight" :inner-margin="innerMargin" :number-size="numberSize")
        p {{ item.content }}
        img(src="http://placehold.it/1000x500" alt="Graphic")
</template>

<script>
import SCard from './SCard.vue'
import SCardHead from './SCardHead.vue'
import SCardBody from './SCardBody.vue'

export default {
  props: {
    data: {
      type: Array,
      required: true
    },
    stagger: {
      type: Number,
      default: 180
    },
    cardHeight: {
      type: Number,
      default: 400
    },
    cardWidth: {
      type: Number,
      default: 450
    },
    outerMargin: {
      type: Number,
      default: 90
    },
    markerSize: {
      type: Number,
      default: 9
    },
    markerColor: {
      type: String,
      default: '#bdbdbd'
    },
    markerBorderColor: {
      type: String,
      default: '#f7f7f7'
    }
  },
  components: { SCard, SCardHead, SCardBody },
  data() {
    return {
      innerMargin: 15,
      numberSize: 35
    }
  },
  computed: {
    cssProps() {
      return {
        '--c-width': `${this.containerWidth()}px`,
        '--c-height': `${this.containerHeight()}px`,
        '--card-width': `${this.cardWidth}px`,
        '--card-height': `${this.cardHeight}px`,
        '--outer-margin': `${this.outerMargin}px`,
        '--half-outer-margin': `${this.outerMargin / 2}px`,
        '--marker-size': `${this.markerSize}px`,
        '--marker-dist-add': `${this.markerDist() + 1}px`,
        '--marker-dist-sub': `${this.markerDist() - 1}px`,
        '--timeline': this.markerColor,
        '--background': this.markerBorderColor,
      }
    }
  },
  methods: {
    pad(n, width, z) {
      z = z || '0'
      n = n + ''
      return n.length >= width ? n : new Array(width - n.length + 1).join(z) + n
    },
    containerHeight() {
      const rows = Math.ceil(this.data.length / 2)
      return rows * (this.cardHeight + this.outerMargin) + this.stagger
    },
    containerWidth() {
      return this.cardWidth * 2 + this.outerMargin * 3
    },
    markerDist() {
      return this.cardWidth + this.outerMargin / 2 - this.markerSize / 2
    }
  }
}
</script>

<style lang="stylus" scoped>
/* Media Queries */
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
  @media (min-width: 1000px)
    {block}

$box-shadow = 0px 1px 22px 4px rgba(0, 0, 0, 0.07)
$border = 1px solid rgba(191, 191, 191, 0.4)

*
  box-sizing: border-box

$oarrow
  position: absolute
  content: ""
  width: 0
  height: 0
  border-top: 15px solid transparent
  border-bottom: 15px solid transparent

$omarker
  position: absolute
  content: ""
  width var(--marker-size)
  height var(--marker-size)
  background-color: var(--timeline)
  border-radius var(--marker-size)
  box-shadow: 0px 0px 2px 8px var(--background)

.s-card-wrapper
  position: relative
  margin: auto
  +mq-lg()
    width var(--c-width)
    height var(--c-height)
    display: flex
    flex-flow: column wrap
    margin: 0 auto
  &::after
    z-index: 1
    content: ""
    position: absolute
    top: 0
    bottom: 0
    left: 50%
    border-left: $border
    +mq-lg()
      border-left: 1px solid var(--timeline)
.s-card
  position: relative
  display: block
  margin: 10px auto 80px
  max-width: 94%
  z-index: 2
  +mq-sm()
    max-width: 60%
    box-shadow: $box-shadow
  +mq-md()
    max-width 40%
  +mq-lg()
    height var(--card-height)
    max-width var(--card-width)
    margin var(--outer-margin)
    margin-top var(--half-outer-margin)
    margin-bottom var(--half-outer-margin)
    &:nth-child(odd)
      margin-right var(--half-outer-margin)
      .head::after
        @extend $oarrow
        border-left-width: 15px
        border-left-style: solid
        left: 100%
      .head::before
        @extend $omarker
        left var(--marker-dist-add)
    &:nth-child(even)
      margin-left var(--half-outer-margin)
      .head::after
        @extend $oarrow
        border-right-width: 15px
        border-right-style: solid
        right: 100%
      .head::before
        @extend $omarker
        right: var(--marker-dist-sub)
</style>

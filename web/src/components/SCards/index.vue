<template lang="pug">
  #s-card-list
    div(v-if="!data.length")
      span.title No data
    .s-card-wrapper(v-else :style="cssProps")
      s-card(
        v-for="(item, i) in data"
        :color="item.color"
        :total="data.length"
        :step="i + 1"
        :stagger="i===1?stagger:0"
        :key="item.id"
        :item="item"
        :marker-color="markerColor"
        :marker-border-color="markerBorderColor"
      )
</template>

<script>
import SCard from './SCard.vue'

export default {
  components: { SCard },
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
      default: 400,
      validator: val => val && val >= 190
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
    timeLineColor: {
      type: String,
      default: '#bdbdbd'
    },
    markerColor: {
      type: String,
      default: '#bdbdbd'
    },
    markerBorderColor: {
      type: String,
      default: '#f7f7f7'
    },
    arrowSize: {
      type: Number,
      default: 15
    },
    numberSize: {
      type: Number,
      default: 35
    },
    innerMargin: {
      type: Number,
      default: 15
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
      return {
        '--c-width': `${this.containerWidth()}px`,
        '--c-height': `${this.containerHeight()}px`,
        '--card-width': `${this.cardWidth}px`,
        '--card-height': `${this.cardHeight}px`,
        '--outer-margin': `${this.outerMargin}px`,
        '--half-outer-margin': `${this.outerMargin / 2}px`,
        '--marker-size': `${this.markerSize}px`,
        '--half-marker-size': `-${this.markerSize / 2 - 1}px`,
        '--marker-dist-add': `${this.markerDist() + 1}px`,
        '--marker-dist-sub': `${this.markerDist() - 1}px`,
        '--time-line-color': this.timeLineColor,
        '--arrow-size': `${this.arrowSize}px`,
        '--inner-margin': `${this.innerMargin}px`,
        '--number-size': `${this.numberSize}px`,
        '--body-height': `${this.bodyHeight()}px`,
        // this is variable for the css
        '--left-marker-dist': '',
        '--title-size': `${this.titleSize}rem`,
        '--sub-title-size': `${this.subTitleSize}rem`
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
    },
    bodyHeight() {
      const headHeight = this.numberSize + 50
      return this.cardHeight - headHeight
    }
  }
}
</script>

<style lang="stylus" src="@/styles/scard.styl">
</style>

<template lang="pug">
  .s-card-wrapper(:style="cssProps")
    s-card(v-for="(item, i) in data" :color="item.color" :total="data.length" :step="i + 1" :stagger="i===1?stagger:0" :key="item.title")
      s-card-head(
        :title="item.title"
        :sub-title="item.subtitle"
        :inner-margin="innerMargin"
        :number-size="numberSize"
        :color="item.color"
        :title-size="titleSize"
        :sub-title-size="subTitleSize"
      )
        .number-box
          span {{ item.dc }}
      s-card-body(:card-height="cardHeight" :card-width="cardWidth" :inner-margin="innerMargin" :number-size="numberSize")
        el-container
          el-card.no-margin
            p {{ item.content }}
          el-footer(height="20px") {{ item.arrival || new Date() }}
        //- img(src="http://placehold.it/1000x500" alt="Graphic")
</template>

<script>
import SCard from './SCard.vue'
import SCardHead from './SCardHead.vue'
import SCardBody from './SCardBody.vue'

export default {
  components: { SCard, SCardHead, SCardBody },
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
    },
    arrowSize: {
      type: Number,
      default: 15,
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
        '--marker-dist-add': `${this.markerDist() + 1}px`,
        '--marker-dist-sub': `${this.markerDist() - 1}px`,
        '--timeline': this.markerColor,
        '--background': this.markerBorderColor,
        '--arrow-size': `${this.arrowSize}px`
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

<style lang="stylus" scoped src="@/styles/scard.styl">
</style>

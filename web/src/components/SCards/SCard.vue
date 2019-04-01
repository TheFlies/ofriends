<template lang="pug">
  //- el-container(v-loading="loading")
  .s-card(:class="classProps" :style="cssProps" v-loading="loading")
    s-card-head(
      v-if="data"
      :id="data.friendId"
      :title="`${data.title}. ${data.name}`"
      :sub-title="`${data.position} at ${data.company}`"
      :icon="data.company"
      key="head"
      :arrival="!todayIsAfter(data.arrivedTime)"
      :attime="todayIsAfter(data.arrivedTime)? data.departureTime : data.arrivedTime"
    )
      .number-box
        span {{ data.lab.toUpperCase() }}
    s-card-body(v-if="data" key="body")
      span {{ data.familyNote }}
      br
      span {{ data.footNote }}
      br
      span {{ data.nextVisitNote }}
      .footer
        el-row(type="flex" justify="space-between")
          el-col(:span="6")
            el-row
              el-col(:span="8")
                el-popover(
                  placement="bottom-start"
                  trigger="click"
                  title="gifts"
                  :offset="-17"
                )
                  div gifts gifts ...
                  svg-icon(name="gift" @click.native="gifts" slot="reference")
              el-col(:span="8")
                el-popover(
                  placement="bottom-start"
                  trigger="click"
                  title="places"
                  :offset="-17"
                )
                  div activity activity ...
                  svg-icon(name="places" @click.native="activities" slot="reference")
              el-col(:span="8")
                el-popover(
                  placement="bottom-start"
                  trigger="click"
                  title="hobbies"
                  :offset="-17"
                )
                  div family food etc ...
                  svg-icon(name="hobby" @click.native="notes" slot="reference")
          el-col(:span="6")
            span {{ data.project }}
          el-col(:span="4")
            el-row(justify="end" :gutter="5")
              el-col(:span="20" style="text-align: right")
                div {{data.city}}
              el-col(:span="4")
                svg-icon(:name="data.country")
</template>

<script>
import SCardHead from './SCardHead.vue'
import SCardBody from './SCardBody.vue'

import { getFriendById } from '@/api/friend'

import { todayIsAfter } from '@/utils/convert'

export default {
  components: { SCardHead, SCardBody },
  props: {
    item: {
      type: Object,
      required: true
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
  data() {
    return {
      loading: false,
      data: null
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
  mounted() {
    this.loading = true
    getFriendById(this.item.friendId)
      .then(res => {
        this.data = Object.assign({}, this.item, res.data)
      })
      .finally(() => {
        this.loading = false
      })
  },
  methods: {
    todayIsAfter,
    calculateOrder() {
      const items = this.total
      const counter = Math.ceil(items / 2)

      let ord = Math.ceil(this.step / 2)
      if (this.step % 2 === 0) {
        ord += counter
      }
      return ord
    },
    gifts() {
      // TODO
      console.log('give him gifts, show previous gifts')
    },
    activities() {
      // TODO
      console.log('add/remove activities, show planned activities')
    },
    notes() {
      // TODO
      console.log('not sure :D')
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
.s-card
  &:hover
    box-shadow 0 6px 8px 2px rgba(0, 0, 0, 0.56)
.svg-icon
  cursor pointer
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


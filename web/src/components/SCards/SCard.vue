<template lang="pug">
.s-card(:class="classProps" :style="cssProps")
  slot
</template>

<script>
export default {
  props: {
    step: {
      type: Number,
      required: true
    },
    color: {
      type: String,
      default: '#332211'
    },
    bgColor: {
      type: String,
      default: '#ffffff'
    },
    total: {
      type: Number,
      required: true
    },
    stagger: {
      type: Number,
      default: 180
    },
    cardWidth: {
      type: Number,
      default: 450
    }
  },
  computed: {
    cssProps() {
      const props = {
        backgroundColor: this.color,
        // this is variable for the css
        '--bg-color': this.color,
        '--margin': `${(this.stagger)}px`,
        '--left-marker-dist': '',
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
@media (min-width: 1000px) {
  .mgt2 {
    margin-top: var(--margin) !important;
  }
}
.head::after {
  border-color: var(--bg-color);
}
</style>


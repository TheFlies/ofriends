<template>
  <svg :class="className" xmlns="http://www.w3.org/2000/svg" :style="cssVars">
    <title v-if="title">{{ title }}</title>
    <use :xlink:href="iconPath" xmlns:xlink="http://www.w3.org/1999/xlink" />
  </svg>
</template>

<script>
export default {
  name: 'SvgIcon',
  props: {
    name: {
      type: String,
      required: true
    },
    title: {
      type: String,
      default: null
    },
    width: {
      type: Number,
      required: false,
      default: 1
    },
    height: {
      type: Number,
      required: false,
      default: 1
    }
  },
  computed: {
    iconPath() {
      let data
      try {
        data = require(`@/icons/svg/${this.name}.svg`).default.url
      } catch(err) {
        data = require(`@/icons/svg/example.svg`).default.url
      }
      return data
    },

    className() {
      return `svg-icon svg-icon--${this.name}`
    },

    cssVars() {
      return {
        '--icon-width': `${this.width}em`,
        '--icon-height': `${this.height}em`,
      }
    }
  }
}
</script>

<style>
.svg-icon {
  width: var(--icon-width);
  height: var(--icon-width);
  vertical-align: -0.15em;
  fill: currentColor;
  overflow: hidden;
}
</style>

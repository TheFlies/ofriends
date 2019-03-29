<template lang="pug">
#timeline(v-loading="loading")
  s-cards(:data="items"
    :outer-margin="30"
    :card-height="220"
    :card-width="540"
    marker-color="rgb(0,230,10)"
    :number-size="15"
    marker-border-color="rgba(0,230,10,.32)" :stagger="120" :arrow-size="12")
</template>

<script>
import SCards from '@/components/SCards'

import { getAllVisits } from '@/api/visit'

const colorTable = {
  lab3: '#777777'
}

export default {
  components: { SCards },
  data() {
    return {
      items: [],
      loading: false
    }
  },
  mounted() {
    this.loading = true
    getAllVisits()
      .then(res => {
        this.items = res.data
          .map(i => Object.assign(i, { color: colorTable[i.lab] || '#111122' }))
      })
      .finally(() => {
        this.loading = false
      })
  }
}
</script>

<style scoped lang="stylus">
#timeline
  margin: 50px 0
  background: #f7f7f7
  border-top: 1px solid rgba(191, 191, 191, 0.4)
  border-bottom: 1px solid rgba(191, 191, 191, 0.4)
</style>

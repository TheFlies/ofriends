<template lang="pug">
#timeline
  s-cards(:data="items")
</template>

<script>
import SCards from '@/components/SCards'

export default {
  components: { SCards },
  data() {
    return {
      items: [
        {
          title: 'Mr. Park Heo',
          color: '#3ee9d1',
          subtitle: 'CEO from Samsung',
          content: 'He working with TNTVU and ppu a',
          dc: 'DC14'
        },
        {
          title: 'Kin Khung Ah',
          color: '#4d92eb',
          subtitle: 'VP',
          content: 'Scary boss and dog lover',
          dc: 'DC12A'
        },
        {
          title: 'Tam Tran',
          color: '#ce43eb',
          subtitle: 'Junior Engineer',
          content: 'NA',
          dc: 'DC1'
        },
        {
          title: 'Hai 3b',
          color: '#3ee9d1',
          subtitle: 'Junior Engineer',
          content: 'NA',
          dc: 'DC14'
        },
      ]
    }
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

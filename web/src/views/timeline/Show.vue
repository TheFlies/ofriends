<template>
  <div class="show-timeline">
    <TimeLine />
  </div>
</template>

<script>
// @ is an alias to /src
import TimeLine from '@/components/timeline/TimeLine.vue'

export default {
  name: 'Show',
  components: {
    TimeLine
  }
}
</script>

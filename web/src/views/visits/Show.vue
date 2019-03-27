<template>
  <div class="show-visit">
    <VisitList />
  </div>
</template>

<script>
// @ is an alias to /src
import VisitList from '@/components/visits/VisitList.vue'

export default {
  name: 'Show',
  components: {
    VisitList
  }
}
</script>

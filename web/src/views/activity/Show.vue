<template>
  <div class="show-activity">
    <ActivityList />
  </div>
</template>

<script>
// @ is an alias to /src
import ActivityList from '@/components/activity/ActivityList.vue'

export default {
  name: 'Show',
  components: {
    ActivityList
  }
}
</script>

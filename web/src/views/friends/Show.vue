<template>
  <div class="show-customer">
    <ShowFriends/>
  </div>
</template>

<script>
// @ is an alias to /src
import ShowFriends from '@/components/ShowFriends.vue'

export default {
  name: 'show',
  components: {
    ShowFriends
  }
}
</script>

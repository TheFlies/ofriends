<template>
  <div class="add-customer">
    <AddFriend/>
  </div>
</template>

<script>
// @ is an alias to /src
import AddFriend from '@/components/friends/AddFriend.vue'

export default {
  name: 'add',
  components: {
    AddFriend
  }
}
</script>

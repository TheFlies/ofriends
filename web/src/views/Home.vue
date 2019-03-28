<template>
  <div class="home">
    <img alt="Vue logo" src="../assets/logo.png">
    <AddFriend msg="Welcome d tosss Your Vue.js App" />
  </div>
</template>

<script>
// @ is an alias to /src
import AddFriend from '@/components/friends/FriendAdd.vue'

export default {
  name: 'Home',
  components: {
    AddFriend
  }
}
</script>

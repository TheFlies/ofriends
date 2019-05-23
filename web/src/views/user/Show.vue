<template>
  <div class="show-user">
    <UserList />
  </div>
</template>

<script>
// @ is an alias to /src
import UserList from '@/components/user/UserList.vue'

export default {
  name: 'Show',
  components: {
    UserList
  }
}
</script>

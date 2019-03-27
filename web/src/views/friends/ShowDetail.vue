<template>
  <div class="add-customer">
    <FriendShowDetail />
    <VisitListByFriend />
  </div>
</template>

<script>
// @ is an alias to /src
import FriendShowDetail from '@/components/friends/FriendShowDetail.vue'
import VisitListByFriend from '@/components/visits/VisitListByFriend.vue'

export default {
  name: 'Show',
  components: {
    FriendShowDetail,
    VisitListByFriend
  }
}
</script>

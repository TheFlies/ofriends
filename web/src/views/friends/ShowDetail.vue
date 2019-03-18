<template>
  <div class="show-friend-detail">
    <FriendShowDetail/>
    <VisitList/>
  </div>
</template>

<script>
// @ is an alias to /src
import FriendShowDetail from '@/components/friends/FriendShowDetail.vue'
import VisitList from '@/components/visits/VisitList.vue'

export default {
  name: 'showDetail',
  components: {
    FriendShowDetail,
    VisitList
  }
}
</script>

<template>
  <div class="home">
    <GiftList msg="Welcome to Your Vue.js App"/>
  </div>
</template>

<script>
// @ is an alias to /src
import GiftList from '@/components/gifts/GiftList.vue'

export default {
  name: 'gifts',
  components: {
    GiftList
  }
}
</script>

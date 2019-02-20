<template>
  <div>
    <v-toolbar color="#4B71FB" dense dark>
      <v-toolbar-title class="white--text" align-content-center align-center>OFriends</v-toolbar-title>
      <v-btn flat to="/" class="text-capitalize font-weight-bold">
        Home
      </v-btn>
      <v-btn flat to="/upcomming" class="text-capitalize font-weight-bold">
        Upcomming  
      </v-btn>
      <v-spacer></v-spacer>
      <v-btn outline icon class="subheading">
        <v-icon>search</v-icon>
      </v-btn>
      <v-btn outline icon to="/friend/new" class="subheading">
        <v-icon>add</v-icon>
      </v-btn>
      <v-btn outline v-if="!auth" @click="onLogin" >Login</v-btn>
      <v-menu offset-y v-if="auth">
        <v-avatar
          size="36px"
          slot="activator"
          >
          <v-img
            :src="user.image"
          ></v-img>
        </v-avatar>
        <v-list>
          <v-list-tile>
            <v-list-tile-title :style="{ cursor: 'pointer'}">Profile</v-list-tile-title>
          </v-list-tile>
          <v-divider></v-divider>
          <v-list-tile>
            <v-list-tile-title :style="{ cursor: 'pointer'}" @click="onLogout">Logout</v-list-tile-title>
          </v-list-tile>
        </v-list>
      </v-menu> 
    </v-toolbar>
  </div>
</template>

<script>
  export default {
  }
</script>

<style scoped>

</style>
<template>
<v-content>
  <toolbar></toolbar>
  <v-container>
    <v-layout row wrap align-center class="mx-5">
      <v-flex xs12 sm6 offset-sm3> 
       <p class="text-xs-center">
        <img src="../static/talkaneast.png" />
       </p>
       <v-card class="text-xs-center pa-3 mx-5 mt-0 d-block">
        <v-text-field
        v-model="username"
        label="Username"
        :rules="userRules"
        >
        </v-text-field>
        <v-text-field
        v-model="password"
        type="password"
        label="Password"
        :rules="passwordRules"
        >
        </v-text-field>
        <v-btn>
          LOG IN
        </v-btn>
       </v-card>
      </v-flex>
    </v-layout>
  </v-container>
</v-content>
</template>


<script>
import Toolbar from '@/components/Toolbar'
export default {
  components: { Toolbar },
  data() {
    return {
      error: null,
      valid: false,
      email: '',
      username: '',
      userRules: [v => !!v || 'Username is required'],
      emailRules: [v => !!v || 'Email is required'],
      password: '',
      passwordRules: [v => !!v || 'Password is required']
    }
  }
}
</script>

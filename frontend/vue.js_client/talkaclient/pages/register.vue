<template>
<v-content>
    <toolbar></toolbar>
    <v-container>
    <v-layout row wrap align-center class="mx-5 mt-0">
      <v-flex xs12 sm6 offset-sm3>
          <p class="text-xs-center">
        <img src="../static/talkaneast.png" />
       </p>
       <v-card class="text-xs-center pa-3 mx-5 d-block">
        <v-text-field
        v-model="username" 
        label="Username"
        >
        </v-text-field>
        <v-text-field
        v-model="email" 
        type="email" 
        :rules="emailRules"
        label="E-mail"
        >
        </v-text-field>
        <v-text-field
        v-model="password" 
        type="password" 
        :rules="passwordRules"
        label="Password"
        >
        </v-text-field>
        <v-text-field
        type="password" 
        :rules="passwordRules"
        label="Repeat password"
        >
        </v-text-field>
        <v-btn
        @click="login">
          SIGN IN
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
      emailRules: [v => !!v || 'Email is required'],
      password: '',
      passwordRules: [v => !!v || 'Password is required']
    }
  },
  methods: {
    async login() {
      this.$options.sockets.onmessage = (event) => {
        const json = JSON.parse(event.data);
        if (json["event"] == "UserRegistered") {
          delete this.$options.sockets.onmessage
          this.$router.push({
            path: '/login'
          })
        }
      }

      this.$socket.sendObj({
        "event": "RegisterUser",
        "data": {
          "username": this.username,
          "password": this.password,
          "email": this.email,
        }}
        )
    }
  }
}
</script>
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
        label="Username"
        >
        </v-text-field>
        <v-text-field
        label="Password"
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
  components: { Toolbar }
}
</script>

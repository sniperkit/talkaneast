<template>
<v-content>
    <toolbar></toolbar>
    <v-container>
    <v-layout row wrap align-center>
      <v-flex>
       <v-card class="text-xs-center pa-3">
        <v-text-field
        label="Username"
        >
        </v-text-field>
        <v-text-field
        label="E-mail"
        >
        </v-text-field>
        <v-text-field
        label="Password"
        >
        </v-text-field>
        <v-text-field
        label="Repeat password"
        >
        </v-text-field>
        <v-btn>
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
  components: { Toolbar }
}
</script>

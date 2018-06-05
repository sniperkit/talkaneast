<template>
  <v-container fluid>
    <v-layout row wrap justify-center>
      <v-flex xs12 sm6>
          <v-container fluid>
            <v-layout row wrap>
              <v-flex xs12>
                <v-text-field
                  v-model="login"
                  label="Username"
                  counter
                  max="15"
                ></v-text-field>
                </v-flex>
                <v-flex xs12>
                <v-text-field
                  v-model="password"
                  label="Password"
                ></v-text-field>
                <a
                class="button--grey">LOG IN</a>
              </v-flex>
            </v-layout>
          </v-container>
      </v-flex>
    </v-layout>
  </v-container>
</template>
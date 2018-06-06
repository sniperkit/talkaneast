<template>
<div>
  <v-toolbar>
    <v-toolbar-title>Talka Neast</v-toolbar-title>
    <v-spacer></v-spacer>
    <v-toolbar-items class="hidden-sm-and-down">
      <v-btn flat to="/" class='black--text'>HOME</v-btn>
      <v-btn flat to="/register" class='black--text'>SIGN IN</v-btn>
      <v-btn flat to="/login" class='black--text'>LOG IN</v-btn>
      <v-btn flat>Dark Theme Coming Soon</v-btn>
    </v-toolbar-items>
  </v-toolbar>
  </div>
</template>

<script>
export default {
  name: 'Toolbar'
}
</script>

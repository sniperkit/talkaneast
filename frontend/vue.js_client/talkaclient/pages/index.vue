<template>
  <section class="container">
    <div>
      <h2 class="subtitle">
        Talka Neast
      </h2>
      <div class="links">
        <nuxt-link to='/register' class="button--grey">SIGN IN</nuxt-link>
         <nuxt-link to='/login' class="button--grey">LOG IN</nuxt-link>
      </div>
    </div>
  </section>
</template>
<script>
export default {
  name: 'index',
  data: _ => ({
    
  })
}
</script>

<style>
.container {
  min-height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  text-align: center;
}

.title {
  font-family: "Quicksand", "Source Sans Pro", -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, sans-serif; /* 1 */
  display: block;
  font-weight: 300;
  font-size: 100px;
  color: #35495e;
  letter-spacing: 1px;
}

.subtitle {
  font-weight: 300;
  font-size: 42px;
  color: #526488;
  word-spacing: 5px;
  padding-bottom: 15px;
}

.links {
  padding-top: 15px;
}
</style>

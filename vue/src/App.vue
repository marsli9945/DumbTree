<template>
  <div id="app">
    <img alt="Vue logo" src="./assets/logo.png">
    <User/>
    <p>
      <router-link to="/vue/helloworld/222">Go to HelloWorld</router-link>
    </p>
    <p>
      <router-link to="/vue/user/222">Go to User</router-link>
    </p>
    <router-view></router-view>
  </div>
</template>

<script>
export default {
  name: 'app',
  mounted() {
    // this.$router.push('helloworld/123')
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>

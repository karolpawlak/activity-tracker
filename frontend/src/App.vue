<template>
  <web-header/>
  <router-view/>
  <web-footer/>
</template>

<script>
import WebHeader from '@/components/WebHeader.vue'
import WebFooter from '@/components/WebFooter.vue'

export default {
  components: {
    WebHeader,
    WebFooter
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

nav {
  padding: 30px;
}

nav a {
  font-weight: bold;
  color: #2c3e50;
}

nav a.router-link-exact-active {
  color: #42b983;
}
</style>

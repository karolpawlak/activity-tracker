<template>
  <div class="home">
    <img id="tanzuLogo" alt="Tanzu logo" src="../assets/logo.png">
    <HelloWorld msg="Tanzu Let's Get Moving"/>
  </div>
</template>

<script>
// @ is an alias to /src
import HelloWorld from '@/components/HelloWorld.vue'

export default {
  name: 'HomeView',
  components: {
    HelloWorld
  }
}
</script>

<style scoped>
#tanzuLogo {
  width: 5%;
  height: auto;
}
</style>

<template>
  <div class="home">
      <Header name="数据面板"></Header>
    <img alt="Vue logo" src="../assets/logo.png">
    <HelloWorld msg="Welcome to Your Vue.js App"/>
  </div>
</template>

<script>
// @ is an alias to /src
import HelloWorld from '@/components/HelloWorld.vue'
import Header from "../components/Header";

export default {
  name: 'Home',
  components: {
      Header,
    HelloWorld
  },
    data() {
      return {
      }
    }
}
</script>

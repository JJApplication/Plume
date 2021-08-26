<template>
  <div id="app">
      <component :is="Comps"></component>
      <Bottom></Bottom>
  </div>
</template>

<script>
import Bottom from "./components/Bottom";
import Home from "./views/Home";
import About from "./views/About";
import Container from "./views/Container";
import Setting from "./views/Setting";
export default {
  name: 'app',
  components: {
      Bottom, Home, About, Container, Setting
  },
  data(){
      return {
          Comps: this.$store.state.comps
      }
  },
  watch: {
      "$store.state.comps": function () {
          console.log(this.$store.state.comps);
          this.Comps = this.$store.state.comps;
      }
  }
}
</script>

<style>
#app {
  height: calc(100% - 40px);
  background-color: white;
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  padding: 20px 40px;
  max-width: 720px;
  margin: 0 auto;
}
* {
    padding: 0;
    margin: 0;
}
html, body {
    height: 100%;
    background-color: #f5f5f5;
}
</style>

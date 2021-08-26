<template>
  <div class="about">
      <Header :name="app"></Header>
        <h3>about {{app}}</h3>
  </div>
</template>

<script>

import Header from "../components/Header";
export default {
    name: 'About',
    components: {Header},
    data(){
        return {
            app: this.$store.state.app
        }
    }
}
</script>

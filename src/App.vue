<template>
  <div id="app">
      <!--      懒加载-->
      <van-skeleton title avatar :row="8" :loading="loading">
          <component :is="Comps"></component>
      </van-skeleton>
      <Bottom></Bottom>
  </div>
</template>

<script>
import Bottom from "./components/Bottom";
import Home from "./views/Home";
import About from "./views/About";
import Container from "./views/Container";
import Setting from "./views/Setting";
import Detail from "./views/Detail";
export default {
    name: 'app',
        components: {
          Bottom, Home, About, Container, Setting, Detail
        },
    data(){
        return {
            Comps: this.$store.state.comps,
            loading: true
        }
    },
    watch: {
        "$store.state.comps": function () {
            this.Comps = this.$store.state.comps;
            // 持久化标签路径
            this.$store.commit('changeComps', this.Comps);
        },
    },
    mounted() {
        setTimeout(() => {
            this.loading = false;
        }, 1000)
        // 挂载时即监听主题变化
        this.$store.commit('changeTheme', this.$store.state.theme);
    }
}
</script>

<style>
#app {
  height: calc(100% - 40px);
  background-color: var(--light-bg-container-color);
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  padding: 20px 40px;
  max-width: 680px;
  margin: 0 auto;
  overflow: hidden;
  line-height: 1.4rem;
  user-select: none;
  -ms-user-select: none;
  -moz-user-select: none;
}

@media (max-width: 480px) {
    #app {
        padding: 20px;
    }
}
* {
    padding: 0;
    margin: 0;
}
html, body {
    height: 100%;
    background-color: var(--light-bg-color, #101010);
    color: var(--light-text-color, #ffffff);
}
/*    隐藏滚动条*/
::-webkit-scrollbar-thumb {
    background-color: transparent;
    color: transparent;
    width: 0;
    height: 0;
}
::-webkit-scrollbar {
    background-color: transparent;
    color: transparent;
    width: 0;
    height: 0;
}
</style>

<style scoped>
    #app /deep/ .van-skeleton__row, #app /deep/ .van-skeleton__title, #app /deep/ .van-skeleton__avatar{
        background-color: var(--light-loading-color, #252525);
    }
</style>

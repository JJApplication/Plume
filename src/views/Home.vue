<template>
  <div class="home">
      <Header name="数据面板"></Header>
      <!--  服务器uname信息    -->
      <div id="data-server-title">
          <p>Ubuntu 20.04LTS</p>
      </div>
<!--    <img alt="Vue logo" src="../assets/logo.png">-->
<!--    <HelloWorld msg="Welcome to Your Vue.js App"/>-->
      <div class="cell">
          <p class="cell-title">负载</p>
      </div>

      <div class="cell">
          <p class="cell-title">内存</p>
      </div>

      <div class="cell">
          <p class="cell-title">网络IO</p>
      </div>

      <div class="cell">
          <p class="cell-title">磁盘IO</p>
      </div>
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

<style scoped>
    #data-server-title {
        text-align: left;
        margin-top: 4px;
        margin-bottom: 16px;
    }
    #data-server-title p {
        font-size: .85rem;
        color: #404040;
    }
    .cell {
        text-align: left;
        background-color: white;
        padding: 10px;
        border-radius: 10px;
        margin-bottom: 20px;
    }
    .cell .cell-title {
        font-size: .9rem;
        font-weight: bold;
    }
</style>
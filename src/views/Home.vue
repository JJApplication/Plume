<template>
  <div class="home">
      <Header name="数据面板"></Header>
      <!--  服务器uname信息    -->
      <div id="data-server-title">
          <p>Ubuntu 20.04LTS</p>
      </div>
      <div class="cell">
          <p class="cell-title">负载</p>
          <van-row gutter="20" justify="center" style="margin-top: .5rem">
              <van-col span="6" class="cell-head cell-head-cpu">{{ cpu_usage }} %</van-col>
              <van-col span="6">
                  <div>
                      <span class="cell-head">系统</span><br>
                      <span>0%</span>
                  </div>
              </van-col>
              <van-col span="6">
                  <div>
                      <span class="cell-head">用户</span><br>
                      <span>0%</span>
                  </div>
              </van-col>
              <van-col span="6">
                  <div>
                      <span class="cell-head">IO等待</span><br>
                      <span>0%</span>
                  </div>
              </van-col>
          </van-row>
          <van-progress :percentage="50" stroke-width="8" style="margin-top: .8rem;margin-bottom: .8rem" color="var(--light-text-active-color, #ff4a9e)"/>
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
import Header from "../components/Header";

export default {
  name: 'Home',
  components: {
      Header,
  },
    data() {
      return {
          cpu_usage: 0
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
        color: var(--light-text-title-color);
    }
    .cell {
        text-align: left;
        background-color: var(--light-bg-color, #2f2f2f);
        padding: 10px;
        border-radius: 10px;
        margin-bottom: 20px;
    }
    .cell .cell-title {
        font-size: 1.2rem;
        font-weight: bold;
        color: var(--light-text-color, #ffffff);
    }
    .cell .cell-head {
        color: var(--light-text-bold-color);
        font-weight: bold;
        font-size: .9rem;
    }
    .cell-head.cell-head-cpu {
        font-size: 1.6rem;
    }
</style>
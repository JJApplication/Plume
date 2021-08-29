<template>
    <div class="scroll-body">
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
                            <span class="cell-head-cpu-pre cell-head-cpu-system"></span><span class="cell-head">系统</span><br>
                            <span class="cell-data">0</span><span style="margin-left: 2px;font-size: .9rem">%</span>
                        </div>
                    </van-col>
                    <van-col span="6">
                        <div>
                            <span class="cell-head-cpu-pre cell-head-cpu-user"></span><span class="cell-head">用户</span><br>
                            <span class="cell-data">0</span><span style="margin-left: 2px;font-size: .9rem">%</span>
                        </div>
                    </van-col>
                    <van-col span="6">
                        <div>
                            <span class="cell-head-cpu-pre cell-head-cpu-wait"></span><span class="cell-head">IO等待</span><br>
                            <span class="cell-data">0</span><span style="margin-left: 2px;font-size: .9rem">%</span>
                        </div>
                    </van-col>
                </van-row>
                <van-progress :percentage="cpu_usage" stroke-width="8" style="margin-top: .8rem;margin-bottom: .8rem" color="var(--light-text-active-color, #ff4a9e)"/>
                <br/>
                <van-row gutter="20" justify="center" style="margin-top: .2rem">
                    <van-col span="6" class="cell-head">
                        <div>
                            <span class="cell-head">核心数</span><br>
                            <span class="cell-data">2</span>
                        </div>
                    </van-col>
                    <van-col span="6">
                        <div>
                            <span class="cell-head">空闲</span><br>
                            <span class="cell-data">0</span><span style="margin-left: 2px;font-size: .9rem">%</span>
                        </div>
                    </van-col>
                    <van-col span="6">
                        <div>
                            <span class="cell-head">运行时间</span><br>
                            <span class="cell-data">0</span><span style="margin-left: 2px;font-size: .9rem">Day</span>
                        </div>
                    </van-col>
                    <van-col span="6">
                        <div>
                            <span class="cell-head">平均负载</span><br>
                            <span class="cell-data">0</span>
                        </div>
                    </van-col>
                </van-row>
            </div>

            <div class="cell">
                <p class="cell-title">内存</p>
                <van-row gutter="20" justify="center" style="margin-top: .5rem">
                    <van-col span="6" class="cell-head cell-head-mem">{{ mem_usage }} %</van-col>
                    <van-col span="6">
                        <div>
                            <span class="cell-head">可用</span><br>
                            <span class="cell-data">0</span><span style="margin-left: 2px;font-size: .9rem">G</span>
                        </div>
                    </van-col>
                    <van-col span="6">
                        <div>
                            <span class="cell-head">已用</span><br>
                            <span class="cell-data">0</span><span style="margin-left: 2px;font-size: .9rem">G</span>
                        </div>
                    </van-col>
                    <van-col span="6">
                        <div>
                            <span class="cell-head">页面缓存</span><br>
                            <span class="cell-data">0</span><span style="margin-left: 2px;font-size: .9rem">G</span>
                        </div>
                    </van-col>
                </van-row>
            </div>

            <div class="cell">
                <p class="cell-title">网络IO</p>
                <van-row gutter="10" justify="center" style="margin-top: .5rem">
                    <van-col span="6">
                        <div>
                            <span class="cell-head">上传</span><br>
                            <span class="cell-data">{{net_upload}} /s</span>
                        </div>
                    </van-col>
                    <van-col span="6">
                        <div>
                            <span class="cell-head">下载</span><br>
                            <span class="cell-data">{{net_download}} /s</span>
                        </div>
                    </van-col>
                    <van-col span="8">
                        <div class="cell-data">
                            <font-awesome-icon icon="arrow-up" style="color: #a0a0a0" />&nbsp;<span>{{network_upload}}</span><span>G</span><span style="background-color: #ff7f50;height: 12px;width: 6px;display: inline-block;margin-left: 4px;border-radius: 4px"></span><br>
                            <font-awesome-icon icon="arrow-down" style="color: #a0a0a0" />&nbsp;<span>{{network_download}}</span><span>G</span><span style="background-color: #7fff00;height: 12px;width: 6px;display: inline-block;margin-left: 4px;border-radius: 4px"></span>
                        </div>
                    </van-col>
                    <van-col span="4">
                        <div>
                            <van-circle
                                    v-model:current-rate="network_init"
                                    :rate="network_percent" :speed="100"
                                    :stroke-width="160"
                                    size="50px"
                                    color="var(--light-circle-color, #7fff00)"
                                    layer-color="var(--light-circle-bg-color, #ff7f50)"/>
                        </div>
                    </van-col>
                </van-row>
                <van-divider />
                <van-row gutter="20" justify="center" style="margin-top: .5rem">
                    <van-col span="6">
                        <div>
                            <span class="cell-head">重传率</span><br>
                            <span class="cell-data">0</span><span style="margin-left: 2px;font-size: .9rem">%</span>
                        </div>
                    </van-col>
                    <van-col span="6">
                        <div>
                            <span class="cell-head">主动建连</span><br>
                            <span class="cell-data">0</span>
                        </div>
                    </van-col>
                    <van-col span="6">
                        <div>
                            <span class="cell-head">被动建连</span><br>
                            <span class="cell-data">0</span>
                        </div>
                    </van-col>
                    <van-col span="6">
                        <div>
                            <span class="cell-head">建连失败</span><br>
                            <span class="cell-data">0</span>
                        </div>
                    </van-col>
                </van-row>
                <van-divider />
                <div style="font-size: .85rem"><font-awesome-icon icon="wifi" style="color: #23fb1a" />&emsp;NETWORK ⇌ {{ip}}</div>
            </div>

            <div class="cell">
                <p class="cell-title">磁盘IO</p>
                <van-row gutter="20" justify="center" style="margin-top: .5rem">
                    <van-col span="12">
                        <div>
                            <span class="cell-head" style="font-weight: bold">/</span><br>
                            <span class="cell-data">/dev/vda1</span>
                        </div>
                    </van-col>
                    <van-col span="12">
                        <div>
                            <span class="cell-head">20G</span>
                            <span>/30G</span>
                            <div class="data-bar"><span></span></div>
                        </div>
                    </van-col>
                </van-row>
                <van-divider />
                <van-row gutter="20" justify="center" style="margin-top: .5rem">
                    <van-col span="6">
                        <div>
                            <span class="cell-head cell-io"></span><br>
                            <span class="cell-data">读</span><br>
                            <span class="cell-data">写</span>
                        </div>
                    </van-col>
                    <van-col span="6">
                        <div>
                            <span class="cell-head">速率</span><br>
                            <span class="cell-data">0 MB/s</span><br>
                            <span class="cell-data">0 MB/s</span>
                        </div>
                    </van-col>
                    <van-col span="6">
                        <div>
                            <span class="cell-head">字节</span><br>
                            <span class="cell-data">0 B</span><br>
                            <span class="cell-data">0 B</span>
                        </div>
                    </van-col>
                    <van-col span="6">
                        <div>
                            <span class="cell-head">延迟</span><br>
                            <span class="cell-data">0</span><br>
                            <span class="cell-data">0</span>
                        </div>
                    </van-col>
                </van-row>
            </div>
        </div>
    </div>
</template>

<script>
// @ is an alias to /src
import Header from "../components/Header";
import "@/assets/animate_io.css";

export default {
  name: 'Home',
  components: {
      Header,
  },
    data() {
      return {
          cpu_usage: 0,
          mem_usage: 0,
          network_init: 0,
          net_upload: '100m',
          net_download: '100m',
          network_upload: 100,
          network_download: 100,
          network_percent: 0,
          ip: '127.0.0.1'
      }
    },
    mounted() {
      this.calcNetPercent();
    },
    methods: {
      calcNetPercent() {
          this.network_percent =  (this.network_download / (this.network_download + this.network_upload) * 100).toFixed(0);
      }
    }
}
</script>

<style scoped>
    .home {
        overflow-y: auto;
        padding-bottom: 40px;
        padding-top: 8px;
    }
    .scroll-body {
        overflow: auto;
        height: 100%
    }
    .scroll-body::-webkit-scrollbar {
        display: none;
        width: 0;
        height: 0;
    }
    .home::-webkit-scrollbar {
        display: none;
        width: 0;
        height: 0;
    }
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
        padding: 10px 20px;
        border-radius: 10px;
        margin-bottom: 20px;
    }
    .cell .cell-title {
        font-size: 1.1rem;
        font-weight: bold;
        color: var(--light-text-color, #ffffff);
    }
    .cell .cell-head {
        color: var(--light-text-bold-color, #7d7d7d);
        font-weight: bold;
        font-size: .9rem;
    }
    .cell .cell-data {
        color: var(--light-text-color, #fff);
        font-size: .9rem;
        font-weight: bold;
    }
    .cell .cell-head-cpu-pre {
        height: 10px;
        width: 4px;
        display: inline-block;
        border-radius: 4px;
        margin-right: 4px
    }
    .cell .cell-head-cpu-system {
        background-color: #ff0000;
    }
    .cell .cell-head-cpu-user {
        background-color: #7fff00;
    }
    .cell .cell-head-cpu-wait {
        background-color: #9370db;
    }
    .cell-head.cell-head-cpu {
        font-size: 1.6rem;
    }
    .cell-head.cell-head-mem {
        font-size: 1.6rem;
    }
    .data-bar {
        height: 40px;
        width: 24px;
        display: inline-block;
        background-color: var(--light-bg-container-color, #1f1f1f);
        position: relative;
        border-radius: 6px;
        margin-left: 8px;
    }
    .data-bar span {
        background-color: #1989fa;
        height: 10%;
        width: 24px;
        display: inline-block;
        position: absolute;
        bottom: 0;
        border-bottom-left-radius: 6px;
        border-bottom-right-radius: 6px;
    }
    .cell-head.cell-io {
        height: 10px;
        width: 10px;
        display: inline-block;
        background-color: #23fb1a;
        border-radius: 50%;
        animation-name: io;
        animation-iteration-count: infinite;
        animation-timing-function: ease;
        animation-duration: 2s;
    }
</style>
<template>
    <div class="scroll-body">
        <div class="home">
            <Header name="数据面板"></Header>
            <!--  服务器uname信息    -->
            <div id="data-server-title">
                <p>{{server}}</p>
            </div>
            <div class="cell">
                <p class="cell-title">负载</p>
                <van-row gutter="20" justify="center" style="margin-top: .5rem">
                    <van-col span="6" class="cell-head cell-head-cpu">{{ cpu_usage }} %</van-col>
                    <van-col span="6">
                        <div>
                            <span class="cell-head-cpu-pre cell-head-cpu-system"></span><span class="cell-head">系统</span><br>
                            <span class="cell-data">{{cpu_usage_system}}</span><span style="margin-left: 2px;font-size: .9rem">%</span>
                        </div>
                    </van-col>
                    <van-col span="6">
                        <div>
                            <span class="cell-head-cpu-pre cell-head-cpu-user"></span><span class="cell-head">用户</span><br>
                            <span class="cell-data">{{cpu_usage_user}}</span><span style="margin-left: 2px;font-size: .9rem">%</span>
                        </div>
                    </van-col>
                    <van-col span="6">
                        <div>
                            <span class="cell-head-cpu-pre cell-head-cpu-wait"></span><span class="cell-head">IO等待</span><br>
                            <span class="cell-data">{{cpu_io_wait}}</span><span style="margin-left: 2px;font-size: .9rem">%</span>
                        </div>
                    </van-col>
                </van-row>
                <van-progress :percentage="cpu_usage" stroke-width="8" style="margin-top: .8rem;margin-bottom: .8rem" color="var(--light-text-active-color, #ff4a9e)"/>
                <br/>
                <van-row gutter="20" justify="center" style="margin-top: .2rem">
                    <van-col span="6" class="cell-head">
                        <div>
                            <span class="cell-head">核心数</span><br>
                            <span class="cell-data">{{cpu_count}}</span>
                        </div>
                    </van-col>
                    <van-col span="6">
                        <div>
                            <span class="cell-head">空闲</span><br>
                            <span class="cell-data">{{cpu_free}}</span><span style="margin-left: 2px;font-size: .9rem">%</span>
                        </div>
                    </van-col>
                    <van-col span="6">
                        <div>
                            <span class="cell-head">运行时间</span><br>
                            <span class="cell-data">{{cpu_run}}</span><span style="margin-left: 2px;font-size: .9rem">Day</span>
                        </div>
                    </van-col>
                    <van-col span="6">
                        <div>
                            <span class="cell-head">平均负载</span><br>
                            <span class="cell-data">{{cpu_load}}</span>
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
                            <span class="cell-data">{{mem_free}}</span>
                        </div>
                    </van-col>
                    <van-col span="6">
                        <div>
                            <span class="cell-head">已用</span><br>
                            <span class="cell-data">{{mem_used}}</span>
                        </div>
                    </van-col>
                    <van-col span="6">
                        <div>
                            <span class="cell-head">页面缓存</span><br>
                            <span class="cell-data">{{mem_cache}}</span>
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
                            <font-awesome-icon icon="arrow-up" style="color: #a0a0a0" />&nbsp;<span>{{network_upload}}</span><span style="background-color: #ff7f50;height: 12px;width: 6px;display: inline-block;margin-left: 4px;border-radius: 4px"></span><br>
                            <font-awesome-icon icon="arrow-down" style="color: #a0a0a0" />&nbsp;<span>{{network_download}}</span><span style="background-color: #7fff00;height: 12px;width: 6px;display: inline-block;margin-left: 4px;border-radius: 4px"></span>
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
                            <span class="cell-data">{{net_retry}}</span><span style="margin-left: 2px;font-size: .9rem">%</span>
                        </div>
                    </van-col>
                    <van-col span="6">
                        <div>
                            <span class="cell-head">主动建连</span><br>
                            <span class="cell-data">{{net_active}}</span>
                        </div>
                    </van-col>
                    <van-col span="6">
                        <div>
                            <span class="cell-head">被动建连</span><br>
                            <span class="cell-data">{{net_passive}}</span>
                        </div>
                    </van-col>
                    <van-col span="6">
                        <div>
                            <span class="cell-head">建连失败</span><br>
                            <span class="cell-data">{{net_fail}}</span>
                        </div>
                    </van-col>
                </van-row>
                <van-divider />
                <div style="font-size: .85rem"><font-awesome-icon icon="wifi" style="color: #23fb1a" />&emsp;<span class="cell-data">ipv4</span> ⇌ {{ipv4}}</div>
                <div style="font-size: .85rem"><font-awesome-icon icon="wifi" style="color: #23fb1a" />&emsp;<span class="cell-data">ipv6</span> ⇌ {{ipv6}}</div>
            </div>

            <div class="cell">
                <p class="cell-title">磁盘IO</p>
                <van-row gutter="20" justify="center" style="margin-top: .5rem">
                    <van-col span="12">
                        <div>
                            <span class="cell-head" style="font-weight: bold">/</span><br>
                            <span class="cell-data">{{ disk_mount }}</span>
                        </div>
                    </van-col>
                    <van-col span="12">
                        <div>
                            <span class="cell-head">{{disk_used}}</span>
                            <span class="cell-head">/{{disk_all}}</span>
                            <div class="data-bar" id="data-bar"><span id="data-bar-inner"></span></div>
                        </div>
                    </van-col>
                </van-row>
                <van-divider />
                <van-row gutter="20" justify="center" style="margin-top: .5rem">
                    <van-col span="4">
                        <div>
                            <span class="cell-head cell-io"></span><br>
                            <span class="cell-data">读</span><br>
                            <span class="cell-data">写</span>
                        </div>
                    </van-col>
                    <van-col span="8">
                        <div>
                            <span class="cell-head">速率</span><br>
                            <span class="cell-data">{{disk_read_rate}}</span><br>
                            <span class="cell-data">{{disk_write_rate}}</span>
                        </div>
                    </van-col>
                    <van-col span="6">
                        <div>
                            <span class="cell-head">字节</span><br>
                            <span class="cell-data">{{disk_read_byte}}</span><br>
                            <span class="cell-data">{{disk_write_byte}}</span>
                        </div>
                    </van-col>
                    <van-col span="6">
                        <div>
                            <span class="cell-head">延迟</span><br>
                            <span class="cell-data">{{disk_read_delay}}ms</span><br>
                            <span class="cell-data">{{disk_write_delay}}ms</span>
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
import apis from "../actions/api";

export default {
  name: 'Home',
  components: {
      Header,
  },
    data() {
      return {
          global_timer: null,

          server: 'Ubuntu 20.04LTS',
          cpu_usage: 0,
          cpu_usage_system: 0,
          cpu_usage_user: 0,
          cpu_io_wait: 0,
          cpu_count: 2,
          cpu_free: 0,
          cpu_run: 0,
          cpu_load: '0,0,0',

          mem_usage: 0,
          mem_free: 0,
          mem_used: 0,
          mem_cache: 0,

          network_init: 0,
          net_upload: '100mb',
          net_download: '100mb',
          network_upload: '100G',
          network_download: '100G',
          network_percent: 0,
          net_retry: 0,
          net_active: 0,
          net_passive: 0,
          net_fail: 0,
          ipv4: '127.0.0.1',
          ipv6: '::1',

          disk_mount: '/dev/vda1',
          disk_used: '0G',
          disk_all: '0G',
          disk_usage: 50,
          disk_read_rate: '0 KB/s',
          disk_write_rate: '0 KB/s',
          disk_read_byte: 0,
          disk_write_byte: 0,
          disk_read_delay: 0,
          disk_write_delay: 0
      }
    },
    watch: {
        cpu_usage: function () {
            if (this.cpu_usage >= this.$store.state.limit) {
                this.notifyDanger("cpu使用量超出限制" + this.$store.state.limit + "%")
            }
        },
        mem_usage: function () {
            if (this.mem_usage >= this.$store.state.limit) {
                this.notifyDanger("内存使用量超出限制" + this.$store.state.limit + "%")
            }
        }
    },
    mounted() {
      this.global_timer = setInterval(() => {
          if (this.$store.state.comps === 'Home' && this.$store.state.watchdog === 'false' || this.$store.state.watchdog === false) {
              this.getServerData();
              this.getCPUdata();
              this.getMemData();
              this.getNetData();
              this.getDiskData();
          }else {
              this.calcNetPercent();
              this.calcBar();
          }
      }, parseInt(this.$store.state.duration, 10) * 1000)
    },
    beforeDestroy() {
      if (this.global_timer) {
          clearInterval(this.global_timer);
      }
    },
    methods: {
      calcNetPercent() {
          this.network_percent = (this.network_download / (this.network_download + this.network_upload) * 100).toFixed(0);
          return (this.network_download / (this.network_download + this.network_upload) * 100).toFixed(0);
      },
      calcBar() {
          let bar = document.getElementById("data-bar-inner");
          if (bar) {
              let height = this.disk_usage;
              if (height >= 100) {
                  bar.style.borderRadius = '6px';
              }
              bar.style.height = height + '%';
          }
      },
      // 告警提示
      notifyDanger(message) {
          this.$notify({ type: 'danger', message: message });
      },
      // 服务器信息
      getServerData() {
          this.$axios.post(apis.api_server)
          .then(res => {
              this.server = res.data.data ? res.data.data : 'Unrecognizable server type';
          }).catch(()=>{
              this.notifyDanger("接口" + apis.api_server + "请求失败");
              this.server = 'Unrecognizable server type';
          })
      },
      // 获取cpu信息
      getCPUdata() {
        this.$axios.post(apis.api_cpu)
        .then(res => {
            let data = res.data.data;
            this.cpu_usage = data.cpu_usage;
            this.cpu_usage_system = data.cpu_usage_system;
            this.cpu_usage_user = data.cpu_usage_user;
            this.cpu_io_wait = data.cpu_io_wait;
            this.cpu_count = data.cpu_count;
            this.cpu_free = data.cpu_free;
            this.cpu_load = data.cpu_load;
            this.cpu_run = data.cpu_run;
        }).catch(() => {
            this.notifyDanger("接口" + apis.api_cpu + "请求失败");
        });
      },
      // 获取内存信息
      getMemData() {
          this.$axios.post(apis.api_mem)
              .then(res => {
                  let data = res.data.data;
                  this.mem_usage = data.mem_usage;
                  this.mem_free = data.mem_free;
                  this.mem_used = data.mem_used;
                  this.mem_cache = data.mem_cache;
              }).catch(() => {
              this.notifyDanger("接口" + apis.api_mem + "请求失败");
          });
      },
      getNetData() {
          this.$axios.post(apis.api_net)
              .then(res => {
                  let data = res.data.data;
                  this.net_upload = data.net_upload;
                  this.net_download = data.net_download;
                  this.network_upload = data.network_upload;
                  this.network_download = data.network_download;
                  this.network_percent = this.calcNetPercent();
                  this.net_retry = data.net_retry;
                  this.net_active = data.net_active;
                  this.net_passive = data.net_passive;
                  this.net_fail = data.net_fail;
                  this.ipv4 = data.ipv4;
                  this.ipv6 = data.ipv6;
              }).catch(() => {
              this.notifyDanger("接口" + apis.api_net + "请求失败");
          });
      },
      getDiskData() {
          this.$axios.post(apis.api_disk)
              .then(res => {
                  let data = res.data.data;
                  this.disk_mount = data.disk_mount;
                  this.disk_used = data.disk_used;
                  this.disk_all = data.disk_all;
                  this.disk_usage = data.disk_usage;
                  this.disk_read_rate = data.disk_read_rate;
                  this.disk_write_rate = data.disk_write_rate;
                  this.disk_read_byte = data.disk_read_byte;
                  this.disk_write_byte = data.disk_write_byte;
                  this.disk_read_delay = data.disk_read_delay;
                  this.disk_write_delay = data.disk_write_delay;

                  this.calcBar();
              }).catch(() => {
              this.notifyDanger("接口" + apis.api_disk + "请求失败");
          });
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
        background-color: var(--light-bg-color, #202020);
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
    .data-bar #data-bar-inner {
        background-color: #1989fa;
        height: 0;
        width: 24px;
        display: inline-block;
        position: absolute;
        bottom: 0;
        border-bottom-left-radius: 6px;
        border-bottom-right-radius: 6px;
        transition: height .3s ease;
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
    .cell /deep/ .van-divider {
        border-color: var(--light-line-border-color, #707070);
    }
</style>
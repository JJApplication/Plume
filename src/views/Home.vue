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
                            <div class="load-bar"><span id="m1"></span></div>
                            <div class="load-bar"><span id="m5"></span></div>
                            <div class="load-bar"><span id="m15"></span></div>
                        </div>
                    </van-col>
                </van-row>
                <van-divider />
                <van-row>
                    <van-col span="24">
                        <font-awesome-icon icon="network-wired" style="font-size: .9rem;margin-right: 4px"/>
                            <span class="cell-head" style="margin-right: 16px">平均负载</span>
                            <span class="cell-data">{{cpu_load}}</span>
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

            <div class="cell"  @click="show_progress = true">
                <p class="cell-title">进程</p>
                <van-row style="margin-top: .5rem">
                    <van-col span="24">
                        <font-awesome-icon :icon="['fab', 'ubuntu']" v-if="kernel_os === 'ubuntu'" style="font-size: .9rem;margin-right: 4px"/>
                        <font-awesome-icon :icon="['fab', 'centos']" v-else-if="kernel_os === 'centos'" style="font-size: .9rem;margin-right: 4px"/>
                        <font-awesome-icon :icon="['fab', 'suse']" v-else-if="kernel_os === 'suse'" style="font-size: .9rem;margin-right: 4px"/>
                        <font-awesome-icon :icon="['fab', 'redhat']" v-else-if="kernel_os === 'redhat'" style="font-size: .9rem;margin-right: 4px"/>
                        <font-awesome-icon :icon="['fab', 'fedora']" v-else-if="kernel_os === 'fedora'" style="font-size: .9rem;margin-right: 4px"/>
                        <font-awesome-icon :icon="['fab', 'linux']" v-else style="font-size: .9rem;margin-right: 4px"/>
                        <span class="cell-data" style="margin-right: 8px">{{ kernel_type }}</span>
                        <span class="cell-data">{{kernel_version}}</span>
                    </van-col>
                </van-row>
                <van-divider />
                <van-row gutter="20" justify="center">
                    <van-col span="6">
                        <div>
                            <span class="cell-head">总进程</span><br>
                            <span class="cell-data">{{progress_all}}</span>
                        </div>
                    </van-col>
                    <van-col span="6">
                        <div>
                            <span class="cell-head">运行进程</span><br>
                            <span class="cell-data">{{progress_run}}</span>
                        </div>
                    </van-col>
                    <van-col span="6">
                        <div>
                            <span class="cell-head">僵尸进程</span><br>
                            <span class="cell-data">{{progress_dead}}</span>
                        </div>
                    </van-col>
                    <van-col span="6">
                        <div>
                            <span class="cell-head">休眠进程</span><br>
                            <span class="cell-data">{{progress_sleep}}</span>
                        </div>
                    </van-col>
                </van-row>
            </div>
            <!--配套的进程详情-->
            <van-action-sheet v-model:show="show_progress" title="进程详情" style="padding: 10px 0">
                <div class="content" style="padding: 0 10px;text-align: left">
                    <van-row gutter="20" justify="center" style="margin-top: .2rem">
                        <van-col span="4">
                            <span class="bold" style="margin-right: 16px;font-weight: bold">PID</span>
                        </van-col>
                        <van-col span="4">
                            <span style="margin-right: 16px;font-weight: bold">cpu(%)</span>
                        </van-col>
                        <van-col span="4">
                            <span class="cell-head" style="margin-right: 16px;font-weight: bold">mem(%)</span>
                        </van-col>
                        <van-col span="12" align="center">
                            <span class="cell-head" style="margin-right: 16px;font-weight: bold">命令</span>
                        </van-col>
                    </van-row>
                    <div v-for="p in progress_list">
                        <van-row gutter="20" justify="center" style="margin-top: .2rem">
                            <van-col span="4">
                                <span class="bold" style="margin-right: 16px;font-size: .9rem">{{ p.pid }}</span>
                            </van-col>
                            <van-col span="4">
                                <span style="margin-right: 16px;font-size: .9rem">{{ p.cpu }}</span>
                            </van-col>
                            <van-col span="4">
                                <span class="cell-head" style="margin-right: 16px;font-size: .9rem">{{ p.mem }}</span>
                            </van-col>
                            <van-col span="12">
                                <span class="cell-head"
                                      style="margin-right: 16px;
                                      width: 90%;
                                      font-size: .9rem;
                                      overflow: hidden;
                                      text-overflow: ellipsis;
                                      display: inline-block;
                                      white-space: nowrap"
                                      :title="p.cmd"
                                >{{ p.cmd }}</span>
                            </van-col>
                        </van-row>
                    </div>
                </div>
            </van-action-sheet>

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
                    <van-col span="6">
                        <div class="cell-data">
                            <font-awesome-icon icon="arrow-up" style="color: #a0a0a0" />&nbsp;<span>{{network_upload}}</span><span style="background-color: #ff7f50;height: 12px;width: 6px;display: inline-block;margin-left: 4px;border-radius: 4px"></span><br>
                            <font-awesome-icon icon="arrow-down" style="color: #a0a0a0" />&nbsp;<span>{{network_download}}</span><span style="background-color: #7fff00;height: 12px;width: 6px;display: inline-block;margin-left: 4px;border-radius: 4px"></span>
                        </div>
                    </van-col>
                    <van-col span="6" align="center">
                            <van-circle
                                    v-model:current-rate="network_init"
                                    :rate="network_percent" :speed="100"
                                    :stroke-width="180"
                                    size="42px"
                                    color="var(--light-circle-color, #7fff00)"
                                    layer-color="var(--light-circle-bg-color, #ff7f50)"/>
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
                <van-row gutter="10" justify="center" style="margin-top: .5rem">
                    <van-col span="14">
                        <div>
                            <span class="cell-head" style="font-weight: bold">/</span><br>
                            <span class="cell-data">{{ disk_mount }}</span>
                        </div>
                    </van-col>
                    <van-col span="10" align="center">
                        <div>
                            <span class="cell-head">{{disk_used}}</span>
                            <span class="cell-head">/{{disk_all}}</span>
                            <div class="data-bar" id="data-bar"><span id="data-bar-inner"></span></div>
                        </div>
                    </van-col>
                </van-row>
                <van-divider />
                <van-row gutter="20" justify="end" style="margin-top: .5rem">
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

          mem_usage: '10.00',
          mem_free: 0,
          mem_used: 0,
          mem_cache: 0,

          kernel_os: 'linux',
          kernel_type: 'x86_64',
          kernel_version: 'unknown version',
          progress_all: 0,
          progress_run: 0,
          progress_dead: 0,
          progress_sleep: 0,
          show_progress: false,
          progress_list: [{pid: 0, cpu: 0.4, mem: 0.1, cmd: '/bin/echo hello'}],

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
          // 是否使用整合接口获取数据
          if (this.$store.state.comps === 'Home' && this.$store.state.watchdog === 'false' || this.$store.state.watchdog === false) {
              if (this.$store.state.comb === 'true' || this.$store.state.comb === true) {
                  this.getCombinationData();
                  if (this.show_progress) {
                      this.getProgressListData();
                  }
              }else {
                  this.getServerData();
                  this.getCPUdata();
                  this.getMemData();
                  this.getKernelData();
                  this.getProgressData();
                  this.getNetData();
                  this.getDiskData();
                  if (this.show_progress) {
                      this.getProgressListData();
                  }
              }
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
          // 根据流量计算 换算规则G = 1024 * 1024 M=1024 K=1
          let download = 0;
          let upload = 0;
          if (this.network_upload.indexOf("G")) {
              upload = this.network_upload.replace("G", "") * 1024 * 1024;
          } else if (this.network_upload.indexOf("M")) {
              upload = this.network_upload.replace("M", "") * 1024;
          } else if (this.network_upload.indexOf("K")) {
              upload = this.network_upload.replace("K", "");
          }

          if (this.network_download.indexOf("G")) {
              download = this.network_download.replace("G", "") * 1024 * 1024;
          } else if (this.network_download.indexOf("M")) {
              download = this.network_download.replace("M", "") * 1024;
          } else if (this.network_download.indexOf("K")) {
              download = this.network_download.replace("K", "");
          }
          this.network_percent = (download / (download + upload) * 100).toFixed(0);
          return (download / (download + upload) * 100).toFixed(0);
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
      calcLoad(load, cpu_count) {
          let m1 = document.getElementById("m1");
          let m5 = document.getElementById("m5");
          let m15 = document.getElementById("m15");

          let loads = load.split(" ");
          if (m1) {
              let height = loads[0] * 100 / cpu_count / 2;
              height = Math.floor(height);
              if (height >= 100) {
                  height = 100;
                  m1.style.borderRadius = '4px';
              }
              m1.style.height = height + '%';
          }
          if (m5) {
              let height = loads[1] * 100 / cpu_count / 2;
              height = Math.floor(height);
              if (height >= 100) {
                  height = 100;
                  m5.style.borderRadius = '4px';
              }
              m5.style.height = height + '%';
          }
          if (m15) {
              let height = loads[2] * 100 / cpu_count / 2;
              height = Math.floor(height);
              if (height >= 100) {
                  height = 100;
                  m15.style.borderRadius = '4px';
              }
              m15.style.height = height + '%';
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
          }).catch((e)=>{
              if (e.response.data !== "key error") {
                  this.notifyDanger("接口" + apis.api_server + "请求失败");
              }
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

            this.calcLoad(this.cpu_load, this.cpu_count);
        }).catch((e) => {
            if (e.response.data !== "key error") {
                this.notifyDanger("接口" + apis.api_cpu + "请求失败");
            }
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
              }).catch((e) => {
              if (e.response.data !== "key error") {
                  this.notifyDanger("接口" + apis.api_mem + "请求失败");
              }
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
              }).catch((e) => {
              if (e.response.data !== "key error") {
                  this.notifyDanger("接口" + apis.api_net + "请求失败");
              }
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
              }).catch((e) => {
              if (e.response.data !== "key error") {
                  this.notifyDanger("接口" + apis.api_disk + "请求失败");
              }
          });
      },
        getKernelData() {
            this.$axios.post(apis.api_kernel)
                .then(res => {
                    let data = res.data.data;
                    this.kernel_os = data.kernel_os;
                    this.kernel_type = data.kernel_type;
                    this.kernel_version = data.kernel_version;
                }).catch((e) => {
                if (e.response.data !== "key error") {
                    this.notifyDanger("接口" + apis.api_kernel + "请求失败");
                }
            });
        },
        getProgressData() {
            this.$axios.post(apis.api_progress)
                .then(res => {
                    let data = res.data.data;
                    this.progress_all = data.progress_all;
                    this.progress_dead = data.progress_dead;
                    this.progress_run = data.progress_run;
                    this.progress_sleep = data.progress_sleep;
                }).catch((e) => {
                if (e.response.data !== "key error") {
                    this.notifyDanger("接口" + apis.api_progress + "请求失败");
                }
            });
        },
      // 获得进程列表
      getProgressListData() {
          this.$axios.post(apis.api_progress_list + "?num=" + this.$store.state.progress)
              .then(res => {
                  this.progress_list = res.data.data;
              }).catch((e) => {
              if (e.response.data !== "key error") {
                  this.notifyDanger("接口" + apis.api_progress_list + "请求失败");
              }
          });
      },
        getCombinationData() {
            this.$axios.post(apis.api_comb)
                .then(res => {
                    let server_data = res.data.data.server_info;
                    this.server = server_data? server_data: 'Unrecognizable server type';

                    let cpu_data = res.data.data.cpu_info;
                    this.cpu_usage = cpu_data.cpu_usage;
                    this.cpu_usage_system = cpu_data.cpu_usage_system;
                    this.cpu_usage_user = cpu_data.cpu_usage_user;
                    this.cpu_io_wait = cpu_data.cpu_io_wait;
                    this.cpu_count = cpu_data.cpu_count;
                    this.cpu_free = cpu_data.cpu_free;
                    this.cpu_load = cpu_data.cpu_load;
                    this.cpu_run = cpu_data.cpu_run;

                    this.calcLoad(this.cpu_load, this.cpu_count);

                    let mem_data = res.data.data.mem_info;
                    this.mem_usage = mem_data.mem_usage;
                    this.mem_free = mem_data.mem_free;
                    this.mem_used = mem_data.mem_used;
                    this.mem_cache = mem_data.mem_cache;

                    let kernel_data = res.data.data.kernel_info;
                    this.kernel_os = kernel_data.kernel_os;
                    this.kernel_type = kernel_data.kernel_type;
                    this.kernel_version = kernel_data.kernel_version;

                    let progress_data = res.data.data.progress_info;
                    this.progress_all = progress_data.progress_all;
                    this.progress_dead = progress_data.progress_dead;
                    this.progress_run = progress_data.progress_run;
                    this.progress_sleep = progress_data.progress_sleep;

                    let net_data = res.data.data.net_info;
                    this.net_upload = net_data.net_upload;
                    this.net_download = net_data.net_download;
                    this.network_upload = net_data.network_upload;
                    this.network_download = net_data.network_download;
                    this.network_percent = this.calcNetPercent();
                    this.net_retry = net_data.net_retry;
                    this.net_active = net_data.net_active;
                    this.net_passive = net_data.net_passive;
                    this.net_fail = net_data.net_fail;
                    this.ipv4 = net_data.ipv4;
                    this.ipv6 = net_data.ipv6;

                    let disk_data = res.data.data.disk_info;
                    this.disk_mount = disk_data.disk_mount;
                    this.disk_used = disk_data.disk_used;
                    this.disk_all = disk_data.disk_all;
                    this.disk_usage = disk_data.disk_usage;
                    this.disk_read_rate = disk_data.disk_read_rate;
                    this.disk_write_rate = disk_data.disk_write_rate;
                    this.disk_read_byte = disk_data.disk_read_byte;
                    this.disk_write_byte = disk_data.disk_write_byte;
                    this.disk_read_delay = disk_data.disk_read_delay;
                    this.disk_write_delay = disk_data.disk_write_delay;

                    this.calcBar();
                }).catch((e) => {
                    if (e.response.data !== "key error") {
                        this.notifyDanger("接口" + apis.api_comb + "请求失败");
                    }
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

    @media (max-width: 580px) {
        .cell-head.cell-head-mem {
            font-size: 1.4rem;
        }
    }
    @media (max-width: 540px) {
        .cell-head.cell-head-mem {
            font-size: 1.2rem;
        }
    }
    @media (max-width: 440px) {
        .cell-head.cell-head-mem {
            font-size: 1rem;
        }
    }
    @media (max-width: 380px) {
        .cell-head.cell-head-mem {
            font-size: .85rem;
        }
    }
    .data-bar {
        height: 40px;
        width: 24px;
        display: inline-block;
        background-color: var(--light-bg-container-color, #1a1a1a);
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
        left: 0;
        border-bottom-left-radius: 6px;
        border-bottom-right-radius: 6px;
        transition: height .3s ease;
    }

    .load-bar {
        height: 30px;
        width: 16px;
        display: inline-block;
        background-color: var(--light-bg-container-color, #1a1a1a);
        position: relative;
        border-radius: 4px;
        margin-right: 4px;
    }

    .load-bar:last-child {
        margin-right: 0;
    }

    .load-bar span {
        height: 30%;
        width: 16px;
        display: inline-block;
        position: absolute;
        bottom: 0;
        left: 0;
        border-bottom-left-radius: 4px;
        border-bottom-right-radius: 4px;
        transition: height .3s ease;
    }

    .load-bar #m1 {
        background-color: #65ff1c;
    }
    .load-bar #m5 {
        background-color: #ff9b1b;
    }
    .load-bar #m15 {
        background-color: #ff4e06;
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
    .home /deep/ .van-divider {
        margin: 10px 0;
    }
    .home /deep/ .van-popup {
        background-color: var(--light-bg-container-color, #1a1a1a);
    }
    .home /deep/ .van-action-sheet {
        color: var(--light-text-color, #a0a0a0);
    }
</style>
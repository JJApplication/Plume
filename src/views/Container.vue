<template>
    <div class="container">
        <Header name="容器"></Header>
        <span id="create-btn"><font-awesome-icon icon="plus" @click="open_create"></font-awesome-icon></span>
        <div class="docker-top">
            <font-awesome-icon :icon="['fab', 'docker']"></font-awesome-icon>
        </div>
        <van-notice-bar
            left-icon="info-o"
            mode="closeable"
            text="Docker容器展示基于Docker API"
        />
        <div class="scroll">
            <van-skeleton title :row="6" :loading="!visible" style="margin-top: 2rem"/>
            <van-empty description="容器功能将在后台支持Docker API时自动开启" image="network" v-show="!containers" />
            <transition name="van-slide-down">
                <div v-show="visible" class="container-list">
                    <div class="container-body" v-for="c in containers" :key="c.id" @click="showDetail(c.id)">
                        <p class="container-title">{{c.name && c.name.length == 1 ? c.name[0] : c.name}}</p>
                        <van-row justify="space-between" class="container-info">
                            <van-col span="12"><span style="font-weight: bold">镜像</span><br><span class="container-info-data">{{c.image}}</span></van-col>
                            <van-col span="9"><span style="font-weight: bold">创建</span><br><span class="container-info-data">{{c.date}}</span></van-col>
                            <van-col span="3" align="center"><span style="font-weight: bold">状态</span><br>
                                <font-awesome-icon icon="stop" style="color: red" v-if="c.status === 'paused'"/>
                                <font-awesome-icon icon="skull" style="color: grey" v-if="c.status === 'exited'"/>
                                <font-awesome-icon icon="play" style="color: forestgreen" v-if="c.status !== 'paused' && c.status !== 'exited'"/>
                            </van-col>
                        </van-row>
                    </div>
                </div>
            </transition>
        </div>
        <van-action-sheet v-model="show_create" title="创建容器" @close="clear">
            <van-form @submit="create">
                <van-field
                    v-model="container_name"
                    name="容器名"
                    label="容器名"
                    placeholder="容器名"
                />
                <van-field
                    v-model="container_user"
                    name="运行用户"
                    label="运行用户"
                    placeholder="运行用户"
                />
                <van-field
                    v-model="container_image"
                    name="镜像"
                    label="镜像"
                    placeholder="镜像"
                    :rules="[{ required: true, message: '镜像必须指定' }]"
                />
                <van-field
                    v-model="container_ports"
                    name="暴露端口"
                    label="暴露端口"
                    placeholder="暴露端口(使用空格分隔80:80)"
                />
                <van-field
                    v-model="container_volumes"
                    name="挂载目录"
                    label="挂载目录"
                    placeholder="挂载目录(使用空格分隔/root:/home)"
                />
                <van-field
                    v-model="container_cmds"
                    name="命令"
                    label="命令"
                    placeholder="命令(使用逗号分隔)"
                />
                <van-field
                    v-model="container_envs"
                    name="环境变量"
                    label="环境变量"
                    placeholder="环境变量(使用空格分隔foo=foo)"
                />
                <van-field name="switch" label="使用tty">
                    <template #input>
                        <van-switch v-model="container_tty" size="20" />
                    </template>
                </van-field>
                <van-field name="switch" label="后台运行">
                    <template #input>
                        <van-switch v-model="container_daemon" size="20" />
                    </template>
                </van-field>
                <van-field name="switch" label="自动删除">
                    <template #input>
                        <van-switch v-model="container_rm" size="20" />
                    </template>
                </van-field>
                <van-field
                    v-model="final_cmd"
                    name="预览"
                    label="预览"
                    placeholder="最终命令形式"
                />
                <div style="margin: 16px;">
                    <van-button round block type="info" native-type="submit">创建</van-button>
                </div>
            </van-form>
        </van-action-sheet>
    </div>
</template>

<script>
import Header from "../components/Header";
import apis from "../actions/api";

export default {
    name: "Container",
    components: {Header},
    data() {
        return {
            visible: false,
            show_create: false,
            containers: [
                {
                    id: '0',
                    name: 'plume',
                    image: 'landers1037/plume',
                    date: '2021-01-01 12:00',
                    status: 'exit'
                },
                {
                    id: '1',
                    name: 'plume',
                    image: 'landers1037/plume',
                    date: '2021-01-01 12:00',
                    status: 'stop'
                },
                {
                    id: '2',
                    name: 'plume',
                    image: 'landers1037/plume',
                    date: '2021-01-01 12:00',
                    status: 'start'
                }
            ],
            container_name: '',
            container_user: '',
            container_image: '',
            container_ports: '',
            container_volumes: '',
            container_cmds: '',
            container_envs: '',
            container_tty: false,
            container_daemon: false,
            container_rm: true,
            final_cmd: ''
        }
    },
    computed: {
      dataRange() {
          const {container_name, container_user, container_image,
              container_ports, container_volumes, container_cmds,
              container_envs, container_tty, container_daemon, container_rm} = this
          return {
              container_name, container_user, container_image,
              container_ports, container_volumes, container_cmds,
              container_envs, container_tty, container_daemon, container_rm
          }
      }
    },
    watch: {
        dataRange() {
            this.generate_final_cmd();
        }
    },
    mounted() {
        if (this.$store.state.watchdog === 'false' || this.$store.state.watchdog === false) {
            this.getContainers()
        }else {
            setTimeout(() => {
                this.visible = true
            }, 1000)
        }
    },
    methods: {
        showDetail(id) {
            this.$store.commit('changeID', id);
            this.$store.commit('changeComps', 'Detail');
        },
        // 告警提示
        notifySuccess(message) {
            this.$notify({ type: 'success', message: message });
        },
        notifyWarning(message) {
            this.$notify({ type: 'warning', message: message });
        },
        notifyDanger(message) {
            this.$notify({ type: 'danger', message: message });
        },
        getContainers() {
            this.$axios.post(apis.api_containers)
            .then(res => {
                this.containers = res.data.data;
                this.visible = true;
            }).catch(()=>{
                this.notifyDanger("获取容器列表失败")
            });
        },
        // 创建容器
        open_create() {
            this.show_create = true;
        },
        create() {
            let data = {};
            data.HostConfig = {}
            if (this.container_user) {
                data.User = this.container_user
            }
            if (this.container_image) {
                data.Image = this.container_image
            }
            if (this.container_tty) {
                data.Tty = this.container_tty
            }
            if (this.container_daemon) {
                data.HostConfig.Init = true
            }
            if (this.container_rm) {
                data.HostConfig.AutoRemove = true
            }
            if (this.container_ports) {
                data.HostConfig = {"PortBindings": {}}
                data.ExposedPorts = {}
                let ports = this.container_ports.split(" ");
                console.log(ports)
                for(let p of ports) {
                    let pd = p.split(':')
                    let key = pd[1] + '/tcp'
                    data.ExposedPorts[key] = {}
                    // for host port
                    data.HostConfig.PortBindings[key] = [{"HostPort": pd[0]}]
                }
            }
            if (this.container_volumes) {
                data.HostConfig.Binds = this.container_volumes.split(" ");
            }
            if (this.container_envs) {
                data.Env = this.container_envs.split(" ")
            }
            if (this.container_cmds) {
                data.Cmd = this.container_cmds.split(" ")
            }
            this.$axios.post(apis.api_container_create + "?name=" + this.container_name, data)
            .then(res=>{
                let code = res.data.data;
                console.log(code)
                switch (code) {
                    case "201":
                        this.notifySuccess("容器创建成功");
                        break
                    case "400":
                        this.notifyWarning("参数错误");
                        break
                    case "409":
                        this.notifyWarning("容器创建冲突");
                        break
                    case "404":
                        this.notifyDanger("镜像不存在");
                        break
                    default:
                        this.notifyDanger("容器创建失败")
                }
                this.getContainers();
            })
        },
        generate_final_cmd() {
            let data = 'docker run';
            if (this.container_name) {
                data += ' --name ' + this.container_name;
            }
            if (this.container_user) {
                data += ' --user ' + this.container_user
            }
            if (this.container_tty) {
                data += ' -it '
            }
            if (this.container_daemon) {
                data += ' -d '
            }
            if (this.container_rm) {
                data += ' --rm '
            }
            if (this.container_volumes) {
                let volumes = this.container_volumes.split(" ");
                for (let v of volumes) {
                    data += ' -v ' + v
                }
            }
            if (this.container_ports) {
                let ports = this.container_ports.split(" ");
                for (let p of ports) {
                    data += ' -p ' + p
                }
            }
            if (this.container_envs) {
                let envs = this.container_envs.split(" ");
                for(let e of envs) {
                    data += ' -e ' + e
                }
            }
            if (this.container_image) {
                data += ' ' + this.container_image
            }
            if (this.container_cmds) {
                data += ' ' + this.container_cmds
            }
            this.final_cmd = data;
        },
        clear() {
            this.container_name = '';
            this.container_user = '';
            this.container_image = '';
            this.container_ports = '';
            this.container_volumes = '';
            this.container_cmds = '';
            this.container_envs = '';
            this.container_tty = false;
            this.container_daemon = false;
            this.container_rm = true;
            this.final_cmd = '';
        }
    }
}
</script>

<style scoped>
    .container {
        padding: 8px 0;
        height: 100%;
        position: relative;
    }
    #create-btn {
        position: absolute;
        top: 6px;
        right: 0;
        cursor: pointer;
    }
    .docker-top {
        padding: 10px;
        background-color: var(--light-bg-docker-color, #2a2b2b);
        border-radius: 4px;
        margin: 10px 0;
        font-size: 1.2rem;
    }
    .container .scroll {
        overflow-y: auto;
        height: calc(100% - 180px);
    }
    .container .scroll::-webkit-scrollbar {
        display: none;
        width: 0;
        height: 0;
    }
    .container .container-list {
        margin-top: 1rem;
        text-align: left;
        overflow-y: auto;
    }
    .container .container-body {
        cursor: pointer;
        background-color: var(--light-bg-docker-color, #2a2b2b);
        padding: 10px;
        border-radius: 8px;
        margin-bottom: 10px;
    }
    .container .container-title {
        font-weight: bold;
        font-size: .9rem;
        color: var(--light-text-title-color, #707070);
    }
    .container .container-info {
        font-size: .85rem;
        margin-top: 8px;
    }
    .container-info .container-info-data {
        color: var(--light-text-bold-color, #7d7d7d);
        word-break: break-all;
        word-wrap: break-word;
    }
    .container /deep/ .van-skeleton__row, .container /deep/ .van-skeleton__title, .container /deep/ .van-skeleton__avatar{
        background-color: var(--light-loading-color, #252525);
    }
</style>
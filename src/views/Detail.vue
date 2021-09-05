<template>
    <div class="detail">
        <Header name="容器信息"></Header>
        <p class="container-id">
            <font-awesome-icon @click="back" :icon="['fab', 'docker']" style="margin-right: 10px;cursor: pointer"/>
            <span>ID: </span><span class="id">{{id}}</span>
        </p>
        <div class="scroll">
            <van-skeleton title :row="6" :loading="!visible" style="margin-top: 2rem"/>
            <div class="info-group" v-show="visible">
                <div class="info">
                    <div class="info-title">
                        <span>名称</span>
                        <span class="icon"><font-awesome-icon icon="globe-americas"/></span>
                    </div>
                    <div class="info-data">
                        <span>{{ name ? name : 'none' }}</span>
                    </div>
                </div>
                <div class="info">
                    <div class="info-title">
                        <span>命令</span>
                        <span class="icon"><font-awesome-icon icon="terminal"/></span>
                    </div>
                    <div class="info-data">
                        <span>{{ cmd }}</span>
                    </div>
                </div>
                <div class="info">
                    <div class="info-title">
                        <span>镜像</span>
                        <span class="icon"><font-awesome-icon icon="cloud"/></span>
                    </div>
                    <div class="info-data">
                        <span>{{ image }}</span>
                    </div>
                </div>
                <div class="info">
                    <div class="info-title">
                        <span>用户</span>
                        <span class="icon"><font-awesome-icon icon="user-tie"/></span>
                    </div>
                    <div class="info-data">
                        <span>{{ user ? user : 'unknown' }}</span>
                    </div>
                </div>
                <div class="info">
                    <div class="info-title">
                        <span>挂载卷</span>
                        <span class="icon"><font-awesome-icon icon="database"/></span>
                    </div>
                    <div class="info-data">
                        <span>{{ calc_volume }}</span>
                    </div>
                </div>
                <div class="info">
                    <div class="info-title">
                        <span>暴露端口</span>
                        <span class="icon"><font-awesome-icon icon="external-link-square-alt"/></span>
                    </div>
                    <div class="info-data">
                        <span>{{ calc_port }}</span>
                    </div>
                </div>
                <div class="info">
                    <div class="info-title">
                        <span>工作路径</span>
                        <span class="icon"><font-awesome-icon icon="wave-square"/></span>
                    </div>
                    <div class="info-data">
                        <span>{{ wrkdir ? wrkdir : 'no workdir' }}</span>
                    </div>
                </div>
                <div class="info">
                    <div class="info-title">
                        <span>创建时间</span>
                        <span class="icon"><font-awesome-icon icon="calendar-alt"/></span>
                    </div>
                    <div class="info-data">
                        <span>{{ date }}</span>
                    </div>
                </div>
                <div class="info">
                    <div class="info-title">
                        <span>运行状态</span>
                        <span class="icon"><font-awesome-icon icon="eye"/></span>
                    </div>
                    <div class="info-data">
                        <span>{{ status ? status : 'unknown' }}</span>
                    </div>
                </div>
                <div class="info">
                    <div class="info-title">
                        <span>CPU占用</span>
                        <span class="icon"><font-awesome-icon icon="microchip"/></span>
                    </div>
                    <div class="info-data">
                        <span>{{ cpu ? cpu : 0 }} %</span>
                    </div>
                </div>
                <div class="info">
                    <div class="info-title">
                        <span>内存占用</span>
                        <span class="icon"><font-awesome-icon icon="memory"/></span>
                    </div>
                    <div class="info-data">
                        <span>{{ mem ? mem : 0 }} %</span>
                    </div>
                </div>
                <div class="info">
                    <div class="info-title">
                        <span>容器操作</span>
                        <span class="icon"><font-awesome-icon icon="cat"/></span>
                    </div>
                    <div class="info-data">
                        <van-button type="primary" size="small" style="margin-right: 4px" @click="start">启动容器</van-button>
                        <van-button type="danger" size="small" style="margin-right: 4px" @click="stop">停止容器</van-button>
                        <van-button type="info" size="small" style="margin-right: 4px" @click="restart">重启容器</van-button>
                        <van-button type="warning" size="small" @click="del">删除容器</van-button>
                    </div>
                </div>

            </div>
        </div>
    </div>
</template>

<script>
// 容器的详细信息
import Header from "../components/Header";
import apis from "../actions/api";
export default {
    name: "Detail",
    components: {Header},
    data() {
        return {
            visible: false,
            id: this.$store.state.container_id !== null ? this.$store.state.container_id : 'unknown',
            name: '',
            cmd: ['/bin/bash'],
            image: 'landers1037/plume',
            user: '',
            volume: ['/dev'],
            port: [],
            wrkdir: '',
            date: '2021-01-01',
            status: '',
            cpu: 0,
            mem: 0
        }
    },
    mounted() {
        if (this.$store.state.watchdog === 'false' || this.$store.state.watchdog === false) {
            this.getContainerInfo()
        }else {
            setTimeout(() => {
                this.visible = true
            }, 1000)
        }
    },
    beforeDestroy() {
        this.$store.commit('changeID', null);
    },
    computed: {
        calc_port: function () {
            let map = this.port;
            if (map) {
                let data = [];
                for (let k in map) {
                    // 本地端口数组
                    let d = map[k];
                    let s = '';
                    for (let port of d) {
                        s += port["HostPort"]
                    }
                    data.push(k + ':' + s)
                }
                return data.join(', ')
            }

            return ''
        },
        calc_volume: function () {
            let volumes = this.volume;
            if (volumes) {
                return volumes.join(', ')
            }
            return ''
        }
    },
    methods: {
        back() {
            this.$store.commit('changeComps', 'Container');
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
        getContainerInfo() {
            this.$axios.post(apis.api_container + "?id=" + this.id)
            .then(res => {
                let data = res.data.data;
                this.name = data.name;
                this.cmd = data.cmd;
                this.image = data.image;
                this.user = data.user;
                this.volume = data.volume;
                this.port = data.port;
                this.wrkdir = data.wrkdir;
                this.date = data.date;
                this.status = data.status;
                this.cpu = data.cpu;
                this.mem = data.mem;

                this.visible = true;
            }).catch(()=>{
               this.notifyDanger("获取容器信息失败");
            });
        },
        start() {
            this.$axios.post(apis.api_container_start + "?id=" + this.id)
            .then(res=>{
                let code = res.data.data;
                switch (code) {
                    case "204":
                        this.notifySuccess("容器启动成功");
                        break
                    case "304":
                        this.notifyWarning("容器已启动");
                        break
                    case "404":
                        this.notifyDanger("容器不存在");
                        break
                    default:
                        this.notifyDanger("容器启动失败")
                }
                this.getContainerInfo();
            })
        },
        stop() {
            this.$axios.post(apis.api_container_stop + "?id=" + this.id)
                .then(res=>{
                    let code = res.data.data;
                    switch (code) {
                        case "204":
                            this.notifySuccess("容器停止成功")
                            break
                        case "304":
                            this.notifyWarning("容器已停止");
                            break
                        case "404":
                            this.notifyDanger("容器不存在");
                            break
                        default:
                            this.notifyDanger("容器停止失败")
                    }
                    this.getContainerInfo();
                })
        },
        restart() {
            this.$axios.post(apis.api_container_restart + "?id=" + this.id)
                .then(res=>{
                    let code = res.data.data;
                    switch (code) {
                        case "204":
                            this.notifySuccess("容器重启成功")
                            break
                        case "304":
                            this.notifyWarning("容器已重启");
                            break
                        case "404":
                            this.notifyDanger("容器不存在");
                            break
                        default:
                            this.notifyDanger("容器重启失败")
                    }
                    this.getContainerInfo();
                })
        },
        del() {
            this.$axios.post(apis.api_container_delete + "?id=" + this.id)
                .then(res=>{
                    let code = res.data.data;
                    switch (code) {
                        case "204":
                            this.notifySuccess("容器删除成功")
                            break
                        case "400":
                            this.notifyWarning("删除参数错误");
                            break
                        case "404":
                            this.notifyDanger("容器不存在");
                            break
                        default:
                            this.notifyDanger("容器删除失败")
                    }
                })
        }
    }
}
</script>

<style scoped>
    .detail {
        padding: 8px 0;
        height: 100%;
    }
    .container-id {
        text-align: left;
        font-size: 1.2rem;
        margin-top: 20px;
        font-weight: bold;
        color: var(--light-text-bold-color, #7d7d7d);
        overflow: hidden;
        text-overflow: ellipsis;
        white-space: nowrap;
    }
    .container-id span {
        color: var(--light-text-color, #a0a0a0);
    }
    .container-id .id {

    }
    .detail .scroll {
        overflow-y: auto;
        height: calc(100% - 150px);
        border-radius: 10px;
        margin-top: 10px;
    }
    .detail .scroll::-webkit-scrollbar {
        display: none;
        width: 0;
        height: 0;
    }

    @media (max-width: 480px) {
        .detail /deep/ .van-cell-group--inset {
            margin: 0;
        }
    }
    .detail /deep/ .van-cell__title span {
        font-weight: bold;
        font-size: 1rem;
    }
    .detail /deep/ .van-cell__label span {
        font-size: .8rem;
        font-weight: normal;
    }
    .detail /deep/ .van-cell-group {
        background-color: var(--light-bg-color, #202020);
    }
    .detail /deep/ .van-cell {
        background-color: var(--light-bg-color, #202020);
        color: var(--light-text-color, #a0a0a0);
    }
    .detail /deep/ .van-hairline--top-bottom::after, .detail /deep/ .van-hairline-unset--top-bottom::after {
        border: none;
    }
    .detail /deep/ .van-cell::after {
        border-bottom-color: var(--light-line-border-color, #707070);
    }
    .detail /deep/ .van-skeleton__row, .detail /deep/ .van-skeleton__title, .detail /deep/ .van-skeleton__avatar{
        background-color: var(--light-loading-color, #252525);
    }

    .detail .info-group {
        text-align: left;
        margin-top: 10px;
        padding: 0 20px 4px 20px;
        background-color: var(--light-bg-color, #202020);
        border-radius: 10px;
    }
    @media (max-width: 520px) {
        .detail .info-group {
            padding: 0 10px 4px 10px;
        }
    }
    .info-group .info {
        padding: 10px 8px;
        position: relative;
    }
    .info-group .info::after {
        position: absolute;
        content: ' ';
        left: 10px;
        right: 10px;
        bottom: 0;
        transform: scaleY(.5);
        box-sizing: border-box;
        border-bottom: 1px solid var(--light-line-border-color, #707070);
    }
    .info-group .info:last-child::after {
        border: none;
    }
    .info-group .info .info-title {
        font-size: 1rem;
        font-weight: bold;
        color: var(--light-text-color, #a0a0a0);
        position: relative;
        margin-bottom: 6px;
    }
    .info-title .icon {
        position: absolute;
        right: 0;
    }
    .info-group .info-data {
        font-size: .9rem;
        color: var(--light-text-title-color, #404040);
        word-break: break-all;
    }
</style>
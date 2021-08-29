<template>
    <div class="container">
        <Header name="容器"></Header>
        <div class="docker-top">
            <font-awesome-icon :icon="['fab', 'docker']"></font-awesome-icon>
        </div>
        <van-notice-bar
            left-icon="info-o"
            mode="closeable"
            text="Docker容器展示基于Docker API"
        />
        <div class="scroll">
            <div class="container-list">
                <div class="container-body" v-for="c in containers" :key="c.name">
                    <p class="container-title">{{c.name}}</p>
                    <van-row justify="space-between" class="container-info">
                        <van-col span="9"><span style="font-weight: bold">镜像</span><br><span class="container-info-data">{{c.image}}</span></van-col>
                        <van-col span="6"><span style="font-weight: bold">创建</span><br><span class="container-info-data">{{c.date}}</span></van-col>
                        <van-col span="6"><span style="font-weight: bold">端口</span><br><span class="container-info-data">{{c.port}}</span></van-col>
                        <van-col span="3"><span style="font-weight: bold">状态</span><br>
                            <font-awesome-icon icon="stop" style="color: red" v-if="c.status === 'stop'"/>
                            <font-awesome-icon icon="skull" style="color: grey" v-if="c.status === 'exit'"/>
                            <font-awesome-icon icon="play" style="color: forestgreen" v-if="c.status !== 'stop' && c.status !== 'exit'"/>
                        </van-col>
                    </van-row>
                </div>
            </div>
        </div>

    </div>
</template>

<script>
import Header from "../components/Header";
export default {
    name: "Container",
    components: {Header},
    data() {
        return {
            containers: [
                {
                    name: 'plume',
                    image: 'landers1037/plume',
                    date: '2021-01-01',
                    port: '',
                    status: 'exit'
                },
                {
                    name: 'plume',
                    image: 'landers1037/plume',
                    date: '2021-01-01',
                    port: [80],
                    status: 'stop'
                },
                {
                    name: 'plume',
                    image: 'landers1037/plume',
                    date: '2021-01-01',
                    port: [80, 8080],
                    status: 'start'
                }
            ]
        }
    }
}
</script>

<style scoped>
    .container {
        padding: 8px 0;
        height: 100%;
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
        background-color: var(--light-bg-docker-color, #2f2f2f);
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
</style>
<template>
    <div class="detail">
        <Header name="容器信息"></Header>
        <p class="container-id">
            <font-awesome-icon @click="back" :icon="['fab', 'docker']" style="margin-right: 10px;cursor: pointer"/>
            <span>ID: </span>{{id}}
        </p>
        <div class="scroll">
            <van-skeleton title :row="6" :loading="!visible" style="margin-top: 2rem"/>
            <van-cell-group v-show="visible" inset>
                <van-cell title="名称">
                    <template #extra>
                        <font-awesome-icon icon="globe-americas"/>
                    </template>
                    <template #label>
                        <span>{{ name }}</span>
                    </template>
                </van-cell>
                <van-cell title="命令">
                    <template #extra>
                        <font-awesome-icon icon="terminal"/>
                    </template>
                    <template #label>
                        <span>{{ cmd }}</span>
                    </template>
                </van-cell>
                <van-cell title="镜像">
                    <template #extra>
                        <font-awesome-icon icon="cloud"/>
                    </template>
                    <template #label>
                        <span>{{ image }}</span>
                    </template>
                </van-cell>
                <van-cell title="用户">
                    <template #extra>
                        <font-awesome-icon icon="user-tie"/>
                    </template>
                    <template #label>
                        <span>{{ user }}</span>
                    </template>
                </van-cell>
                <van-cell title="挂载卷">
                    <template #extra>
                        <font-awesome-icon icon="database"/>
                    </template>
                    <template #label>
                        <span>{{ volume }}</span>
                    </template>
                </van-cell>
                <van-cell title="工作路径">
                    <template #extra>
                        <font-awesome-icon icon="wave-square"/>
                    </template>
                    <template #label>
                        <span>{{ wrkdir }}</span>
                    </template>
                </van-cell>
                <van-cell title="创建时间">
                    <template #extra>
                        <font-awesome-icon icon="calendar-alt"/>
                    </template>
                    <template #label>
                        <span>{{ date }}</span>
                    </template>
                </van-cell>
                <van-cell title="运行状态">
                    <template #extra>
                        <font-awesome-icon icon="eye"/>
                    </template>
                    <template #label>
                        <span>{{ status }}</span>
                    </template>
                </van-cell>
                <van-cell title="CPU占用">
                    <template #extra>
                        <font-awesome-icon icon="microchip"/>
                    </template>
                    <template #label>
                        <span>{{ cpu }} %</span>
                    </template>
                </van-cell>
                <van-cell title="内存占用">
                    <template #extra>
                        <font-awesome-icon icon="memory"/>
                    </template>
                    <template #label>
                        <span>{{ mem }}</span>
                    </template>
                </van-cell>
            </van-cell-group>
        </div>
    </div>
</template>

<script>
// 容器的详细信息
import Header from "../components/Header";
export default {
    name: "Detail",
    components: {Header},
    data() {
        return {
            visible: false,
            id: this.$store.state.container_id !== null ? this.$store.state.container_id : 'unknown',
            name: 'none',
            cmd: ['/bin/bash'],
            image: 'landers1037/plume',
            user: 'root',
            volume: ['/dev'],
            wrkdir: '/root',
            date: '2021-01-01',
            status: 'running',
            cpu: 0,
            mem: 0
        }
    },
    mounted() {
      setTimeout(() => {
          this.visible = true
      }, 1000)
    },
    beforeDestroy() {
        this.$store.commit('changeID', null);
    },
    methods: {
        back() {
            this.$store.commit('changeComps', 'Container');
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
        font-size: 1.4rem;
        margin-top: 20px;
        font-weight: bold;
        color: var(--light-text-bold-color, #7d7d7d);
    }
    .container-id span {
        color: var(--light-text-color, #a0a0a0);
    }
    .detail .scroll {
        overflow-y: auto;
        height: calc(100% - 140px);
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
        background-color: var(--light-bg-color, #2f2f2f);
    }
    .detail /deep/ .van-cell {
        background-color: var(--light-bg-color, #2f2f2f);
        color: var(--light-text-color, #a0a0a0);
    }
    .detail /deep/ .van-hairline--top-bottom::after, .detail /deep/ .van-hairline-unset--top-bottom::after {
        border: none;
    }
    .detail /deep/ .van-cell::after {
        border-bottom-color: var(--light-line-border-color, #707070);
    }
</style>
<template>
    <div class="bottom">
        <van-tabbar v-model="active">
            <van-tabbar-item name="Home" @click="changeIt('Home')"><font-awesome-icon icon="border-all" class="bottom-icon"></font-awesome-icon>面板</van-tabbar-item>
            <van-tabbar-item name="Container" @click="changeIt('Container')"><font-awesome-icon icon="atom" class="bottom-icon"></font-awesome-icon>容器</van-tabbar-item>
            <van-tabbar-item name="Setting" @click="changeIt('Setting')"><font-awesome-icon icon="cog" class="bottom-icon"></font-awesome-icon>设置</van-tabbar-item>
            <van-tabbar-item name="About" @click="changeIt('About')"><font-awesome-icon icon="copyright" class="bottom-icon"></font-awesome-icon>关于</van-tabbar-item>
        </van-tabbar>
        <div class="van-safe-area-bottom"></div>
    </div>
</template>

<script>
// 底部的按钮，绝对布局位置
import {mapMutations} from "vuex";

export default {
    name: "Bottom",
    data(){
        return {
            active: this.$store.state.comps ? this.$store.state.comps : 'Home',
            watch: {
                "$store.state.comps": function () {
                    this.active = this.$store.state.comps;
                }
            }
        }
    },
    methods: {
        ...mapMutations({
            change: 'changeComps',
        }),
        changeIt(c) {
            this.change(c);
        }
    }
}
</script>

<style scoped>
    .bottom {
        position: absolute;
        bottom: 0;
    }
    .bottom-icon {
        position: relative;
        margin: 0 auto 6px auto;
        display: block;
        font-size: 20px;
    }
    .bottom /deep/ .van-tabbar.van-tabbar--fixed {
        max-width: 760px;
        left: 50%;
        transform: translateX(-50%);
    }
    .bottom /deep/ .van-tabbar, .van-tabbar-item--active {
        background-color: var(--light-bg-color, #202020);
    }
    .bottom /deep/ .van-tabbar-item {
        color: var(--light-text-color, #a0a0a0);
    }
    .bottom /deep/ .van-tabbar-item--active {
        color: var(--light-text-active-color, #ff4a9e);
    }
    .bottom /deep/ .van-hairline--top-bottom::after, .van-hairline-unset--top-bottom::after {
        border-width: 0;
    }

    .bottom /deep/ .van-tabbar.van-tabbar--fixed {
        padding-bottom: 10px;
    }
    /*针对移动端设备*/
    @media (max-width: 480px) {
        .bottom /deep/ .van-tabbar.van-tabbar--fixed {
            padding-bottom: 20px;
        }
    }
</style>
<template>
    <div class="bottom">
        <van-tabbar v-model="active">
            <van-tabbar-item replace to="/" @click="changeIt('Home')"><font-awesome-icon icon="border-all" class="bottom-icon"></font-awesome-icon>面板</van-tabbar-item>
            <van-tabbar-item replace to="/container" @click="changeIt('Container')"><font-awesome-icon icon="atom" class="bottom-icon"></font-awesome-icon>容器</van-tabbar-item>
            <van-tabbar-item replace to="/setting" @click="changeIt('Setting')"><font-awesome-icon icon="cog" class="bottom-icon"></font-awesome-icon>设置</van-tabbar-item>
            <van-tabbar-item replace to="/about" @click="changeIt('About')"><font-awesome-icon icon="copyright" class="bottom-icon"></font-awesome-icon>关于</van-tabbar-item>
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
            active: 1
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
        max-width: 720px;
        margin: 0 auto;
        left: auto;
    }
</style>
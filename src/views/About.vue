<template>
  <div class="about">
      <Header :name="app"></Header>
        <h3>about {{app}}</h3>
        <p>{{app}} is a project for watching your own linux server.</p>
        <van-divider dashed>by renj.io</van-divider>
        <van-cell-group inset>
            <van-cell title="Github" label="github.com/landers1037/plume" is-link url="https://github.com/landers1037/plume" />
            <van-cell title="在线示例" label="plume.renj.io" is-link url="http://plume.renj.io" />
            <van-cell title="贡献issue" label="github.com/landers1037/plume" is-link url="https://github.com/landers1037/plume" />
        </van-cell-group>

        <div style="margin-top: 2rem">
            <a style="font-size: 1.6rem;margin-bottom: 1rem;display: block;color: var(--light-text-color, #f0f0f0)" href="https://github.com/landers1037/plume">
                <font-awesome-icon :icon="['fab', 'github']"></font-awesome-icon>
            </a>
            <van-rate
                v-model="rate"
                :size="25"
                color="#ffd21e"
                void-icon="star"
                void-color="#eee"
            />
        </div>

  </div>
</template>

<script>

import Header from "../components/Header";
export default {
    name: 'About',
    components: {Header},
    data(){
        return {
            app: this.$store.state.app,
            rate: 5
        }
    }
}
</script>
<style scoped>
    .about {
        line-height: 2rem;
    }
    .about /deep/ .van-cell {
        background-color: var(--light-bg-color, #2f2f2f);
    }
    .about /deep/ .van-cell .van-cell__title span{
        color: var(--light-cell-title-color, #cfcfcf);
    }
    .about /deep/ .van-hairline--top-bottom::after, .van-hairline-unset--top-bottom::after {
        border-width: 0;
    }
</style>
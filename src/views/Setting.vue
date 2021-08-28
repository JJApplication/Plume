<template>
    <div>
        <Header name="设置"></Header>
        <van-switch v-model="theme" active-color="var(--light-text-active-color, #ff4a9e)"/>
    </div>
</template>

<script>
import Header from "../components/Header";
export default {
    name: "Setting",
    components: {Header},
    data() {
        return {
            theme: this.$store.state.theme === 'dark' // true: dark
        }
    },
    computed: {
        theme_now: function() {
            // return this.theme === 'light';
            return false
        }
    },
    watch: {
        theme: function () {
            console.log(this.theme);
            let t = null;
            if (this.theme) {t='dark'} else {t='light'}
            this.$store.commit('changeTheme', t);
        },
        "$store.state.theme": function () {
            if (this.$store.state.theme === 'light') {
                this.theme = false
            }else {
                this.theme = true
            }
        }
    }
}
</script>

<style scoped>

</style>
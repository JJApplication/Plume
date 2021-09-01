<template>
    <div class="setting">
        <Header name="设置"></Header>
        <div class="scroll">
            <div class="setting-body">
                <van-cell-group inset>
                    <van-field name="switch" label="选择主题">
                        <template #input>
                            <van-switch v-model="theme" size="20" active-color="var(--light-text-active-color, #ff4a9e)" />
                        </template>
                    </van-field>
                    <van-field
                        v-model="app"
                        name="app"
                        label="应用名称"
                        placeholder="plume"
                        :rules="[{ required: true, message: '请填写应用名称' }]"
                    />
                    <van-field
                        v-model="key"
                        name="key"
                        label="通信密钥"
                        placeholder="input key if you have"
                    />
                    <van-field
                        v-model="api"
                        name="api"
                        label="API接口"
                        placeholder="input plume api"
                    />
                    <van-field name="stepper" label="触发器(s)">
                        <template #input>
                            <van-stepper v-model="duration" />
                        </template>
                    </van-field>
                    <van-field name="stepper" label="告警阈值(%)">
                        <template #input>
                            <van-stepper max="100" v-model="limit" />
                        </template>
                    </van-field>
                    <van-field name="stepper" label="进程数量">
                        <template #input>
                            <van-stepper max="50" v-model="progress" />
                        </template>
                    </van-field>
                    <van-field name="switch" label="暂停监控">
                        <template #input>
                            <van-switch v-model="watchdog" size="20" active-color="var(--light-text-active-color, #ff4a9e)" />
                        </template>
                    </van-field>
                    <van-field name="switch" label="重置配置" @click="resetStore">
                        <template #input>
                            <van-button type="info" size="mini"><font-awesome-icon icon="sync-alt" style="padding: 0 4px"/>重置</van-button>
                        </template>
                    </van-field>
                    <van-field name="switch" label="清除缓存" @click="clearStore">
                        <template #input>
                            <van-button type="danger" size="mini"><font-awesome-icon icon="trash" style="padding: 0 4px"/>清除</van-button>
                        </template>
                    </van-field>
                    <van-field name="switch" label="导出配置" @click="exportStore">
                        <template #input>
                            <van-button type="primary" size="mini"><font-awesome-icon icon="file-import" style="padding: 0 4px"/>导出</van-button>
                        </template>
                    </van-field>
                    <van-field name="uploader" label="导入配置">
                        <template #input>
                            <van-uploader :before-read="loadConfig" :max-size="1024" accept=".json"/>
                        </template>
                    </van-field>
                    <van-field label="查看日志" @click="showLog">
                        <template #input>
                            <van-button type="default" size="mini"><font-awesome-icon icon="code" style="padding: 0 4px"/>查看</van-button>
                        </template>
                    </van-field>
                    <van-field label="清空日志" @click="clearLog">
                        <template #input>
                            <van-button type="default" size="mini"><font-awesome-icon icon="truck" style="padding: 0 4px"/>清空</van-button>
                        </template>
                    </van-field>
                </van-cell-group>
            </div>
            <van-action-sheet v-model:show="show_log" title="运行日志" style="padding: 10px 0">
                <div class="content" style="padding: 0 10px;text-align: left">
                    <van-skeleton title :row="6" :loading="show_load && show_log"/>
                    <pre
                    style="word-break: break-all;
                        word-spacing: normal;
                        white-space: pre-line;
                        border-radius: 8px;
                        padding: 10px 4px;
                        background-color: var(--light-bg-docker-color, #2a2b2b);"
                    >{{log}}</pre>
                </div>
            </van-action-sheet>
        </div>
    </div>
</template>

<script>
import Header from "../components/Header";
import {delStore, initStore} from "../plugins/store";
import {exportConfig, readConfig} from "../plugins/config";
import apis from "../actions/api";
export default {
    name: "Setting",
    components: {Header},
    data() {
        return {
            theme: this.$store.state.theme === 'dark', // true: dark
            app: this.$store.state.app,
            key: this.$store.state.key,
            duration: this.$store.state.duration,
            limit: this.$store.state.limit,
            watchdog: this.$store.state.watchdog === 'true' || this.$store.state.watchdog === true,
            api: this.$store.state.api,
            progress: this.$store.state.progress,
            show_log: false,
            show_load: true,
            log: '',
        }
    },
    computed: {
        theme_now: function() {
            // return this.theme === 'light';
            return false
        }
    },
    mounted() {
        this.theme = this.$store.state.theme === 'dark';
        this.app = this.$store.state.app;
        this.key = this.$store.state.key;
        this.duration = this.$store.state.duration;
        this.limit = this.$store.state.limit;
        this.watchdog = this.$store.state.watchdog === 'true' || this.$store.state.watchdog === true;
        this.api = this.$store.state.api;
        this.progress = this.$store.state.progress;
    },
    watch: {
        theme: function () {
            console.log('theme changed', this.theme);
            let t = null;
            if (this.theme) {t='dark'} else {t='light'}
            this.$store.commit('changeTheme', t);
        },
        "$store.state.theme": function () {
            this.theme = this.$store.state.theme !== 'light';
        },
        app: function () {
            console.log('app name changed', this.app);
            this.$store.commit('changeApp', this.app);
        },
        key: function () {
            console.log('app key changed', this.key);
            this.$store.commit('changeKey', this.key);
        },
        duration: function () {
            console.log('request duration changed', this.duration);
            this.$store.commit('changeDuration', this.duration);
        },
        limit: function () {
            console.log('limit changed', this.limit);
            this.$store.commit('changeLimit', this.limit);
        },
        watchdog: function () {
            console.log('state watch changed', this.watchdog);
            this.$store.commit('changeWatch', this.watchdog);
        },
        api: function () {
            console.log('api changed', this.api);
            this.$store.commit('changeApi', this.api);
        },
        progress: function () {
            console.log('progress num changed', this.progress);
            this.$store.commit('changeProgress', this.progress);
        }
    },
    methods: {
        resetStore() {
          delStore();
          initStore();
          this.notifyReset();
          this.reloadState();
        },
        clearStore() {
            delStore();
            this.$store.commit('clearAll');
            this.reloadState();
            this.notifyClear();
        },
        reloadState() {
            this.$store.commit('reloadAll');
            this.theme = this.$store.state.theme === 'dark';
            this.app = this.$store.state.app;
            this.key = this.$store.state.key;
            this.duration = this.$store.state.duration;
            this.limit = this.$store.state.limit;
            this.watchdog = this.$store.state.watchdog;
            this.api = this.$store.state.api;
            this.progress = this.$store.state.progress;
        },
        exportStore() {
          // 暂时仅做json转换显示
          let data = exportConfig(this.$store.state);
            this.$dialog.alert({
                title: 'JSON配置',
                messageAlign: 'left',
                theme: 'round-button',
                message: JSON.stringify(data, ' ', 4),
            });
        },
        // 导入配置，监听文件变更函数
        loadConfig(file) {
            let reader = new FileReader();
            let that = this;
            reader.readAsText(file, 'utf8')
            reader.onload = function () {
                let e = readConfig(this.result)
                if (this.error || e === 'error') {
                    that.$notify({ type: 'danger', message: '文件解析失败' });
                }else {
                    that.reloadState();
                    that.$notify({ type: 'success', message: '配置更新成功' });
                }
            }
        },
        showLog() {
            this.show_log = true;
            this.$axios.post(apis.api_log)
            .then(res => {
                if (res.data.data && res.data.data !== '') {
                    this.log = res.data.data;
                    this.show_load = false;
                } else {
                    this.log = '暂时没有日志';
                    this.show_load = false;
                }
            }).catch(()=>{
                this.$notify({ type: 'danger', message: '日志获取失败' });
            })
        },
        clearLog() {
            this.$axios.post(apis.api_log_del)
                .then(res => {
                    if (res.data.data && res.data.data === "ok" ) {
                            this.$notify({ type: 'success', message: '日志清空完毕' });
                    } else {
                            this.$notify({ type: 'danger', message: '日志清空失败' });
                    }
                }).catch(()=>{
                this.$notify({ type: 'danger', message: '日志清空失败' });
            })
        },
        notifyReset() {
            this.$notify({ type: 'primary', message: '配置重置完毕' });
        },
        notifyClear() {
            this.$notify({ type: 'success', message: '缓存清理完毕' });
        }
    }
}
</script>

<style scoped>
    .setting {
        padding: 8px 0;
        height: 100%;
    }
    .setting .scroll {
        margin-top: 20px;
        overflow-y: auto;
        height: calc(100% - 110px);
        border-radius: 10px;
    }
    .setting .scroll::-webkit-scrollbar {
        display: none;
        width: 0;
        height: 0;
    }
    .setting /deep/ .van-cell-group, .van-calendar {
        background-color: var(--light-bg-color, #202020);
    }
    .setting /deep/ .van-cell {
        background-color: var(--light-bg-color, #202020);
        color: var(--light-text-color, #a0a0a0);
    }
    .setting /deep/ .van-calendar__month-mark {
        color: var(--light-month-color, rgba(242,243,245,.1));
    }
    .setting /deep/ .van-stepper__minus, .setting /deep/ .van-stepper__plus, .setting /deep/ .van-stepper__minus--disabled, .setting /deep/ .van-stepper__plus--disabled,.setting /deep/  .van-stepper__input {
        background-color: var(--light-step-bg-color, #101010);
        color: var(--light-step-color, #cfcfcf);
    }
    .setting /deep/ .van-hairline--top-bottom::after, .setting /deep/ .van-hairline-unset--top-bottom::after {
        border: none;
    }
    .setting /deep/ .van-field__control {
        color: var(--light-text-color, #f0f0f0);
    }
    .setting /deep/ .van-button--mini {
        padding: 2px 8px;
        font-size: 12px;
    }
    @media (max-width: 480px) {
        .setting /deep/ .van-cell-group--inset {
            margin: 0;
        }
    }
    .setting /deep/ .van-cell::after {
        border-bottom-color: var(--light-line-border-color, #707070);
    }
    .setting /deep/ .van-uploader__upload {
        background-color: var(--light-bg-container-color, #101010);
    }
    .setting /deep/ .van-popup {
        background-color: var(--light-bg-container-color, #1a1a1a);
    }
    .setting /deep/ .van-action-sheet {
        color: var(--light-text-color, #a0a0a0);
    }
</style>

<style>
    .van-dialog {
        background-color: var(--light-bg-color, #1f1f1f);
    }
    .van-dialog__header {
        color: var(--light-text-title-color, #7d7d7d);
    }
    .van-dialog--round-button .van-dialog__message {
        color: var(--light-text-title-color, #a0a0a0);
        font-family: monospace;
    }
    .van-goods-action {
        background-color: var(--light-bg-color, #1f1f1f);
    }
</style>
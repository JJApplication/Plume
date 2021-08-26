import Vue from 'vue'
import './plugins/fontawesome'
import './plugins/axios'
import App from './App.vue'
import './registerServiceWorker'
import router from './router'
import store from './store'
import './plugins/vant.js'
import {initStore} from "./plugins/store";

Vue.config.productionTip = false

// 初始化store
initStore();

new Vue({
  router,
  store,
  render: function (h) { return h(App) }
}).$mount('#app')

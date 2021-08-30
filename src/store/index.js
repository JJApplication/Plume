import Vue from 'vue'
import Vuex from 'vuex'
import {getStore, initStore, saveStore} from "../plugins/store";
import {change_theme} from "../plugins/change_theme";

Vue.use(Vuex)

initStore();

export default new Vuex.Store({
  state: {
    app: getStore("app"),
    theme: getStore("theme"),
    comps: getStore("comps"),
    watchdog: getStore("watchdog"),
    duration: getStore("duration"),
    limit: getStore("limit"),
    key: getStore("key"),
    api: getStore("api"),
    // 容器id无需持久化
    container_id: ''
  },
  mutations: {
    changeComps(state, comp) {
      saveStore("comps", comp);
      state.comps = comp;
    },
    changeTheme(state, theme) {
      change_theme(theme);
      state.theme = theme;
    },
    changeApp(state, app) {
      saveStore("app", app);
      state.app = app;
    },
    changeKey(state, key) {
      saveStore("key", key);
      state.key = key;
    },
    changeDuration(state, duration) {
      saveStore("duration", duration);
      state.duration = duration;
    },
    changeLimit(state, limit) {
      if (limit > 100) {
        limit = 100;
      }
      if (limit < 0) {
        limit = 0
      }
      saveStore("limit", limit);
      state.limit = limit;
    },
    changeWatch(state, watch) {
      saveStore("watchdog", watch);
      state.watchdog = watch;
    },
    changeApi(state, api) {
      saveStore("api", api);
      state.api = api;
    },
    changeID(state, id) {
      state.container_id = id;
    },
    clearAll(state) {
      state.app = null;
      state.theme = null;
      state.watchdog = null;
      state.duration = null;
      state.limit = null;
      state.key = null;
      state.api = null;
    },
    reloadAll(state) {
      state.app = getStore("app");
      state.theme = getStore("theme");
      state.watchdog = getStore("watchdog");
      state.duration = getStore("duration");
      state.limit = getStore("limit");
      state.key = getStore("key");
      state.api = getStore("api");
    }
  },
  actions: {
  },
  modules: {
  }
})

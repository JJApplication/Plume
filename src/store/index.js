import Vue from 'vue'
import Vuex from 'vuex'
import {getStore, initStore} from "../plugins/store";
import {change_theme} from "../plugins/change_theme";

Vue.use(Vuex)

initStore();

export default new Vuex.Store({
  state: {
    app: getStore("app"),
    theme: getStore("theme"),
    comps: getStore("comps")
  },
  mutations: {
    changeComps(state, comp) {
      state.comps = comp;
    },
    changeTheme(state, theme) {
      change_theme(theme);
      state.theme = theme;
    }
  },
  actions: {
  },
  modules: {
  }
})

import Vue from 'vue'
import Vuex from 'vuex'
import {getStore} from "../plugins/store";

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    app: getStore("app"),
    theme: getStore("theme"),
    comps: getStore("comps")
  },
  mutations: {
    changeComps(state, comp) {
      state.comps = comp;
    }
  },
  actions: {
  },
  modules: {
  }
})

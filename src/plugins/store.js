// 初始化localstorage
// 使用storage进行持久化 在每次重新打开时优先从store中读取配置参数

// 每次启动创建变量到store中 然后vuex加载变量
import consts from "../actions/consts";

export function initStore() {
    saveStore("app", consts.app)
    saveStore("theme", consts.default_theme)
    saveStore("comps", consts.comps)
}

// 从store中取值 不存在返会null
export function getStore(key) {
    return localStorage.getItem(key)
}

// 更新值 在vuex更新操作中 值会接下来更新到store中
export function saveStore(key, value) {
    localStorage.setItem(key, value)
}

// 删除全部缓存
export function delStore() {
    localStorage.clear()
}
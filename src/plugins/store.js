// 初始化localstorage
// 使用storage进行持久化 在每次重新打开时优先从store中读取配置参数

// 每次启动创建变量到store中 然后vuex加载变量
import consts from "../actions/consts";

export function initStore() {
    initKey("app", consts.app)
    initKey("theme", consts.default_theme)
    initKey("comps", consts.comps)
    initKey("watchdog", consts.watchdog)
    initKey("duration", consts.duration)
    initKey("limit", consts.limit)
    initKey("key", consts.key)
    initKey("api", consts.api)
    initKey("progress", consts.progress)
    initKey("theme_url", consts.theme_url)
}

function initKey(key, value) {
    if (!localStorage.getItem(key)) {
        saveStore(key, value)
    }
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
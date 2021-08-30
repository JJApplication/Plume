// 读取与导出配置文件
// c是配置json字符串
import {saveStore} from "./store";

const configMap = [
    "app",
    "theme",
    "comps",
    "watchdog",
    "duration",
    "limit",
    "key",
    "api"
]
export function readConfig(c) {
    let raw = JSON.parse(c);
    // 容器id非必要项无需
    for (let key of configMap) {
        if (!(key in raw)) {
            return  "error"
        }
    }

    // 全部遍历完毕 开始存储刷新配置
    for (let key of configMap) {
        saveStore(key, raw[key]);
    }
    return "ok"
}

export function exportConfig(d) {
    return d
}

export function changeConfig() {

}


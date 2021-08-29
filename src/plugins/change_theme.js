// 根据css root方式加载新的主题样式
// 在light和dark之间切换 只需要取否
import {getStore, saveStore} from "./store";
import consts from "../actions/consts";

export function change_theme(current_theme) {
    switch (current_theme) {
        case "light": {
            console.log(current_theme)
            // 从store中取出url
            let url = getStore("theme-url");
            if (url) {
                let links = document.head.getElementsByTagName("link");
                for (let l of links) {
                    let id = l.getAttribute("id");
                    if (id && id === 'theme-style') {
                        l.setAttribute("href", url);
                    }
                }
            } else {
                // 不存在url时需要直接使用默认值
                url = consts.theme_url;
                let links = document.head.getElementsByTagName("link");
                for (let l of links) {
                    let id = l.getAttribute("id");
                    if (id && id === 'theme-style') {
                        l.setAttribute("href", url);
                    }
                }
            }

            saveStore("theme", 'light');
            break;
        }
        case "dark": {
            console.log(current_theme)
            let links = document.head.getElementsByTagName("link");
            for (let l of links) {
                let id = l.getAttribute("id");
                if (id && id === 'theme-style' && l.getAttribute("href")) {
                    // 在替换前 保存至store
                    saveStore("theme-url", l.getAttribute("href"));
                    l.setAttribute("href", "");
                }
            }
            saveStore("theme", 'dark');
            break;
        }
        default: {
            console.log(current_theme)
            break;
        }
    }
}
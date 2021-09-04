# plume-app

一个服务器可视化webapp

模仿ios应用 [serverCat](https://servercat.app/)编写

## Project setup
```
npm install
```

### Compiles and hot-reloads for development
```
npm run serve
```

### Compiles and minifies for production
```
npm run build
```

## 依赖

本服务尽量使用linux自带的命令获取系统信息

对于一些特殊需求，需要安装额外的程序

对于ubuntu

```bash
apt install sysstat
```

## 启动参数

```bash
# ./plume -h
Usage of ./plume:
  -api string
        set docker api (default "http://127.0.0.1:2375")
  -debug
        enable debug
  -disk string
        set disk
  -eth string
        set eth (default "eth0")
  -host string
        set host (default "127.0.0.1")
  -key string
        set key
  -log string
        set log file path
  -port int
        set port (default 5000)
```

`debug` 调试模式，此模式下会输出路由日志和命令行日志

`host` 设置监听的本地地址

`port` 设置监听的端口

`log` 并非写日志的参数 而是指定一个可读取的日志文件路径

`key` 双方通信密钥 plume和plume-app使用相同密钥才能通信

`eth` 指定要监控的网卡 默认eth0网卡， 使用`ifconfig`查看全部网卡

`disk` 指定要监控的磁盘，使用`lsblk -d`方式查看全部磁盘

`api` 指定docker api接口地址 不建议开放公共api而是直接由plume内部调用

## 示例

![](demo/demo1.png)

![](demo/demo2.png)

![](demo/demo3.png)

**同时支持暗黑模式切换**

![](demo/demo4.png)

![](demo/demo5.png)

![](demo/demo6.png)

![](demo/demo7.png)

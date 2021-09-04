/*
Name: plume
File: plume.go
Author: Landers

Plume app of plume webapp
plume配套的简单后端服务 单文件 基于gin和exec和操作系统交互
*/

package main

import (
	"flag"
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	// 处理命令行
	var c CmdArgs
	debug := flag.Bool("debug", false, "enable debug")
	eth := flag.String("eth", "eth0", "set eth")
	disk := flag.String("disk", "", "set disk")
	port := flag.Int("port", 5000, "set port")
	host := flag.String("host", "127.0.0.1", "set host")
	log := flag.String("log", "", "set log file path")
	key := flag.String("key", "", "set key")
	api := flag.String("api", "http://127.0.0.1:2375", "set docker api")

	flag.Parse()

	c = CmdArgs{
		Debug:     *debug,
		Host:      *host,
		Port:      *port,
		Log:       *log,
		Key:       *key,
		Eth:       *eth,
		Disk:      *disk,
		DockerAPI: *api,
	}
	app := engine(c)
	fmt.Printf("running at %s:%d\n", *host, *port)
	e := app.Run(fmt.Sprintf("%s:%d", *host, *port))
	if e != nil {
		fmt.Printf("plume exit because: %s\n", e.Error())
	}
}

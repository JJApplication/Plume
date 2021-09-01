/*
Name: plume
File: utils_server.go
Author: Landers
*/

package main

import (
	"os/exec"
	"strings"
)

// 系统信息
func getServerInfo() string {
	// lsb_release -a
	// head -n 1 /etc/issue | sed 's/\\n//g' | sed 's/\\l//g'
	sh := "lsb_release -d|sed 's/'$(lsb_release -d|awk '{print $1}')'//g'"
	cmd := exec.Command("bash", "-c", sh)
	res, e := cmd.Output()
	if e != nil {
		return ""
	}else {
		return strings.Trim(string(res), "\n\\n\\l\t\\t")
	}
}

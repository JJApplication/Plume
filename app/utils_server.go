/*
Name: plume
File: utils_server.go
Author: Landers
*/

package main

import (
	"strings"
)

// 系统信息
func getServerInfo(debug bool) string {
	// lsb_release -a
	// head -n 1 /etc/issue | sed 's/\\n//g' | sed 's/\\l//g'
	sh := "lsb_release -d|sed 's/'$(lsb_release -d|awk '{print $1}')'//g'"
	res, e := cmdRun(sh, debug)
	if e != nil {
		return ""
	} else {
		return strings.Trim(string(res), "\n\\n\\l\t\\t")
	}
}

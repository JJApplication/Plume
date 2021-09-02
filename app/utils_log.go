/*
Name: plume
File: utils_log.go
Author: Landers
*/

package main

import (
	"fmt"
)

// 只会根据log获取log路径下的日志 不会自动生成日志
func getLogData(log string, debug bool) string {
	if log == "" {
		return ""
	}
	sh := fmt.Sprintf("tail -n 100 %s", log)
	res, e := cmdRun(sh, debug)
	if e != nil {
		return ""
	}

	return string(res)
}

func delLog(log string, debug bool) string {
	if log == "" {
		return "ok"
	}
	sh := fmt.Sprintf(":> %s", log)
	_, e := cmdRun(sh, debug)
	if e != nil {
		return "fail"
	}

	return "ok"
}
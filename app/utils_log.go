/*
Name: plume
File: utils_log.go
Author: Landers
*/

package main

import (
	"fmt"
	"os/exec"
)

// 只会根据log获取log路径下的日志 不会自动生成日志
func getLogData(log string) string {
	if log == "" {
		return ""
	}
	sh := fmt.Sprintf("tail -n 100 %s", log)
	cmd := exec.Command("bash", "-c", sh)
	res, e := cmd.Output()
	if e != nil {
		return ""
	}

	return string(res)
}

func delLog(log string) string {
	if log == "" {
		return "ok"
	}
	sh := fmt.Sprintf(":> %s", log)
	cmd := exec.Command("bash", "-c", sh)
	_, e := cmd.Output()
	if e != nil {
		return "fail"
	}

	return "ok"
}
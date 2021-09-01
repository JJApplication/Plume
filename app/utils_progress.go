/*
Name: plume
File: utils_progress.go
Author: Landers
*/

package main

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

// 进程函数
func getProgressData() ProgressInfo {
	var pro ProgressInfo

	sh := "ps ax | wc -l"
	cmd := exec.Command("bash", "-c", sh)
	res, e := cmd.Output()
	if e != nil {
		pro.ProgressAll = "0"
	} else {
		pro.ProgressAll = strings.Trim(string(res), "\n")
	}

	sh = "ps ax | awk '{print $3}' | grep R | wc -l"
	cmd = exec.Command("bash", "-c", sh)
	res, e = cmd.Output()
	if e != nil {
		pro.ProgressRun = "0"
	} else {
		pro.ProgressRun = strings.Trim(string(res), "\n")
	}

	sh = "ps ax | awk '{print $3}' | grep Z | wc -l"
	cmd = exec.Command("bash", "-c", sh)
	res, e = cmd.Output()
	if e != nil {
		pro.ProgressDead = "0"
	} else {
		pro.ProgressDead = strings.Trim(string(res), "\n")
	}

	sh = "ps ax | awk '{print $3}' | grep S | wc -l"
	cmd = exec.Command("bash", "-c", sh)
	res, e = cmd.Output()
	if e != nil {
		pro.ProgressSleep = "0"
	} else {
		pro.ProgressSleep = strings.Trim(string(res), "\n")
	}

	return pro
}

// 格式化进程信息
func fmtData(s string) (ProgressListInfo, error) {
	data := strings.Fields(s)
	if len(data) < 4 {
		return ProgressListInfo{}, errors.New("bad progress")
	}
	return ProgressListInfo{
		PID: data[0],
		CPU: data[1],
		Mem: data[2],
		Cmd: data[3],
	}, nil
}

func getProgressListData(num string) []ProgressListInfo {
	var pro []ProgressListInfo

	if num == "" {
		num = "10"
	}

	sh := fmt.Sprintf("ps aux | grep -v PID | awk '{print $2, $3, $4, $11}' | sort -rn -k +3 | head -n %s | tr '\n' ','", num)
	cmd := exec.Command("bash", "-c", sh)
	res, e := cmd.Output()
	if e != nil {
		return []ProgressListInfo{}
	}

	list := strings.Split(strings.Trim(string(res), "\n"), ",")
	for _, l := range list {

		if p, e := fmtData(l); e == nil {
			pro = append(pro, p)
		}
	}

	return pro
}


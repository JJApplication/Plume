/*
Name: plume
File: utils_mem.go
Author: Landers
*/

package main

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

// 内存信息
func getMemInfo() MemInfo {
	var m MemInfo
	sh := "free -m | grep Mem | sed 's/Mem://g' | awk '{print $1, $2, $5, $6}'"
	cmd := exec.Command("bash", "-c", sh)
	res, e := cmd.Output()
	mem_info := strings.Fields(string(res))

	if e != nil || len(mem_info) < 4 {
		return MemInfo{}
	}
	// 以mb为单位
	total, _ := strconv.Atoi(mem_info[1])
	used, _ := strconv.Atoi(mem_info[0])

	m.MemUsage = fmt.Sprintf("%.2f", float64(total) / float64(used) * 100)
	m.MemUsed = mem_info[1] + "M"
	m.MemCache = mem_info[2] + "M"
	m.MemFree = mem_info[3] + "M"
	return m
}

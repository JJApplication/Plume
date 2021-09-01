/*
Name: plume
File: utils_cpu.go
Author: Landers
*/

package main

import (
	"os/exec"
	"strconv"
	"strings"
)

// cpu函数
// 转0函数
func equalZero(s string) string {
	if s == "0" || s == "0.0" {
		return "0"
	}
	return s
}

func getCPUInfo() CPUInfo {
	var cpu CPUInfo
	// 根据top获取占用情况
	sh := "top -bn 1 -i -c | grep %Cpu | sed 's/%Cpu(s)://g'"
	cmd := exec.Command("bash", "-c", sh)
	res, e := cmd.Output()
	cpu_info := strings.Split(string(res), ",")

	if e != nil || len(cpu_info) < 5 {
		cpu.CpuUsage , cpu.CpuUsageSys , cpu.CpuUsageUser = "0", "0", "0"
		cpu.CpuFree, cpu.CpuIOWait = "0", "0"
	} else {
		cpu_us := equalZero(strings.Fields(cpu_info[0])[0])
		cpu_sy := equalZero(strings.Fields(cpu_info[1])[0])
		cpu_free := equalZero(strings.Fields(cpu_info[3])[0])
		t, _ := strconv.Atoi(strings.Split(cpu_free, ".")[0])
		cpu_usage := strconv.Itoa(100 - t)
		cpu_io_wait := equalZero(strings.Fields(cpu_info[4])[0])

		cpu.CpuUsage = cpu_usage
		cpu.CpuUsageSys = cpu_sy
		cpu.CpuUsageUser = cpu_us
		cpu.CpuFree = cpu_free
		cpu.CpuIOWait = cpu_io_wait
	}

	// load
	sh = "cat /proc/loadavg | awk '{print $1,$2,$3}'"
	cmd = exec.Command("bash", "-c", sh)
	res, e = cmd.Output()
	if e != nil {
		cpu.CpuLoad = "0 0 0"
	}else {
		cpu.CpuLoad = strings.Trim(string(res), "\n")
	}

	// cpu count
	sh = "cat /proc/cpuinfo | grep \"cpu cores\" | uniq | awk '{print $4}'"
	cmd = exec.Command("bash", "-c", sh)
	res, e = cmd.Output()
	if e != nil {
		cpu.CpuCount = "0"
	}else {
		cpu.CpuCount = strings.Trim(string(res), "\n")
	}

	// cpu run
	sh = "uptime | awk '{print $3}'"
	cmd = exec.Command("bash", "-c", sh)
	res, e = cmd.Output()
	if e != nil {
		cpu.CpuRun = "0"
	}else {
		cpu.CpuRun = strings.Trim(string(res), "\n")
	}

	return cpu
}

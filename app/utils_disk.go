/*
Name: plume
File: utils_disk.go
Author: Landers
*/

package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)

// 获取磁盘信息
func getDiskInfo(d string) DiskInfo {
	var disk DiskInfo
	// Size  Used Use%
	sh := "df -hl / | awk 'NR==2{print $1, $2, $3, $5}'"
	cmd := exec.Command("bash", "-c", sh)
	res, e := cmd.Output()
	disk_info := strings.Fields(string(res))

	if e != nil || len(disk_info) < 4 {
		disk.DiskMount = "/"
		disk.DiskUsage = "0"
		disk.DiskUsed = "0"
		disk.DiskAll = "0"
	} else {
		// 容量取整
		disk.DiskMount = disk_info[0]
		disk.DiskAll = strings.Trim(disk_info[1], "\n")
		disk.DiskUsed = strings.Trim(disk_info[2], "\n")
		disk.DiskUsage = strings.Trim(disk_info[3], "%")
	}

	// 使用默认磁盘
	if d != "" {
		sh = fmt.Sprintf("iostat -d %s -x -o JSON", d)
	} else {
		sh = "iostat -d -x -o JSON"
	}

	type Disk struct {
		Read string `json:"kB_read"`
		Write string `json:"kB_wrtn"`
	}
	type Statistics struct {
		Disk []map[string]interface{} `json:"disk"`
	}
	type Hosts struct {
		Statistics []Statistics `json:"statistics"`
	}
	type SysStat struct {
		Hosts []Hosts `json:"hosts"`
	}
	type stat struct {
		SysStat SysStat `json:"sysstat"`
	}

	// 读速率 读延迟 写速率k/s 写延迟
	var td stat
	cmd = exec.Command("bash", "-c", sh)
	res, e = cmd.Output()

	e = json.Unmarshal(res, &td)

	readDelay := td.SysStat.Hosts[0].Statistics[0].Disk[0]["r_await"]
	writeDelay := td.SysStat.Hosts[0].Statistics[0].Disk[0]["w_await"]
	readRate := td.SysStat.Hosts[0].Statistics[0].Disk[0]["rkB/s"]
	writeRate := td.SysStat.Hosts[0].Statistics[0].Disk[0]["wkB/s"]

	if e != nil {
		disk.DiskReadRate = "0 KB/s"
		disk.DiskWriteRate = "0 KB/s"
	}else {
		disk.DiskReadRate = calcRate(readRate)
		disk.DiskReadDelay = fmt.Sprintf("%.2f", readDelay.(float64))
		disk.DiskWriteRate = calcRate(writeRate)
		disk.DiskWriteDelay = fmt.Sprintf("%.2f", writeDelay.(float64))
	}

	// 读写量kb 因为版本问题不一致 不能使用列判断
	if d != "" {
		sh = fmt.Sprintf("iostat -d %s -o JSON", d)
	}else {
		sh = "iostat -d -o JSON"
	}
	cmd = exec.Command("bash", "-c", sh)
	res, e = cmd.Output()

	var t stat
	e = json.Unmarshal(res, &t)

	readByte := t.SysStat.Hosts[0].Statistics[0].Disk[0]["kB_read"]
	writeByte := t.SysStat.Hosts[0].Statistics[0].Disk[0]["kB_wrtn"]

	if e != nil {
		disk.DiskReadByte = "0"
		disk.DiskWriteByte = "0"
	}else {
		disk.DiskReadByte = calcByte(readByte)
		disk.DiskWriteByte = calcByte(writeByte)
	}

	return disk
}

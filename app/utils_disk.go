/*
Name: plume
File: utils_disk.go
Author: Landers
*/

package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

// 获取磁盘信息
func getDiskInfo(d string, debug bool) DiskInfo {
	var disk DiskInfo
	// Size  Used Use%
	sh := "df -hl / | awk 'NR==2{print $1, $2, $3, $5}'"
	res, e := cmdRun(sh, debug)
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
	res, e = cmdRun(sh, debug)

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

	res, e = cmdRun(sh, debug)

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

func getDiskInfoDetail(debug bool) DiskInfoDetail {
	var disk DiskInfoDetail

	sh := "fdisk -l|grep Disklabel"
	res, e := cmdRun(sh, debug)
	if e != nil {
		disk.Label = "unknown"
	} else {
		disk.Label = strings.Split(string(res), ":")[1]
	}

	sh = "fdisk -l|grep 'model'"
	res, e = cmdRun(sh, debug)
	if e != nil {
		disk.Model = "unknown"
	} else {
		disk.Model = strings.Split(string(res), ":")[1]
	}

	sh = "df -h|grep -v Filesystem|tr '\n' ','"
	res, e = cmdRun(sh, debug)
	if e != nil {
		disk.List = []map[string]string{}
	} else {
		var d []map[string]string
		data := strings.Split(string(res), ",")
		for _, di := range data {
			dis := strings.Fields(di)
			if len(dis) >= 6 {
				d = append(d, map[string]string{
					"system": dis[0],
					"size": dis[1],
					"used": dis[2],
					"avail": dis[3],
					"use": dis[4],
					"mount": dis[5],
				})
			}
		}
		disk.List = d
	}
	return disk
}
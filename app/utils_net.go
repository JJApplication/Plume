/*
Name: plume
File: utils_net.go
Author: Landers
*/

package main

import (
	"fmt"
	"math"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

// 计算读写速率转换kb mb
// 默认kb/s
func calcRate(s interface{}) string {
	f := s.(float64)
	f2 := f / 1024
	if f2 >= 1 {
		return fmt.Sprintf("%.2f MB/s", f2)
	}
	return fmt.Sprintf("%.2f KB/s", s)
}

// 计算读写量单位换算
func calcByte(s interface{}) string {
	data := int(math.Ceil(s.(float64)))

	// 转换为gb
	gb := data / (1024 * 1024)
	mb := data / 1024
	if gb >= 1 {
		return fmt.Sprintf("%d GB", gb)
	}else if mb >= 1 {
		return fmt.Sprintf("%d MB", mb)
	}else {
		return fmt.Sprintf("%s KB", s)
	}
}

// 计算流量
func calcNet(s string) string {
	s = strings.Trim(s, "\n")
	data, e := strconv.Atoi(s)
	if e != nil {
		return "0B"
	}
	gb := data / (1024*1024*1024)
	mb := data / (1024*1024)
	kb := data / 1024

	if gb >= 1 {
		return fmt.Sprintf("%dG", gb)
	} else if mb >= 1 {
		return fmt.Sprintf("%dM", mb)
	} else if kb >= 1 {
		return fmt.Sprintf("%dK", kb)
	} else {
		return fmt.Sprintf("%dB", data)
	}
}

// 时间差200ms
func calcNetRate(old, new string, offset int) string {
	old = strings.Trim(old, "\n")
	new = strings.Trim(new, "\n")
	oldData, e := strconv.Atoi(old)

	if e != nil {
		return  "0b"
	}

	newData, e := strconv.Atoi(new)
	if e != nil {
		return  "0b"
	}

	data := (newData - oldData) * 1000 / offset
	gb := data / (1024*1024*1024)
	mb := data / (1024*1024)
	kb := data / 1024

	if gb >= 1 {
		return fmt.Sprintf("%dgb", gb)
	} else if mb >= 1 {
		return fmt.Sprintf("%dmb", mb)
	} else if kb >= 1 {
		return fmt.Sprintf("%dkb", kb)
	} else {
		return fmt.Sprintf("%db", data)
	}
}

// 计算重传差值
func calcRetrans(send1, send2 string, r1, r2 string, offset int) string {
	send1Int, _ := strconv.Atoi(strings.Trim(send1, "\n"))
	send2Int, _ := strconv.Atoi(strings.Trim(send2, "\n"))

	r1Int, _ := strconv.Atoi(strings.Trim(r1, "\n"))
	r2Int, _ := strconv.Atoi(strings.Trim(r2, "\n"))

	send := send2Int - send1Int
	r := r2Int - r1Int
	if send == 0 || r == 0 {
		return "0"
	}
	return fmt.Sprintf("%d", r / send * 1000 / offset)
}

// 获取网络信息
func getNetInfo(eth string) NetInfo {
	var net NetInfo
	// 获取ip地址
	sh := fmt.Sprintf("ifconfig %s | grep inet | awk '{print $2}' | tr '\n' ' '", eth)
	cmd := exec.Command("bash", "-c", sh)
	res, e := cmd.Output()
	net_info := strings.Fields(string(res))

	if e != nil || len(net_info) < 1 {
		net.IPV4 = "127.0.0.1"
		net.IPV6 = "::1"
	} else {
		if len(net_info) == 1 {
			net.IPV4 = net_info[0]
			net.IPV6 = "::1"
		} else {
			net.IPV4 = net_info[0]
			net.IPV6 = net_info[1]
		}
	}

	// 获取网络流量
	sh = fmt.Sprintf("ifconfig %s | grep \"RX packets\" | awk '{print $5}'", eth)
	cmd = exec.Command("bash", "-c", sh)
	rx, e := cmd.Output()
	if e != nil {
		net.NetWorkDownload = "0"
	} else {
		net.NetWorkDownload = calcNet(string(rx))
	}

	sh = fmt.Sprintf("ifconfig %s | grep \"TX packets\" | awk '{print $5}'", eth)
	cmd = exec.Command("bash", "-c", sh)
	tx, e := cmd.Output()
	if e != nil {
		net.NetWorkUpload = "0"
	} else {
		net.NetWorkUpload = calcNet(string(tx))
	}

	// 获取网络速率 (200ms计算)
	time.Sleep(time.Millisecond * 200)
	sh = fmt.Sprintf("ifconfig %s | grep \"RX packets\" | awk '{print $5}'", eth)
	cmd = exec.Command("bash", "-c", sh)
	res, e = cmd.Output()
	if e != nil {
		net.NetDownload = "0"
	} else {
		net.NetDownload = calcNetRate(string(rx), string(res), 200)
	}

	sh = fmt.Sprintf("ifconfig %s | grep \"TX packets\" | awk '{print $5}'", eth)
	cmd = exec.Command("bash", "-c", sh)
	res, e = cmd.Output()
	if e != nil {
		net.NetUpload = "0"
	} else {
		net.NetUpload = calcNetRate(string(tx), string(res), 200)
	}

	// 获取建连信息
	sh = "netstat -na | grep ESTABLISHED | wc -l"
	cmd = exec.Command("bash", "-c", sh)
	res, e = cmd.Output()
	if e != nil {
		net.NetActive = "0"
	} else {
		net.NetActive = strings.Trim(string(res), "\n")
	}

	sh = "netstat -na | grep LISTEN | wc -l"
	cmd = exec.Command("bash", "-c", sh)
	res, e = cmd.Output()
	if e != nil {
		net.NetPassive = "0"
	} else {
		net.NetPassive = strings.Trim(string(res), "\n")
	}

	sh = "netstat -na | grep TIME_WAIT | wc -l"
	cmd = exec.Command("bash", "-c", sh)
	res, e = cmd.Output()
	if e != nil {
		net.NetFail = "0"
	} else {
		net.NetFail = strings.Trim(string(res), "\n")
	}

	// 重传计算
	sh = "netstat -s -t | grep \"segments sent out\" | awk '{print $1}'"
	cmd = exec.Command("bash", "-c", sh)
	out1, _ := cmd.Output()

	sh = "netstat -s -t | grep \"segments retransmitted\" | awk '{print $1}'"
	cmd = exec.Command("bash", "-c", sh)
	re1, _ := cmd.Output()

	var out2, re2 string
	c := make(chan string, 2)
	go func() {
		time.Sleep(time.Second * 1)
		sh := "netstat -s -t | grep \"segments sent out\" | awk '{print $1}'"
		cmd := exec.Command("bash", "-c", sh)
		o, _ := cmd.Output()

		sh = "netstat -s -t | grep \"segments retransmitted\" | awk '{print $1}'"
		cmd = exec.Command("bash", "-c", sh)
		r, _ := cmd.Output()

		c<-string(o)
		c<-string(r)

	}()
	out2 = <-c
	re2 = <-c

	net.NetRetry = calcRetrans(string(out1), string(out2), string(re1), string(re2), 1000)


	return net
}

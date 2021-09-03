/*
Name: plume
File: utils_net.go
Author: Landers
*/

package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

var cache cacheNet

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
func getNetInfo(eth string, debug bool) NetInfo {
	var net NetInfo
	// 获取ip地址
	sh := fmt.Sprintf("ifconfig %s | grep inet | awk '{print $2}' | tr '\n' ' '", eth)
	res, e := cmdRun(sh, debug)
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
	sh = fmt.Sprintf("ifconfig %s | grep -E \"RX packets|TX packets\" | awk '{print $5}' | tr '\n' ','", eth)
	res, e = cmdRun(sh, debug)
	if e != nil {
		net.NetWorkDownload = "0"
		net.NetWorkUpload = "0"
	} else {
		d := strings.Split(string(res), ",")
		if len(d) < 2 {
			net.NetWorkDownload = "0"
			net.NetWorkUpload = "0"
		} else {
			net.NetWorkDownload = calcNet(d[0])
			net.NetWorkUpload = calcNet(d[1])
		}
	}
	// 获取网络速率 (200ms计算)
	net.NetDownload, net.NetUpload = calcNetworkRate(eth, debug)

	// 获取建连信息  为提升效率使用ss
	//sh = "netstat -na | grep ESTABLISHED | wc -l"
	sh = "ss -a | grep ESTAB | wc -l;ss -a | grep LISTEN | wc -l;ss -a | grep UNCONN | wc -l"
	res, e = cmdRun(sh, debug)
	if e != nil {
		net.NetActive = "0"
		net.NetPassive = "0"
		net.NetFail = "0"
	} else {
		d := strings.Split(string(res), "\n")
		net.NetActive = strings.Trim(d[0], "\n ")
		net.NetPassive = strings.Trim(d[1], "\n ")
		net.NetFail = strings.Trim(d[2], "\n ")
	}

	// 重传计算
	//sh = "netstat -s -t | grep \"segments sent out\" | awk '{print $1}'"
	sh = "netstat -s -t | grep \"segments retransmitted\" | awk '{print $1}'"
	res, _ = cmdRun(sh, debug)
	net.NetRetry = strings.Trim(string(res), "\n")

	return net
}

// 轮询计算网速
// 规则 20ms 循环10次 取平均值
// 使用goroutine同时计算上传和下载
func calcNetworkRateByLoop(eth string, debug bool) (string, string) {
	shDown := fmt.Sprintf("ifconfig %s | grep \"RX packets\" | awk '{print $5}'", eth)
	shUp := fmt.Sprintf("ifconfig %s | grep \"TX packets\" | awk '{print $5}'", eth)

	chU := make(chan int)
	chD := make(chan int)

	go func() {
		var total int
		for i:=0; i<10; i++ {
			res, _ := cmdRun(shDown, debug)
			byte1, _ := strconv.Atoi(strings.Trim(string(res), "\n\\n"))
			time.Sleep(time.Millisecond * 20)
			res, _ = cmdRun(shDown, debug)
			byte2, _ := strconv.Atoi(strings.Trim(string(res), "\n\\n"))

			byteDownLoad := byte2 - byte1
			total += byteDownLoad
		}

		chD<-total
	}()

	go func() {
		var total int
		for i:=0; i<10; i++ {
			res, _ := cmdRun(shUp, debug)
			byte1, _ := strconv.Atoi(strings.Trim(string(res), "\n\\n"))
			time.Sleep(time.Millisecond * 20)
			res, _ = cmdRun(shUp, debug)
			byte2, _ := strconv.Atoi(strings.Trim(string(res), "\n\\n"))

			byteUpload := byte2 - byte1
			total += byteUpload
		}

		chU<-total
	}()

	download := <-chD
	upload := <-chU


	return calcByteAll(download/10*1000/20), calcByteAll(upload/10*1000/20)
}

// 根据缓存时间计算网速
func calcNetworkRate(eth string, debug bool) (string, string) {
	var empty cacheNet

	sh := fmt.Sprintf("ifconfig %s | grep -E \"RX packets|TX packets\" | awk '{print $5}' | tr '\n' ','", eth)
	res, e := cmdRun(sh, debug)
	if e != nil {
		return "0", "0"
	}

	data := strings.Split(string(res), ",")
	if len(data) < 2 {
		return "0", "0"
	}
	download, _ := strconv.Atoi(strings.Trim(data[0], "\n\\n"))
	upload, _ := strconv.Atoi(strings.Trim(data[1], "\n\\n"))
	timeNow := time.Now().Unix()
	offset := int(timeNow - cache.UpdateTime)

	if offset <= 0 {
		offset = 1
	}
	// 首次刷新缓存
	if cache == empty {
		cache.UpdateTime = timeNow
		cache.Download = download
		cache.Upload = upload
		return "0", "0"
	}

	rateDownload := calcByteAll((download-cache.Download)/offset)
	rateUpload := calcByteAll((upload-cache.Upload)/offset)

	// cache it
	cache.Download = download
	cache.Upload = upload
	cache.UpdateTime = timeNow

	return rateDownload, rateUpload
}

func calcByteAll(data int) string {
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

// 获取网络详情
func getNetInfoDetail(num string, debug bool) []NetInfoDetail {
	// Netid State Recv-Q Send-Q                   Local Address:Port     Peer Address:Port
	var net []NetInfoDetail
	sh := fmt.Sprintf("ss -txu  | grep -v Netid |sort -rk 3| head -n %s|awk '{print $1, $2, $3, $4, $5}'|tr '\n' ','", num)
	res, e := cmdRun(sh, debug)
	if e != nil {
		return []NetInfoDetail{}
	}

	d := strings.Split(string(res), ",")
	for _, n := range d {
		ni := strings.Fields(n)
		if len(ni) >= 5 {
			net = append(net, NetInfoDetail{
				ID: ni[0],
				State: ni[1],
				R: ni[2],
				S: ni[3],
				Address: ni[4],
			})
		}
	}

	return net
}
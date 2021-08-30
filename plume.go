/*
Name: plume
File: plume.go
Author: Landers

Plume app of plume webapp
plume配套的简单后端服务 单文件 基于gin和exec和操作系统交互
*/

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

// 需要外部参数网卡 默认为eth0

// CPUInfo 结构体定义
type CPUInfo struct {
	CpuUsage string `json:"cpu_usage"`
	CpuUsageSys string `json:"cpu_usage_system"`
	CpuUsageUser string `json:"cpu_usage_user"`
	CpuIOWait string `json:"cpu_io_wait"`
	CpuCount string `json:"cpu_count"`
	CpuFree string `json:"cpu_free"`
	CpuLoad string `json:"cpu_load"`
	CpuRun string `json:"cpu_run"`
}

type MemInfo struct {
	MemUsage string `json:"mem_usage"`
	MemUsed string `json:"mem_used"`
	MemFree string `json:"mem_free"`
	MemCache string `json:"mem_cache"`
}

type NetInfo struct {
	NetUpload string `json:"net_upload"`
	NetDownload string `json:"net_download"`
	NetWorkUpload string `json:"network_upload"`
	NetWorkDownload string `json:"network_download"`
	NetRetry string `json:"net_retry"`
	NetActive string `json:"net_active"`
	NetPassive string `json:"net_passive"`
	NetFail string `json:"net_fail"`
	IPV4 string `json:"ipv4"`
	IPV6 string `json:"ipv6"`

}

// DiskInfo 磁盘信息使用iostat
type DiskInfo struct {
	DiskMount string `json:"disk_mount"`
	DiskUsed string `json:"disk_used"`
	DiskAll string `json:"disk_all"`
	DiskUsage string `json:"disk_usage"`
	DiskReadRate string `json:"disk_read_rate"`
	DiskReadByte string `json:"disk_read_byte"`
	DiskReadDelay string `json:"disk_read_delay"`
	DiskWriteRate string `json:"disk_write_rate"`
	DiskWriteByte string `json:"disk_write_byte"`
	DiskWriteDelay string `json:"disk_write_delay"`
}

// Containers 容器列表
type Containers struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Image string `json:"image"`
	Date string `json:"date"`
	Status string `json:"status"`
}

// 中间件
func middlewareCors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, DELETE, PUT, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}	
}

// 服务app
func engine(eth, disk string) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	s := gin.New()
	s.Use(gin.Recovery())
	s.Use(middlewareCors())

	wrapApi(s, eth, disk)
	
	return s
}

// 路由规则
func wrapApi(s *gin.Engine, eth, disk string) {
	s.POST("/api/cpu", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": getCPUInfo()})
	})

	s.POST("/api/server", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": getServerInfo()})
	})

	s.POST("/api/mem", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": getMemInfo()})
	})

	s.POST("/api/net", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": getNetInfo(eth)})
	})

	s.POST("/api/disk", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": getDiskInfo(disk)})
	})

	s.POST("/api/container", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": getContainers()})
	})
}

// 功能函数
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

// 系统信息
func getServerInfo() string {
	sh := "head -n 1 /etc/issue | sed 's/\\\\n//g' | sed 's/\\\\l//g'"
	cmd := exec.Command("bash", "-c", sh)
	res, e := cmd.Output()
	if e != nil {
		return ""
	}else {
		return strings.Trim(string(res), "\\n\\l")
	}
}

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

// 容器接口代理请求
func getContainers() []Containers {
	var containers []Containers

	return containers
}

func main() {
	// 处理命令行
	eth := flag.String("eth", "eth0", "set eth")
	disk := flag.String("disk", "", "set disk")
	port := flag.Int("port", 5000, "set port")
	host := flag.String("host", "127.0.0.1", "set host")

	flag.Parse()
	app := engine(*eth, *disk)
	fmt.Printf("running at %s:%d\n", *host, *port)
	e := app.Run(fmt.Sprintf("%s:%d", *host, *port))
	if e != nil {
		fmt.Printf("plume exit because: %s\n", e.Error())
	}
}

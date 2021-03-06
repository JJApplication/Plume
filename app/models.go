/*
Name: plume
File: models.go
Author: Landers
*/

package main

// 对象结构体定义

type CmdArgs struct {
	Debug     bool
	Host      string
	Port      int
	Log       string
	Key       string
	Eth       string
	Disk      string
	DockerAPI string
}

type Combination struct {
	ServerInfo   string       `json:"server_info"`
	CPUInfo      CPUInfo      `json:"cpu_info"`
	MemInfo      MemInfo      `json:"mem_info"`
	KernelInfo   KernelInfo   `json:"kernel_info"`
	NetInfo      NetInfo      `json:"net_info"`
	DiskInfo     DiskInfo     `json:"disk_info"`
	ProgressInfo ProgressInfo `json:"progress_info"`
}

// CPUInfo 结构体定义
type CPUInfo struct {
	CpuUsage     string `json:"cpu_usage"`
	CpuUsageSys  string `json:"cpu_usage_system"`
	CpuUsageUser string `json:"cpu_usage_user"`
	CpuIOWait    string `json:"cpu_io_wait"`
	CpuCount     string `json:"cpu_count"`
	CpuPhysical  string `json:"cpu_physical"`
	CpuFree      string `json:"cpu_free"`
	CpuLoad      string `json:"cpu_load"`
	CpuRun       string `json:"cpu_run"`
}

type CPUInfoDetail struct {
	Info  string `json:"info"`
	Freq  string `json:"freq"`
	Cache string `json:"cache"`
}

type MemInfo struct {
	MemUsage string `json:"mem_usage"`
	MemUsed  string `json:"mem_used"`
	MemFree  string `json:"mem_free"`
	MemCache string `json:"mem_cache"`
}

type MemInfoDetail struct {
	ManuFacturer string `json:"manufacturer"`
	Product      string `json:"product"`
	Size         string `json:"size"`
	Speed        string `json:"speed"`
	Width        string `json:"width"`
}

type KernelInfo struct {
	KernelOs      string `json:"kernel_os"`
	KernelType    string `json:"kernel_type"`
	KernelVersion string `json:"kernel_version"`
}

type ProgressInfo struct {
	ProgressAll   string `json:"progress_all"`
	ProgressRun   string `json:"progress_run"`
	ProgressDead  string `json:"progress_dead"`
	ProgressSleep string `json:"progress_sleep"`
}

type ProgressListInfo struct {
	PID string `json:"pid,omitempty"`
	CPU string `json:"cpu"`
	Mem string `json:"mem"`
	Cmd string `json:"cmd"`
}

type NetInfo struct {
	NetUpload       string `json:"net_upload"`
	NetDownload     string `json:"net_download"`
	NetWorkUpload   string `json:"network_upload"`
	NetWorkDownload string `json:"network_download"`
	NetRetry        string `json:"net_retry"`
	NetActive       string `json:"net_active"`
	NetPassive      string `json:"net_passive"`
	NetFail         string `json:"net_fail"`
	IPV4            string `json:"ipv4"`
	IPV6            string `json:"ipv6"`
}

type NetInfoDetail struct {
	ID      string `json:"id"`
	State   string `json:"state"`
	R       string `json:"r"`
	S       string `json:"s"`
	Address string `json:"address"`
}

// DiskInfo 磁盘信息使用iostat
type DiskInfo struct {
	DiskMount      string `json:"disk_mount"`
	DiskUsed       string `json:"disk_used"`
	DiskAll        string `json:"disk_all"`
	DiskUsage      string `json:"disk_usage"`
	DiskReadRate   string `json:"disk_read_rate"`
	DiskReadByte   string `json:"disk_read_byte"`
	DiskReadDelay  string `json:"disk_read_delay"`
	DiskWriteRate  string `json:"disk_write_rate"`
	DiskWriteByte  string `json:"disk_write_byte"`
	DiskWriteDelay string `json:"disk_write_delay"`
}

type DiskInfoDetail struct {
	Label string              `json:"label"`
	Model string              `json:"model"`
	List  []map[string]string `json:"list"`
}

// Containers 容器列表

type ContainersRaw struct {
	ID      string
	Image   string
	Created int64
	Names   interface{}
	State   string
}

type Containers struct {
	ID     string      `json:"id"`
	Name   interface{} `json:"name"`
	Image  string      `json:"image"`
	Date   string      `json:"date"`
	Status string      `json:"status"`
}

type ContainerRaw struct {
	Name       string
	Config     map[string]interface{}
	Created    string
	HostConfig map[string]interface{}
	State      map[string]interface{}
}

type Container struct {
	Name   string      `json:"name"`
	Cmd    interface{} `json:"cmd"`
	Image  interface{} `json:"image"`
	User   interface{} `json:"user"`
	Volume interface{} `json:"volume"`
	Port   interface{} `json:"port"`
	WrkDir interface{} `json:"wrkdir"`
	Date   string      `json:"date"`
	Status interface{} `json:"status"`
	Cpu    interface{} `json:"cpu"`
	Mem    interface{} `json:"mem"`
}

type ImagesRaw struct {
	ID         string
	RepoTags   []string
	Created    int64
	Size       int
	Containers int
}

type Images struct {
	ID         string   `json:"id"`
	Tags       []string `json:"tags"`
	Date       string   `json:"date"`
	Size       string   `json:"size"`
	Containers int      `json:"containers"`
}

// 缓存
type cacheNet struct {
	Download   int
	Upload     int
	UpdateTime int64
}

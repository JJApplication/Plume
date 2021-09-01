/*
Name: plume
File: models.go
Author: Landers
*/

package main

// 对象结构体定义

type Combination struct {
	ServerInfo string `json:"server_info"`
	CPUInfo CPUInfo `json:"cpu_info"`
	MemInfo MemInfo `json:"mem_info"`
	KernelInfo KernelInfo `json:"kernel_info"`
	NetInfo NetInfo `json:"net_info"`
	DiskInfo DiskInfo `json:"disk_info"`
	ProgressInfo ProgressInfo `json:"progress_info"`
}

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

type KernelInfo struct {
	KernelOs string `json:"kernel_os"`
	KernelType string `json:"kernel_type"`
	KernelVersion string `json:"kernel_version"`
}

type ProgressInfo struct {
	ProgressAll string `json:"progress_all"`
	ProgressRun string `json:"progress_run"`
	ProgressDead string `json:"progress_dead"`
	ProgressSleep string `json:"progress_sleep"`
}

type ProgressListInfo struct {
	PID string `json:"pid,omitempty"`
	CPU string `json:"cpu"`
	Mem string `json:"mem"`
	Cmd string `json:"cmd"`
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


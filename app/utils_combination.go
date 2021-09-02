/*
Name: plume
File: utils_combination.go
Author: Landers
*/

package main

// 整合接口使用协程处理
func getCombinationData(eth, disk string, debug bool) Combination {
	var comb Combination
	// 设计通道 接受所有接口返回值
	chServer := make(chan string)
	chCPU := make(chan CPUInfo)
	chMem := make(chan MemInfo)
	chKernel := make(chan KernelInfo)
	chPro := make(chan ProgressInfo)
	chNet := make(chan NetInfo)
	chDisk := make(chan DiskInfo)

	go func() { chServer<-getServerInfo(debug) }()
	go func() { chCPU<-getCPUInfo(debug) }()
	go func() { chMem<-getMemInfo(debug) }()
	go func() { chKernel<-getKernelData(debug) }()
	go func() { chNet<-getNetInfo(eth, debug) }()
	go func() { chPro<-getProgressData(debug) }()
	go func() { chDisk<-getDiskInfo(disk, debug) }()

	comb.ServerInfo = <-chServer
	comb.CPUInfo = <-chCPU
	comb.MemInfo = <-chMem
	comb.KernelInfo = <-chKernel
	comb.ProgressInfo = <-chPro
	comb.NetInfo = <-chNet
	comb.DiskInfo = <-chDisk

	return comb
}

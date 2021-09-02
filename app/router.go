/*
Name: plume
File: router.go
Author: Landers
*/

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 服务app
func engine(eth, disk, log, key string, debug bool) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	s := gin.New()
	s.Use(gin.Recovery())
	s.Use(middlewareCors())
	s.Use(middlewareAuthKey(key))
	if log != "" {
		s.Use(gin.Logger())
	}

	wrapApi(s, eth, disk, log, debug)

	return s
}

// 路由规则
func wrapApi(s *gin.Engine, eth, disk, log string, debug bool) {
	// 整合接口
	s.POST("/api/comb", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": getCombinationData(eth, disk, debug)})
	})

	s.POST("/api/cpu", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": getCPUInfo(debug)})
	})

	s.POST("/api/cpu_info", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": getCpuInfoDetail(debug)})
	})

	s.POST("/api/server", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": getServerInfo(debug)})
	})

	s.POST("/api/mem", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": getMemInfo(debug)})
	})

	s.POST("/api/mem_info", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": getMemInfoDetail(debug)})
	})

	s.POST("/api/kernel", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": getKernelData(debug)})
	})

	s.POST("/api/progress", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": getProgressData(debug)})
	})

	s.POST("/api/progress_list", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": getProgressListData(c.Query("num"), debug)})
	})

	s.POST("/api/net", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": getNetInfo(eth, debug)})
	})

	s.POST("/api/net_info", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": getNetInfoDetail(c.Query("num"), debug)})
	})

	s.POST("/api/disk", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": getDiskInfo(disk, debug)})
	})

	s.POST("/api/disk_info", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": getDiskInfoDetail(debug)})
	})

	s.POST("/api/container", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": getContainers(debug)})
	})

	s.POST("/api/log", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": getLogData(log, debug)})
	})

	s.POST("/api/log_del", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": delLog(log, debug)})
	})
}

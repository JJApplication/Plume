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
func engine(eth, disk, log string) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	s := gin.New()
	s.Use(gin.Recovery())
	s.Use(middlewareCors())
	if log != "" {
		s.Use(gin.Logger())
	}

	wrapApi(s, eth, disk, log)

	return s
}

// 路由规则
func wrapApi(s *gin.Engine, eth, disk, log string) {
	s.POST("/api/cpu", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": getCPUInfo()})
	})

	s.POST("/api/server", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": getServerInfo()})
	})

	s.POST("/api/mem", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": getMemInfo()})
	})

	s.POST("/api/kernel", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": getKernelData()})
	})

	s.POST("/api/progress", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": getProgressData()})
	})

	s.POST("/api/progress_list", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": getProgressListData(c.Query("num"))})
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

	s.POST("/api/log", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": getLogData(log)})
	})

	s.POST("/api/log_del", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": delLog(log)})
	})
}

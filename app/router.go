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
func engine(c CmdArgs) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	s := gin.New()
	s.Use(gin.Recovery())
	s.Use(middlewareCors())
	s.Use(middlewareAuthKey(c.Key))
	if c.Log != "" {
		s.Use(gin.Logger())
	}

	wrapApi(s, c)

	return s
}

// 路由规则
func wrapApi(s *gin.Engine, cmd CmdArgs) {
	// 整合接口
	s.POST("/api/comb", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": getCombinationData(cmd.Eth, cmd.Disk, cmd.Debug)})
	})

	s.POST("/api/cpu", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": getCPUInfo(cmd.Debug)})
	})

	s.POST("/api/cpu_info", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": getCpuInfoDetail(cmd.Debug)})
	})

	s.POST("/api/server", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": getServerInfo(cmd.Debug)})
	})

	s.POST("/api/mem", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": getMemInfo(cmd.Debug)})
	})

	s.POST("/api/mem_info", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": getMemInfoDetail(cmd.Debug)})
	})

	s.POST("/api/kernel", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": getKernelData(cmd.Debug)})
	})

	s.POST("/api/progress", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": getProgressData(cmd.Debug)})
	})

	s.POST("/api/progress_list", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": getProgressListData(c.Query("num"), cmd.Debug)})
	})

	s.POST("/api/net", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": getNetInfo(cmd.Eth, cmd.Debug)})
	})

	s.POST("/api/net_info", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": getNetInfoDetail(c.Query("num"), cmd.Debug)})
	})

	s.POST("/api/disk", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": getDiskInfo(cmd.Disk, cmd.Debug)})
	})

	s.POST("/api/disk_info", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": getDiskInfoDetail(cmd.Debug)})
	})

	s.POST("/api/containers", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": getContainers(cmd.DockerAPI)})
	})

	s.POST("/api/container", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": getContainer(c.Query("id"), cmd.DockerAPI)})
	})

	s.POST("/api/container/start", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": startContainer(c.Query("id"), cmd.DockerAPI)})
	})

	s.POST("/api/container/stop", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": stopContainer(c.Query("id"), cmd.DockerAPI)})
	})

	s.POST("/api/container/restart", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": restartContainer(c.Query("id"), cmd.DockerAPI)})
	})

	s.POST("/api/container/create", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": createContainer(c.Query("name"), cmd.DockerAPI, c.Request.Body)})
	})

	s.POST("/api/container/delete", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": delContainer(c.Query("id"), cmd.DockerAPI)})
	})

	s.POST("/api/log", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": getLogData(cmd.Log, cmd.Debug)})
	})

	s.POST("/api/log_del", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": delLog(cmd.Log, cmd.Debug)})
	})
}

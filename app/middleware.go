/*
Name: plume
File: middleware.go
Author: Landers
*/

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

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

// 通信密钥中间件仅在key参数存在时使用
func middlewareAuthKey(key string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if key != "" {
			if token != key {
				c.JSON(http.StatusForbidden, "key error")
				c.AbortWithStatus(http.StatusForbidden)
			}
			c.Next()
		} else {
			c.Next()
		}
	}
}
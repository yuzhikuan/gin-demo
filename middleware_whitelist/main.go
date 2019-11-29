package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// IpAuthMiddleware 白名单中间件
func IpAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ipList := [1]string{
			"127.0.0.2",
		}
		flag := false
		clientIp := c.ClientIP()
		for _, host := range ipList {
			if clientIp == host {
				flag = true
				break
			}
		}
		if !flag {
			c.String(http.StatusForbidden, "%s, not in iplist", clientIp)
			c.Abort()
		}
	}
}

func main() {
	r := gin.Default()

	r.Use(IpAuthMiddleware())

	r.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "hello, world!")
	})

	r.Run()
}

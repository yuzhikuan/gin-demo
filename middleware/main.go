package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// Logger 自定义
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// 在gin上下文中定义变量
		c.Set("example", "12345")

		// 请求前

		c.Next() // 处理中

		// 请求后
		latency := time.Since(t) // time.Now().Sub(t)的缩写
		log.Print(latency)

		status := c.Writer.Status()
		log.Print(status)
	}
}

func main() {
	r := gin.New()                  // Default() 自动包含了Logger/Recovery中间件
	r.Use(Logger(), gin.Recovery()) // Recover中间件能恢复panic

	r.GET("/test", func(c *gin.Context) {
		// 获取gin上下文中的变量example
		example := c.MustGet("example").(string)
		log.Print(example)
	})

	r.Run()
}

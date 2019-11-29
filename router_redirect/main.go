package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/home", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to home !")
	})

	r.GET("/redirect", func(c *gin.Context) {
		// 可以重定向内部地址或者外部地址
		c.Redirect(http.StatusTemporaryRedirect, "/home") // 307
		// c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com/") // 301
	})

	r.Run()
}

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/autotls"
)

func main() {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {

		c.String(http.StatusOK, "hello test")
	})

	autotls.Run(r, "www.itpp.tk")
}

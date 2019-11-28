package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/:name/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name": c.Param("name"),
			"id":   c.Param("id"),
		})
	})
	r.Run()
}

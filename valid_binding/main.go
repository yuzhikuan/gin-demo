package main

import "github.com/gin-gonic/gin"

import "net/http"

type Person struct {
	Age     int    `form:"age" binding:"required,gt=10"`
	Name    string `form:"name" binding:"required"`
	Address string `form:"address" binding:"required"`
}

func main() {
	r := gin.Default()
	r.GET("/testing", func(c *gin.Context) {
		var person Person
		if err := c.ShouldBind(&person); err != nil {
			c.String(http.StatusBadRequest, "%v", err)
			c.Abort()
			return
		}
		c.String(http.StatusOK, "%v", person)
	})
	r.Run()
}

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/test", func(c *gin.Context) {
		firstName := c.PostForm("first_name")
		lastName := c.DefaultPostForm("last_name", "defaut_last_name")
		c.String(http.StatusOK, "%s,%s", firstName, lastName)
	})
	r.Run(":8081")
}

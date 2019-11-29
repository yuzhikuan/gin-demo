package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 加载模板
	// r.LoadHTMLGlob("tmpls/*")
	r.LoadHTMLFiles("tmpls/index.tmpl")

	// 定义路由
	r.GET("/index", func(c *gin.Context) {
		//根据完整文件名渲染模板，并传递参数
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main Website Index",
		})
	})

	// 加载模板
	r.LoadHTMLGlob("tmpls/**/*")
	// 以下会将两个模板内容合并输出
	r.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
			"title": "Posts Index Title",
		})
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			"title": "Users Index Title",
		})
	})

	r.Run(":8080")
}

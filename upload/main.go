package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/upload", func(c *gin.Context) {
		file, header, err := c.Request.FormFile("upload")
		fileName := header.Filename
		fmt.Printf("Upload file name is: %s", fileName)

		// tmp目录需存在，不然会报错
		out, err := os.Create("./tmp/" + fileName)
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()

		_, err = io.Copy(out, file)
		if err != nil {
			log.Fatal(err)
		}
	})
	r.Run()
}

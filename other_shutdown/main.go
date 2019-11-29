package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 监听/test路由，并在10秒后返回结果
	r.GET("/test", func(c *gin.Context) {
		time.Sleep(10 * time.Second)
		c.String(http.StatusOK, "test shutdown")
	})

	srv := &http.Server{
		Addr:    ":8085",
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen Error: %s\n", err)
		}
	}()

	// 监听关闭服务事件，阻塞channel
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown server ...")

	// 倒计时10秒
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil { // 正式关闭服务
		log.Fatalf("Shutdown Error: %s\n", err)
	}

	log.Println("Exiting ...")
}

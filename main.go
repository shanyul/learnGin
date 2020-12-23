package main

import (
	"context"
	"fmt"
	"learngo/pkg/setting"
	"learngo/routers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	/* router := gin.Default()
	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	}) */

	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
        Handler:        router,
        ReadTimeout:    setting.ReadTimeout,
        WriteTimeout:   setting.WaitTimeout,
        MaxHeaderBytes: 1 << 20, // 1M
	}

	go func ()  {
		if err := s.ListenAndServe(); err != nil {
			log.Printf("Listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<- quit

	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")
}

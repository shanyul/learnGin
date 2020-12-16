package main

import (
	"fmt"
	"learngo/pkg/setting"
	"net/http"
	"learngo/routers"

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

	s.ListenAndServe()
}

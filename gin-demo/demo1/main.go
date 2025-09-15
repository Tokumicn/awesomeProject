package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"runtime"
	"time"
)

func main() {

	fmt.Println(time.Now().Format(time.DateTime))

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		fmt.Println("当前的 Goroutine 数量: ", runtime.NumGoroutine())
		c.JSONP(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8081")
}

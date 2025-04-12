package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {

	fmt.Println(time.Now().Format(time.DateTime))

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSONP(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8081")
}

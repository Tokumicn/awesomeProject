package main

import (
	"github.com/gin-gonic/gin"
	"github.com/playwright-community/playwright-go"
	"log"
	"net/http"
)

func main() {

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSONP(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/playwright/version", func(c *gin.Context) {

		err := playwright.Install()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
			log.Printf("无法安装 playwright: %v\n", err)
			return
		}

		// 初始化Playwright
		pw, err := playwright.Run()
		if err != nil {
			log.Printf("无法启动 playwright: %v\n", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
			return
		}
		defer pw.Stop()

		c.JSON(http.StatusOK, gin.H{
			"message": "playwright 安装完成,环境完备。",
		})
	})

	err := r.Run(":8080")
	if err != nil {
		log.Fatalf("启动服务器失败: %v", err)
	}
}

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func main() {
	engine := gin.New()
	engine.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "煎鱼")
	})
	engine.GET("/metrics", gin.WrapH(promhttp.Handler()))
	engine.Run(":10001")
}

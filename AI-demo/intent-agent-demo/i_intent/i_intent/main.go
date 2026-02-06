package main

import (
	"context"
	"i_intent/data/scene_data"
	"i_intent/handler"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
)

func main() {
	ctx := context.Background()

	scene_data.LoadAllSceneConfigs()

	err := g.Cfg().GetAdapter().(*gcfg.AdapterFile).SetPath("./config")
	if err != nil {
		panic(err)
	}

	router := gin.Default()
	router.Use(handler.CORSMiddleware())

	handler.RegisterRoutes(ctx, router)

	router.GET("/", func(c *gin.Context) {
		c.File("./config/user_input.html")
	})

	log.Printf("🚀 后端服务启动中...")
	log.Printf("📍 地址: %s", g.Cfg().MustGet(ctx, "server.BackendURL").String())
	log.Printf("🌍 环境: %s", g.Cfg().MustGet(ctx, "server.ENV").String())
	log.Printf("🔗 允许跨域: *")

	host := g.Cfg().MustGet(ctx, "server.HOST").String()
	port := g.Cfg().MustGet(ctx, "server.BACKEND_PORT").String()

	if err := router.Run(host + ":" + port); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}

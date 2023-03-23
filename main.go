package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"somnus-gin/bootstrap"
	"somnus-gin/global"
)

func main() {

	// 初始化配置
	bootstrap.InitConfig()

	// 初始化日志
	global.App.Log = bootstrap.InitLog()
	global.App.Log.Info("log init success!")

	r := gin.Default()

	// 测试路由
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// 启动服务器
	r.Run(":" + global.App.Config.App.Port)
}

package bootstrap

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"somnus-gin/global"
	"somnus-gin/routes"
	"syscall"
	"time"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	// 前端项目静态资源
	router.StaticFile("/", "./static/dist/index.html")
	router.Static("/assets", "./static/dist/assets")
	router.StaticFile("/favicon.ico", "./static/dist/favicon.ico")
	// 其他静态资源
	router.Static("/public", "./static")
	router.Static("/storage", "./storage/app/public")

	// 注册api分组路由
	apiGroup := router.Group("/api")
	routes.SetApiGroupRoutes(apiGroup)
	return router
}

// RunServer 启动服务器
func RunServer() {
	r := setupRouter()
	srv := &http.Server{
		Addr:    ":" + global.App.Config.App.Port,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待终端信号以优雅的关闭服务器(设置5秒的超时时间)
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("shutdown Server ....")

	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancle()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}

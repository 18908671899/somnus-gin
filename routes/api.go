package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"somnus-gin/app/common/request"
	"somnus-gin/app/controller/app"
	"somnus-gin/app/middleware"
	"somnus-gin/app/services"
	"time"
)

func SetApiGroupRoutes(router *gin.RouterGroup) {
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	router.GET("/test", func(ctx *gin.Context) {
		time.Sleep(5 * time.Second)
		ctx.String(http.StatusOK, "success")
	})

	router.POST("/user/register", func(c *gin.Context) {
		var form request.Register
		if err := c.ShouldBindJSON(&form); err != nil {
			fmt.Println("err:", err)
			c.JSON(http.StatusOK, gin.H{
				"error": request.GetErrorMsg(form, err),
			})
			return
		}
		// fmt.Println("form的值:", form)
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	})

	router.POST("/auth/register", app.Register)

	router.POST("/auth/login", app.Login)

	authRouter := router.Group("").Use(middleware.JWTAuth(services.AppGuardName))
	{
		authRouter.POST("/auth/info", app.Info)
	}
}

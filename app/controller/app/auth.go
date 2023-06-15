package app

import (
	"github.com/gin-gonic/gin"
	"somnus-gin/app/common/request"
	"somnus-gin/app/common/response"
	"somnus-gin/app/services"
)

func Login(ctx *gin.Context) {
	var form request.Login

	if err := ctx.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(ctx, request.GetErrorMsg(form, err))
		return
	}

	if err, user := services.UserService.Login(form); err != nil {
		response.BusinessFail(ctx, err.Error())
	} else {
		tokenData, err, _ := services.JwtService.CreateToken(services.AppGuardName, user)
		if err != nil {
			response.BusinessFail(ctx, err.Error())
			return
		}
		response.Success(ctx, tokenData)
	}
}

func Info(ctx *gin.Context) {
	err, user := services.UserService.GetUserInfo(ctx.Keys["id"].(string))
	if err != nil {
		response.BusinessFail(ctx, err.Error())
		return
	}
	response.Success(ctx, user)
}

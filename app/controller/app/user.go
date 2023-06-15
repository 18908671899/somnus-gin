package app

import (
	"github.com/gin-gonic/gin"
	"somnus-gin/app/common/request"
	"somnus-gin/app/common/response"
	"somnus-gin/app/services"
)

func Register(ctx *gin.Context) {
	var form request.Register
	if err := ctx.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(ctx, request.GetErrorMsg(form, err))
		return
	}

	if err, user := services.UserService.Register(form); err != nil {
		response.BusinessFail(ctx, err.Error())
	} else {
		response.Success(ctx, user)
	}
}

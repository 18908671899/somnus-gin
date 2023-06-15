package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"somnus-gin/global"
)

// 响应结构体
type Response struct {
	ErrorCode int         `json:"error_code"` // 自定义错误码
	Data      interface{} `json:"data"`
	Message   string      `json:"message"`
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		0,
		data,
		"ok",
	})
}

func Fail(ctx *gin.Context, errorCode int, msg string) {
	ctx.JSON(http.StatusOK, Response{
		errorCode,
		nil,
		msg,
	})
}

func FailByError(ctx *gin.Context, error global.CustomError) {
	Fail(ctx, error.ErrorCode, error.ErrorMsg)
}

func ValidateFail(ctx *gin.Context, msg string) {
	Fail(ctx, global.Errors.ValidatorError.ErrorCode, msg)
}

func BusinessFail(ctx *gin.Context, msg string) {
	Fail(ctx, global.Errors.BusinessError.ErrorCode, msg)
}

func TokenFail(c *gin.Context) {
	FailByError(c, global.Errors.TokenError)
}

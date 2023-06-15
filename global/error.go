package global

type CustomError struct {
	ErrorCode int
	ErrorMsg  string
}

type CustomErrors struct {
	BusinessError  CustomError
	ValidatorError CustomError
	TokenError     CustomError
}

var Errors = CustomErrors{
	BusinessError:  CustomError{40000, "业务错误"},
	ValidatorError: CustomError{42200, "请求参数错误"},
	TokenError: CustomError{
		ErrorCode: 40100,
		ErrorMsg:  "登录授权失败",
	},
}

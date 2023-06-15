package services

import (
	"errors"
	"somnus-gin/app/common/request"
	"somnus-gin/app/models"
	"somnus-gin/global"
	"somnus-gin/utils"
	"strconv"
)

type userService struct {
}

var UserService = new(userService)

func (u *userService) Register(param request.Register) (err error, user *models.User) {
	user = new(models.User)
	result := global.App.DB.Where("mobile = ?", param.Mobile).Select("id").First(user)
	if result.RowsAffected != 0 {
		err = errors.New("手机号已存在")
		return
	}
	user = &models.User{Name: param.Name, Mobile: param.Mobile, Password: utils.BcryptMake([]byte(param.Password))}
	err = global.App.DB.Create(user).Error
	return
}

func (u *userService) Login(params request.Login) (err error, user *models.User) {
	err = global.App.DB.Where("mobile = ?", params.Mobile).First(&user).Error
	if err != nil || !utils.BcryptMakeCheck([]byte(params.Password), user.Password) {
		err = errors.New("用户名不存在或密码错误")
	}
	return
}

func (u *userService) GetUserInfo(id string) (err error, user models.User) {
	intId, err := strconv.Atoi(id)
	err = global.App.DB.First(&user, intId).Error
	if err != nil {
		err = errors.New("数据不存在")
	}
	return
}

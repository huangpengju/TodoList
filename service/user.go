package service

import (
	"TodoList/model"
	"TodoList/pkg/utils"
	"TodoList/serializer"

	"github.com/jinzhu/gorm"
)

type UserService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=3,max=15"`
	Password string `form:"password" json:"password" binding:"required,min=5,max=16"`
}

// 注册服务，返回Json
func (service *UserService) Register() serializer.Response {
	var user model.User
	var count int
	model.DB.Model(&model.User{}).Where("user_name=?", service.UserName).First(&user).Count(&count)
	if count == 1 {
		return serializer.Response{
			Status: 400,
			Msg:    "已经有这个人了，无需再注册",
		}
	}
	user.UserName = service.UserName
	//加密密码
	if err := user.SetPassword(service.Password); err != nil {
		return serializer.Response{
			Status: 400,
			Msg:    err.Error(),
		}
	}
	//创建用户
	if err := model.DB.Create(&user).Error; err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "数据库操作错误",
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "用户注册成功",
	}
}

// 用户登录
func (service *UserService) Login() serializer.Response {
	var user model.User
	//先去验证用户
	if err := model.DB.Where("user_name=?", service.UserName).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return serializer.Response{
				Status: 400,
				Msg:    "用户不存在，请先登录",
			}
		}
		//用户存在，是其他因素的错误
		return serializer.Response{
			Status: 500,
			Msg:    "数据库错误",
		}
	}
	//去验证登录用户的密码
	if !user.CheckPassword(service.Password) {
		return serializer.Response{
			Status: 400,
			Msg:    "密码错误",
		}
	}
	//发一个token ,为了其他功能需要身份验证，所给前端存储的。

	token, err := utils.GenerateToken(user.ID, service.UserName, service.Password)
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "Token签发错误",
		}
	}
	//返回用户信息和token
	return serializer.Response{
		Status: 200,
		Data:   serializer.TokenData{User: serializer.BuildUser(user), Token: token},
		Msg:    "登录成功",
	}
}

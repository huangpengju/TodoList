package api

import (
	"TodoList/service"
	"fmt"

	"github.com/gin-gonic/gin"
)

// 把路由注册命令，传给服务
func UserRegister(c *gin.Context) {
	var userRegister service.UserService
	if err := c.ShouldBind(&userRegister); err == nil {
		res := userRegister.Register()
		fmt.Println("注册绑定成功====", res)

		c.JSON(200, res)
	} else {
		fmt.Println("注册绑定失败====", err)
		c.JSON(400, err)
	}
}

// 用户登录
func UserLogin(c *gin.Context) {
	var userLogin service.UserService
	if err := c.ShouldBind(&userLogin); err == nil {
		res := userLogin.Login()
		fmt.Println("登录绑定成功====", res)

		c.JSON(200, res)
	} else {
		fmt.Println("登录绑定错误====", err)
		c.JSON(400, err)
	}
}

package routes

import (
	"TodoList/api"
	"TodoList/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default() //需要学习
	store := cookie.NewStore([]byte("something-very-secret"))
	r.Use(sessions.Sessions("mysession", store))
	v1 := r.Group("api/v1")
	{
		//用户操作
		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)
		authed := v1.Group("/")
		authed.Use(middleware.JWT()) //验证中间件，是否有访问权限
		{
			//创建备忘录
			authed.POST("task", api.CreateTask)
			authed.GET("task/:id", api.ShowTask)
			authed.POST("tasks", api.ListTask)
			authed.GET("tasks1/:page_num/:page_size", api.ListTask1)
			authed.PUT("task/:id", api.UpdateTask)
			authed.POST("search", api.SearchTask)
			authed.DELETE("task/:id", api.DeleteTask)
		}
	}
	return r
}

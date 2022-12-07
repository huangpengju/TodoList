package middleware

import (
	"TodoList/pkg/utils"
	"time"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := 200
		// var data interface{}//包没写
		token := c.GetHeader("Authorization")
		// fmt.Println("token===", token)
		if token == "" {
			code = 404 //参数不对 没有token
		} else {
			claim, err := utils.ParseToken(token) //生成token
			if err != nil {                       //解析toke有错误
				code = 403 //无权限，是假的

			} else if time.Now().Unix() > claim.ExpiresAt {
				code = 401 //token无效了
			}
		}
		if code != 200 {
			c.JSON(200, gin.H{
				"status": code,
				"msg":    "Token解析错误",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

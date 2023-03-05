package middlerware

import (
	"gin-learn/gin-todolist/api-gateway/pkg/e"
	"gin-learn/gin-todolist/api-gateway/pkg/utils"
	"time"

	"github.com/gin-gonic/gin"
)

// JWT token验证中间件
func JWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var code int
		var data interface{}
		code = 200
		token := ctx.GetHeader("Authorization")
		if token == "" {
			code = 404
		} else {
			if claims, err := utils.ParseToken(token); err != nil {
				code = e.ErrorAuthCheckTokenFail
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ErrorAuthCheckTokenTimeout
			}
		}
		if code != e.SUCCESS {
			ctx.JSON(200, gin.H{
				"status": code,
				"msg":    e.GetMsg(uint(code)),
				"data":   data,
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}

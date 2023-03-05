package middlerware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// 接收服务实例, 并存到gin.Key中
func InitMiddleware(service []interface{}) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 将实例存在gin.Keys中
		ctx.Keys = make(map[string]interface{})
		ctx.Keys["user"] = service[0]
		ctx.Keys["task"] = service[1]
		ctx.Next()
	}
}

// 错误处理中间件
func ErrorMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				ctx.JSON(200, gin.H{
					"code": 404,
					"msg":  fmt.Sprintf("%s", r),
				})
				ctx.Abort()
			}
		}()
		ctx.Next()
	}
}

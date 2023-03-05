package routes

import (
	"gin-learn/gin-todolist/api-gateway/internal/handler"
	middleware "gin-learn/gin-todolist/api-gateway/middlerware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func NewRouter(service ...interface{}) *gin.Engine {
	ginRouter := gin.Default()
	// 全局使用中间件
	ginRouter.Use(middleware.Cors(), middleware.InitMiddleware(service), middleware.ErrorMiddleware())
	store := cookie.NewStore([]byte("something-very-secret"))
	ginRouter.Use(sessions.Sessions("mysession", store))
	v1 := ginRouter.Group("/api/v1")
	{
		v1.GET("ping", func(ctx *gin.Context) {
			ctx.JSON(200, "success")
		})
		// user 模块
		v1.POST("/user/register", handler.UserRegister)
		v1.POST("/user/login", handler.UserLogin)

		// 需要登录保护, 组中间件
		authed := v1.Group("/")
		authed.Use(middleware.JWT())
		{
			// task 模块
			authed.GET("task", handler.GetTaskList)
			authed.POST("task", handler.CreateTask)
			authed.PUT("task", handler.UpdateTask)
			authed.DELETE("task", handler.DeleteTask)
		}
	}
	return ginRouter
}

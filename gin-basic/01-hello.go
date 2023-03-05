package main

import "github.com/gin-gonic/gin"

func main() {
	// 创建一个默认路由
	r := gin.Default()
	// GET: 请求方式; /hello: 请求路径
	// 当客户端使用GET方式请求/hello路径时，会执行后面的函数
	r.GET("/hello", func(context *gin.Context) {
		// context.JSON返回json格式的数据
		context.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})
	// 启动服务，默认在0.0.0.0:8080启动
	r.Run()
}

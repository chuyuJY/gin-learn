package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: false,
		AllowOrigins:    []string{"https://foo.com"}, // 允许跨域发来请求的网站,
		AllowOriginFunc: func(origin string) bool { // 自定义过滤源站的方法
			return origin == "https://github.com"
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // 允许的请求方法,
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	r.Run()
}

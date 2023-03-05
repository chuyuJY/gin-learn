package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// gin.H是map[string]interface{}的缩写
	r.GET("/someXML", func(context *gin.Context) {
		// 1. 自己拼接xml
		context.XML(http.StatusOK, gin.H{"message": "Hello World!"})
	})
	r.GET("/moreXML", func(context *gin.Context) {
		// 2. 使用结构体
		type Msg struct {
			Name   string `xml:"name"`
			Gender string `xml:"gender"`
			Age    int    `xml:"age"`
		}
		var msg Msg
		msg.Name = "小王子"
		msg.Gender = "男"
		msg.Age = 18
		context.XML(http.StatusOK, msg)
	})
	r.Run(":8080")
}

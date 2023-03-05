package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// gin.H是map[string]interface{}的缩写
	r.GET("/someJSON", func(context *gin.Context) {
		// 1. 自己拼接json
		context.JSON(http.StatusOK, gin.H{"message": "Hello World!"})
	})
	r.GET("/moreJSON", func(context *gin.Context) {
		// 2. 使用结构体
		type Msg struct {
			Name   string `json:"name,omitempty"`
			Gender string `json:"gender,omitempty"`
			Age    int    `json:"age,omitempty"`
		}
		msg := Msg{}
		msg.Name = "小王子"
		msg.Gender = "男"
		msg.Age = 18
		context.JSON(http.StatusOK, msg)
	})
	r.Run(":8080")
}

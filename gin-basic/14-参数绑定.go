package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Login struct {
	Username string `json:"username,omitempty" form:"username" binding:"required"`
	Password string `json:"password,omitempty" form:"password" binding:"required"`
}

func main() {
	r := gin.Default()
	// 1. 绑定json ({"user": "q1mi", "password": "123456"})
	r.POST("/loginJSON", func(context *gin.Context) {
		login := Login{}
		if err := context.ShouldBind(&login); err == nil {
			log.Printf("login info:%#v\n", login)
			context.JSON(http.StatusOK, gin.H{
				"username": login.Username,
				"password": login.Password,
			})
		} else {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})
	// 2. 绑定form表单 (user=q1mi&password=123456)
	r.POST("/loginForm", func(context *gin.Context) {
		login := Login{}
		if err := context.ShouldBind(&login); err == nil {
			log.Printf("login info:%#v\n", login)
			context.JSON(http.StatusOK, gin.H{
				"username": login.Username,
				"password": login.Password,
			})
		} else {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})
	// 3. 绑定queryString (/loginQuery?user=q1mi&password=123456)
	r.GET("/loginQuery", func(context *gin.Context) {
		login := Login{}
		// ShouldBind()会根据请求的Content-Type自行选择绑定器
		if err := context.ShouldBind(&login); err == nil {
			context.JSON(http.StatusOK, gin.H{
				"user":     login.Username,
				"password": login.Password,
			})
		} else {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})
	r.Run()
}

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/user/search", func(context *gin.Context) {
		username := context.DefaultQuery("username", "小王子")
		address := context.Query("address")
		context.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
			"address":  address,
		})
	})
	r.Run()
}

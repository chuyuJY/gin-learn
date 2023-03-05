package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/test1", func(context *gin.Context) {
		context.Request.URL.Path = "/test2"
		r.HandleContext(context)
	})
	r.GET("test2", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"hello": "world"})
	})
	r.Run()
}

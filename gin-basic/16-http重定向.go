package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/test", func(context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently, "http://www.baidu.com/")
	})
	r.Run()
}

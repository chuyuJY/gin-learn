package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("./templates/**/*")
	//r.LoadHTMLFiles("templates/posts/index.html", "templates/users/index.html")
	r.GET("posts/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "posts/index.html", gin.H{
			"title": "posts/index",
		})
	})
	r.GET("users/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "users/index.html", gin.H{
			"title": "users/index",
		})
	})
	r.Run(":8080")
}

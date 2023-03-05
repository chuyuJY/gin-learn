package main

import "github.com/gin-gonic/gin"

func rest(r *gin.Engine) {
	r.GET("/book", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "GET",
		})
	})
	r.POST("/book", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "POST",
		})
	})
	r.PUT("/book", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "PUT",
		})
	})
	r.DELETE("/book", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "DELETE",
		})
	})
	r.Run()
}

func main() {
	r := gin.Default()
	rest(r)
}

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/someYAML", func(context *gin.Context) {
		context.YAML(http.StatusOK, gin.H{"message": "ok", "status": http.StatusOK})
	})
	r.Run(":8080")
}

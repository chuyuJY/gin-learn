package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
)

func main() {
	r := gin.Default()
	r.POST("/json", func(context *gin.Context) {
		data, _ := context.GetRawData()
		m := map[string]interface{}{}
		_ = json.Unmarshal(data, &m)
		context.JSON(http.StatusOK, m)
	})
	r.Run()
}

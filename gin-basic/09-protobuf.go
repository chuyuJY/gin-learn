package main

import (
	"net/http"

	"github.com/gin-gonic/gin/testdata/protoexample"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/someProtobuf", func(context *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		label := "test"
		data := &protoexample.Test{
			Label: &label,
			Reps:  reps,
		}
		context.ProtoBuf(http.StatusOK, data)
	})
	r.Run(":8080")
}

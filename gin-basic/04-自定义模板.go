package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"safe": func(s string) template.HTML {
			return template.HTML(s)
		},
	})
	r.LoadHTMLFiles("./index.tmpl")
	r.GET("/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.tmpl", "<a href='https://liwenzhou.com'>李文周的博客</a>")
	})
	r.Run(":8080")
}

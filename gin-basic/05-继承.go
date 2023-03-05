package main

import (
	"net/http"
	"path/filepath"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	layouts, err := filepath.Glob(templatesDir + "/layouts/*.tmpl")
	if err != nil {
		panic(err.Error())
	}
	includes, err := filepath.Glob(templatesDir + "/includes/*.tmpl")
	if err != nil {
		panic(err.Error())
	}
	// 为layouts/和includes/目录生成templates map
	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		r.AddFromFiles(filepath.Base(include), files...)
	}

	return r
}

func indexFunc(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", nil)
}

func homeFunc(c *gin.Context) {
	c.HTML(http.StatusOK, "home.tmpl", nil)
}

func main() {
	r := gin.Default()
	r.HTMLRender = loadTemplates("./templates")
	r.GET("/index", indexFunc)
	r.GET("/home", homeFunc)
	r.Run()
}

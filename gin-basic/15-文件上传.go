package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func uploadSingleFile(r *gin.Engine) {
	r.POST("/uploadSingle", func(context *gin.Context) {
		// 单个文件
		file, err := context.FormFile("file")
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		log.Println(file.Filename)
		dst := fmt.Sprintf("./upload_files/%s", file.Filename)
		// 上传文件到指定文件夹
		err = context.SaveUploadedFile(file, dst)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"message": "upload failed."})
		}
		context.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("%s upload success.", file.Filename)})
	})
	//r.Run()
}

func uploadMultiFiles(r *gin.Engine) {
	r.POST("/uploadMulti", func(context *gin.Context) {
		form, _ := context.MultipartForm()
		files := form.File["files[]"]
		for index, file := range files {
			log.Println(file.Filename)
			dst := fmt.Sprintf("./upload_files/%d_%s", index, file.Filename)
			// 上传文件到指定目录
			context.SaveUploadedFile(file, dst)
		}
		context.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("%d files upload success.", len(files))})
	})
	//r.Run()
}

func main() {
	r := gin.Default()
	// 1. 静态页面	默认显示./upload下的index.html页面。注意，必须是index命名
	//r.Static("/", "./upload_files")
	// 2. 动态显示
	r.LoadHTMLFiles("./upload_files/upload_files.html")
	r.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "upload_files.html", gin.H{
			"title": "upload_files",
		})
	})

	// 处理multipart forms提交文件时默认的内存限制是32 MiB
	// 可以通过下面的方式修改
	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	uploadSingleFile(r)
	uploadMultiFiles(r)
	r.Run()
}

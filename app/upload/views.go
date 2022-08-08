package upload

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadDemo(ctx *gin.Context) {
	// file, _ := ctx.FormFile("file")
	// ctx.SaveUploadedFile(file, "./"+file.Filename)
	// ctx.FileAttachment("./"+file.Filename, file.Filename)

	form, _ := ctx.MultipartForm()
	files := form.File["upload"]
	for _, file := range files {
		fmt.Println(file.Filename)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

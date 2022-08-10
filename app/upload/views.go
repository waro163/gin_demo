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

func UploadToThird(ctx *gin.Context) {
	url := "http://127.0.0.1:8081/upload/multi"
	req, err := http.NewRequest(ctx.Request.Method, url, ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	req.Header = ctx.Request.Header
	transport := http.DefaultTransport
	res, err := transport.RoundTrip(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	// ctx.Data(res.StatusCode, res.Header.Get("Content-Type"), res.Body)
	ctx.DataFromReader(res.StatusCode, res.ContentLength, res.Header.Get("Content-Type"), res.Body, nil)
}

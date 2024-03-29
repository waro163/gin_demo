package download

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"path"

	"github.com/gin-gonic/gin"
)

func ShowFileDemo(ctx *gin.Context) {
	dir, err := os.Getwd()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"err_msg": err,
		})
		return
	}
	filePath := path.Join(dir, "/static/tesla-shiming.wav")
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"err_msg": "read file error:" + err.Error(),
		})
		return
	}
	// ctx.Data(http.StatusOK, "audio/wav", fileContent)
	reader := bytes.NewReader(fileContent)
	extraHeaders := map[string]string{
		"Content-Disposition": `attachment; filename="gopher.wav"`,
	}
	ctx.DataFromReader(http.StatusOK, int64(len(fileContent)), "audio/wav", reader, extraHeaders)
}

func DownloadDemo(ctx *gin.Context) {
	dir, err := os.Getwd()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"err_msg": err,
		})
		return
	}
	file := path.Join(dir, "/static/go-demo.png")
	info, err := os.Stat(file)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"err_msg": "os stat file error: " + err.Error(),
		})
		return
	}
	if info.Mode().IsRegular() {
		ctx.FileAttachment(file, "demo.png")
	} else {
		ctx.JSON(http.StatusNotFound, gin.H{
			"err_msg": "not found" + file,
		})
	}

}

func RemoteDownload(ctx *gin.Context) {
	// response, err := http.Get("https://raw.githubusercontent.com/gin-gonic/logo/master/color.png")
	response, err := http.Get("http://127.0.0.1:8080/api/download/demo")
	if err != nil || response.StatusCode != http.StatusOK {
		ctx.Status(http.StatusServiceUnavailable)
		return
	}

	reader := response.Body
	defer reader.Close()
	contentLength := response.ContentLength
	contentType := response.Header.Get("Content-Type")

	extraHeaders := map[string]string{
		"Content-Disposition": `attachment; filename="gopher.png"`,
	}

	ctx.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
}

func DownloadAndCreateFile(ctx *gin.Context) {
	response, err := http.Get("http://127.0.0.1:8080/api/download/demo")
	if err != nil || response.StatusCode != http.StatusOK {
		ctx.Status(http.StatusServiceUnavailable)
		return
	}

	reader := response.Body
	defer reader.Close()
	// we can get file name ant type from Content-Disposition response header
	fileName := "demo.png"
	file, err := os.Create(fileName)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"message": "create file error"})
		return
	}
	_, err = io.Copy(file, reader)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"message": "copy file error"})
		return
	}
	if err = file.Close(); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"message": "close file error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
}

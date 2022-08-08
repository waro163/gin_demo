package download

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DownloadDemo(ctx *gin.Context) {
	// ctx.JSON(200, gin.H{
	// 	"hi": "world",
	// })
	ctx.FileAttachment("/Users/waro/work/gin_demo/main.go", "main.go")

}

func RemoteDownload(ctx *gin.Context) {
	response, err := http.Get("https://raw.githubusercontent.com/gin-gonic/logo/master/color.png")
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

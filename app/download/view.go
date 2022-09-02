package download

import (
	"net/http"
	"os"
	"path"

	"github.com/gin-gonic/gin"
)

func DownloadDemo(ctx *gin.Context) {
	dir, err := os.Getwd()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"err_msg": err,
		})
		return
	}
	file := path.Join(dir, "main.go")
	info, _ := os.Stat(file)
	if info.Mode().IsRegular() {
		ctx.FileAttachment(file, "main.go")
	} else {
		ctx.JSON(http.StatusNotFound, gin.H{
			"err_msg": "not found" + file,
		})
	}

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

package download

import (
	"github.com/gin-gonic/gin"
)

func DownloadDemo(ctx *gin.Context) {
	// ctx.JSON(200, gin.H{
	// 	"hi": "world",
	// })
	ctx.FileAttachment("/Users/waro/work/gin_demo/main.go", "main.go")

}

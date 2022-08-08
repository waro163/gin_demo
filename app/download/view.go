package download

import (
	"github.com/gin-gonic/gin"
)

func DownloadDemo(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"name": "",
	})
}

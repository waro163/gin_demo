package download

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.RouterGroup) {
	r.GET("/demo", DownloadDemo)
	r.GET("/get_from_remote", RemoteDownload)
}

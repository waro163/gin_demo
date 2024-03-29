package download

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.RouterGroup) {
	r.GET("/show", ShowFileDemo)
	r.GET("/demo", DownloadDemo)
	r.GET("/get_from_remote", RemoteDownload)
	r.GET("/get_create", DownloadAndCreateFile)
}

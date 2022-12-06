package upload

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.RouterGroup) {
	r.POST("/demo", UploadDemo)
	r.Any("/multi", UploadToThird)
	r.Any("/local", UploadLocalFile)
}

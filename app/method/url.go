package method

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.RouterGroup) {
	r.GET("/get", GetMethod)
	r.POST("/post", PostMethod)
	r.PUT("/put", PutMethod)
	r.PATCH("/patch", PatchMethod)
	r.DELETE("/delete", DeleteMethod)
}

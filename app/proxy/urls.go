package proxy

import "github.com/gin-gonic/gin"

func RegisterRouter(r *gin.RouterGroup) {
	r.Any("/*", ProxyDemo)
}

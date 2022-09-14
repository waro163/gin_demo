package proxy

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.RouterGroup) {
	r.Any("/normal/*ping", ProxyDemo)
	r.Any("/tracing/*any", TracingDemo)
}

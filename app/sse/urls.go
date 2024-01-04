package sse

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.RouterGroup) {
	r.Any("/event", gin.WrapF(sseHandle))
	r.Any("/gin", ServerSendEvents)
}

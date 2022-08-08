package cookie

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func CookieDemo(ctx *gin.Context) {
	name, err := ctx.Cookie("name")
	if err != nil {
		ctx.SetCookie("name", "waro", 3600, "/cookie", "127.0.0.1", true, false)
		ctx.SetCookie("age", "18", 3600, "/", "127.0.0.1", false, true)
		ctx.SetCookie("phone", "18888888", 3600, "/", "127.0.0.1", true, true)
	}
	ctx.JSON(200, gin.H{
		"name": name,
	})
}

func generate4096byte() string {
	var builder strings.Builder
	for i := 0; i < 4090; i++ {
		builder.WriteByte('a')
	}
	return builder.String()
}

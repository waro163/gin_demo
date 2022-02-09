package app_user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserLogin(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

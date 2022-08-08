package main

import (
	"gin_demo/app/cookie"
	"gin_demo/app/download"
	"gin_demo/app/upload"
	_ "gin_demo/config"
	_ "gin_demo/log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Input struct {
	Phone string `json:"phone" binding:"required,min=11,max=11"`
}

func main() {
	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})

	cookieApp := router.Group("api/cookie")
	cookie.RegisterRouter(cookieApp)

	uploadApp := router.Group("api/upload")
	upload.RegisterRouter(uploadApp)

	downloadApp := router.Group("api/download")
	download.RegisterRouter(downloadApp)

	router.Run(":8080")
}

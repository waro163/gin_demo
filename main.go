package main

import (
	"flag"
	"fmt"
	"gin_demo/app/cookie"
	"gin_demo/app/download"
	"gin_demo/app/method"
	"gin_demo/app/proxy"
	"gin_demo/app/upload"
	_ "gin_demo/config"

	// _ "gin_demo/log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var port uint64

func init() {
	flag.Uint64Var(&port, "Port", 8080, "web server listen port")
	flag.Parse()
}

func main() {
	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})
	methodApp := router.Group("api/method")
	method.RegisterRouter(methodApp)

	cookieApp := router.Group("api/cookie")
	cookie.RegisterRouter(cookieApp)

	uploadApp := router.Group("api/upload")
	upload.RegisterRouter(uploadApp)

	downloadApp := router.Group("api/download")
	download.RegisterRouter(downloadApp)

	proxyApp := router.Group("api/proxy")
	proxy.RegisterRouter(proxyApp)

	router.Run(fmt.Sprintf(":%d", port))
}

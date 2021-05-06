package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/ok", hello)

	router.Run(":8080")
}

func hello(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, map[string]string{"message": "ok"})
}

package main

import (
	"fmt"
	"gin_test/app_user"
	"log"
	"net/http"
	"time"

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
	router.POST("/upload", func(ctx *gin.Context) {
		// file, _ := ctx.FormFile("file")
		// ctx.SaveUploadedFile(file, "./"+file.Filename)
		// ctx.FileAttachment("./"+file.Filename, file.Filename)

		form, _ := ctx.MultipartForm()
		files := form.File["upload"]
		for _, file := range files {
			fmt.Println(file.Filename)
		}
		ctx.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	api := router.Group("/api")
	{
		api.Use(jwtMiddleware, authMiddleware)
		api.GET("/test", hello)
		api.POST("/login", login)
	}

	user := router.Group("/api/user/")
	app_user.RegisterRouter(user)

	router.Run(":8080")
}

func hello(ctx *gin.Context) {
	re := ctx.GetString("hi")
	log.Println("inside...", re)
	ctx.JSON(http.StatusOK, map[string]string{"message": "ok"})
}

func login(ctx *gin.Context) {
	var input Input
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"phone": input.Phone})
}

func jwtMiddleware(ctx *gin.Context) {
	t := time.Now()
	log.Println("jwt middleware before")
	ctx.Set("hi", "hello")
	// status := ctx.Writer.Status()
	// log.Println(status)
	latency := time.Since(t)
	log.Println(latency)
	ctx.Next()
	log.Println("jwt middleware after")
}

func authMiddleware(ctx *gin.Context) {
	log.Println("auth middleware before")
	ctx.Next()
	log.Println("auth middleware after")
}

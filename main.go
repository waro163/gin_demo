package main

import (
	"fmt"
	"gin_demo/app_user"
	_ "gin_demo/config"
	_ "gin_demo/log"
	"log"
	"net/http"
	"strings"
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

	router.GET("/cookie", func(ctx *gin.Context) {
		name, err := ctx.Cookie("name")
		if err != nil {
			ctx.SetCookie("name", "waro", 3600, "/cookie", "127.0.0.1", true, false)
			ctx.SetCookie("age", "18", 3600, "/", "127.0.0.1", false, true)
			ctx.SetCookie("phone", "18888888", 3600, "/", "127.0.0.1", true, true)
		}
		ctx.JSON(200, gin.H{
			"name": name,
		})
	})

	api := router.Group("/api")
	{
		api.Use(jwtMiddleware, authMiddleware)
		api.GET("/hello", hello)
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

func generate4096byte() string {
	var builder strings.Builder
	for i := 0; i < 4090; i++ {
		builder.WriteByte('a')
	}
	return builder.String()
}

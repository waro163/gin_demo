package upload

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func UploadDemo(ctx *gin.Context) {
	// file, err := ctx.FormFile("file")
	// if err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{"msg": err})
	// 	return
	// }
	// ctx.SaveUploadedFile(file, "./"+file.Filename)
	// ctx.JSON(http.StatusOK, gin.H{
	// 	"message": "ok",
	// })
	// return
	// ctx.FileAttachment("./"+file.Filename, file.Filename)

	form, err := ctx.MultipartForm()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}

	fmt.Println("---------form value---------")
	for fieldName, value := range form.Value {
		fmt.Println(fieldName, " : ", value)
	}

	fmt.Println("---------form file---------")
	for fieldName, files := range form.File {
		for _, file := range files {
			fmt.Println(fieldName, " : ", file.Filename)
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func UploadToThird(ctx *gin.Context) {
	url := "http://127.0.0.1:8080/api/upload/demo"
	req, err := http.NewRequest(ctx.Request.Method, url, ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	req.Header = ctx.Request.Header
	// transport := http.DefaultTransport
	// resp, err := transport.RoundTrip(req)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	// ctx.Data(res.StatusCode, res.Header.Get("Content-Type"), res.Body)
	ctx.DataFromReader(resp.StatusCode, resp.ContentLength, resp.Header.Get("Content-Type"), resp.Body, nil)
}

type FormData struct {
	IsFile     bool
	FieldName  string
	FileName   string
	FieldValue []byte
}

func UploadLocalFile(ctx *gin.Context) {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)

	formData1 := FormData{FieldName: "ttl", FieldValue: []byte("3600")}
	formData2 := FormData{IsFile: true, FieldName: "file", FileName: "./static/go-demo.png"}
	formDatas := []FormData{formData1, formData2}

	for _, formData := range formDatas {
		if formData.IsFile {
			w, err := writer.CreateFormFile(formData.FieldName, formData.FileName)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"msg": err})
				return
			}
			file, err := os.Open(formData.FileName)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"msg": err})
				return
			}
			defer file.Close()
			_, err = io.Copy(w, file)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"msg": err})
				return
			}
		} else {
			// err := writer.WriteField(formData.FieldName, string(formData.FieldValue))
			w, err := writer.CreateFormField(formData.FieldName)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"msg": err})
				return
			}
			_, err = w.Write(formData.FieldValue)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"msg": err})
				return
			}
		}
	}
	contentType := writer.FormDataContentType()
	writer.Close()

	req, err := http.NewRequest(http.MethodPost, "http://127.0.0.1:8080/api/upload/demo", payload)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	req.Header.Set("Content-Type", contentType)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	ctx.DataFromReader(resp.StatusCode, resp.ContentLength, resp.Header.Get("Content-Type"), resp.Body, nil)
}

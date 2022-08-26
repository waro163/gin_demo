package proxy

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProxyDemo(ctx *gin.Context) {
	url := "http://127.0.0.1:8081" + ctx.Request.RequestURI
	req, err := http.NewRequest(ctx.Request.Method, url, ctx.Request.Body)
	req.Header = ctx.Request.Header
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	transport := http.DefaultTransport //new(http.Transport)
	resp, err := transport.RoundTrip(req)
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, gin.H{"msg": err})
		return
	}
	contentType := resp.Header.Get("Content-Type")
	ctx.DataFromReader(resp.StatusCode, resp.ContentLength, contentType, resp.Body, nil)
	// ctx.Data(resp.StatusCode, contentType, resp.Body)

}

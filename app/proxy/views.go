package proxy

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProxyDemo(ctx *gin.Context) {
	req := ctx.Copy()
	req.Request.Host = "127.0.0.1:8080"
	transport := new(http.Transport)
	resp, err := transport.RoundTrip(req.Request)
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, gin.H{})
		return
	}
	contentType := resp.Header.Get("Content-Type")
	ctx.DataFromReader(resp.StatusCode, resp.ContentLength, contentType, resp.Body, nil)
	// ctx.Data(resp.StatusCode, contentType, resp.Body)

}

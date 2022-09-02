package proxy

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

const prefix = "/api/proxy"

func ProxyDemo(ctx *gin.Context) {
	baseUrl, _ := url.Parse("http://127.0.0.1:8080")
	path := strings.TrimPrefix(ctx.Request.RequestURI, prefix)
	url := baseUrl.ResolveReference(&url.URL{Path: path})

	req, err := http.NewRequest(ctx.Request.Method, url.String(), ctx.Request.Body)
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

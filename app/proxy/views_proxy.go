package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

const (
	ProxyHost = "http://127.0.0.1:8080/api/demo"
)

var (
	proxy *httputil.ReverseProxy
)

func init() {
	var err error
	// if we don't use init function, we should use sync.Once.Do to create
	proxy, err = genProxy()
	if err != nil {
		proxy = nil
	}
}

func genProxy() (*httputil.ReverseProxy, error) {
	originUrl, err := url.Parse(ProxyHost)
	if err != nil {
		return nil, err
	}
	// if ProxyHost contains path, we should remove it, else we should use targetUrl:=originUrl instead
	targetUrl := &url.URL{
		Scheme: originUrl.Scheme,
		Host:   originUrl.Host,
	}
	proxy := httputil.NewSingleHostReverseProxy(targetUrl)
	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)
		req.Host = targetUrl.Host
	}
	proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
		http.Error(w, "Service unavailable", http.StatusBadGateway)
	}
	return proxy, nil
}

func ProxyApi() gin.HandlerFunc {
	// proxy, err := genProxy()
	// if err != nil {
	// 	return func(ctx *gin.Context) {
	// 		ctx.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
	// 	}
	// }
	if proxy == nil {
		return func(ctx *gin.Context) {
			ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "internal error"})
		}
	}
	// if we use all proxyApi path
	// /http/proxy/*proxyPath, 这个代理逻辑是会代理所有的path;
	// 例如/http/proxy/abc/test, 会代理到http://127.0.0.1:8080/http/proxy/abc/test
	return gin.WrapH(proxy)
	// if we only use partial path;
	// /http/proxy/*proxyPath,这个代理逻辑是会去掉前缀部分:/http/proxy/,只代理proxyPath部分;
	// 例如例如/http/proxy/abc/test, 只会代理到http://127.0.0.1:8080/abc/test
	return func(ctx *gin.Context) {
		proxyPath := ctx.Param("proxyPath")
		ctx.Request.URL.Path = proxyPath
		proxy.ServeHTTP(ctx.Writer, ctx.Request)
	}
}

package proxy

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

const prefix = "/api/proxy/normal"

func ProxyDemo(ctx *gin.Context) {
	baseUrl, _ := url.Parse("http://127.0.0.1:8080")
	// Reqeust.RequestURI include query parameter, Request.URL.Path only include path
	path := strings.TrimPrefix(ctx.Request.URL.Path, prefix)

	// get query parameter
	values := ctx.Request.URL.Query()
	// add customer query parameter
	// values.Add("key", "value")
	url := baseUrl.ResolveReference(&url.URL{Path: path, RawQuery: values.Encode()})

	// Request Body as reader, for `not Get method` could be 412 status response
	// req, err := http.NewRequest(ctx.Request.Method, url.String(), ctx.Request.Body)
	// so replace bytes.NewReader function
	body, _ := ioutil.ReadAll(ctx.Request.Body)
	reader := bytes.NewReader(body)
	req, err := http.NewRequestWithContext(ctx.Request.Context(), ctx.Request.Method, url.String(), reader)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	// add query header
	req.Header = ctx.Request.Header

	// transport := http.DefaultTransport //new(http.Transport)
	// resp, err := transport.RoundTrip(req)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, gin.H{"msg": err})
		return
	}
	contentType := resp.Header.Get("Content-Type")
	ctx.DataFromReader(resp.StatusCode, resp.ContentLength, contentType, resp.Body, nil)
	// ctx.Data(resp.StatusCode, contentType, resp.Body)

}

func TracingDemo(ctx *gin.Context) {
	baseUrl, _ := url.Parse("http://127.0.0.1:8080")
	// Reqeust.RequestURI include query parameter, Request.URL.Path only include path
	path := strings.TrimPrefix(ctx.Request.URL.Path, "/api/proxy/tracing")

	// get query parameter
	values := ctx.Request.URL.Query()
	// add customer query parameter
	// values.Add("key", "value")
	url := baseUrl.ResolveReference(&url.URL{Path: path, RawQuery: values.Encode()})

	// Request Body as reader, for `not get method` could be 412 status response
	// req, err := http.NewRequest(ctx.Request.Method, url.String(), ctx.Request.Body)
	// so replace bytes.NewReader function
	body, _ := ioutil.ReadAll(ctx.Request.Body)
	reader := bytes.NewReader(body)
	req, err := http.NewRequestWithContext(ctx.Request.Context(), ctx.Request.Method, url.String(), reader)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	otel.GetTextMapPropagator().Inject(req.Context(), propagation.HeaderCarrier(req.Header))

	tr := otel.GetTracerProvider().Tracer("remote-call")
	_, span := tr.Start(req.Context(), "proxy")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		span.End()
		span.RecordError(err)
		ctx.JSON(http.StatusServiceUnavailable, gin.H{"msg": err})
		return
	}
	span.AddEvent(resp.Status)
	span.End()
	contentType := resp.Header.Get("Content-Type")
	ctx.DataFromReader(resp.StatusCode, resp.ContentLength, contentType, resp.Body, nil)
}

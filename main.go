package main

import (
	"context"
	"flag"
	"fmt"
	"gin_demo/app/cookie"
	"gin_demo/app/download"
	"gin_demo/app/method"
	"gin_demo/app/proxy"
	"gin_demo/app/sse"
	"gin_demo/app/upload"
	_ "gin_demo/config"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
)

var (
	port       uint64
	servername string
)

func init() {
	flag.Uint64Var(&port, "port", 8080, "web server listen port")
	flag.StringVar(&servername, "name", "gin_demo", "web server name")
	flag.Parse()
}

func main() {

	tp := initOteleMetry(servername)
	defer tp.Shutdown(context.Background())

	router := gin.Default()
	router.Use(cors.Default())
	router.Use(otelgin.Middleware("gin_demo"))

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

	sseApp := router.Group("api/sse")
	sse.RegisterRouter(sseApp)

	router.Run(fmt.Sprintf(":%d", port))
}

func initOteleMetry(name string) *tracesdk.TracerProvider {
	tp, err := tracerProvider("http://localhost:14268/api/traces", name)
	if err != nil {
		log.Fatal(err)
	}
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	return tp
}

func tracerProvider(url, name string) (*tracesdk.TracerProvider, error) {
	// Create the Jaeger exporter
	// using http directly send
	// exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))

	// using udp send agent, all setting use default value
	exp, err := jaeger.New(jaeger.WithAgentEndpoint())
	if err != nil {
		return nil, err
	}
	tp := tracesdk.NewTracerProvider(
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exp),
		// tracesdk.WithSampler(tracesdk.AlwaysSample()),
		tracesdk.WithSampler(tracesdk.ParentBased(tracesdk.TraceIDRatioBased(0.5))),
		// Record information about this application in a Resource.
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(name),
			attribute.String("environment", "dev"),
			// attribute.Int64("ID", id),
		)),
	)
	return tp, nil
}

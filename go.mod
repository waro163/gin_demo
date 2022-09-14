module gin_demo

go 1.16

require (
	github.com/gin-gonic/gin v1.8.1
	github.com/spf13/viper v1.11.0
	go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin v0.34.0
	go.opentelemetry.io/otel v1.10.0
	go.opentelemetry.io/otel/exporters/jaeger v1.10.0
	go.opentelemetry.io/otel/sdk v1.10.0
)

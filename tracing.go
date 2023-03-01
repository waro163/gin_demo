package main

import (
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/attribute"
	oteltrace "go.opentelemetry.io/otel/trace"
)

const (
	requestID = "X-Request-ID"
)

func AddCustomSpanOption(ctx *gin.Context) (sso []oteltrace.SpanStartOption) {
	var attrs []attribute.KeyValue
	rid := ctx.GetHeader(requestID)
	attrs = append(attrs, attribute.String(requestID, rid))
	attrs = append(attrs, attribute.StringSlice("tracing", []string{"demo", "test"}))

	sso = append(sso,
		oteltrace.WithAttributes(attrs...),
	)
	return
}

package main

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/openzipkin/zipkin-go/idgenerator"
	"go.opentelemetry.io/otel/trace"
)

const defaultTraceIDHeader = "x-trace-id"

func WithTraceID(name string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if name == "" {
			name = defaultTraceIDHeader
		}

		// 只拿traceID，供日志上下文使用
		if header := c.GetHeader(name); header != "" {
			traceID, _ := trace.TraceIDFromHex(strings.Split(header, ":")[0])
			c.Set(name, traceID.String())
		} else {
			gen := idgenerator.NewRandom128()
			traceID := gen.TraceID()
			c.Set(name, traceID.String())
		}
	}
}

func main() {
	r := gin.Default()
	r.Use(WithTraceID(defaultTraceIDHeader))
	r.GET("/", func(c *gin.Context) {
		traceID := c.GetString(defaultTraceIDHeader)
		c.String(200, "hello %s", traceID)
	})

	//nolint
	r.Run()
}

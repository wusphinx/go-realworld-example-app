package otel

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/openzipkin/zipkin-go/idgenerator"
	"go.opentelemetry.io/otel/trace"
)

const defaultTraceIDHeader = "x-trace-id"

func WithTraceID(agrs ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var name string
		if len(agrs) == 0 {
			name = defaultTraceIDHeader
		} else {
			name = agrs[0]
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

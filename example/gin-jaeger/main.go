package main

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/openzipkin/zipkin-go/idgenerator"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/trace"
)

const defaultTraceIDHeader = "x-trace-id"

type formatter struct {
	Formatter logrus.Formatter
}

func (l formatter) Format(entry *logrus.Entry) ([]byte, error) {
	if entry.Context != nil {
		traceID, _ := entry.Context.Value(defaultTraceIDHeader).(string)
		entry.Data[defaultTraceIDHeader] = traceID
	}

	return l.Formatter.Format(entry)
}

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
	logger := logrus.New()
	logger.SetFormatter(&formatter{Formatter: &logrus.JSONFormatter{}})

	r := gin.Default()
	r.Use(WithTraceID(defaultTraceIDHeader))
	r.GET("/", func(c *gin.Context) {
		traceID := c.GetString(defaultTraceIDHeader)
		logger.WithContext(c).Infof("hello world, traceID: %s", traceID)
		c.String(200, "hello %s", traceID)
	})

	//nolint
	r.Run()
}

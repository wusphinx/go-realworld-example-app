package main

import (
	"github.com/wusphinx/go-realworld-example-app/example/libs/otel"

	"github.com/gin-gonic/gin"
)

func main() {
	logger := otel.NewLogger()

	r := gin.Default()
	r.Use(otel.WithTraceID())
	r.GET("/", func(c *gin.Context) {
		logger.WithContext(c).Infof("hello world")
		c.String(200, "hello trace")
	})

	//nolint
	r.Run()
}

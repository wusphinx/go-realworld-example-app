package main

import (
	"github.com/wusphinx/go-realworld-example-app/example/gin-test/apis"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	apis.Register(route)
	//nolint
	route.Run()
}

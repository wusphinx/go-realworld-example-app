package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	r := gin.New()

	r.GET("/one", gin.WrapF(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `HandlerFunc Wrapper`)
	}))

	// http.Handler wrapper
	r.GET("/two", gin.WrapH(promhttp.Handler()))

	//nolint
	r.Run()
}

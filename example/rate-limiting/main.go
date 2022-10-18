package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

// refer: https://twitter.com/tebeka/status/1582252256379879426?s=20&t=H_VJD1BMP6HyaivYRWA1Vw
// 100/sec, burst of 200
var limiter = rate.NewLimiter(rate.Limit(100), 200)

func rated(w http.ResponseWriter, r *http.Request) {
	if !limiter.Allow() {
		status := http.StatusTooManyRequests
		http.Error(w, "too fast", status)
		return
	}

	fmt.Fprintln(w, "OK")
}

func main() {
	r := gin.New()

	r.GET("/rate", gin.WrapF(rated))

	//nolint
	r.Run()
}

package main

import (
	"fmt"
	"log"
	"net"

	"github.com/gin-gonic/gin"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
)

type Engine struct {
	port       int
	grpcServer *grpc.Server
	ginServer  *gin.Engine
}

//nolint:errcheck
func (e *Engine) Serve() {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", e.port))
	if err != nil {
		log.Fatalf("start server failed :%v", err)
	}

	m := cmux.New(l)
	grpcL := m.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))

	// grpc server
	go func() {
		if e.ginServer != nil {
			defer func() {
				if rec := recover(); rec != nil {
					log.Printf("grpc server failed :%v", rec)
				}
			}()

			e.grpcServer.Serve(grpcL)
		}
	}()

	// http server
	go func() {
		defer func() {
			if rec := recover(); rec != nil {
				log.Printf("http server failed :%v", rec)
			}
		}()

		httpL := m.Match(cmux.HTTP1Fast())
		e.ginServer.RunListener(httpL)
	}()

	m.Serve()
}

func main() {
	e := &Engine{
		port:       8080,
		grpcServer: grpc.NewServer(),
		ginServer:  gin.New(),
	}

	e.Serve()
}

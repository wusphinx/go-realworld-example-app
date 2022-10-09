package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/gin-gonic/gin"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

// copy from https://github.com/grpc/grpc-go/blob/master/examples/helloworld/greeter_server/main.go
// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

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
	grpcS := grpc.NewServer()
	pb.RegisterGreeterServer(grpcS, &server{})

	e := &Engine{
		port:       8080,
		grpcServer: grpcS,
		ginServer:  gin.New(),
	}

	e.Serve()
}

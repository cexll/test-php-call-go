package server

import (
	"context"
	"net/http"

	hello "service/api/hello/v1"
	"service/internal/config"
	"service/internal/service"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

func NewHttpServer(c *config.Config, srv *service.HelloService) *http.Server {
	httpServer := &http.Server{}
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	err := hello.RegisterHelloHandlerServer(ctx, mux, srv)
	if err != nil {
		panic(err)
	}
	httpServer.Addr = c.ApiConf.Addr
	httpServer.Handler = mux

	return httpServer
}

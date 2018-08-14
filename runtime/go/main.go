package main

import (
	"context"
	"errors"
	"log"
	"net"
	"os"
	"os/signal"
	"plugin"
	"syscall"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/socketfunc/faas/runtime/go/server"
	"github.com/socketfunc/faas/runtime/proto"
	"google.golang.org/grpc"
)

func main() {
	port := os.Getenv("_RUNTIME_PORT")

	handler, err := loadPlugin("handler.so", "Handler")
	if err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}

	srv := server.New(handler)
	s := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_recovery.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_recovery.UnaryServerInterceptor(),
		)),
	)

	runtime.RegisterRuntimeServer(s, srv)

	closed := make(chan struct{}, 1)
	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGTERM)
		<-sig

		log.Println("SIGTERM received... shutdown server")
		srv.Stop()
		s.GracefulStop()
		<-time.Tick(time.Duration(10) * time.Second)

		close(closed)
	}()

	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}

	<-closed
}

func loadPlugin(codePath, entryPoint string) (server.Handler, error) {
	p, err := plugin.Open(codePath)
	if err != nil {
		return nil, err
	}
	sym, err := p.Lookup(entryPoint)
	if err != nil {
		return nil, err
	}
	if handler, ok := sym.(func(ctx context.Context, req server.Request, res server.Response)); ok {
		return handler, nil
	}
	return nil, errors.New("Entry Point is not found: bad type")
}

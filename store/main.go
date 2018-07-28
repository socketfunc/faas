package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/socketfunc/faas/store/kvs"
	"github.com/socketfunc/faas/store/proto"
)

func main() {
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_recovery.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_recovery.UnaryServerInterceptor(),
		)),
	)
	store.RegisterKvsServer(s, kvs.New())
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

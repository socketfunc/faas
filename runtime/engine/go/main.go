package main

import (
	"context"
	"errors"
	"log"
	"net"
	"os"
	"plugin"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/socketfunc/faas/runtime/engine/go/handler"
	pb "github.com/socketfunc/faas/runtime/proto"
	"google.golang.org/grpc"
)

type Handler func(*handler.Request, *handler.Response) error

type service struct {
	handler Handler
}

var _ pb.InternalServer = (*service)(nil)

func (s *service) Pipe(in *pb.PipeRequest, stream pb.Internal_PipeServer) error {
	req := handler.NewRequest()
	res := handler.NewResponse(stream)
	if err := s.handler(req, res); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (s *service) Specialize(ctx context.Context, in *pb.SpecializeRequest) (*pb.SpecializeResponse, error) {
	var err error
	s.handler, err = loadPlugin(in.CodePath, in.EntryPoint)
	if err != nil {
		return nil, err
	}
	return &pb.SpecializeResponse{}, nil
}

func (s *service) Healthz(ctx context.Context, _ *pb.Empty) (*pb.HealthzResponse, error) {
	return &pb.HealthzResponse{
		StatusCode: 1,
	}, nil
}

func main() {
	port := os.Getenv("_RUNTIME_PORT")
	lis, err := net.Listen("tcp", "localhost:"+port)
	if err != nil {
		log.Fatal(err)
	}

	h, err := loadPlugin("handler.so", "Handler")
	if err != nil {
		log.Fatal(err)
	}
	srv := &service{
		handler: h,
	}

	s := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_recovery.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_recovery.UnaryServerInterceptor(),
		)),
	)
	pb.RegisterInternalServer(s, srv)
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

func loadPlugin(codePath, entryPoint string) (Handler, error) {
	p, err := plugin.Open(codePath)
	if err != nil {
		return nil, err
	}
	sym, err := p.Lookup(entryPoint)
	if err != nil {
		return nil, err
	}
	if fn, ok := sym.(func(*handler.Request, *handler.Response) error); ok {
		return fn, err
	}
	return nil, errors.New("Entry Point is not found: bad type")
}

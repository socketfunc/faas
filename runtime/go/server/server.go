package server

import (
	"context"
	"fmt"
	"time"

	"github.com/socketfunc/faas/runtime/go/store"
	pb "github.com/socketfunc/faas/runtime/proto"
)

var _ pb.RuntimeServer = (*server)(nil)

type Handler func(ctx context.Context, req Request, res Response)

type server struct {
	handler Handler
	status  pb.HealthCheckResponse_ServingStatus
}

func (s *server) Stream(stream pb.Runtime_StreamServer) error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("%+v\n", err)
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	sc := &store.Client{
		Stream: stream,
	}
	ctx = context.WithValue(ctx, store.CtxKey, sc)

	in, err := stream.Recv()
	if err != nil {
		return err
	}

	// TODO: timeout cancel
	timer := time.NewTimer(time.Duration(30) * time.Second)
	go func() {
		<-timer.C
		cancel()
	}()

	if in.Cmd == pb.Cmd_STREAM {
		req := newRequest(in.StreamRequest.Topic, in.StreamRequest.Event, in.StreamRequest.Payload)
		res := newResponse(in.StreamRequest.Topic, in.StreamRequest.Event, stream)
		s.handler(ctx, req, res)
	}

	timer.Stop()

	return nil
}

func (s *server) HealthCheck(ctx context.Context, req *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
	res := &pb.HealthCheckResponse{
		Status: s.status,
	}
	return res, nil
}

func (s *server) Stop() {
	s.status = pb.HealthCheckResponse_NOT_SERVING
}

func New(handler Handler) *server {
	return &server{
		handler: handler,
		status:  pb.HealthCheckResponse_SERVING,
	}
}

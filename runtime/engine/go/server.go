package runtime

import (
	"context"
	"net"
	"os"

	"github.com/socketfunc/faas/runtime/engine/go/store"
	pb "github.com/socketfunc/faas/runtime/proto"
	"google.golang.org/grpc"
)

type Handler func(ctx context.Context, req *Request, res Response)

type Request struct {
	topic   string
	event   string
	payload []byte
}

func (r *Request) Topic() string {
	return r.topic
}

func (r *Request) Event() string {
	return r.event
}

type Response interface {
	Send(topic, event string, payload []byte) error
	Reply()
	Broadcast()
}

type response struct {
	stream pb.Runtime_StreamServer
}

func (r *response) Send(topic, event string, payload []byte) error {
	send := &pb.Send{
		Cmd: pb.Cmd_Stream,
		StreamSend: &pb.StreamSend{
			Topic:   topic,
			Event:   event,
			Payload: payload,
		},
	}
	if err := r.stream.Send(send); err != nil {
		return err
	}
	return nil
}

func (r *response) Reply() {}

func (r *response) Broadcast() {}

type server struct {
	handler Handler
}

var _ pb.RuntimeServer = (*server)(nil)

func (s *server) Stream(stream pb.Runtime_StreamServer) error {
	ctx := context.Background()
	storeClient := &store.Client{
		Stream: stream,
	}
	ctx = context.WithValue(ctx, store.CtxKey, storeClient)
	in, err := stream.Recv()
	if err != nil {
		return err
	}
	if in.Cmd == pb.Cmd_Stream {
		req := &Request{}
		res := &response{
			stream: stream,
		}
		s.handler(ctx, req, res)
	}
	return nil
}

func Start(handler Handler) error {
	port := os.Getenv("_RUNTIME_PORT")
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}
	s := grpc.NewServer()
	srv := &server{
		handler: handler,
	}
	pb.RegisterRuntimeServer(s, srv)
	if err := s.Serve(lis); err != nil {
		return err
	}
	return nil
}

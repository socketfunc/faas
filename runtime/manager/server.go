package manager

import (
	"context"
	"io"
	"log"
	"strconv"
	"time"

	"github.com/socketfunc/faas/proto"
	"github.com/socketfunc/faas/runtime/proto"
	"google.golang.org/grpc"
)

type server struct {
	port int
}

var _ faas.RuntimeServer = (*server)(nil)

func (s *server) Stream(in *faas.StreamRequest, stream faas.Runtime_StreamServer) error {
	conn, err := grpc.Dial("localhost:"+strconv.Itoa(s.port), grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	client := runtime.NewInternalClient(conn)

	req := &runtime.PipeRequest{
		Version: in.Version,
		Id:      in.Id,
		Packet: &runtime.Packet{
			Topic:   in.Packet.Topic,
			Event:   in.Packet.Event,
			Id:      in.Packet.Id,
			Payload: in.Packet.Payload,
		},
	}

	ctx, cancel := context.WithCancel(context.Background())
	pipe, err := client.Pipe(ctx, req)
	if err != nil {
		log.Printf("Err client: %v", err)
		return err
	}

	timer := time.NewTimer(time.Duration(60) * time.Second)
	go func() {
		<-timer.C
		log.Println("timeout")
		cancel()
	}()
	defer timer.Stop()

	for {
		res, err := pipe.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(err)
			return err
		}

		out := &faas.StreamResponse{
			Version: res.Version,
			Cmd:     faas.Cmd(res.Cmd),
			Packet: &faas.Packet{
				Topic:   res.Packet.Topic,
				Event:   res.Packet.Event,
				Id:      res.Packet.Id,
				Payload: res.Packet.Payload,
			},
		}
		if err := stream.Send(out); err != nil {
			return err
		}
	}

	return nil
}

func New(port int) *server {
	return &server{
		port: port,
	}
}

package main

import (
	"context"
	"log"

	"fmt"
	"io"

	"github.com/socketfunc/faas/runtime/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	client := runtime.NewRuntimeClient(conn)
	stream, err := client.Stream(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan struct{})

	go receive(done, stream)

	rec := &runtime.Receive{
		Cmd: runtime.Cmd_Stream,
		StreamRequest: &runtime.StreamRequest{
			Topic:   "topic",
			Event:   "event",
			Payload: []byte("message test"),
		},
	}
	if err := stream.Send(rec); err != nil {
		log.Fatal(err)
	}

	<-done
}

func receive(done chan struct{}, stream runtime.Runtime_StreamClient) {
	resp, err := stream.Recv()
	fmt.Println(resp, err)
	if err == io.EOF {
		close(done)
		return
	}
	if err != nil {
		log.Fatal(err)
	}
}

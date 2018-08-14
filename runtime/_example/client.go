package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/socketfunc/faas/runtime/proto"
	"github.com/socketfunc/faas/store/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
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
		Cmd: runtime.Cmd_STREAM,
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
	for {
		resp, err := stream.Recv()
		fmt.Println(resp, err)
		if err == io.EOF {
			close(done)
			return
		}
		if err != nil {
			log.Fatal(err)
		}
		if resp.Cmd == runtime.Cmd_STORE {
			switch resp.StoreRequest.Cmd {
			case runtime.Store_Cmd_GET:
				// time.Sleep(time.Duration(10) * time.Second)
				recv := &runtime.Receive{
					Cmd: runtime.Cmd_STORE,
					StoreResponse: &runtime.StoreResponse{
						Cmd: runtime.Store_Cmd_GET,
						Entity: &store.Entity{
							Data: map[string]*store.Value{
								"id": {
									Type: store.Type_String,
									Data: []byte("test"),
								},
								"no": {
									Type: store.Type_Int,
									Data: []byte{0, 0, 0, 0, 0, 0, 0, 100},
								},
							},
						},
					},
				}
				stream.Send(recv)
			case runtime.Store_Cmd_PUT:
			case runtime.Store_Cmd_MODIFY:
			case runtime.Store_Cmd_DEL:
			}
		}
	}
}

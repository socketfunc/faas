package main

import (
	"context"
	"log"

	"encoding/binary"
	"io"

	"github.com/socketfunc/faas/store/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:9090", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := store.NewKvsClient(conn)

	stream, err := client.Command(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan struct{})

	go receive(done, stream)

	for i := 0; i < 10; i++ {
		num := int64(i + 1)
		data := make([]byte, 8)
		binary.BigEndian.PutUint64(data, uint64(num))
		cmd := &store.CommandRequest{
			Command: store.Command_Modify,
			Modify: &store.ModifyRequest{
				Namespace: "ns",
				Key:       "k",
				Updates: []*store.Update{
					{
						Op:   store.Op_Inc,
						Name: "field",
						Value: &store.Value{
							Type: store.Type_Int,
							Data: data,
						},
					},
				},
			},
		}
		if err := stream.Send(cmd); err != nil {
			log.Fatal(err)
		}
	}

	<-done
}

func receive(done chan struct{}, stream store.Kvs_CommandClient) {
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			close(done)
			return
		}
		if err != nil {
			log.Fatal(err)
		}
		log.Println(int64(binary.BigEndian.Uint64(resp.Modify.Entity.Data["field"].Data)))
	}
}

package main

import (
	"context"
	"fmt"

	"github.com/socketfunc/faas/runtime/go/server"
	"github.com/socketfunc/faas/runtime/go/store"
)

type Value struct {
	ID string `store:"id"`
	No int    `store:"no"`
}

func Handler(ctx context.Context, req server.Request, res server.Response) {
	fmt.Println("handler")

	value := &Value{}
	err := store.Get(ctx, "key", value)
	fmt.Println(value, err)

	fmt.Println(req.Topic(), req.Event(), string(req.Payload()))

	value = &Value{}
	err = store.Get(ctx, "key", value)
	fmt.Println(value, err)

	res.Send("topic1", "event1", []byte("message1"))

	res.Send("topic2", "event2", []byte("message2"))
}

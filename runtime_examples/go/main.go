package main

import (
	"fmt"
	"time"

	"github.com/socketfunc/faas/runtime/engine/go/handler"
)

func Handler(req *handler.Request, res *handler.Response) error {
	fmt.Println("handler plugin")
	res.Send("topic", "event", 1, []byte("message 01"))
	time.Sleep(time.Second)
	panic("panic go plugin")
	res.Send("topic", "event", 1, []byte("message 02"))
	time.Sleep(time.Second)
	res.Send("topic", "event", 1, []byte("message 03"))
	return nil
}

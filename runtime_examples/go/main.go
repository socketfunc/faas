package main

import (
	"context"

	"github.com/socketfunc/faas/runtime/engine/go"
)

func handler(ctx context.Context, req *runtime.Request, res runtime.Response) {
	res.Send("topic", "event", []byte("message data"))
}

func main() {
	runtime.Start(handler)
}

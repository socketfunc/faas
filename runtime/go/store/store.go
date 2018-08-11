package store

import (
	"context"
	"errors"
	"fmt"

	pb "github.com/socketfunc/faas/runtime/proto"
)

const (
	CtxKey = 0
)

type Client struct {
	Stream pb.Runtime_StreamServer
}

type Comp int32

const (
	Eq Comp = iota
	Gt
	Gte
	Lt
	Lte
	Ne
	Exists
)

type Filter struct {
	Comp  Comp
	Name  string
	value interface{}
}

type Op int32

const (
	Inc Op = iota
	Set
	Unset
)

type Update struct {
	Op    Op
	Name  string
	Value interface{}
}

func Get(ctx context.Context, key string, value interface{}) error {
	client, ok := ctx.Value(CtxKey).(*Client)
	if !ok {
		return errors.New("cannot be client")
	}

	req := &pb.Send{
		Cmd: pb.Cmd_STORE,
		StoreRequest: &pb.StoreRequest{
			Cmd: pb.Store_Cmd_GET,
			Key: key,
		},
	}

	if err := client.Stream.Send(req); err != nil {
		return err
	}
	res, err := client.Stream.Recv()
	if err != nil {
		return err
	}

	fmt.Println(res.StoreResponse.Entity)

	return nil
}

func Put(ctx context.Context, key string, value interface{}) error {
	client, ok := ctx.Value(CtxKey).(*Client)
	if !ok {
		return errors.New("cannot be client")
	}

	req := &pb.Send{
		Cmd: pb.Cmd_STORE,
		StoreRequest: &pb.StoreRequest{
			Cmd: pb.Store_Cmd_PUT,
			Key: key,
		},
	}

	if err := client.Stream.Send(req); err != nil {
		return err
	}
	res, err := client.Stream.Recv()
	if err != nil {
		return err
	}

	fmt.Println(res.StoreResponse.Successful)

	return nil
}

func Modify(ctx context.Context, key string, filters []*Filter, updates []*Update, value interface{}) (bool, error) {
	return false, nil
}

func Del(ctx context.Context, key string) error {
	return nil
}

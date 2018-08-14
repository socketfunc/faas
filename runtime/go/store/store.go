package store

import (
	"context"
	"errors"
	"fmt"

	pb "github.com/socketfunc/faas/runtime/proto"
	"github.com/socketfunc/faas/store/proto"
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
	Value interface{}
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
		return errors.New("cannot client")
	}

	req := &pb.Send{
		Cmd: pb.Cmd_STORE,
		StoreRequest: &pb.StoreRequest{
			Cmd: pb.Store_Cmd_GET,
			Key: key,
		},
	}

	done := make(chan struct{}, 1)

	var res *pb.Receive
	var err error
	go func() {
		defer close(done)
		res, err = client.Stream.Recv()
		fmt.Println("store get finish")
	}()

	if err := client.Stream.Send(req); err != nil {
		return err
	}

	select {
	case <-done:
		if err != nil {
			return err
		}
		if err := decodeEntity(res.StoreResponse.Entity, value); err != nil {
			return err
		}
	case <-ctx.Done():
		return errors.New("timeout")
	}

	return nil
}

func Put(ctx context.Context, key string, value interface{}) error {
	client, ok := ctx.Value(CtxKey).(*Client)
	if !ok {
		return errors.New("cannot client")
	}

	entity, err := encodeEntity(value)
	if err != nil {
		return err
	}

	req := &pb.Send{
		Cmd: pb.Cmd_STORE,
		StoreRequest: &pb.StoreRequest{
			Cmd:    pb.Store_Cmd_PUT,
			Key:    key,
			Entity: entity,
		},
	}

	done := make(chan struct{}, 1)

	var res *pb.Receive
	go func() {
		defer close(done)
		res, err = client.Stream.Recv()
	}()

	if err := client.Stream.Send(req); err != nil {
		return err
	}

	select {
	case <-done:
		if err != nil {
			return err
		}
		if !res.StoreResponse.Successful {
			return errors.New("")
		}
	case <-ctx.Done():
		return errors.New("timeout")
	}

	return nil
}

func Modify(ctx context.Context, key string, filters []*Filter, updates []*Update, value interface{}) (bool, error) {
	client, ok := ctx.Value(CtxKey).(*Client)
	if !ok {
		return false, errors.New("cannot client")
	}

	fl := make([]*store.Filter, 0, len(filters))
	for _, filter := range filters {
		value, err := encodeValue(filter.Value)
		if err != nil {
			return false, err
		}
		fl = append(fl, &store.Filter{
			Comp:  store.Comp(filter.Comp),
			Name:  filter.Name,
			Value: value,
		})
	}
	up := make([]*store.Update, 0, len(updates))
	for _, update := range updates {
		value, err := encodeValue(update.Value)
		if err != nil {
			return false, err
		}
		up = append(up, &store.Update{
			Op:    store.Op(update.Op),
			Name:  update.Name,
			Value: value,
		})
	}

	req := &pb.Send{
		Cmd: pb.Cmd_STORE,
		StoreRequest: &pb.StoreRequest{
			Cmd:     pb.Store_Cmd_MODIFY,
			Key:     key,
			Filters: fl,
			Updates: up,
		},
	}

	done := make(chan struct{}, 1)

	var res *pb.Receive
	var err error
	go func() {
		defer close(done)
		res, err = client.Stream.Recv()
	}()

	if err := client.Stream.Send(req); err != nil {
		return false, err
	}

	select {
	case <-done:
		if err != nil {
			return false, err
		}
		if err := decodeEntity(res.StoreResponse.Entity, value); err != nil {
			return false, err
		}
	case <-ctx.Done():
		return false, errors.New("timeout")
	}

	return res.StoreResponse.Successful, nil
}

func Del(ctx context.Context, key string) error {
	client, ok := ctx.Value(CtxKey).(*Client)
	if !ok {
		return errors.New("cannot client")
	}

	req := &pb.Send{
		Cmd: pb.Cmd_STORE,
		StoreRequest: &pb.StoreRequest{
			Cmd: pb.Store_Cmd_DEL,
			Key: key,
		},
	}

	done := make(chan struct{}, 1)

	var res *pb.Receive
	var err error
	go func() {
		defer close(done)
		res, err = client.Stream.Recv()
	}()

	if err := client.Stream.Send(req); err != nil {
		return err
	}

	select {
	case <-done:
		if err != nil {
			return err
		}
	case <-ctx.Done():
		return errors.New("timeout")
	}

	return nil
}

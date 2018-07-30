package client

import (
	"bytes"
	"context"
	"encoding/binary"
	"errors"
	"math"
	"reflect"

	pb "github.com/socketfunc/faas/store/proto"
	"google.golang.org/grpc"
)

const (
	tagKey = "store"
)

func encodeEntity(value interface{}) (*pb.Entity, error) {
	rv := reflect.ValueOf(value)

	if rv.Kind() == reflect.Ptr && !rv.IsNil() {
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Struct {
		return nil, errors.New("invalid")
	}

	entity := &pb.Entity{
		Data: map[string]*pb.Value{},
	}

	for i := 0; i < rv.NumField(); i++ {
		val := rv.Field(i)
		field := rv.Type().Field(i)
		key := field.Tag.Get(tagKey)

		value := &pb.Value{}
		switch val.Kind() {
		case reflect.String:
			value.Type = pb.Type_String
			value.Data = []byte(val.String())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			data := make([]byte, 8)
			binary.BigEndian.PutUint64(data, uint64(val.Int()))
			value.Type = pb.Type_Int
			value.Data = data
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			data := make([]byte, 8)
			binary.BigEndian.PutUint64(data, val.Uint())
			value.Type = pb.Type_Uint
			value.Data = data
		case reflect.Float32, reflect.Float64:
			data := make([]byte, 8)
			binary.BigEndian.PutUint64(data, math.Float64bits(val.Float()))
			value.Type = pb.Type_Float
			value.Data = data
		case reflect.Bool:
			value.Type = pb.Type_Bool
			if val.Bool() {
				value.Data = []byte{1}
			} else {
				value.Data = []byte{0}
			}
		default:
		}
		entity.Data[key] = value
	}

	return entity, nil
}

func decodeEntity(entity *pb.Entity, value interface{}) error {
	rv := reflect.ValueOf(value)

	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return errors.New("invalid")
	}
	rv = rv.Elem()

	// create struct mapping
	mapping := map[string]int{}
	for i := 0; i < rv.NumField(); i++ {
		field := rv.Type().Field(i)
		key := field.Tag.Get(tagKey)
		mapping[key] = i
	}

	for key, val := range entity.Data {
		idx, ok := mapping[key]
		if !ok {
			continue
		}
		field := rv.Field(idx)
		switch val.Type {
		case pb.Type_String:
			field.SetString(string(val.Data))
		case pb.Type_Int:
			x := binary.BigEndian.Uint64(val.Data)
			field.SetInt(int64(x))
		case pb.Type_Uint:
			x := binary.BigEndian.Uint64(val.Data)
			field.SetUint(x)
		case pb.Type_Float:
			bits := binary.BigEndian.Uint64(val.Data)
			x := math.Float64frombits(bits)
			field.SetFloat(x)
		case pb.Type_Bool:
			field.SetBool(bytes.Equal(val.Data, []byte{1}))
		default:
		}
	}

	return nil
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

func (c Comp) Convert() pb.Comp {
	switch c {
	case Eq:
		return pb.Comp_Eq
	case Gt:
		return pb.Comp_Gt
	case Gte:
		return pb.Comp_Gte
	case Lt:
		return pb.Comp_Lt
	case Lte:
		return pb.Comp_Lte
	case Ne:
		return pb.Comp_Ne
	case Exists:
		return pb.Comp_Exists
	}
	return -1
}

type Filter struct {
	Comp  Comp
	Name  string
	Value interface{}
}

func encodeFilters(filters []*Filter) []*pb.Filter {
	values := make([]*pb.Filter, 0, len(filters))
	for _, filter := range filters {
		rv := reflect.ValueOf(filter.Value)
		switch rv.Kind() {
		case reflect.String:
			values = append(values, &pb.Filter{
				Comp: filter.Comp.Convert(),
				Name: filter.Name,
				Value: &pb.Value{
					Type: pb.Type_String,
					Data: []byte(rv.String()),
				},
			})
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			data := make([]byte, 8)
			binary.BigEndian.PutUint64(data, uint64(rv.Int()))
			values = append(values, &pb.Filter{
				Comp: filter.Comp.Convert(),
				Name: filter.Name,
				Value: &pb.Value{
					Type: pb.Type_Int,
					Data: data,
				},
			})
		case reflect.Float32, reflect.Float64:
			data := make([]byte, 8)
			binary.BigEndian.PutUint64(data, uint64(math.Float64bits(rv.Float())))
			values = append(values, &pb.Filter{
				Comp: filter.Comp.Convert(),
				Name: filter.Name,
				Value: &pb.Value{
					Type: pb.Type_Float,
					Data: data,
				},
			})
		case reflect.Bool:
			if rv.Bool() {
				values = append(values, &pb.Filter{
					Comp: filter.Comp.Convert(),
					Name: filter.Name,
					Value: &pb.Value{
						Type: pb.Type_Bool,
						Data: []byte{1},
					},
				})
			} else {
				values = append(values, &pb.Filter{
					Comp: filter.Comp.Convert(),
					Name: filter.Name,
					Value: &pb.Value{
						Type: pb.Type_Bool,
						Data: []byte{0},
					},
				})
			}
		default:
		}
	}
	return values
}

type Op int32

const (
	Inc Op = iota
	Set
	Unset
	CurrentDate
	Min
	Max
	Mul
	SetOrInsert
)

func (o Op) Convert() pb.Op {
	switch o {
	case Inc:
		return pb.Op_Inc
	case Set:
		return pb.Op_Set
	case Unset:
		return pb.Op_Unset
	case CurrentDate:
		return pb.Op_CurrentDate
	case Min:
		return pb.Op_Min
	case Max:
		return pb.Op_Max
	case Mul:
		return pb.Op_Mul
	case SetOrInsert:
		return pb.Op_SetOnInsert
	}
	return -1
}

type Update struct {
	Op    Op
	Name  string
	Value interface{}
}

func encodeUpdates(updates []*Update) []*pb.Update {
	values := make([]*pb.Update, 0, len(updates))
	for _, update := range updates {
		rv := reflect.ValueOf(update.Value)
		switch rv.Kind() {
		case reflect.String:
			values = append(values, &pb.Update{
				Op:   update.Op.Convert(),
				Name: update.Name,
				Value: &pb.Value{
					Type: pb.Type_String,
					Data: []byte(rv.String()),
				},
			})
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			data := make([]byte, 8)
			binary.BigEndian.PutUint64(data, uint64(rv.Int()))
			values = append(values, &pb.Update{
				Op:   update.Op.Convert(),
				Name: update.Name,
				Value: &pb.Value{
					Type: pb.Type_Int,
					Data: data,
				},
			})
		case reflect.Float32, reflect.Float64:
			data := make([]byte, 8)
			binary.BigEndian.PutUint64(data, math.Float64bits(rv.Float()))
			values = append(values, &pb.Update{
				Op:   update.Op.Convert(),
				Name: update.Name,
				Value: &pb.Value{
					Type: pb.Type_Float,
					Data: data,
				},
			})
		case reflect.Bool:
			if rv.Bool() {
				values = append(values, &pb.Update{
					Op:   update.Op.Convert(),
					Name: update.Name,
					Value: &pb.Value{
						Type: pb.Type_Bool,
						Data: []byte{1},
					},
				})
			} else {
				values = append(values, &pb.Update{
					Op:   update.Op.Convert(),
					Name: update.Name,
					Value: &pb.Value{
						Type: pb.Type_Bool,
						Data: []byte{0},
					},
				})
			}
		default:
		}
	}
	return values
}

type storeClient struct {
	client pb.KvsClient
	Ctx    context.Context
}

func (sc *storeClient) Get(namespace, key string, value interface{}) error {
	req := &pb.GetRequest{
		Namespace: namespace,
		Key:       key,
	}
	res, err := sc.client.Get(sc.Ctx, req)
	if err != nil {
		return err
	}
	if err := decodeEntity(res.Entity, value); err != nil {
		return err
	}
	return nil
}

func (sc *storeClient) Put(namespace, key string, value interface{}) error {
	entity, err := encodeEntity(value)
	if err != nil {
		return err
	}
	req := &pb.PutRequest{
		Namespace: namespace,
		Key:       key,
		Entity:    entity,
	}
	res, err := sc.client.Put(sc.Ctx, req)
	if err != nil {
		return err
	}
	if !res.Successful {
		return errors.New("") // TODO...
	}
	return nil
}

func (sc *storeClient) Modify(namespace, key string, filters []*Filter, updates []*Update, value interface{}) (bool, error) {
	req := &pb.ModifyRequest{
		Namespace: namespace,
		Key:       key,
		Filters:   encodeFilters(filters),
		Updates:   encodeUpdates(updates),
	}
	res, err := sc.client.Modify(sc.Ctx, req)
	if err != nil {
		return false, err
	}
	if err := decodeEntity(res.Entity, value); err != nil {
		return res.Successful, err
	}
	return res.Successful, nil
}

func (sc *storeClient) Del(namespace, key string) error {
	req := &pb.DelRequest{
		Namespace: namespace,
		Key:       key,
	}
	_, err := sc.client.Del(context.Background(), req)
	if err != nil {
		return err
	}
	return nil
}

func (sc *storeClient) DelAll(namespace string) error {
	req := &pb.DelRequest{
		Namespace: namespace,
	}
	_, err := sc.client.DelAll(context.Background(), req)
	if err != nil {
		return err
	}
	return nil
}

func (sc *storeClient) Keys() ([]string, error) {
	req := &pb.KeysRequest{}
	res, err := sc.client.Keys(sc.Ctx, req)
	if err != nil {
		return nil, err
	}
	return res.Keys, nil
}

func New(addr string) (*storeClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := pb.NewKvsClient(conn)
	sc := &storeClient{
		client: client,
		Ctx:    context.Background(),
	}
	return sc, nil
}

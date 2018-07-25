package kvs

import (
	"bytes"
	"encoding/binary"
	"log"
	"math"
	"strings"
	"sync"

	pb "github.com/socketfunc/faas/store/proto"
)

type Entity map[string]*pb.Value

func (e Entity) hasField(name string) bool {
	_, ok := e[name]
	return ok
}

func (e Entity) getValue(name string) *pb.Value {
	return e[name]
}

func (e Entity) setValue(name string, value *pb.Value) {
	e[name] = value
}

func (e Entity) delValue(name string) {
	delete(e, name)
}

type Kvs struct {
	entities sync.Map
}

func (kvs *Kvs) getEntity(namespace, key string) (*Entity, error) {
	entity, ok := kvs.entities.Load(namespace + ":" + key)
	if !ok || entity == nil {
		return nil, ErrNotFound
	}
	return entity.(*Entity), nil
}

func (kvs *Kvs) setEntity(namespace, key string, entity *Entity) {
	kvs.entities.Store(namespace+":"+key, entity)
}

func (kvs *Kvs) delEntity(namespace, key string) {
	kvs.entities.Delete(namespace + ":" + key)
}

func (kvs *Kvs) Get(namespace, key string) (*Entity, error) {
	return kvs.getEntity(namespace, key)
}

func (kvs *Kvs) Put(namespace, key string, filters []*pb.Filter, updates []*pb.Update) (entity *Entity, err error) {
	entity, err = kvs.Get(namespace, key)
	if err != nil && err != ErrNotFound {
		return
	}
	if entity == nil {
		entity = &Entity{}
		err = nil
	}

	for _, filter := range filters {
		if filter == nil {
			continue
		}

		value := entity.getValue(filter.Name)
		if value == nil || value.Type != filter.Value.Type {
			return
		}
		switch filter.Comp {
		case pb.Comp_Eq:
			if !bytes.Equal(value.Data, filter.Value.Data) {
				return
			}
		case pb.Comp_Gt:
			if !(bytes.Compare(value.Data, filter.Value.Data) > 0) {
				return
			}
		case pb.Comp_Gte:
			if !(bytes.Compare(value.Data, filter.Value.Data) >= 0) {
				return
			}
		case pb.Comp_Lt:
			if !(bytes.Compare(value.Data, filter.Value.Data) < 0) {
				return
			}
		case pb.Comp_Lte:
			if !(bytes.Compare(value.Data, filter.Value.Data) <= 0) {
				return
			}
		case pb.Comp_Ne:
			if bytes.Equal(value.Data, filter.Value.Data) {
				return
			}
		case pb.Comp_Exists:
			if !entity.hasField(filter.Name) {
				return
			}
		default:
		}
	}

	for _, update := range updates {
		if update == nil {
			continue
		}

		switch update.Op {
		case pb.Op_Inc:
			switch update.Value.Type {
			case pb.Type_Int: // int64
				value := entity.getValue(update.Name)
				data := make([]byte, 8)
				addNum := binary.BigEndian.Uint64(update.Value.Data)
				if value == nil {
					binary.BigEndian.PutUint64(data, addNum)
				} else {
					num := binary.BigEndian.Uint64(value.Data)
					binary.BigEndian.PutUint64(data, num+addNum)
				}
				entity.setValue(update.Name, &pb.Value{
					Type: pb.Type_Int,
					Data: data,
				})
			case pb.Type_Float: // float64
				value := entity.getValue(update.Name)
				data := make([]byte, 8)
				addNum := float64(binary.BigEndian.Uint64(update.Value.Data))
				if value == nil {
					binary.BigEndian.PutUint64(data, math.Float64bits(addNum))
				} else {
					num := float64(binary.BigEndian.Uint64(value.Data))
					binary.BigEndian.PutUint64(data, math.Float64bits(num+addNum))
				}
				entity.setValue(update.Name, &pb.Value{
					Type: pb.Type_Float,
					Data: data,
				})
			default:
				err = ErrInvalidType
				return
			}
		case pb.Op_Set:
			entity.setValue(update.Name, update.Value)
		case pb.Op_Unset:
			entity.delValue(update.Name)
		default:
			err = ErrInvalidType
			return
		}
	}

	kvs.setEntity(namespace, key, entity)

	return
}

func (kvs *Kvs) Del(namespace, key string) {
	kvs.delEntity(namespace, key)
}

func (kvs *Kvs) DelAll(namespace string) {
	kvs.entities.Range(func(key, value interface{}) bool {
		k, ok := key.(string)
		if !ok {
			log.Println("invalid key")
			return true
		}
		tmp := strings.SplitN(k, ":", 2)
		if tmp[0] == namespace {
			kvs.Del(tmp[0], tmp[1])
		}
		return true
	})
}

func (kvs *Kvs) Keys() []string {
	var keys []string
	kvs.entities.Range(func(key, value interface{}) bool {
		k, ok := key.(string)
		if ok {
			keys = append(keys, k)
		}
		return true
	})
	return keys
}

func NewKvs() *Kvs {
	return &Kvs{
		entities: sync.Map{},
	}
}

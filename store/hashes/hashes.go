package hashes

import (
	"bytes"
	"encoding/binary"
	"errors"
	"math"
	"strings"
	"sync"

	"github.com/bkaradzic/go-lz4"
	"github.com/golang/protobuf/proto"
	pb "github.com/socketfunc/faas/store/proto"
)

var (
	ErrNotFound      = errors.New("store: not found")
	ErrInvalidFilter = errors.New("store: invalid filter")
	ErrInvalidUpdate = errors.New("store: invalid update")
	ErrInvalidType   = errors.New("store: invalid type")
	ErrMaximumEntity = errors.New("store: entity is maximum body")
)

type Hashes struct {
	sync.RWMutex
	entities map[string][]byte
}

func (h *Hashes) get(namespace, key string) (*pb.Entity, error) {
	value, ok := h.entities[namespace+":"+key]
	if !ok {
		return nil, ErrNotFound
	}

	var err error
	value, err = lz4.Decode(nil, value)
	if err != nil {
		return nil, err
	}

	entity := &pb.Entity{}
	if err := proto.Unmarshal(value, entity); err != nil {
		return nil, err
	}

	return entity, nil
}

func (h *Hashes) put(namespace, key string, entity *pb.Entity) error {
	value, err := proto.Marshal(entity)
	if err != nil {
		return err
	}

	if len(value) > 1024*1024 {
		return ErrMaximumEntity
	}

	value, err = lz4.Encode(nil, value)
	if err != nil {
		return err
	}

	h.entities[namespace+":"+key] = value

	return nil
}

func (h *Hashes) del(namespace, key string) {
	delete(h.entities, namespace+":"+key)
}

func (h *Hashes) Get(namespace, key string) (*pb.Entity, error) {
	h.RLock()
	defer h.RUnlock()

	return h.get(namespace, key)
}

func (h *Hashes) Put(namespace, key string, entity *pb.Entity) error {
	h.Lock()
	defer h.Unlock()

	return h.put(namespace, key, entity)
}

func (h *Hashes) Modify(namespace, key string, filters []*pb.Filter, updates []*pb.Update) (*pb.Entity, bool, error) {
	h.Lock()
	defer h.Unlock()

	entity, err := h.get(namespace, key)
	if err != nil && err != ErrNotFound {
		return nil, false, err
	}
	if entity == nil {
		entity = &pb.Entity{
			Data: map[string]*pb.Value{},
		}
	}

	for _, filter := range filters {
		if filter == nil {
			continue
		}

		switch filter.Comp {
		case pb.Comp_Eq: // ==
			value := entity.Data[filter.Name]
			if value == nil || !bytes.Equal(value.Data, filter.Value.Data) {
				return entity, false, nil
			}
		case pb.Comp_Gt: // >
			value := entity.Data[filter.Name]
			if value == nil || !(bytes.Compare(value.Data, filter.Value.Data) > 0) {
				return entity, false, nil
			}
		case pb.Comp_Gte: // >=
			value := entity.Data[filter.Name]
			if value == nil || !(bytes.Compare(value.Data, filter.Value.Data) >= 0) {
				return entity, false, nil
			}
		case pb.Comp_Lt: // <
			value := entity.Data[filter.Name]
			if value == nil || !(bytes.Compare(value.Data, filter.Value.Data) < 0) {
				return entity, false, nil
			}
		case pb.Comp_Lte: // <=
			value := entity.Data[filter.Name]
			if value == nil || !(bytes.Compare(value.Data, filter.Value.Data) <= 0) {
				return entity, false, nil
			}
		case pb.Comp_Ne: // !=
			value := entity.Data[filter.Name]
			if value == nil || bytes.Equal(value.Data, filter.Value.Data) {
				return entity, false, nil
			}
		case pb.Comp_Exists:
			if _, ok := entity.Data[filter.Name]; !ok {
				return entity, false, nil
			}
		default:
			return nil, false, ErrInvalidFilter
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
				value := entity.Data[update.Name]
				data := make([]byte, 8)
				addNum := binary.BigEndian.Uint64(update.Value.Data)
				if value == nil {
					binary.BigEndian.PutUint64(data, addNum)
				} else {
					num := binary.BigEndian.Uint64(value.Data)
					binary.BigEndian.PutUint64(data, num+addNum)
				}
				entity.Data[update.Name] = &pb.Value{
					Type: pb.Type_Int,
					Data: data,
				}
			case pb.Type_Float: // float64
				value := entity.Data[update.Name]
				data := make([]byte, 8)
				addNum := float64(binary.BigEndian.Uint64(update.Value.Data))
				if value == nil {
					binary.BigEndian.PutUint64(data, math.Float64bits(addNum))
				} else {
					num := float64(binary.BigEndian.Uint64(value.Data))
					binary.BigEndian.PutUint64(data, math.Float64bits(num+addNum))
				}
				entity.Data[update.Name] = &pb.Value{
					Type: pb.Type_Float,
					Data: data,
				}
			default:
				return nil, false, ErrInvalidType
			}
		case pb.Op_Set:
			entity.Data[update.Name] = update.Value
		case pb.Op_Unset:
			delete(entity.Data, update.Name)
		default:
			return nil, false, ErrInvalidUpdate
		}
	}

	if err := h.put(namespace, key, entity); err != nil {
		return nil, false, err
	}

	return entity, true, nil
}

func (h *Hashes) Del(namespace, key string) {
	h.Lock()
	defer h.Unlock()

	h.del(namespace, key)
}

func (h *Hashes) DelAll(namespace string) {
	h.Lock()
	defer h.Unlock()

	for index := range h.entities {
		idx := strings.SplitN(index, ":", 2)
		if idx[0] == namespace {
			h.del(namespace, idx[1])
		}
	}
}

func (h *Hashes) Keys() []string {
	h.RLock()
	defer h.RUnlock()

	keys := make([]string, 0, len(h.entities))
	for k := range h.entities {
		keys = append(keys, k)
	}
	return keys
}

func New() *Hashes {
	return &Hashes{
		entities: map[string][]byte{},
	}
}

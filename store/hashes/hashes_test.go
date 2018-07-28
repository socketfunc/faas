package hashes

import (
	"encoding/binary"
	"fmt"
	"testing"

	pb "github.com/socketfunc/faas/store/proto"
)

func TestNew(t *testing.T) {
	h := New()

	entity := &pb.Entity{
		Data: map[string]*pb.Value{
			"field": {
				Type: pb.Type_String,
				Data: []byte("test message"),
			},
			"field2": {
				Type: pb.Type_String,
				Data: []byte("test message"),
			},
		},
	}
	h.Set("ns", "k", entity)

	value, err := h.Get("ns", "k")
	fmt.Println(err)
	fmt.Println(value)

	num := int64(10)
	data := make([]byte, 8)
	binary.BigEndian.PutUint64(data, uint64(num))
	updates := []*pb.Update{
		{
			Op:   pb.Op_Inc,
			Name: "field3",
			Value: &pb.Value{
				Type: pb.Type_Int,
				Data: data,
			},
		},
	}
	entity, success, err := h.Modify("ns", "k", nil, updates)
	fmt.Println(err)
	fmt.Println(success)
	fmt.Println(entity)

	num = int64(10)
	data = make([]byte, 8)
	binary.BigEndian.PutUint64(data, uint64(num))
	filters := []*pb.Filter{
		{
			Comp: pb.Comp_Gt,
			Name: "field3",
			Value: &pb.Value{
				Type: pb.Type_Int,
				Data: data,
			},
		},
	}
	updates = []*pb.Update{
		{
			Op:   pb.Op_Inc,
			Name: "field3",
			Value: &pb.Value{
				Type: pb.Type_Int,
				Data: data,
			},
		},
	}
	entity, success, err = h.Modify("ns", "k", filters, updates)
	fmt.Println(err)
	fmt.Println(success)
	fmt.Println(int64(binary.BigEndian.Uint64(entity.Data["field3"].Data)))
}

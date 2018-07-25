package kvs

import (
	"encoding/binary"
	"fmt"
	"testing"

	pb "github.com/socketfunc/faas/store/proto"
	"github.com/stretchr/testify/assert"
)

func TestKvs_Put(t *testing.T) {
	num := int64(100)
	data := make([]byte, 8)
	binary.BigEndian.PutUint64(data, uint64(num))
	updates := []*pb.Update{
		{
			Op:   pb.Op_Set,
			Name: "field",
			Value: &pb.Value{
				Type: pb.Type_Int,
				Data: data,
			},
		},
	}

	kvs := NewKvs()
	entity, err := kvs.Put("test", "sample", nil, updates)
	fmt.Println(err)
	fmt.Println(entity.getValue("field").Data)

	filters := []*pb.Filter{
		{
			Name: "field",
			Comp: pb.Comp_Exists,
			Value: &pb.Value{
				Type: pb.Type_Int,
				Data: data,
			},
		},
	}
	updates = []*pb.Update{
		{
			Op:   pb.Op_Inc,
			Name: "field",
			Value: &pb.Value{
				Type: pb.Type_Int,
				Data: data,
			},
		},
	}
	entity, err = kvs.Put("test", "sample", filters, updates)
	assert.NoError(t, err)
	fmt.Println(entity.getValue("field").Data)

	val, err := kvs.Get("test", "sample")
	assert.NoError(t, err)
	fmt.Println(val)

	keys := kvs.Keys()
	fmt.Println(keys)
}

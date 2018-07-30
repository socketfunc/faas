package client

import (
	"fmt"
	"testing"

	"encoding/binary"
	"math"

	pb "github.com/socketfunc/faas/store/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_encodeEntity(t *testing.T) {
	type Data struct {
		ID    string  `store:"id"`
		Num   int     `store:"num"`
		Flag  bool    `store:"flag"`
		Float float64 `store:"float"`
		Xu    uint    `store:"xu"`
	}

	data := &Data{
		ID:    "message",
		Num:   100,
		Flag:  true,
		Float: 3.1415,
		Xu:    1,
	}
	entity, err := encodeEntity(data)
	require.NoError(t, err)

	float := make([]byte, 8)
	binary.BigEndian.PutUint64(float, math.Float64bits(3.1415))
	expected := &pb.Entity{
		Data: map[string]*pb.Value{
			"id": {
				Type: pb.Type_String,
				Data: []byte("message"),
			},
			"num": {
				Type: pb.Type_Int,
				Data: []byte{0, 0, 0, 0, 0, 0, 0, 100},
			},
			"flag": {
				Type: pb.Type_Bool,
				Data: []byte{1},
			},
			"float": {
				Type: pb.Type_Float,
				Data: float,
			},
			"xu": {
				Type: pb.Type_Uint,
				Data: []byte{0, 0, 0, 0, 0, 0, 0, 1},
			},
		},
	}
	assert.Equal(t, expected, entity)
}

func Test_decodeEntity(t *testing.T) {
	type Data struct {
		ID    string  `store:"id"`
		Num   int     `store:"num"`
		Flag  bool    `store:"flag"`
		Float float64 `store:"float"`
		Xu    uint    `store:"xu"`
	}

	float := make([]byte, 8)
	binary.BigEndian.PutUint64(float, math.Float64bits(3.1415))
	entity := &pb.Entity{
		Data: map[string]*pb.Value{
			"id": {
				Type: pb.Type_String,
				Data: []byte("message"),
			},
			"num": {
				Type: pb.Type_Int,
				Data: []byte{0, 0, 0, 0, 0, 0, 0, 100},
			},
			"flag": {
				Type: pb.Type_Bool,
				Data: []byte{1},
			},
			"float": {
				Type: pb.Type_Float,
				Data: float,
			},
			"xu": {
				Type: pb.Type_Uint,
				Data: []byte{0, 0, 0, 0, 0, 0, 0, 1},
			},
		},
	}

	data := &Data{}
	err := decodeEntity(entity, data)
	require.NoError(t, err)

	expected := &Data{
		ID:    "message",
		Num:   100,
		Flag:  true,
		Float: 3.1415,
		Xu:    1,
	}
	assert.Equal(t, expected, data)
}

func Test_convertUpdates(t *testing.T) {
	updates := []*Update{
		{
			Op:    Set,
			Name:  "field01",
			Value: "message",
		},
		{
			Op:    Set,
			Name:  "field2",
			Value: 1,
		},
		{
			Op:    Set,
			Name:  "field03",
			Value: 3.1415,
		},
		{
			Op:    Set,
			Name:  "field04",
			Value: true,
		},
	}
	values := convertUpdates(updates)
	fmt.Println(values)
}

func Test_convertFilters(t *testing.T) {
	filters := []*Filter{
		{
			Comp:  Eq,
			Name:  "field01",
			Value: "message",
		},
	}
	values := convertFilters(filters)
	fmt.Println(values)
}

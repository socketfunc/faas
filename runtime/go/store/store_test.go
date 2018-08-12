package store

import (
	"testing"

	"github.com/socketfunc/faas/store/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_encodeEntity(t *testing.T) {
	type Sample struct {
		ID  string `store:"id"`
		Num int    `store:"num"`
	}

	s := &Sample{
		ID:  "test",
		Num: 100,
	}
	entity, err := encodeEntity(s)
	require.NoError(t, err)
	assert.Equal(t, &store.Entity{
		Data: map[string]*store.Value{
			"id": {
				Type: store.Type_String,
				Data: []byte("test"),
			},
			"num": {
				Type: store.Type_Int,
				Data: []byte{0, 0, 0, 0, 0, 0, 0, 100},
			},
		},
	}, entity)
}

func Test_decodeEntity(t *testing.T) {
	type Sample struct {
		ID  string `store:"id"`
		Num int    `store:"num"`
	}

	entity := &store.Entity{
		Data: map[string]*store.Value{
			"id": {
				Type: store.Type_String,
				Data: []byte("test"),
			},
			"num": {
				Type: store.Type_Int,
				Data: []byte{0, 0, 0, 0, 0, 0, 0, 100},
			},
		},
	}

	s := &Sample{}
	err := decodeEntity(entity, s)
	require.NoError(t, err)
	assert.Equal(t, &Sample{
		ID:  "test",
		Num: 100,
	}, s)
}

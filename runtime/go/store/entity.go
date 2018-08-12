package store

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"math"
	"reflect"
	"unsafe"

	"github.com/socketfunc/faas/store/proto"
)

func encodeValue(v interface{}) (*store.Value, error) {
	value := &store.Value{}
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.String:
		str := rv.String()
		value.Type = store.Type_String
		value.Data = *(*[]byte)(unsafe.Pointer(&str))
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		data := make([]byte, 8)
		binary.BigEndian.PutUint64(data, uint64(rv.Int()))
		value.Type = store.Type_Int
		value.Data = data
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		data := make([]byte, 8)
		binary.BigEndian.PutUint64(data, rv.Uint())
		value.Type = store.Type_Uint
		value.Data = data
	case reflect.Float32, reflect.Float64:
		data := make([]byte, 8)
		binary.BigEndian.PutUint64(data, math.Float64bits(rv.Float()))
		value.Type = store.Type_Float
		value.Data = data
	case reflect.Bool:
		value.Type = store.Type_Bool
		if rv.Bool() {
			value.Data = []byte{1}
		} else {
			value.Data = []byte{0}
		}
	default:
		return nil, fmt.Errorf("unsupported value type: type=%s", rv.Kind().String())
	}
	return value, nil
}

func encodeEntity(v interface{}) (*store.Entity, error) {
	rv := reflect.ValueOf(v)

	if rv.Kind() == reflect.Ptr && !rv.IsNil() {
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Struct {
		return nil, errors.New("invalid")
	}

	entity := &store.Entity{
		Data: map[string]*store.Value{},
	}

	for i := 0; i < rv.NumField(); i++ {
		val := rv.Field(i)
		field := rv.Type().Field(i)
		key := field.Tag.Get("store")
		if key == "-" {
			continue
		}
		if key == "" {
			key = field.Name
		}

		value, err := encodeValue(val.Interface())
		if err != nil {
			return nil, err
		}
		entity.Data[key] = value
	}

	return entity, nil
}

func decodeEntity(entity *store.Entity, v interface{}) error {
	rv := reflect.ValueOf(v)

	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return errors.New("invalid value")
	}
	rv = rv.Elem()

	mapping := map[string]int{}
	for i := 0; i < rv.NumField(); i++ {
		field := rv.Type().Field(i)
		key := field.Tag.Get("store")
		if key == "-" {
			continue
		}
		if key == "" {
			key = field.Name
		}
		mapping[key] = i
	}

	for key, val := range entity.Data {
		idx, ok := mapping[key]
		if !ok {
			continue
		}
		field := rv.Field(idx)
		switch val.Type {
		case store.Type_String:
			field.SetString(string(val.Data))
		case store.Type_Int:
			x := binary.BigEndian.Uint64(val.Data)
			field.SetInt(int64(x))
		case store.Type_Uint:
			x := binary.BigEndian.Uint64(val.Data)
			field.SetUint(x)
		case store.Type_Float:
			bits := binary.BigEndian.Uint64(val.Data)
			x := math.Float64frombits(bits)
			field.SetFloat(x)
		case store.Type_Bool:
			field.SetBool(bytes.Equal(val.Data, []byte{1}))
		default:
		}
	}

	return nil
}

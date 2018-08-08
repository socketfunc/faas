package main

import (
	"fmt"
	"reflect"
	"sync"
	"time"
)

type Emitter struct {
	sync.Mutex
	events map[string]reflect.Value
}

func NewEmitter() *Emitter {
	return &Emitter{
		events: map[string]reflect.Value{},
	}
}

func (e *Emitter) On(eventID string, listener interface{}) {
	e.Lock()
	defer e.Unlock()

	fn := reflect.ValueOf(listener)

	if reflect.Func != fn.Kind() {
		return
	}

	e.events[eventID] = fn
}

func (e *Emitter) Emit(eventID string) {
	e.Lock()
	fn, ok := e.events[eventID]
	if !ok {
		return
	}
	delete(e.events, eventID)
	e.Unlock()

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func(fn reflect.Value) {
		defer wg.Done()

		fn.Call(nil)
	}(fn)
	wg.Wait()
}

func Get() (string, error) {
	e := NewEmitter()
	done := make(chan struct{}, 0)

	e.On("test", func() {
		defer close(done)
		fmt.Println("called test")
	})

	time.Sleep(time.Second)

	e.Emit("test")

	<-done

	return "value", nil
}

func main() {
	fmt.Println("---> 1")
	value, err := Get()
	fmt.Println(value, err)
	fmt.Println("---> 2")
}

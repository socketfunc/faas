package main

import (
	"fmt"
	"reflect"
	"sync"
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

func Get() {
	e := NewEmitter()
	done := make(chan struct{}, 0)

	e.On("test", func() {
		defer close(done)
		fmt.Println("called test")
	})
	e.Emit("test")

	<-done
}

func main() {
	fmt.Println("---> 1")
	Get()
	fmt.Println("---> 2")
}

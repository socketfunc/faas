package main

import (
	"fmt"
	"log"

	"github.com/socketfunc/faas/store/client"
)

type Data struct {
	ID  string `store:"id"`
	Num int    `store:"num"`
}

func main() {
	c, err := client.New("localhost:9090")
	if err != nil {
		log.Fatal(err)
	}

	data := &Data{
		ID:  "test",
		Num: 100,
	}
	err = c.Put("ns", "k", data)
	fmt.Println(err)

	val := &Data{}
	err = c.Get("ns", "k", val)
	fmt.Println(err)
	fmt.Println(val)

	updates := []*client.Update{
		{
			Name:  "num",
			Op:    client.Inc,
			Value: 100,
		},
	}
	val = &Data{}
	successful, err := c.Modify("ns", "k", nil, updates, val)
	fmt.Println(successful)
	fmt.Println(err)
	fmt.Println(val)

	keys, err := c.Keys()
	fmt.Println(err)
	fmt.Println(keys)
}

# Store

atomic hashes store

gRPCでの`Simple RPC`を行う

## client

### Get

```go
type Data struct {
	ID  string `store:"id"`
	Num int    `store:"num"`
}

data := &Data{}
err := client.Get(namespace, key, data)
```

### Put

```go
type Data struct {
	ID  string `store:"id"`
	Num int    `store:"num"`
}

data := &Data{
	ID:  "id",
	Num: 100,
}
err := client.Put(namespace, key, data)
```

### Modify

```go
type Data struct {
	ID  string `store:"id"`
	Num int    `store:"num"`
}

filters := []*Filter{
	{
		Name:  "id",
		Op:    Equal,
		Value: "id",
	},
}

updates := []*Update{
	{
		Name:  "id",
		Value: "id2",
	},
}

data := &Data{}
err := client.Modify(namespace, key, filters, updates, data)
```

### Del

```go
client.Del(namespace, key)
```

### DelAll

```go
client.DellAll(namespace)
```

### Keys

```go
client.Keys()
```

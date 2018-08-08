# runtime

## Architecture

gRPCサーバ

## Runtime Config

```yaml
port: 9000
codePath: /path/to/plugin
```

### Go

version: 1.10 over

```yaml
runtime: go1.10
```

```go
import (
	"github.com/socketfunc/faas/runtime/engine/go"
	"github.com/socketfunc/faas/runtime/engine/go/store"
	"github.com/socketfunc/faas/runtime/engine/go/log"
)

type Body struct {
	UserID string `json:"userId"`
}

type Value struct {
	ID  string `store:"id"`
	Num int    `store:"num"`
}

func Handler(ctx context.Context, req *runtime.Request, res runtime.Response) {
	// request
	topic := req.Topic() // read only
	event := req.Event() // read only
	
	body := &Body{}
	err := req.Decode(body)
	
	// store get
	value := &Value{}
	err := store.Get(ctx, key, value)
	// store put
	value := &Value{}
	err := store.Put(ctx, key, value)
	// store modify
	value := &Value{}
	successful, err := store.Modify(ctx, key, filters, updates, value)
	// del
	err := store.Del(ctx, key)
	
	// response send
	res.Send()
	// response broadcast
	res.Broadcast()
	// response reply
	res.Reply()
	
	// log
	log.Debug()
	log.Info()
	log.Warn()
	log.Error()
}

runtime.Start(Handler)
```

### Node.js

version: 8.10 over

拡張子: js

```yaml
runtime: nodejs10.6
```

```javascript
const runtime = require('socketfunc')
const store = runtime.store
const log = runtime.log

async function handler (ctx, req, res) {
  const topic = req.topic()
  const event = req.event()
  
  const body = req.decode()
  
  // store get
  const value = await store.get(ctx, key)
  // store put
  await store.put(ctx, key, value)
  // store modify
  const {successful, value} = await store.modify(ctx, key, filters, updates)
  // store del
  await store.del(ctx, key)
  
  // response
  res.send()
  res.broadcast()
  res.reply()
  
  // log
  log.debug()
  log.info()
  log.warn()
  log.error()
}

runtime.start(handler)
```

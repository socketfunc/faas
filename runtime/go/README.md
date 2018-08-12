# runtime for Go

`1 event` = `1 docker`で設計する

### config

```yaml
runtime: go
handler:
    event: room.join
env:
    # カスタム環境変数
```

### handler

Handler関数を作成

```go
import (
	"github.com/socketfunc/faas/runtime/go/server"
	"github.com/socketfunc/faas/runtime/go/store"
)

func Handler(ctx context.Context, req runtime.Request, res runtime.Response) {
	res.Send("topic", "event", []byte("message"))
}
```

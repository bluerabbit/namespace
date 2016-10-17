# store

## Import

```go
import "github.com/bluerabbit/namespace"
```

## Usage

```go
var client *redis.Client

client = redis.NewClient(&redis.Options{
  Addr:       "localhost:6379",
  Password:   "",
  PoolSize:    10,
  PoolTimeout: 10,
})

var n = namespace.New(client, "user_name")

n.Set(1, "hoge")
n.Set(2, "foo")
n.Get(1) // hoge
n.Get(2) // foo
n.Get(3) // error
```

## License

MIT

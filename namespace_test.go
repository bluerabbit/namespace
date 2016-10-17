package namespace

import (
	"gopkg.in/redis.v4"
	"testing"
)

func Test(t *testing.T) {
	var client *redis.Client

	client = redis.NewClient(&redis.Options{
		Addr:       "localhost:6379",
		Password:   "",
		PoolSize:    10,
		PoolTimeout: 10,
	})

	c := New(client, "key")
	c.Set(1, "foo")

	v1, error := c.Get(1)
	if error != nil {
		t.Errorf("error")
	}
	if v1 != "foo" {
		t.Errorf("got %v", v1)
	}

	_, error2 := c.Get(2)
	if error2 == nil {
		t.Errorf("error")
	}
}

package namespace

import (
	"gopkg.in/redis.v4"
	"fmt"
)

type Namespace struct {
	client *redis.Client
	key    string
}

func New(client *redis.Client, key string) *Namespace {
	return &Namespace{
		client: client,
		key: key,
	}
}

func (s *Namespace) Set(id int, text string) {
	s.client.Set(s.CreateKey(id), text, 0)
}

func (s *Namespace) Get(id int) (string, error) {
	return s.client.Get(s.CreateKey(id)).Result()
}

func (s *Namespace) CreateKey(id int) string {
	return fmt.Sprintf(s.key + "_%d", id)
}

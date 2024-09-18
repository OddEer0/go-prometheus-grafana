package cache

import "sync"

type (
	Cache[T any] interface {
		Get(key string) (T, bool)
		Add(key string, value T)
		Delete(key string)
	}

	cache[T any] struct {
		data map[string]T
		mu   sync.Mutex
	}
)

func (c *cache[T]) Get(key string) (T, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	res, ok := c.data[key]
	return res, ok
}

func (c *cache[T]) Add(key string, value T) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
}

func (c *cache[T]) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.data, key)
}

func New[T any]() Cache[T] {
	return &cache[T]{
		data: make(map[string]T),
		mu:   sync.Mutex{},
	}
}

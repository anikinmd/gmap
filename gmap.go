package gmap

import (
	"errors"
	"sync"
)

var ErrorUnknownKey = errors.New("the key doesn't exist")

type GMap[K comparable, V any] struct {
	data map[K]V
	mu   sync.RWMutex
}

func NewGMap[K comparable, V any]() *GMap[K, V] {
	var c = GMap[K, V]{}
	c.data = make(map[K]V)
	return &c
}

func (c *GMap[K, V]) Set(key K, value V) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
}

func (c *GMap[K, V]) Get(key K) (V, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	val, ok := c.data[key]
	if ok {
		return val, nil
	}
	return val, ErrorUnknownKey
}

func (c *GMap[K, V]) Delete(key K) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, ok := c.data[key]
	if !ok {
		return ErrorUnknownKey
	}
	delete(c.data, key)
	return nil
}

func (c *GMap[K, V]) GetKeys() []K {
	c.mu.RLock()
	defer c.mu.RUnlock()
	keys := make([]K, len(c.data))

	i := 0
	for k := range c.data {
		keys[i] = k
		i++
	}
	return keys
}

func (c *GMap[K, V]) CheckKeyExists(key K) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	_, ok := c.data[key]
	return ok
}

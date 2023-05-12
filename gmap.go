// Package gmap is a generic-based thread-safe map.
package gmap

import (
	"errors"
	"sync"
)

// ErrorUnknownKey is an unknown key error.
var ErrorUnknownKey = errors.New("the key doesn't exist")

// GMap is a generic thread-safe map.
type GMap[K comparable, V any] struct {
	data map[K]V
	mu   sync.RWMutex
}

// NewGMap creates a new generic map.
func NewGMap[K comparable, V any]() *GMap[K, V] {
	var c = GMap[K, V]{}
	c.data = make(map[K]V)
	return &c
}

// Set new key-value pair to the map or overwrite the actual value.
func (c *GMap[K, V]) Set(key K, value V) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
}

// Get value by key or ErrorUnknownKey if the key is not found.
func (c *GMap[K, V]) Get(key K) (V, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	val, ok := c.data[key]
	if ok {
		return val, nil
	}
	return val, ErrorUnknownKey
}

// Delete key and its value from the map.
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

// GetKeys returns a slice of keys.
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

// CheckKeyExists returns true if the key was added to the map.
func (c *GMap[K, V]) CheckKeyExists(key K) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	_, ok := c.data[key]
	return ok
}

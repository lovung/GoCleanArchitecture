package storage

import (
	"sync"
)

var (
	once  sync.Once
	cache *Cache
)

// Cache defines a simple memory cache for key-value
type Cache struct {
	cacheStorage map[string]interface{}
}

// InitCache to init a cache
func InitCache() {
	once.Do(func() {
		cache = &Cache{
			make(map[string]interface{}),
		}
	})
}

// Get returns value with a key
func (c *Cache) Get(key string) interface{} {
	return c.cacheStorage[key]
}

// Set to store a key-val to cache
func (c *Cache) Set(key string, val interface{}) {
	c.cacheStorage[key] = val
}

// GetCache return data directory
func GetCache() *Cache {
	return cache
}

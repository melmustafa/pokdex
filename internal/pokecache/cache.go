package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	val       []byte
	createdAt time.Time
}

type Cache struct {
	values map[string]cacheEntry
	mu     *sync.RWMutex
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		values: make(map[string]cacheEntry),
		mu:     &sync.RWMutex{},
	}
	go cache.reapLoop(interval)
	return cache
}

func (c Cache) Add(key string, val []byte) {
	newCacheEntry := cacheEntry{
		val:       val,
		createdAt: time.Now().UTC(),
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	c.values[key] = newCacheEntry
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	val, ok := c.values[key]
	return val.val, ok
}

func (c Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC().Add(-interval))
	}
}

func (c Cache) reap(target time.Time) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for key, entry := range c.values {
		if entry.createdAt.Before(target) {
			delete(c.values, key)
		}
	}
}

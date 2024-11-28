package cache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	items map[string]CacheEntry
	mu    *sync.RWMutex
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		items: make(map[string]CacheEntry),
		mu:    &sync.RWMutex{},
	}
	return cache
}

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (c *Cache) Add(key string, value []byte) error {
	if key == "" || len(value) == 0 {
		return fmt.Errorf("Cache entry cannot be empty")
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	fmt.Printf("Adding cache entry with key: %s\n", key)
	c.items[key] = CacheEntry{
		createdAt: time.Now(),
		val:       value,
	}
	return nil
}

func (c *Cache) Get(key string) (val []byte, hit bool, err error) {
	if key == "" {
		return nil, false, fmt.Errorf("Cache hit check cannot be empty")
	}
	c.mu.RLock()
	defer c.mu.RUnlock()
	entry, exists := c.items[key]
	if !exists {
		return nil, false, nil
	}
	fmt.Println("cache hit")
	return entry.val, true, nil
}

func (c *Cache) ReapLoop(interval time.Duration, done chan bool) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			c.mu.Lock()
			for k, v := range c.items {
				if time.Since(v.createdAt) > interval {
					delete(c.items, k)
				}
			}
			c.mu.Unlock()
		case <-done:
			return
		}
	}
}

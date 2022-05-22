package cache

import (
	"time"
)

type cacheItem struct {
	value          string
	expirationTime time.Time
}

func (c cacheItem) isExpired() bool {
	return c.expirationTime != time.Time{} && c.expirationTime.Unix() < time.Now().Unix()
}

type Cache struct {
	items map[string]cacheItem
}

func NewCache() Cache {
	return Cache{
		items: map[string]cacheItem{},
	}
}

func (c *Cache) Get(key string) (string, bool) {
	item, ok := c.items[key]
	if !ok || item.isExpired() {
		return "", false
	}

	return item.value, true
}

func (c *Cache) Put(key, value string) {
	c.PutTill(key, value, time.Time{})
}

func (c *Cache) Keys() []string {
	result := make([]string, 0)
	for key, item := range c.items {
		if !item.isExpired() {
			result = append(result, key)
		}
	}

	if len(result) == 0 {
		return nil
	}

	return result
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	c.items[key] = cacheItem{
		value:          value,
		expirationTime: deadline,
	}
}

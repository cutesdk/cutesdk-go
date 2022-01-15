package cache

import (
	"errors"
	"fmt"
	"time"

	gocache "github.com/patrickmn/go-cache"
)

// Cache interface defined how to use cache
type Cache interface {
	Set(key string, value interface{}, expire time.Duration) error
	Get(key string) (interface{}, error)
}

// CommCahce common use for cache
type CommCache struct {
	c *gocache.Cache
}

// New init cache
func New() *CommCache {
	c := gocache.New(5*time.Minute, 10*time.Minute)

	return &CommCache{c}
}

// Set set value to cache
func (c *CommCache) Set(key string, value interface{}, expire time.Duration) error {
	fmt.Println(key, value, expire)
	c.c.Set(key, value, expire)

	return nil
}

// Get get value from cache
func (c *CommCache) Get(key string) (interface{}, error) {
	v, ok := c.c.Get(key)
	fmt.Println(key, v, ok)
	if !ok {
		return nil, errors.New("cache data not exist")
	}

	return v, nil
}

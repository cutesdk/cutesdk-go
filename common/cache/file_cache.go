package cache

import (
	"errors"
	"time"

	"github.com/liyaojian/cache"
)

// FileCache: file-based cache
type FileCache struct {
	cache *cache.FileCache
}

// NewFileCache: init cache component
func NewFileCache(conf map[string]interface{}) (*FileCache, error) {
	cacheDir := "./cache"
	if v, ok := conf["dir"]; ok {
		cacheDir = v.(string)
	}

	c := cache.NewFileCache(cacheDir)

	return &FileCache{c}, nil
}

// Set: set data to cache
func (c *FileCache) Set(key string, value interface{}, expire time.Duration) error {
	return c.cache.Set(key, value, expire)
}

// Get: get data from cache
func (c *FileCache) Get(key string) (interface{}, error) {
	value := c.cache.Get(key)

	if value == nil {
		return nil, errors.New("cache not exists")
	}

	return value, nil
}

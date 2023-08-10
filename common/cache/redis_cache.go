package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

// RedisCache: redis-based cache component
type RedisCache struct {
	cache   *redis.Client
	timeout time.Duration
}

// NewRedisCache: init cache component
func NewRedisCache(conf map[string]interface{}) (*RedisCache, error) {
	timeout := 5 * time.Second
	if v, ok := conf["timeout"]; ok {
		if td, ok := v.(time.Duration); ok {
			timeout = td
		}
		if ts, ok := v.(string); ok {
			if t, err := time.ParseDuration(ts); err == nil {
				timeout = t
			}
		}
	}

	if v, ok := conf["client"]; ok {
		if cache, ok := v.(*redis.Client); ok {
			return &RedisCache{cache, timeout}, nil
		}
	}

	dsn := ""
	if v, ok := conf["dsn"]; ok {
		dsn = v.(string)
	}

	if dsn == "" {
		return nil, fmt.Errorf("invalid redis dsn: %s", dsn)
	}

	opt, err := redis.ParseURL(dsn)
	if err != nil {
		return nil, fmt.Errorf("invalid redis dsn: %v", err)
	}

	cache := redis.NewClient(opt)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if _, err := cache.Ping(ctx).Result(); err != nil {
		return nil, fmt.Errorf("connect to cache failed: %v", err)
	}

	return &RedisCache{cache, timeout}, nil
}

// Set: set data to cache
func (c *RedisCache) Set(key string, value interface{}, expire time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	_, err := c.cache.Do(ctx, "SETEX", key, expire.Seconds(), value).Result()

	return err
}

// Get: get data from cache
func (c *RedisCache) Get(key string) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	value, err := c.cache.Do(ctx, "GET", key).Result()

	return value, err
}

// Delete: delete data from cache
func (c *RedisCache) Delete(key string) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	_, err := c.cache.Do(ctx, "DEL", key).Result()

	return err
}

package cache

import "time"

// ICache: defined cache interface
type ICache interface {
	Set(key string, value interface{}, expire time.Duration) error
	Get(key string) (interface{}, error)
	Delete(key string) error
}

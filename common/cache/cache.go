package cache

import (
	"errors"
)

// NewCache: new cache handler
func NewCache(iopts IOptions) (ICache, error) {
	opts := iopts.ToOptions()

	if opts == nil {
		return nil, errors.New("invalid cache options")
	}

	driver := opts.Driver
	conf := opts.Conf

	if driver == "redis" {
		cache, err := NewRedisCache(conf)
		if err != nil {
			return nil, errors.New("new redis cache failed: " + err.Error())
		}

		return cache, nil
	}

	if driver == "file" {
		cache, err := NewFileCache(conf)
		if err != nil {
			return nil, errors.New("new file cache failed: " + err.Error())
		}

		return cache, nil
	}

	return nil, errors.New("invalid cache driver")
}

package cache

import "time"

// IOptions: interface of custom cache options
type IOptions interface {
	ToOptions() *Options
}

// Options: cache options
type Options struct {
	Driver string
	Conf   map[string]interface{}
}

// FileOptions: file cache options
type FileOptions struct {
	Dir string
}

// ToOptions: to cache options
func (o *FileOptions) ToOptions() *Options {
	return &Options{
		Driver: "file",
		Conf: map[string]interface{}{
			"dir": o.Dir,
		},
	}
}

// RedisOptions: redis cache options
type RedisOptions struct {
	Dsn     string
	Timeout time.Duration
}

// ToOptions: to cache options
func (o *RedisOptions) ToOptions() *Options {
	return &Options{
		Driver: "redis",
		Conf: map[string]interface{}{
			"dsn":     o.Dsn,
			"timeout": o.Timeout,
		},
	}
}

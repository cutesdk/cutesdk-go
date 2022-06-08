package wxapp

import (
	"github.com/cutesdk/cutesdk-go/common/cache"
	"github.com/cutesdk/cutesdk-go/common/request"
)

// Options: custom options
type Options struct {
	Debug   bool
	Request *request.Options
	Cache   *cache.Options
	Appid   string
	Secret  string
}

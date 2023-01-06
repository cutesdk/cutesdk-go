package wxapp

import (
	"time"

	"github.com/cutesdk/cutesdk-go/common/cache"
)

// Options: custom options
type Options struct {
	Debug   bool
	Timeout time.Duration
	Cache   cache.IOptions
	Appid   string
	Secret  string
	Token   string
	AesKey  string // 43 bit
	aesKey  []byte // 32 bit
}

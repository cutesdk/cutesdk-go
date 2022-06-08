package wxopen

import (
	"github.com/cutesdk/cutesdk-go/common/cache"
	"github.com/cutesdk/cutesdk-go/common/request"
)

// Options: custom options
type Options struct {
	Debug              bool
	Request            *request.Options
	Cache              *cache.Options
	ComponentAppid     string
	ComponentAppsecret string
	VerifyToken        string
	EncodingAesKey     string // 43 bit
	aesKey             []byte // 32 bit
}

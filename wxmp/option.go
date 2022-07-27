package wxmp

import (
	"github.com/cutesdk/cutesdk-go/common/cache"
	"github.com/cutesdk/cutesdk-go/common/request"
	"github.com/cutesdk/cutesdk-go/wxopen"
)

// Options: custom options
type Options struct {
	Debug          bool
	Request        *request.Options
	Cache          *cache.Options
	Appid          string
	Secret         string
	VerifyToken    string
	EncodingAesKey string // 43 bit
	aesKey         []byte // 32 bit
	// authorizer info
	AuthorizerProvider     *wxopen.Instance
	AuthorizerRefreshToken string
}

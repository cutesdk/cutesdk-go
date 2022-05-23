package wxwork

import (
	"github.com/cutesdk/cutesdk-go/common/cache"
	"github.com/cutesdk/cutesdk-go/common/request"
)

// Options: custom options
type Options struct {
	Request        *request.Options
	Cache          *cache.Options
	Corpid         string
	Agentid        string
	Secret         string
	VerifyToken    string
	EncodingAesKey string // 43 bit
	aesKey         []byte // 32 bit
}

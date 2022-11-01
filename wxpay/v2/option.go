package wxpay

import (
	"github.com/cutesdk/cutesdk-go/common/cache"
	"github.com/cutesdk/cutesdk-go/common/request"
)

// Options: custom options
type Options struct {
	Debug      bool
	Request    *request.Options
	Cache      *cache.Options
	MchId      string
	Appid      string
	SubMchId   string
	SubAppid   string
	ApiKey     string
	CertKey    string
	PrivateKey string
	NotifyUrl  string
}

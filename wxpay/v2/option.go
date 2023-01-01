package wxpay

import (
	"time"
)

// Options: custom options
type Options struct {
	Debug    bool
	Timeout  time.Duration
	BaseUri  string
	MchId    string
	SubMchId string
	ApiKey   string
	CertPem  string
	KeyPem   string
	CertFile string
	KeyFile  string
}

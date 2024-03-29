package wxpay

import (
	"crypto/rsa"
	"time"
)

// Options: custom options
type Options struct {
	Debug      bool
	Timeout    time.Duration
	BaseUri    string
	MchId      string
	SubMchId   string
	ApiKey     string
	SerialNo   string
	KeyPem     string
	KeyPath    string
	privateKey *rsa.PrivateKey
}

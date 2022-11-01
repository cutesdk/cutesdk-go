package request

import (
	"crypto/tls"
	"time"
)

// Options: request options
type Options struct {
	BaseUri      string
	Debug        bool
	Timeout      time.Duration
	Certificates []tls.Certificate
}

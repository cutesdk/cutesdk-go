package request

import "time"

// Options: request options
type Options struct {
	BaseUri string
	Debug   bool
	Timeout time.Duration
}

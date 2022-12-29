package wxmp

import "github.com/idoubi/goutils/crypt"

// Server: wxmp server
type Server struct {
	opts *Options
}

// NewServer: new wxmp server
func NewServer(opts *Options) (*Server, error) {
	if opts.EncodingAesKey != "" {
		if v, _ := crypt.Base64Decode(opts.EncodingAesKey + "="); v != nil {
			opts.aesKey = v
		}
	}

	// new server
	server := &Server{opts: opts}

	return server, nil
}

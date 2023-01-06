package wxopen

import (
	"github.com/idoubi/goutils/crypt"
)

// Server: wxopen server
type Server struct {
	opts *Options
	cli  *Client
}

// NewServer: new wxopen server
func NewServer(opts *Options) (*Server, error) {
	cli, err := NewClient(opts)
	if err != nil {
		return nil, err
	}

	if opts.AesKey != "" {
		if v, _ := crypt.Base64Decode(opts.AesKey + "="); v != nil {
			opts.aesKey = v
		}
	}

	// new server
	svr := &Server{opts, cli}

	return svr, nil
}

// GetClient: get wxapp client
func (svr *Server) GetClient() *Client {
	return svr.cli
}

package wxmp

import "github.com/idoubi/goutils/crypt"

// Server: wxmp server
type Server struct {
	opts *Options
	cli  *Client
}

// NewServer: new wxmp server
func NewServer(opts *Options) (*Server, error) {
	cli, err := NewClient(opts)
	if err != nil {
		return nil, err
	}

	if opts.EncodingAesKey != "" {
		if v, _ := crypt.Base64Decode(opts.EncodingAesKey + "="); v != nil {
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

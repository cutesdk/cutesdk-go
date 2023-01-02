package wxpay

import "context"

// Server: wxpay server
type Server struct {
	opts *Options
	cli  *Client
	ctx  context.Context
}

// NewServer: new wxpay server
func NewServer(opts *Options) (*Server, error) {
	cli, err := NewClient(opts)
	if err != nil {
		return nil, err
	}

	s := &Server{opts, cli, context.Background()}

	return s, nil
}

// GetClient: get wxapp client
func (svr *Server) GetClient() *Client {
	return svr.cli
}

// GetOptions: get options
func (s *Server) GetOptions() *Options {
	return s.opts
}

// GetMchId: get mch_id
func (svr *Server) GetMchId() string {
	return svr.opts.MchId
}

package wxpay

// Server: wxpay server
type Server struct {
	opts *Options
}

// NewServer: new wxpay server
func NewServer(opts *Options) (*Server, error) {
	s := &Server{opts: opts}

	return s, nil
}

// GetOptions: get options
func (s *Server) GetOptions() *Options {
	return s.opts
}

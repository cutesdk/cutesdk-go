package wxwork

import "github.com/idoubi/goutils/crypt"

// Server: wxwork server
type Server struct {
	opts *Options
}

// NewServer: new wxwork server
func NewServer(opts *Options) (*Server, error) {
	if opts.EncodingAesKey != "" {
		if v, _ := crypt.Base64Decode(opts.EncodingAesKey + "="); v != nil {
			opts.aesKey = v
		}
	}
	s := &Server{opts: opts}

	return s, nil
}

// GetOptions: get options
func (s *Server) GetOptions() *Options {
	return s.opts
}

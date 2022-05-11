package wxopen

import (
	"github.com/cutesdk/cutesdk-go/common/token"
	"github.com/idoubi/goutils/crypt"
)

// Server: wxopen server
type Server struct {
	opts                          *Options
	componentVerifyTicketCacheKey string
	componentVerifyTicketHandler  token.IToken
}

// NewServer: new wxopen server
func NewServer(opts *Options) (*Server, error) {
	if opts.EncodingAesKey != "" {
		if v, _ := crypt.Base64Decode(opts.EncodingAesKey + "="); v != nil {
			opts.aesKey = v
		}
	}

	s := &Server{opts: opts}

	return s, nil
}

// GetComponentAppid: get component_appid
func (c *Server) GetComponentAppid() string {
	return c.opts.ComponentAppid
}

package wxpay

import (
	"time"

	"github.com/cutesdk/cutesdk-go/common/request"
	"github.com/idoubi/goz"
)

// Client: wxpay client
type Client struct {
	opts          *Options
	requestClient *request.Client
}

// NewClient create wxpay client
func NewClient(opts *Options) (*Client, error) {
	// set default options
	if opts.Request == nil {
		opts.Request = &request.Options{}
	}
	if opts.Request.BaseUri == "" {
		opts.Request.BaseUri = "https://api.mch.weixin.qq.com"
	}
	if opts.Request.Timeout == 0 {
		opts.Request.Timeout = 5 * time.Second
	}

	// new client
	c := &Client{opts: opts}

	// set request client
	c.requestClient = request.NewClient(opts.Request)

	return c, nil
}

// Get: request api with get method
func (c *Client) Get(uri string, args ...map[string]interface{}) (request.Result, error) {
	return c.GetRequestClient().Get(uri, args...)
}

// Post: request api with post method
func (c *Client) Post(uri string, args ...map[string]interface{}) (request.Result, error) {
	return c.GetRequestClient().PostXml(uri, args...)
}

// Request: request api
func (c *Client) Request(method, uri string, opts goz.Options) (request.Result, error) {
	return c.GetRequestClient().Request(method, uri, opts)
}

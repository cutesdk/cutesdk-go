package app

import (
	"fmt"
	"time"

	"github.com/cutesdk/cutesdk-go/common/request"
	"github.com/idoubi/goz"
)

// Client: base app client
type Client struct {
	appid  string
	secret string
	opts   *Options
}

// NewClient create toutiao client
func NewClient(appid, secret string, optfuncs ...OptionFunc) (*Client, error) {
	c := &Client{appid, secret, &Options{
		requestOptions: request.Options{},
	}}

	for _, f := range optfuncs {
		if err := f(c.opts); err != nil {
			return nil, fmt.Errorf("init option failed: %v", err)
		}
	}

	// set default request timeout
	if c.opts.requestOptions.Timeout == 0 {
		c.opts.requestOptions.Timeout = 5 * time.Second
	}

	// set default cache
	if c.opts.cacheHandler == nil {
		c.SetOption(WithCache("file", map[string]interface{}{
			"dir": "./cache",
		}))
	}

	if c.opts.requestClient == nil {
		c.opts.requestClient = request.NewClient(c.opts.requestOptions)
	}

	return c, nil
}

// SetOption: set client option
func (c *Client) SetOption(opt OptionFunc) error {
	return opt(c.opts)
}

// SetRequestOptions: set request options
func (c *Client) SetRequestOptions(opts request.Options) {
	c.opts.requestOptions = opts

	c.opts.requestClient = request.NewClient(c.opts.requestOptions)
}

// GetRequestOptions: get request options
func (c *Client) GetRequestOptions() request.Options {
	return c.opts.requestOptions
}

// GetAppid get client appid
func (c *Client) GetAppid() string {
	return c.appid
}

// GetSecret get client secret
func (c *Client) GetSecret() string {
	return c.secret
}

// GetAccessToken: get access_token from cache or api
func (c *Client) GetAccessToken() (string, error) {
	return c.GetAccessTokenHandler().GetToken()
}

// Get: request api with get method
func (c *Client) Get(uri string, args ...map[string]interface{}) (request.Result, error) {
	return c.GetRequestClient().Get(uri, args...)
}

// Post: request api with post method
func (c *Client) Post(uri string, args ...map[string]interface{}) (request.Result, error) {
	return c.GetRequestClient().Post(uri, args...)
}

// Request: request api
func (c *Client) Request(method, uri string, opts goz.Options) (request.Result, error) {
	return c.GetRequestClient().Request(method, uri, opts)
}

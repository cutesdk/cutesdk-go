package wxmp

import (
	"fmt"
	"time"

	"github.com/cutesdk/cutesdk-go/common/cache"
	"github.com/cutesdk/cutesdk-go/common/request"
	"github.com/cutesdk/cutesdk-go/common/token"
	"github.com/idoubi/goz"
)

// Client: wxmp client
type Client struct {
	opts                *Options
	requestClient       *request.Client
	cacheHandler        cache.ICache
	accessTokenCacheKey string
	accessTokenHandler  token.IToken
	jsapiTicketCacheKey string
	jsapiTicketHandler  token.IToken
}

// NewClient create wxmp client
func NewClient(opts *Options) (*Client, error) {
	// set default options
	if opts.Request == nil {
		opts.Request = &request.Options{}
	}
	if opts.Cache == nil {
		opts.Cache = &cache.Options{}
	}
	if opts.Request.BaseUri == "" {
		opts.Request.BaseUri = "https://api.weixin.qq.com"
	}
	if opts.Request.Timeout == 0 {
		opts.Request.Timeout = 5 * time.Second
	}
	if opts.Cache.Driver == "" {
		opts.Cache.Driver = "file"
		opts.Cache.Conf = map[string]interface{}{
			"dir": "./cache",
		}
	}

	// new client
	c := &Client{opts: opts}

	// set cache handler
	if cache, err := cache.NewCache(opts.Cache); err != nil {
		return nil, fmt.Errorf("new cache handler failed: " + err.Error())
	} else {
		c.SetCacheHandler(cache)
	}

	// set request client
	c.requestClient = request.NewClient(opts.Request)

	// set default access_token cache key
	accessTokenCacheKey := fmt.Sprintf("wxmp.access_token.%s", c.GetAppid())
	c.SetAccessTokenCacheKey(accessTokenCacheKey)

	// set default access_token handler
	c.SetAccessTokenHandler(NewAccessToken(c))

	// set default jsapi_ticket cache key
	jsapiTicketCacheKey := fmt.Sprintf("wxmp.jsapi_ticket.%s", c.GetAppid())
	c.SetJsapiTicketCacheKey(jsapiTicketCacheKey)

	// set default jsapi_ticket handler
	c.SetJsapiTicketHandler(NewJsapiTicket(c))

	return c, nil
}

// GetAccessToken: get access_token from cache or api
func (c *Client) GetAccessToken() (string, error) {
	return c.GetAccessTokenHandler().GetToken()
}

// RefreshAccessToken: refresh access_token
func (c *Client) RefreshAccessToken() (string, error) {
	return c.GetAccessTokenHandler().RefreshToken()
}

// SetAccessToken: set access_token
func (c *Client) SetAccessToken(token string, expire time.Duration) error {
	return c.GetAccessTokenHandler().SetToken(token, expire)
}

// GetJsapiTicket: get jsapi_ticket from cache or api
func (c *Client) GetJsapiTicket() (string, error) {
	return c.GetJsapiTicketHandler().GetToken()
}

// RefreshJsapiTicket: refresh jsapi_ticket
func (c *Client) RefreshJsapiTicket() (string, error) {
	return c.GetJsapiTicketHandler().RefreshToken()
}

// SetJsapiTicket: set jsapi_ticket
func (c *Client) SetJsapiTicket(token string, expire time.Duration) error {
	return c.GetJsapiTicketHandler().SetToken(token, expire)
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

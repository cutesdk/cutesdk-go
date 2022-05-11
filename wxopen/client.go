package wxopen

import (
	"fmt"
	"time"

	"github.com/cutesdk/cutesdk-go/common/cache"
	"github.com/cutesdk/cutesdk-go/common/request"
	"github.com/cutesdk/cutesdk-go/common/token"
	"github.com/idoubi/goz"
)

// Client: wxopen client
type Client struct {
	opts                          *Options
	requestClient                 *request.Client
	cacheHandler                  cache.ICache
	componentAccessTokenCacheKey  string
	componentAccessTokenHandler   token.IToken
	componentVerifyTicketCacheKey string
	componentVerifyTicketHandler  token.IToken
}

// NewClient create wxopen client
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

	// set default component_access_token cache key
	componentAccessTokenCacheKey := fmt.Sprintf("wxopen.component_access_token.%s", c.GetComponentAppid())
	c.SetComponentAccessTokenCacheKey(componentAccessTokenCacheKey)

	// set default component_access_token handler
	c.SetComponentAccessTokenHandler(NewComponentAccessToken(c))

	// set default component_verify_ticket cache key
	componentVerifyTicketCacheKey := fmt.Sprintf("wxopen.component_verify_ticket.%s", c.GetComponentAppid())
	c.SetComponentVerifyTicketCacheKey(componentVerifyTicketCacheKey)

	// set default component_verify_ticket handler
	c.SetComponentVerifyTicketHandler(NewComponentVerifyTicket(c))

	return c, nil
}

// GetComponentAccessToken: get component_access_token from cache or api
func (c *Client) GetComponentAccessToken() (string, error) {
	return c.GetComponentAccessTokenHandler().GetToken()
}

// RefreshComponentAccessToken: refresh component_access_token
func (c *Client) RefreshComponentAccessToken() (string, error) {
	return c.GetComponentAccessTokenHandler().RefreshToken()
}

// SetComponentAccessToken: set component_access_token
func (c *Client) SetComponentAccessToken(token string, expire time.Duration) error {
	return c.GetComponentAccessTokenHandler().SetToken(token, expire)
}

// GetComponentVerifyTicket: get component_verify_ticket from cache
func (c *Client) GetComponentVerifyTicket() (string, error) {
	return c.GetComponentVerifyTicketHandler().GetToken()
}

// RefreshComponentVerifyTicket: refresh component_verify_ticket
func (c *Client) RefreshComponentVerifyTicket() (string, error) {
	return c.GetComponentVerifyTicketHandler().RefreshToken()
}

// SetComponentVerifyTicket: set component_verify_ticket
func (c *Client) SetComponentVerifyTicket(token string, expire time.Duration) error {
	return c.GetComponentVerifyTicketHandler().SetToken(token, expire)
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

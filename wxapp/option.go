package wxapp

import (
	"github.com/cutesdk/cutesdk-go/common/cache"
	"github.com/cutesdk/cutesdk-go/common/request"
	"github.com/cutesdk/cutesdk-go/common/token"
)

// Options: custom options
type Options struct {
	Request *request.Options
	Cache   *cache.Options
	Appid   string
	Secret  string
}

// GetRequestClient: get request handler
func (c *Client) GetRequestClient() *request.Client {
	return c.requestClient
}

// GetCacheHandler: get cache handler
func (c *Client) GetCacheHandler() cache.ICache {
	return c.cacheHandler
}

// SetCacheHandler: set cache handler
func (c *Client) SetCacheHandler(handler cache.ICache) error {
	c.cacheHandler = handler

	return nil
}

// GetAppid: get appid
func (c *Client) GetAppid() string {
	return c.opts.Appid
}

// GetSecret: get secret
func (c *Client) GetSecret() string {
	return c.opts.Secret
}

// GetAccessTokenCacheKey: get access_token cache key
func (c *Client) GetAccessTokenCacheKey() string {
	return c.accessTokenCacheKey
}

// SetAccessTokenCacheKey: set access_token cache key
func (c *Client) SetAccessTokenCacheKey(cacheKey string) error {
	c.accessTokenCacheKey = cacheKey

	return nil
}

// GetAccessTokenHandler: get access_token handler
func (c *Client) GetAccessTokenHandler() token.IToken {
	return c.accessTokenHandler
}

// SetAccessTokenHandler: set access_token handler
func (c *Client) SetAccessTokenHandler(handler token.IToken) error {
	c.accessTokenHandler = handler

	return nil
}

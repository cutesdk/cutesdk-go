package app

import (
	"time"

	"github.com/cutesdk/cutesdk-go/common/cache"
	"github.com/cutesdk/cutesdk-go/common/request"
	"github.com/cutesdk/cutesdk-go/common/token"
)

// Options: custom options
type Options struct {
	requestOptions      request.Options
	requestClient       *request.Client
	cacheHandler        cache.ICache
	accessTokenCacheKey string
	accessTokenHandler  token.IToken
	jsapiTicketCacheKey string
	jsapiTicketHandler  token.IToken
}

// OptionFunc is a function to init options
type OptionFunc func(*Options) error

// WithBaseUri set api base uri
func WithBaseUri(baseUri string) OptionFunc {
	return func(o *Options) error {
		o.requestOptions.BaseUri = baseUri

		return nil
	}
}

// WithDebug open debug mode
func WithDebug(debug bool) OptionFunc {
	return func(o *Options) error {
		o.requestOptions.Debug = debug

		return nil
	}
}

// WithTimeout: set request timeout
func WithTimeout(timeout time.Duration) OptionFunc {
	return func(o *Options) error {
		o.requestOptions.Timeout = timeout

		return nil
	}
}

// WithCache: set cache option
func WithCache(driver string, conf map[string]interface{}) OptionFunc {
	return func(o *Options) error {
		if driver == "redis" {
			if cache, err := cache.NewRedisCache(conf); err != nil {
				return err
			} else {
				o.cacheHandler = cache
			}
		}

		if driver == "file" {
			if cache, err := cache.NewFileCache(conf); err != nil {
				return err
			} else {
				o.cacheHandler = cache
			}
		}

		return nil
	}
}

// SetAccessTokenCacheKey: set access_token cache key
func (c *Client) SetAccessTokenCacheKey(cacheKey string) error {
	c.opts.accessTokenCacheKey = cacheKey

	return nil
}

// GetAccessTokenCacheKey: get access_token cache key
func (c *Client) GetAccessTokenCacheKey() string {
	return c.opts.accessTokenCacheKey
}

// SetAccessTokenHandler: set access_token handler
func (c *Client) SetAccessTokenHandler(handler token.IToken) error {
	c.opts.accessTokenHandler = handler

	return nil
}

// GetAccessTokenHandler: get access_token handler
func (c *Client) GetAccessTokenHandler() token.IToken {
	return c.opts.accessTokenHandler
}

// SetJsapiTicketCacheKey: set jsapi_ticket cache key
func (c *Client) SetJsapiTicketCacheKey(cacheKey string) error {
	c.opts.jsapiTicketCacheKey = cacheKey

	return nil
}

// GetJsapiTicketCacheKey: get jsapi_ticket cache key
func (c *Client) GetJsapiTicketCacheKey() string {
	return c.opts.jsapiTicketCacheKey
}

// SetJsapiTicketHandler: set jsapi_ticket handler
func (c *Client) SetJsapiTicketHandler(handler token.IToken) error {
	c.opts.jsapiTicketHandler = handler

	return nil
}

// GetJsapiTicketHandler: get jsapi_ticket handler
func (c *Client) GetJsapiTicketHandler() token.IToken {
	return c.opts.jsapiTicketHandler
}

// SetCacheHandler: set cache handler
func (c *Client) SetCacheHandler(handler cache.ICache) error {
	c.opts.cacheHandler = handler

	return nil
}

// GetCacheHandler: get cache handler
func (c *Client) GetCacheHandler() cache.ICache {
	return c.opts.cacheHandler
}

// GetRequestClient: get request handler
func (c *Client) GetRequestClient() *request.Client {
	return c.opts.requestClient
}

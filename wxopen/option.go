package wxopen

import (
	"github.com/cutesdk/cutesdk-go/common/cache"
	"github.com/cutesdk/cutesdk-go/common/request"
	"github.com/cutesdk/cutesdk-go/common/token"
)

// Options: custom options
type Options struct {
	Request            *request.Options
	Cache              *cache.Options
	ComponentAppid     string
	ComponentAppsecret string
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

// GetComponentAppid: get component_appid
func (c *Client) GetComponentAppid() string {
	return c.opts.ComponentAppid
}

// GetComponentAppsecret get component_appsecret
func (c *Client) GetComponentAppsecret() string {
	return c.opts.ComponentAppsecret
}

// GetComponentAccessTokenCacheKey: get component_access_token cache key
func (c *Client) GetComponentAccessTokenCacheKey() string {
	return c.componentAccessTokenCacheKey
}

// SetComponentAccessTokenCacheKey: set component_access_token cache key
func (c *Client) SetComponentAccessTokenCacheKey(cacheKey string) error {
	c.componentAccessTokenCacheKey = cacheKey

	return nil
}

// GetComponentAccessTokenHandler: get component_access_token handler
func (c *Client) GetComponentAccessTokenHandler() token.IToken {
	return c.componentAccessTokenHandler
}

// SetComponentAccessTokenHandler: set component_access_token handler
func (c *Client) SetComponentAccessTokenHandler(handler token.IToken) error {
	c.componentAccessTokenHandler = handler

	return nil
}

// GetComponentVerifyTicketCacheKey: get component_verify_ticket cache key
func (c *Client) GetComponentVerifyTicketCacheKey() string {
	return c.componentVerifyTicketCacheKey
}

// SetComponentVerifyTicketCacheKey: set component_verify_ticket cache key
func (c *Client) SetComponentVerifyTicketCacheKey(cacheKey string) error {
	c.componentVerifyTicketCacheKey = cacheKey

	return nil
}

// GetComponentVerifyTicketHandler: get component_verify_ticket handler
func (c *Client) GetComponentVerifyTicketHandler() token.IToken {
	return c.componentVerifyTicketHandler
}

// SetComponentVerifyTicketHandler: set component_verify_ticket handler
func (c *Client) SetComponentVerifyTicketHandler(handler token.IToken) error {
	c.componentVerifyTicketHandler = handler

	return nil
}

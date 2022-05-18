package wxpay

import (
	"github.com/cutesdk/cutesdk-go/common/cache"
	"github.com/cutesdk/cutesdk-go/common/request"
)

// Options: custom options
type Options struct {
	Request   *request.Options
	Cache     *cache.Options
	MchId     string
	Appid     string
	SubMchId  string
	SubAppid  string
	ApiKey    string
	NotifyUrl string
}

// GetRequestClient: get request handler
func (c *Client) GetRequestClient() *request.Client {
	return c.requestClient
}

// GetMchId: get mch_id
func (c *Client) GetMchId() string {
	return c.opts.MchId
}

// GetAppid: get appid
func (c *Client) GetAppid() string {
	return c.opts.Appid
}

// GetApiKey: get api_key
func (c *Client) GetApiKey() string {
	return c.opts.ApiKey
}

// GetNotifyUrl: get notify_url
func (c *Client) GetNotifyUrl() string {
	return c.opts.NotifyUrl
}

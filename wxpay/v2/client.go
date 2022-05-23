package wxpay

import (
	"fmt"
	"time"

	"github.com/cutesdk/cutesdk-go/common/request"
	"github.com/idoubi/goutils"
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

// BuildParams: build params with sign
func (c *Client) BuildParams(params map[string]interface{}) (map[string]interface{}, error) {
	if params == nil {
		return nil, fmt.Errorf("invalid params")
	}

	signType := "MD5"
	if v, ok := params["sign_type"]; ok {
		signType = v.(string)
	}
	if signType != "MD5" && signType != "HMAC-SHA256" {
		return nil, fmt.Errorf("invalid sign_type")
	}

	params["mch_id"] = c.opts.MchId
	params["appid"] = c.opts.Appid
	params["nonce_str"] = goutils.NonceStr(32)
	params["sign_type"] = signType

	if signType == "HMAC-SHA256" {
		params["sign"] = SignWithHmacSha256(params, c.opts.ApiKey)
	} else {
		params["sign"] = SignWithMd5(params, c.opts.ApiKey)
	}

	return params, nil
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

// GetRequestClient: get request handler
func (c *Client) GetRequestClient() *request.Client {
	return c.requestClient
}

// GetOptions: get options
func (c *Client) GetOptions() *Options {
	return c.opts
}

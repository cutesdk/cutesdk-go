package wxpay

import (
	"context"
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"

	"github.com/cutesdk/cutesdk-go/common/request"
	"github.com/spf13/cast"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/validators"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

// Client: wxpay client
type Client struct {
	opts      *Options
	payClient *core.Client
	ctx       context.Context
}

// NewClient: new wxpay client
func NewClient(opts *Options) (*Client, error) {
	// set default options
	if opts.Timeout <= 0 {
		opts.Timeout = 30 * time.Second
	}
	if opts.BaseUri == "" {
		opts.BaseUri = "https://api.mch.weixin.qq.com"
	}

	// new client
	cli := &Client{opts: opts, ctx: context.Background()}

	payClient, err := cli.newPayClient()
	if err != nil {
		return nil, err
	}

	cli.payClient = payClient

	return cli, nil
}

// NewPayClient: new wxpay v3 client
func (cli *Client) newPayClient() (*core.Client, error) {
	opts := cli.opts
	if opts == nil || opts.MchId == "" || opts.ApiKey == "" || opts.SerialNo == "" {
		return nil, fmt.Errorf("invalid mch options")
	}

	// 私钥证书内容和证书路径不能同时为空
	if opts.KeyPem == "" && opts.KeyPath == "" {
		return nil, fmt.Errorf("invalid mch key or path")
	}

	var privateKey *rsa.PrivateKey
	var err error

	if opts.KeyPem != "" {
		privateKey, err = utils.LoadPrivateKey(opts.KeyPem)
	} else {
		privateKey, err = utils.LoadPrivateKeyWithPath(opts.KeyPath)
	}

	opts.privateKey = privateKey

	if err != nil {
		return nil, fmt.Errorf("load mch key failed: %v", err)
	}

	httpClient := &http.Client{
		Timeout: opts.Timeout,
	}

	clientOptions := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(opts.MchId, opts.SerialNo, privateKey, opts.ApiKey),
		option.WithHTTPClient(httpClient),
	}

	payClient, err := core.NewClient(cli.ctx, clientOptions...)
	if err != nil {
		return nil, fmt.Errorf("new wxpay client failed: %v", err)
	}

	return payClient, nil
}

// WithoutValidator: create new client without validate response data
func (cli *Client) WithoutValidator() *Client {
	cli.payClient = core.NewClientWithValidator(cli.payClient, &validators.NullValidator{})

	return cli
}

// GetOptions: return Client options
func (cli *Client) GetOptions() *Options {
	return cli.opts
}

// GetMchId: get mch_id
func (cli *Client) GetMchId() string {
	return cli.opts.MchId
}

// GetApiKey: get api_key
func (cli *Client) GetApiKey() string {
	return cli.opts.ApiKey
}

// AppendParamsToUri: append params to request uri
func (cli *Client) AppendParamsToUri(uri string, params map[string]interface{}) (string, error) {
	if len(params) == 0 {
		return uri, nil
	}

	u, err := url.Parse(uri)
	if err != nil {
		return uri, err
	}

	q := u.Query()
	for k, v := range params {
		if k == "" || v == nil {
			continue
		}
		if str, ok := v.(string); ok {
			q.Set(k, str)
			continue
		}

		if str := cast.ToString(v); str != "" {
			q.Set(k, str)
		}
	}

	u.RawQuery = q.Encode()
	uri = u.String()

	return uri, nil
}

// Get: request api with get method
func (cli *Client) Get(uri string, args ...map[string]interface{}) (*request.Result, error) {
	return cli.Request(http.MethodGet, uri, args...)
}

// Post: request api with post method
func (cli *Client) Post(uri string, args ...map[string]interface{}) (*request.Result, error) {
	return cli.Request(http.MethodPost, uri, args...)
}

// Put: request api with put method
func (cli *Client) Put(uri string, args ...map[string]interface{}) (*request.Result, error) {
	return cli.Request(http.MethodPut, uri, args...)
}

// Patch: request api with patch method
func (cli *Client) Patch(uri string, args ...map[string]interface{}) (*request.Result, error) {
	return cli.Request(http.MethodPatch, uri, args...)
}

// Delete: request api with delete method
func (cli *Client) Delete(uri string, args ...map[string]interface{}) (*request.Result, error) {
	return cli.Request(http.MethodDelete, uri, args...)
}

// Request: request api
func (cli *Client) Request(method, uri string, args ...map[string]interface{}) (*request.Result, error) {
	var res *core.APIResult
	var err error

	params := map[string]interface{}{}
	if len(args) > 0 {
		params = args[0]
	}

	headers := map[string]interface{}{}
	if len(args) > 1 {
		headers = args[1]
	}

	httpHeaders := http.Header{}
	for k, v := range headers {
		if k == "" || v == nil {
			continue
		}
		if str, ok := v.(string); ok {
			httpHeaders.Set(k, str)
			continue
		}
		if str := cast.ToString(v); str != "" {
			httpHeaders.Set(k, str)
		}
	}

	if !strings.HasPrefix(uri, "http") {
		uri = cli.opts.BaseUri + uri
	}

	if method == http.MethodGet {
		uri, _ = cli.AppendParamsToUri(uri, params)
		res, err = cli.payClient.Get(cli.ctx, uri)
	} else {
		res, err = cli.payClient.Request(
			cli.ctx,
			method,
			uri,
			httpHeaders,
			nil,
			params,
			"application/json",
		)
	}

	if cli.opts.Debug {
		var reqInfo, respInfo []byte
		if dump, err := httputil.DumpRequest(res.Request, true); err == nil {
			reqInfo = dump
		}
		if dump, err := httputil.DumpResponse(res.Response, true); err == nil {
			respInfo = dump
		}
		log.Printf("\n------request info------\n%s\n------response info------\n%s\n", reqInfo, respInfo)
	}

	if err != nil {
		return nil, err
	}

	defer res.Response.Body.Close()

	body, err := ioutil.ReadAll(res.Response.Body)
	if err != nil {
		return nil, fmt.Errorf("parse response body failed: %v", err)
	}

	return request.NewResult(body), nil
}

// UnwrapError: unwrap api error
func (cli *Client) UnwrapError(err error) *core.APIError {
	return UnwrapError(err)
}

// UnwrapError: unwrap api error
func UnwrapError(err error) *core.APIError {
	if e, ok := err.(*core.APIError); ok {
		return e
	}

	return nil
}

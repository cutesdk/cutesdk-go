package wxpay

import (
	"context"
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/cutesdk/cutesdk-go/common/request"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

// Instance: wxpay instance
type Instance struct {
	opts          *Options
	requestClient *request.Client
	payClient     *core.Client
	ctx           context.Context
}

// New create wxpay instance
func New(opts *Options) (*Instance, error) {
	// set default options
	if opts.Request == nil {
		opts.Request = &request.Options{
			Debug: opts.Debug,
		}
	}
	if opts.Request.BaseUri == "" {
		opts.Request.BaseUri = "https://api.mch.weixin.qq.com"
	}
	if opts.Request.Timeout == 0 {
		opts.Request.Timeout = 5 * time.Second
	}

	// new instance
	ins := &Instance{opts: opts, ctx: context.Background()}

	// set request client
	ins.requestClient = request.NewClient(opts.Request)

	if payClient, err := NewPayClient(opts); err != nil {
		return nil, err
	} else {
		ins.payClient = payClient
	}

	return ins, nil
}

// NewPayClient: new wxpay v3 client
func NewPayClient(opts *Options) (*core.Client, error) {
	if opts == nil || opts.MchId == "" || opts.ApiKey == "" || opts.SerialNo == "" {
		return nil, fmt.Errorf("invalid mch options")
	}

	// 私钥证书内容和证书路径不能同时为空
	if opts.PrivateKey == "" && opts.PrivateKeyPath == "" {
		return nil, fmt.Errorf("invalid mch key or path")
	}

	var privateKey *rsa.PrivateKey
	var err error

	if opts.PrivateKey != "" {
		privateKey, err = utils.LoadPrivateKey(opts.PrivateKey)
	} else {
		privateKey, err = utils.LoadPrivateKeyWithPath(opts.PrivateKeyPath)
	}

	if err != nil {
		return nil, fmt.Errorf("load mch key failed: %v", err)
	}

	clientOptions := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(opts.MchId, opts.SerialNo, privateKey, opts.ApiKey),
	}

	ctx := context.Background()
	payClient, err := core.NewClient(ctx, clientOptions...)
	if err != nil {
		return nil, fmt.Errorf("new wxpay client failed: %v", err)
	}

	return payClient, nil
}

// GetOptions: return instance options
func (ins *Instance) GetOptions() *Options {
	return ins.opts
}

// Get: request api with get method
func (ins *Instance) Get(uri string) (*request.Result, error) {
	return ins.Request(http.MethodGet, uri, nil)
}

// Post: request api with post method
func (ins *Instance) Post(uri string, params map[string]interface{}) (*request.Result, error) {
	return ins.Request(http.MethodPost, uri, params)
}

// Put: request api with put method
func (ins *Instance) Put(uri string, params map[string]interface{}) (*request.Result, error) {
	return ins.Request(http.MethodPut, uri, params)
}

// Request: request api
func (ins *Instance) Request(method, uri string, params map[string]interface{}) (*request.Result, error) {
	var res *core.APIResult
	var err error

	if !strings.HasPrefix(uri, "http") {
		uri = ins.opts.Request.BaseUri + uri
	}

	// make request
	switch method {
	case http.MethodGet:
		res, err = ins.payClient.Get(ins.ctx, uri)
	case http.MethodPost:
		res, err = ins.payClient.Post(ins.ctx, uri, params)
	case http.MethodPut:
		res, err = ins.payClient.Put(ins.ctx, uri, params)
	case http.MethodPatch:
		res, err = ins.payClient.Patch(ins.ctx, uri, params)
	case http.MethodDelete:
		res, err = ins.payClient.Delete(ins.ctx, uri, params)
	default:
		return nil, fmt.Errorf("method: %s not support yet", method)
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

// UnwrapError 解压缩微信支付 APIError
func UnwrapError(err error) *core.APIError {
	if e, ok := err.(*core.APIError); ok {
		return e
	}

	return nil
}

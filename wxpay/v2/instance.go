package wxpay

import (
	"fmt"
	"time"

	"github.com/cutesdk/cutesdk-go/common/request"
	"github.com/idoubi/goutils"
	"github.com/idoubi/goz"
)

// Instance: wxpay instance
type Instance struct {
	opts          *Options
	requestClient *request.Client
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
	ins := &Instance{opts: opts}

	// set request client
	ins.requestClient = request.NewClient(opts.Request)

	return ins, nil
}

// BuildParams: build params with sign
func (ins *Instance) BuildParams(params map[string]interface{}) (map[string]interface{}, error) {
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

	params["mch_id"] = ins.opts.MchId
	params["appid"] = ins.opts.Appid
	params["nonce_str"] = goutils.NonceStr(32)
	params["sign_type"] = signType

	if signType == "HMAC-SHA256" {
		params["sign"] = SignWithHmacSha256(params, ins.opts.ApiKey)
	} else {
		params["sign"] = SignWithMd5(params, ins.opts.ApiKey)
	}

	return params, nil
}

// Post: request api with post method
func (ins *Instance) Post(uri string, params map[string]interface{}) (*request.Result, error) {
	data, err := ins.BuildParams(params)
	if err != nil {
		return nil, err
	}

	res, err := ins.GetRequestClient().PostXml(uri, data)
	if err != nil {
		return nil, err
	}

	res.XmlParsed()

	return res, nil
}

// Request: request api
func (ins *Instance) Request(method, uri string, opts goz.Options) (*request.Result, error) {
	return ins.GetRequestClient().Request(method, uri, opts)
}

// GetRequestClient: get request handler
func (ins *Instance) GetRequestClient() *request.Client {
	return ins.requestClient
}

// GetOptions: get options
func (ins *Instance) GetOptions() *Options {
	return ins.opts
}

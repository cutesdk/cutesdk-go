package wxpay

import (
	"crypto/tls"
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

// NewClient: new wxpay client
func NewClient(opts *Options) (*Client, error) {
	// set default options
	if opts.Timeout <= 0 {
		opts.Timeout = 10 * time.Second
	}
	if opts.BaseUri == "" {
		opts.BaseUri = "https://api.mch.weixin.qq.com"
	}

	// new client
	cli := &Client{opts: opts}

	reqOpts := &request.Options{
		Debug:   opts.Debug,
		Timeout: opts.Timeout,
		BaseUri: opts.BaseUri,
	}

	certificates := []tls.Certificate{}

	// load cert from pem content
	if opts.CertPem != "" && opts.KeyPem != "" {
		if crt, err := tls.X509KeyPair([]byte(opts.CertPem), []byte(opts.KeyPem)); err == nil {
			certificates = append(certificates, crt)
		}
	}

	// load cert from file path
	if opts.CertFile != "" && opts.KeyFile != "" {
		if crt, err := tls.LoadX509KeyPair(opts.CertFile, opts.KeyFile); err == nil {
			certificates = append(certificates, crt)
		}
	}

	if len(certificates) > 0 {
		reqOpts.Certificates = certificates
	}

	// set request client
	cli.requestClient = request.NewClient(reqOpts)

	return cli, nil
}

// BuildParams: build params with sign
func (cli *Client) BuildParams(uri string, params map[string]interface{}) (map[string]interface{}, error) {
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

	params["mch_id"] = cli.opts.MchId
	params["nonce_str"] = goutils.NonceStr(32)

	if signType == "HMAC-SHA256" {
		params["sign"] = SignWithHmacSha256(params, cli.opts.ApiKey)
	} else {
		params["sign"] = SignWithMd5(params, cli.opts.ApiKey)
	}

	return params, nil
}

// Post: request api with post method
func (cli *Client) Post(uri string, params map[string]interface{}) (*request.Result, error) {
	data, err := cli.BuildParams(uri, params)
	if err != nil {
		return nil, err
	}

	res, err := cli.GetRequestClient().PostXml(uri, data)
	if err != nil {
		return nil, err
	}

	res.XmlParsed()

	return res, nil
}

// Request: request api
func (cli *Client) Request(method, uri string, opts goz.Options) (*request.Result, error) {
	return cli.GetRequestClient().Request(method, uri, opts)
}

// GetRequestClient: get request handler
func (cli *Client) GetRequestClient() *request.Client {
	return cli.requestClient
}

// GetOptions: get options
func (cli *Client) GetOptions() *Options {
	return cli.opts
}

// GetMchId: get mch_id
func (cli *Client) GetMchId() string {
	return cli.opts.MchId
}

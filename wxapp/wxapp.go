package wxapp

import (
	"errors"
	"net/url"

	"github.com/idoubi/goz"
	"github.com/tidwall/gjson"
)

var (
	apiBase = "https://api.weixin.qq.com"
)

// Options 配置项
type Options struct {
	Debug          bool                         // 是否开启调试
	Appid          string                       // 小程序appid
	Secret         string                       // 小程序secret
	GetAccessToken func(*Wxapp) (string, error) // 自定义获取access_token的方法
}

// Wxapp 小程序对象
type Wxapp struct {
	opts Options
}

// Result 响应数据
type Result []byte

// New 实例化对象
func New(opts Options) *Wxapp {
	return &Wxapp{
		opts: opts,
	}
}

// Parsed 转为解析后的数据
func (r Result) Parsed() gjson.Result {
	return gjson.ParseBytes(r)
}

// String 字符串打印
func (r Result) String() string {
	return string(r)
}

// Get 获取值
func (r Result) Get(key string) gjson.Result {
	return r.Parsed().Get(key)
}

// Get 请求GET类型的API
func (w *Wxapp) Get(apiPath string, params map[string]string) (Result, error) {
	return w.RequestWithAccessToken(apiPath, "GET", params, nil)
}

// Post 请求POST类型的API
func (w *Wxapp) Post(apiPath string, params map[string]string, data map[string]interface{}) (Result, error) {
	return w.RequestWithAccessToken(apiPath, "POST", params, data)
}

// RequestWithAccessToken 请求API
func (w *Wxapp) RequestWithAccessToken(apiPath, method string, params map[string]string, data map[string]interface{}) (Result, error) {
	// 仅支持GET、POST两种请求
	if method != "GET" && method != "POST" {
		return nil, errors.New("invalid method")
	}

	if params == nil {
		params = make(map[string]string)
	}

	// 请求参数自动带上access_token
	if _, ok := params["access_token"]; !ok {
		accessToken, err := w.getAccessToken()
		if err != nil {
			return nil, err
		}
		params["access_token"] = accessToken
	}

	queryParams := url.Values{}
	for k, v := range params {
		queryParams.Set(k, v)
	}

	apiUrl := apiBase + apiPath
	if len(queryParams) > 0 {
		apiUrl += "?" + queryParams.Encode()
	}

	var resp *goz.Response
	var err error

	// Get 请求
	if method == "GET" {
		resp, err = goz.Get(apiUrl, goz.Options{
			Debug: w.opts.Debug,
		})
	}

	// Post 请求
	if method == "POST" {
		resp, err = goz.Post(apiUrl, goz.Options{
			Debug: w.opts.Debug,
			JSON:  data,
		})
	}

	// 请求失败
	if err != nil {
		return nil, err
	}

	body, err := resp.GetBody()

	return Result(body), err
}

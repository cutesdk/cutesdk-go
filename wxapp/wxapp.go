package wxapp

import "github.com/tidwall/gjson"

var (
	apiBase = "https://api.weixin.qq.com"
)

// Options 配置项
type Options struct {
	Debug     bool
	Appid     string
	Appsecret string
	GetToken  func(*Wxapp) string
}

// Token 访问凭证
type Token struct {
	AccessToken string
}

// Wxapp 小程序对象
type Wxapp struct {
	token Token
	opts  Options
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

package wxopen

import (
	"github.com/idoubi/goutils/crypt"
	"github.com/tidwall/gjson"
)

var (
	apiBase = "https://api.mch.weixin.qq.com"
)

// Options 微信开放平台参数
type Options struct {
	Debug          bool
	Appid          string
	AppSecret      string
	VerifyToken    string
	EncodingAesKey string // 43位
	aesKey         []byte // 32位
}

// WxOpen 微信对象
type WxOpen struct {
	opts Options
}

// New 初始化
func New(opts Options) *WxOpen {
	if opts.EncodingAesKey != "" {
		if v, _ := crypt.Base64Decode(opts.EncodingAesKey + "="); v != nil {
			opts.aesKey = v
		}
	}
	return &WxOpen{
		opts: opts,
	}
}

// Result 响应结果
type Result []byte

// Parsed 获取解析后的数据
func (r Result) Parsed() gjson.Result {
	return gjson.ParseBytes(r)
}

// XmlParsed 解析xml数据
func (r Result) XmlParsed() {

}

package wxpay

import "github.com/tidwall/gjson"

var (
	apiBase = "https://api.mch.weixin.qq.com"
)

// Options 微信支付参数
type Options struct {
	Debug     bool
	MchID     string
	SubMchID  string
	APIKey    string
	Appid     string
	SubAppid  string
	NotifyURL string
}

// Wxpay 微信支付对象
type Wxpay struct {
	opts Options
}

// Result 响应结果
type Result []byte

// New 初始化
func New(opts Options) *Wxpay {
	return &Wxpay{
		opts: opts,
	}
}

// Parsed 获取解析后的数据
func (r Result) Parsed() gjson.Result {
	return gjson.ParseBytes(r)
}

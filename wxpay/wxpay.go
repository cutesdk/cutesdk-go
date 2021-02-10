package wxpay

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

// New 初始化
func New(opts Options) *Wxpay {
	return &Wxpay{
		opts: opts,
	}
}

var (
	apiBase = "https://api.mch.weixin.qq.com"
)

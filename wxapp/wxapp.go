package wxapp

// Options 配置项
type Options struct {
	Appid     string
	Appsecret string
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

// New 实例化对象
func New(opts Options) *Wxapp {
	return &Wxapp{
		opts: opts,
	}
}

var (
	apiBase = "https://api.weixin.qq.com"
)

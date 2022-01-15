package wxapp

import "github.com/cutesdk/cutesdk-go/comm/cache"

var (
	apiBase = "https://api.weixin.qq.com"
)

// Options 配置项
type Options struct {
	Debug              bool                         // 是否开启调试
	Appid              string                       // 小程序appid
	AppSecret          string                       // 小程序secret
	accessTokenHandler func(*WxApp) (string, error) // 自定义获取access_token的方法
	customCache        cache.Cache                  // 自定义缓存组件
}

// WxApp 微信小程序
type WxApp struct {
	opts Options
}

// New 实例化对象
func New(opts Options) *WxApp {
	// 设置默认的缓存组件
	opts.customCache = cache.New()

	return &WxApp{
		opts: opts,
	}
}

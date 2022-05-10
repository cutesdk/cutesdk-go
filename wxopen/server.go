package wxopen

// import (
// 	"github.com/idoubi/goutils/crypt"
// )

// var (
// 	apiBase = "https://api.weixin.qq.com"
// )

// // Options 微信开放平台参数
// type Options struct {
// 	Debug          bool
// 	Appid          string
// 	AppSecret      string
// 	VerifyToken    string
// 	EncodingAesKey string // 43位
// 	aesKey         []byte // 32位
// }

// // WxOpen 微信开发平台
// type WxOpen struct {
// 	opts Options
// }

// // New 初始化
// func New(opts Options) *WxOpen {
// 	if opts.EncodingAesKey != "" {
// 		if v, _ := crypt.Base64Decode(opts.EncodingAesKey + "="); v != nil {
// 			opts.aesKey = v
// 		}
// 	}
// 	return &WxOpen{
// 		opts: opts,
// 	}
// }

package wxapp

import (
	"fmt"

	"github.com/idoubi/goz"
	"github.com/tidwall/gjson"
)

// GetAccessToken 获取access_token
func (w *Wxapp) GetAccessToken() (*gjson.Result, error) {
	apiURL := fmt.Sprintf(apiBase+"/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", w.opts.Appid, w.opts.Appsecret)
	resp, err := goz.Get(apiURL, goz.Options{
		Debug: w.opts.Debug,
	})

	if err != nil {
		return nil, err
	}

	return resp.GetParsedBody()
}

// getAccessToken 获取access_token 内部调用
func (w *Wxapp) getAccessToken() string {
	// 从缓存读取 todo
	token, err := w.GetAccessToken()
	if err != nil {
		return ""
	}

	return token.Get("access_token").String()
}
